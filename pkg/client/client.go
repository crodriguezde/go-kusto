package client

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type Client struct {
	cred azcore.TokenCredential
}

// main is the entry point of the program.
func main() {
	client := Client{}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Errorf("failed to create credential: %v", err)
	}

	client.cred = cred

}
