package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/crodriguezde/go-kusto/pkg/client"
)

// main is the entry point of the program.
func main() {

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Println(fmt.Errorf("failed to create credential: %v", err))
	}

	c, err := client.New(
		client.WithTokenCredential(cred),
		client.WithEndpoint("https://adeeventmapperprod.eastus2.kusto.windows.net"),
	)

	if err != nil {
		fmt.Printf("failed to create client: %v\n", err)
		return
	}

	err = c.Query(context.Background())
	if err != nil {
		fmt.Printf("failed to query: %v\n", err)
		return
	}
}
