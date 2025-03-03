# \APIAPI

All URIs are relative to *http://https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiClient**](APIAPI.md#GetApiClient) | **Get** /org/{orgId}/apiClient/{apiClientId} | get api client
[**ListApiClients**](APIAPI.md#ListApiClients) | **Get** /org/{orgId}/apiClient | list api clients



## GetApiClient

> GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibIdentityApiClient GetApiClient(ctx, orgId, apiClientId).Execute()

get api client



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sugerio/suger-sdk-go"
)

func main() {
	orgId := "orgId_example" // string | Organization ID
	apiClientId := "apiClientId_example" // string | API client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAPI.GetApiClient(context.Background(), orgId, apiClientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAPI.GetApiClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiClient`: GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibIdentityApiClient
	fmt.Fprintf(os.Stdout, "Response from `APIAPI.GetApiClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**apiClientId** | **string** | API client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibIdentityApiClient
**](GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibIdentityApiClient.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiClients

> []GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibIdentityApiClient ListApiClients(ctx, orgId).Execute()

list api clients



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sugerio/suger-sdk-go"
)

func main() {
	orgId := "orgId_example" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAPI.ListApiClients(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAPI.ListApiClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApiClients`: []GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibIdentityApiClient
	fmt.Fprintf(os.Stdout, "Response from `APIAPI.ListApiClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApiClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibIdentityApiClient
**](GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibIdentityApiClient.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

