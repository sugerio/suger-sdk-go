# \AIUsageAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAIUsage**](AIUsageAPI.md#GetAIUsage) | **Get** /org/{orgId}/ai-usage | Get AI usage data



## GetAIUsage

> ServiceMarketplaceServiceApiGetAIUsageResponse GetAIUsage(ctx, orgId).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).Execute()

Get AI usage data



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
	startTime := "startTime_example" // string | Start time (RFC3339 format)
	endTime := "endTime_example" // string | End time (RFC3339 format)
	limit := int32(56) // int32 | Maximum number of records to return (default: 1000, max: 10000) (optional)
	offset := int32(56) // int32 | Number of records to skip for pagination (default: 0) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIUsageAPI.GetAIUsage(context.Background(), orgId).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIUsageAPI.GetAIUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIUsage`: ServiceMarketplaceServiceApiGetAIUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `AIUsageAPI.GetAIUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **string** | Start time (RFC3339 format) | 
 **endTime** | **string** | End time (RFC3339 format) | 
 **limit** | **int32** | Maximum number of records to return (default: 1000, max: 10000) | 
 **offset** | **int32** | Number of records to skip for pagination (default: 0) | 

### Return type

[**ServiceMarketplaceServiceApiGetAIUsageResponse**](ServiceMarketplaceServiceApiGetAIUsageResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

