# \EntitlementTermAPI

All URIs are relative to *http://https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEntitlementTerm**](EntitlementTermAPI.md#GetEntitlementTerm) | **Get** /org/{orgId}/entitlement/{entitlementId}/entitlementTerm/{entitlementTermId} | get entitlement term
[**ListEntitlementTerms**](EntitlementTermAPI.md#ListEntitlementTerms) | **Get** /org/{orgId}/entitlement/{entitlementId}/entitlementTerm | list entitlement terms



## GetEntitlementTerm

> WorkloadEntitlementTerm GetEntitlementTerm(ctx, orgId, entitlementId, entitlementTermId).Execute()

get entitlement term



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
    entitlementId := "entitlementId_example" // string | Entitlement ID
    entitlementTermId := "entitlementTermId_example" // string | Entitlement Term ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementTermAPI.GetEntitlementTerm(context.Background(), orgId, entitlementId, entitlementTermId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementTermAPI.GetEntitlementTerm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntitlementTerm`: WorkloadEntitlementTerm
    fmt.Fprintf(os.Stdout, "Response from `EntitlementTermAPI.GetEntitlementTerm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 
**entitlementTermId** | **string** | Entitlement Term ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementTermRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkloadEntitlementTerm**](WorkloadEntitlementTerm.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlementTerms

> []WorkloadEntitlementTerm ListEntitlementTerms(ctx, orgId, entitlementId).Execute()

list entitlement terms



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
    entitlementId := "entitlementId_example" // string | Entitlement ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementTermAPI.ListEntitlementTerms(context.Background(), orgId, entitlementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementTermAPI.ListEntitlementTerms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntitlementTerms`: []WorkloadEntitlementTerm
    fmt.Fprintf(os.Stdout, "Response from `EntitlementTermAPI.ListEntitlementTerms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementTermsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WorkloadEntitlementTerm**](WorkloadEntitlementTerm.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

