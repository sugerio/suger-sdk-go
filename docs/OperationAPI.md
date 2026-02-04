# \OperationAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOperationV2**](OperationAPI.md#GetOperationV2) | **Get** /org/{orgId}/v2/operation/{operationId} | get operation details
[**ListOperationHistoryV2**](OperationAPI.md#ListOperationHistoryV2) | **Get** /org/{orgId}/v2/operation/{operationId}/history | get operation history events
[**ListOperationsV2**](OperationAPI.md#ListOperationsV2) | **Post** /org/{orgId}/v2/operation/list | list operations with filters



## GetOperationV2

> Operation GetOperationV2(ctx, orgId, operationId).RunId(runId).HasChildren(hasChildren).Execute()

get operation details



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
	operationId := "operationId_example" // string | Temporal Workflow ID
	runId := "runId_example" // string | Temporal run ID of the workflow (optional)
	hasChildren := true // bool | Include data aggregated from all child workflows (1 level), default false (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.GetOperationV2(context.Background(), orgId, operationId).RunId(runId).HasChildren(hasChildren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.GetOperationV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperationV2`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.GetOperationV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**operationId** | **string** | Temporal Workflow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperationV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **runId** | **string** | Temporal run ID of the workflow | 
 **hasChildren** | **bool** | Include data aggregated from all child workflows (1 level), default false | 

### Return type

[**Operation**](Operation.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationHistoryV2

> OperationEventsResponse ListOperationHistoryV2(ctx, orgId, operationId).Execute()

get operation history events



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
	operationId := "operationId_example" // string | Temporal Workflow ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.ListOperationHistoryV2(context.Background(), orgId, operationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.ListOperationHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationHistoryV2`: OperationEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.ListOperationHistoryV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**operationId** | **string** | Temporal Workflow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationHistoryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationEventsResponse**](OperationEventsResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsV2

> ListOperationsResponse ListOperationsV2(ctx, orgId).Request(request).Execute()

list operations with filters



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
	request := *openapiclient.NewListOperationsV2Request() // ListOperationsV2Request | Supports logical operators (and/or) and comparison operators (=, !=, in, not_in).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.ListOperationsV2(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.ListOperationsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsV2`: ListOperationsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.ListOperationsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**ListOperationsV2Request**](ListOperationsV2Request.md) | Supports logical operators (and/or) and comparison operators (&#x3D;, !&#x3D;, in, not_in). | 

### Return type

[**ListOperationsResponse**](ListOperationsResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

