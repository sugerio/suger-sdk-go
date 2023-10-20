# \OperationAPI

All URIs are relative to *http://https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOperations**](OperationAPI.md#ListOperations) | **Get** /org/{orgId}/operation | list operations



## ListOperations

> []Operation ListOperations(ctx, orgId).OfferId(offerId).EntitlementId(entitlementId).CrmOpportunityId(crmOpportunityId).PartnerOpportunityId(partnerOpportunityId).Execute()

list operations



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
    offerId := "offerId_example" // string | filter by offerId (optional)
    entitlementId := "entitlementId_example" // string | filter by entitlementId (optional)
    crmOpportunityId := "crmOpportunityId_example" // string | filter by crmOpportunityId (optional)
    partnerOpportunityId := "partnerOpportunityId_example" // string | filter by partnerOpportunityId (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperationAPI.ListOperations(context.Background(), orgId).OfferId(offerId).EntitlementId(entitlementId).CrmOpportunityId(crmOpportunityId).PartnerOpportunityId(partnerOpportunityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.ListOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOperations`: []Operation
    fmt.Fprintf(os.Stdout, "Response from `OperationAPI.ListOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offerId** | **string** | filter by offerId | 
 **entitlementId** | **string** | filter by entitlementId | 
 **crmOpportunityId** | **string** | filter by crmOpportunityId | 
 **partnerOpportunityId** | **string** | filter by partnerOpportunityId | 

### Return type

[**[]Operation**](Operation.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

