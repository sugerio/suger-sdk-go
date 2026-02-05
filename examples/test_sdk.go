package main

import (
	"context"
	"fmt"
	"os"

	suger "github.com/sugerio/suger-sdk-go"
)

func main() {
	// Get API key and org ID from environment or command line
	apiKey := os.Getenv("SUGER_API_KEY")
	orgId := os.Getenv("SUGER_ORG_ID")
	apiURL := os.Getenv("SUGER_API_URL") // Optional: defaults to https://api.suger.cloud

	if apiKey == "" {
		fmt.Println("Error: SUGER_API_KEY environment variable is required")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("  Production: SUGER_API_KEY=your_key SUGER_ORG_ID=your_org_id go run test_sdk.go")
		fmt.Println("  Dev:        SUGER_API_KEY=your_key SUGER_ORG_ID=your_org_id SUGER_API_URL=https://api.dev.suger.cloud go run test_sdk.go")
		os.Exit(1)
	}

	if orgId == "" {
		fmt.Println("Error: SUGER_ORG_ID environment variable is required")
		os.Exit(1)
	}

	// Create configuration
	configuration := suger.NewConfiguration()

	// Override server URL if provided (for dev/staging environments)
	if apiURL != "" {
		configuration.Servers = suger.ServerConfigurations{
			{URL: apiURL, Description: "Custom API server"},
		}
	}

	// Create API client
	apiClient := suger.NewAPIClient(configuration)

	// Set up context with API key authentication
	// The API expects: Authorization: Key <api_key>
	ctx := context.WithValue(context.Background(), suger.ContextAPIKeys, map[string]suger.APIKey{
		"APIKeyAuth": {Key: apiKey, Prefix: "Key"},
	})

	fmt.Println("Testing Suger Go SDK v3.128.222")
	fmt.Println("================================")
	serverURL, _ := configuration.ServerURL(0, nil)
	fmt.Printf("API URL: %s\n", serverURL)
	fmt.Printf("Org ID:  %s\n\n", orgId)

	// Test 1: List API Clients
	fmt.Println("Test 1: Listing API Clients...")
	apiClients, httpRes, err := apiClient.APIAPI.ListApiClients(ctx, orgId).Execute()
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
		if httpRes != nil {
			fmt.Printf("  HTTP Status: %d\n", httpRes.StatusCode)
		}
	} else {
		fmt.Printf("  Success! HTTP Status: %d\n", httpRes.StatusCode)
		fmt.Printf("  Found %d API client(s)\n", len(apiClients))
	}

	// Test 2: List Products
	fmt.Println("\nTest 2: Listing Products...")
	products, httpRes, err := apiClient.ProductAPI.ListProducts(ctx, orgId).Execute()
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
		if httpRes != nil {
			fmt.Printf("  HTTP Status: %d\n", httpRes.StatusCode)
		}
	} else {
		fmt.Printf("  Success! HTTP Status: %d\n", httpRes.StatusCode)
		fmt.Printf("  Found %d product(s)\n", len(products))
	}

	// Test 3: List Buyers
	fmt.Println("\nTest 3: Listing Buyers...")
	buyers, httpRes, err := apiClient.BuyerAPI.ListBuyers(ctx, orgId).Execute()
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
		if httpRes != nil {
			fmt.Printf("  HTTP Status: %d\n", httpRes.StatusCode)
		}
	} else {
		fmt.Printf("  Success! HTTP Status: %d\n", httpRes.StatusCode)
		fmt.Printf("  Found %d buyer(s)\n", len(buyers))
	}

	// Test 4: List Entitlements
	fmt.Println("\nTest 4: Listing Entitlements...")
	entitlements, httpRes, err := apiClient.EntitlementAPI.ListEntitlements(ctx, orgId).Execute()
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
		if httpRes != nil {
			fmt.Printf("  HTTP Status: %d\n", httpRes.StatusCode)
		}
	} else {
		fmt.Printf("  Success! HTTP Status: %d\n", httpRes.StatusCode)
		fmt.Printf("  Found %d entitlement(s)\n", len(entitlements))
	}

	fmt.Println("\n================================")
	fmt.Println("SDK test completed!")
}
