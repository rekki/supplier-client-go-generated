package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/rekki/supplier-client-go-generated/client"
)

func main() {
	key := flag.String("k", "", "api key")
	flag.Parse()

	if key == nil || *key == "" {
		fmt.Println("Please specify api key")
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	auth := context.WithValue(context.Background(), client.ContextAPIKey, client.APIKey{
		Key:    *key,
		Prefix: "Bearer",
	})

	configuration := client.NewConfiguration()

	apiClient := client.NewAPIClient(configuration)

	resp, _, err := apiClient.CatalogApi.GetCatalogItemsV3(auth, client.RekkiAuthorizationType)

	fmt.Printf("resp = %#v, err = %#v", resp, err)
}
