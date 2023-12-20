package client

import "github.com/Azure/azure-sdk-for-go/sdk/azidentity"

type ClentOption func(c *Client)

func WithDefaultAzureCredential(cred *azidentity.DefaultAzureCredential) ClentOption {
	return func(c *Client) {
		c.cred = cred
	}
}

func WithChainedTokenCredential(cred *azidentity.ChainedTokenCredential) ClentOption {
	return func(c *Client) {
		c.cred = cred
	}
}

func WithEnvironmentCredential(cred *azidentity.EnvironmentCredential) ClentOption {
	return func(c *Client) {
		c.cred = cred
	}
}

func WithManagedIdentityCredential(cred *azidentity.ManagedIdentityCredential) ClentOption {
	return func(c *Client) {
		c.cred = cred
	}
}

func WithWorkloadIdentityCredential(cred *azidentity.WorkloadIdentityCredential) ClentOption {
	return func(c *Client) {
		c.cred = cred
	}
}

func WithClientAssertionCredential(cred *azidentity.ClientAssertionCredential) ClentOption {
	return func(c *Client) {
		c.cred = cred
	}
}

func WithClientCertificateCredential(cred *azidentity.ClientCertificateCredential) ClentOption {
	return func(c *Client) {
		c.cred = cred
	}
}

func WithClientSecretCredential(cred *azidentity.ClientSecretCredential) ClentOption {
	return func(c *Client) {
		c.cred = cred
	}
}

func WithInteractiveBrowserCredential(cred *azidentity.InteractiveBrowserCredential) ClentOption {
	return func(c *Client) {
		c.cred = cred
	}
}

func WithDeviceCodeCredential(cred *azidentity.DeviceCodeCredential) ClentOption {
	return func(c *Client) {
		c.cred = cred
	}
}

func WithUsernamePasswordCredential(cred *azidentity.UsernamePasswordCredential) ClentOption {
	return func(c *Client) {
		c.cred = cred
	}
}
