# \APIAPI

All URIs are relative to *http://https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiClient**](APIAPI.md#GetApiClient) | **Get** /org/{orgId}/apiClient/{apiClientId} | get api client
[**GetApiClientAccessToken**](APIAPI.md#GetApiClientAccessToken) | **Post** /public/apiClient/accessToken | get api access token
[**ListApiClients**](APIAPI.md#ListApiClients) | **Get** /org/{orgId}/apiClient | list api clients



## GetApiClient

> GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient GetApiClient(ctx, orgId, apiClientId).Execute()

get api client



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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
	// response from `GetApiClient`: GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient
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

[**GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient**](GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiClientAccessToken

> ApiClientAccessToken GetApiClientAccessToken(ctx).Data(data).Execute()

get api access token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	data := *openapiclient.NewGetApiClientAccessTokenParams("Id_example", "OrganizationID_example", "Secret_example") // GetApiClientAccessTokenParams | Suger API Client

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAPI.GetApiClientAccessToken(context.Background()).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAPI.GetApiClientAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiClientAccessToken`: ApiClientAccessToken
	fmt.Fprintf(os.Stdout, "Response from `APIAPI.GetApiClientAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiClientAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**GetApiClientAccessTokenParams**](GetApiClientAccessTokenParams.md) | Suger API Client | 

### Return type

[**ApiClientAccessToken**](ApiClientAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiClients

> []GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient ListApiClients(ctx, orgId).Execute()

list api clients



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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
	// response from `ListApiClients`: []GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient
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

[**[]GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient**](GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

