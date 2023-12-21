package conn

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/crodriguezde/go-kusto/pkg/errors"
	"github.com/crodriguezde/go-kusto/pkg/query"
)

var metadataPath = "/v1/rest/auth/metadata"

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

type Conn struct {
	auth          azcore.TokenCredential
	endpoint      *url.URL
	queryURL      *url.URL
	client        *http.Client
	scope         []string
	clientOptions *azcore.ClientOptions
	appId         string
}

type Metadata struct {
	AzureAD CloudInfo `json:"AzureAD"`
}

type CloudInfo struct {
	LoginEndpoint          string `json:"LoginEndpoint"`
	LoginMfaRequired       bool   `json:"LoginMfaRequired"`
	KustoClientAppID       string `json:"KustoClientAppId"`
	KustoClientRedirectURI string `json:"KustoClientRedirectUri"`
	KustoServiceResourceID string `json:"KustoServiceResourceId"`
	FirstPartyAuthorityURL string `json:"FirstPartyAuthorityUrl"`
}

type QueryMsg struct {
	DB  string `json:"db"`
	CSL string `json:"csl"`
}

func (c *Conn) QueryAzureMetadataEndpoint() (*CloudInfo, error) {
	url := c.endpoint.JoinPath(metadataPath)

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, errors.ErrWrapf(err, "failed to get metadata")
	}

	// Handle internal server error as a special case and return as an error (to be consistent with other SDK's)
	if resp.StatusCode >= 300 && resp.StatusCode != 404 {
		return nil, fmt.Errorf("error %s when querying endpoint %s", resp.Status, url.String())
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(b) == 0 {
		return nil, fmt.Errorf("empty response body when querying endpoint %s", url.String())
	}

	ci, err := UnmarshalCloudInfo(b)

	if err != nil {
		return nil, errors.ErrWrapf(err, "failed to unmarshal cloud info: %v", string(b))
	}

	return ci, nil
}

func UnmarshalCloudInfo(b []byte) (*CloudInfo, error) {
	md := &Metadata{}
	if err := json.Unmarshal(b, md); err != nil {
		return nil, err
	}
	return &md.AzureAD, nil
}

func (c *Conn) NewClientOptions(ci *CloudInfo) *azcore.ClientOptions {
	clientOptions := &azcore.ClientOptions{
		Transport: c.client,
		Retry: policy.RetryOptions{
			MaxRetries: 3,
			TryTimeout: 10,
		},
		Cloud: cloud.Configuration{
			ActiveDirectoryAuthorityHost: ci.LoginEndpoint,
		},
	}

	return clientOptions
}

// NewConn returns a new Conn object with an injected http.Client
func NewConn(endpoint string, cred azcore.TokenCredential, client *http.Client) (*Conn, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.ErrWrapf(err, "failed to parse endpoint")
	}

	if endpoint == "" {
		return nil, errors.ErrWrapf(err, "endpoint cannot be empty")
	}

	if !strings.HasPrefix(u.Path, "/") {
		u.Path = "/" + u.Path
	}

	c := &Conn{
		auth:     cred,
		queryURL: u.JoinPath("/v2/rest/query"),
		client:   client,
		endpoint: u,
	}

	cloudInfo, err := c.QueryAzureMetadataEndpoint()
	if err != nil {
		return nil, err
	}

	c.SetScope(cloudInfo)

	clientOptions := c.NewClientOptions(cloudInfo)

	c.clientOptions = clientOptions

	c.appId = cloudInfo.KustoClientAppID

	return c, nil
}

func (c *Conn) SetScope(ci *CloudInfo) {
	resourceURI := ci.KustoServiceResourceID
	if ci.LoginMfaRequired {
		resourceURI = strings.Replace(resourceURI, ".kusto.", ".kustomfa.", 1)
	}
	c.scope = []string{fmt.Sprintf("%s/.default", resourceURI)}
}

func (c *Conn) Query(ctx context.Context, db string, query string, options *query.QueryOptions) error {
	//c.conn.Query(ctx, "eventmapper", "transaction_logs | take 10", nil)
	token, err := c.auth.GetToken(ctx, policy.TokenRequestOptions{
		Scopes: c.scope,
	})

	if err != nil {
		return err
	}

	headers := c.getHeaders()
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", token.Token))

	buff := bufferPool.Get().(*bytes.Buffer)
	buff.Reset()
	defer bufferPool.Put(buff)

	err = json.NewEncoder(buff).Encode(
		QueryMsg{
			DB:  db,
			CSL: query,
		},
	)
	if err != nil {
		return errors.ErrWrapf(err, "failed to encode query")
	}

	req := &http.Request{
		Method: http.MethodPost,
		URL:    c.queryURL,
		Header: headers,
		Body:   io.NopCloser(buff),
	}

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	switch enc := strings.ToLower(resp.Header.Get("Content-Encoding")); enc {
	case "":
		fmt.Printf("resp.Body: %v\n", resp.Body)
	case "gzip":
		var err error
		wrapper, err := gzip.NewReader(resp.Body)
		if err != nil {
			return fmt.Errorf("gzip reader error: %w", err)
		}
		defer wrapper.Close()

		all, e := io.ReadAll(wrapper)
		if e != nil {
			return fmt.Errorf("gzip reader error: %w", e)
		}
		fmt.Printf("gzip all: %s\n", string(all))

	case "deflate":
		wrapper := flate.NewReader(resp.Body)
		all, e := io.ReadAll(wrapper)
		if err != nil {
			return errors.ErrWrapf(e, "deflate reader error")
		}
		fmt.Printf("deflate all: %v\n", string(all))
		defer wrapper.Close()
	default:
		return fmt.Errorf("Content-Encoding was unrecognized: %s", enc)
	}

	return nil
}

func (c *Conn) Scope() []string {
	return c.scope
}

func (c *Conn) getHeaders( /*properties requestProperties*/ ) http.Header {
	header := http.Header{}
	header.Add("Accept", "application/json")
	header.Add("Accept-Encoding", "gzip, deflate")
	header.Add("Content-Type", "application/json; charset=utf-8")
	header.Add("Connection", "Keep-Alive")
	header.Add("x-ms-version", "2019-02-13")
	/*
		if properties.ClientRequestID != "" {
			header.Add(ClientRequestIdHeader, properties.ClientRequestID)
		} else {
			header.Add(ClientRequestIdHeader, "KGC.execute;"+uuid.New().String())
		}

		if properties.Application != "" {
			header.Add(ApplicationHeader, properties.Application)
		} else {
			header.Add(ApplicationHeader, c.clientDetails.ApplicationForTracing())
		}

		if properties.User != "" {
			header.Add(UserHeader, properties.User)
		} else {
			header.Add(UserHeader, c.clientDetails.UserNameForTracing())
		}

		header.Add(ClientVersionHeader, c.clientDetails.ClientVersionForTracing())
	*/
	return header
}
