# \SearchAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSearchableObjects**](SearchAPI.md#ListSearchableObjects) | **Get** /org/{orgId}/search | Global search



## ListSearchableObjects

> SearchResponse ListSearchableObjects(ctx, orgId).Q(q).Types(types).Limit(limit).Cursor(cursor).Execute()

Global search



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
	q := "q_example" // string | Search query
	types := "types_example" // string | Comma-separated object types (e.g. product,offer) (optional)
	limit := int32(56) // int32 | Page size (default 20) (optional)
	cursor := "cursor_example" // string | Cursor (offset:query base64 encoded) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.ListSearchableObjects(context.Background(), orgId).Q(q).Types(types).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.ListSearchableObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSearchableObjects`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.ListSearchableObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSearchableObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Search query | 
 **types** | **string** | Comma-separated object types (e.g. product,offer) | 
 **limit** | **int32** | Page size (default 20) | 
 **cursor** | **string** | Cursor (offset:query base64 encoded) | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

