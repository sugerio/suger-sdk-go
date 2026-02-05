# \AuditingAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryAuditingEvents**](AuditingAPI.md#QueryAuditingEvents) | **Get** /org/{orgId}/auditingEvent/query | query auditing events



## QueryAuditingEvents

> GithubComSugerioMarketplaceServicePkgCrudListBaseResponseAuditingEvent QueryAuditingEvents(ctx, orgId).PageSize(pageSize).PageNumber(pageNumber).Q(q).S(s).Execute()

query auditing events



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
	pageSize := int32(56) // int32 | Number of items per page (default: 20, max: 1000) (optional)
	pageNumber := int32(56) // int32 | Page number (default: 1) (optional)
	q := "q_example" // string | LISP-style filter expression (e.g., '(= event_type \\ (optional)
	s := "s_example" // string | Sort fields: 'field:asc,field2:desc' or '-field,field2' format (e.g., 'creation_time:desc,event_type:asc' or '-creation_time,event_type') (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditingAPI.QueryAuditingEvents(context.Background(), orgId).PageSize(pageSize).PageNumber(pageNumber).Q(q).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditingAPI.QueryAuditingEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryAuditingEvents`: GithubComSugerioMarketplaceServicePkgCrudListBaseResponseAuditingEvent
	fmt.Fprintf(os.Stdout, "Response from `AuditingAPI.QueryAuditingEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryAuditingEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | Number of items per page (default: 20, max: 1000) | 
 **pageNumber** | **int32** | Page number (default: 1) | 
 **q** | **string** | LISP-style filter expression (e.g., &#39;(&#x3D; event_type \\ | 
 **s** | **string** | Sort fields: &#39;field:asc,field2:desc&#39; or &#39;-field,field2&#39; format (e.g., &#39;creation_time:desc,event_type:asc&#39; or &#39;-creation_time,event_type&#39;) | 

### Return type

[**GithubComSugerioMarketplaceServicePkgCrudListBaseResponseAuditingEvent**](GithubComSugerioMarketplaceServicePkgCrudListBaseResponseAuditingEvent.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

