# \CosellAPI

All URIs are relative to *http://https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCosellOppReferrals**](CosellAPI.md#ListCosellOppReferrals) | **Post** /org/{orgId}/cosell/oppReferral/query | list cosell opp referrals



## ListCosellOppReferrals

> ListCosellOppReferralsResponse ListCosellOppReferrals(ctx, orgId).Type_(type_).Data(data).Execute()

list cosell opp referrals



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
    type_ := "type__example" // string | Referral type: inbox, outbox, insync, archived (optional)
    data := []openapiclient.SqlCondition{*openapiclient.NewSqlCondition()} // []SqlCondition | Query conditions (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CosellAPI.ListCosellOppReferrals(context.Background(), orgId).Type_(type_).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CosellAPI.ListCosellOppReferrals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCosellOppReferrals`: ListCosellOppReferralsResponse
    fmt.Fprintf(os.Stdout, "Response from `CosellAPI.ListCosellOppReferrals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCosellOppReferralsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Referral type: inbox, outbox, insync, archived | 
 **data** | [**[]SqlCondition**](SqlCondition.md) | Query conditions | 

### Return type

[**ListCosellOppReferralsResponse**](ListCosellOppReferralsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

