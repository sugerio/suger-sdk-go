# \NotificationAPI

All URIs are relative to *http://https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotificationMessage**](NotificationAPI.md#GetNotificationMessage) | **Get** /org/{orgId}/notificationMessage/{notificationMessageId} | get notification message
[**ListNotificationEvents**](NotificationAPI.md#ListNotificationEvents) | **Get** /org/{orgId}/notificationEvent | list notification events
[**ListNotificationEventsByEntity**](NotificationAPI.md#ListNotificationEventsByEntity) | **Get** /org/{orgId}/notificationEvent/{entityType}/{entityId} | list notification events by entity
[**ListNotificationMessages**](NotificationAPI.md#ListNotificationMessages) | **Get** /org/{orgId}/notificationMessage | list notification messages



## GetNotificationMessage

> NotificationMessage GetNotificationMessage(ctx, orgId, notificationMessageId).Execute()

get notification message



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
	notificationMessageId := "notificationMessageId_example" // string | Notification Message ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationAPI.GetNotificationMessage(context.Background(), orgId, notificationMessageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.GetNotificationMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationMessage`: NotificationMessage
	fmt.Fprintf(os.Stdout, "Response from `NotificationAPI.GetNotificationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**notificationMessageId** | **string** | Notification Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NotificationMessage**](NotificationMessage.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationEvents

> ListNotificationEventsResponse ListNotificationEvents(ctx, orgId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Priorities(priorities).Execute()

list notification events



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
	startDate := "startDate_example" // string | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate (optional)
	endDate := "endDate_example" // string | end date (UTC) in YYYY-MM-DD format, default is today (optional)
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)
	priorities := "priorities_example" // string | Filter by priorities, empty means all priorities. Valid values are: LOW, MEDIUM, HIGH, CRITICAL. Multiple values are supported, separated by comma. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationAPI.ListNotificationEvents(context.Background(), orgId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Priorities(priorities).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.ListNotificationEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationEvents`: ListNotificationEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationAPI.ListNotificationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate | 
 **endDate** | **string** | end date (UTC) in YYYY-MM-DD format, default is today | 
 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 
 **priorities** | **string** | Filter by priorities, empty means all priorities. Valid values are: LOW, MEDIUM, HIGH, CRITICAL. Multiple values are supported, separated by comma. | 

### Return type

[**ListNotificationEventsResponse**](ListNotificationEventsResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationEventsByEntity

> ListNotificationEventsResponse ListNotificationEventsByEntity(ctx, orgId, entityType, entityId).Limit(limit).Offset(offset).Execute()

list notification events by entity



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
	entityType := "entityType_example" // string | Entity type, valid values are: PRODUCT, OFFER, ENTITLEMENT, INTEGRATION etc.
	entityId := "entityId_example" // string | Entity ID
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationAPI.ListNotificationEventsByEntity(context.Background(), orgId, entityType, entityId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.ListNotificationEventsByEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationEventsByEntity`: ListNotificationEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationAPI.ListNotificationEventsByEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entityType** | **string** | Entity type, valid values are: PRODUCT, OFFER, ENTITLEMENT, INTEGRATION etc. | 
**entityId** | **string** | Entity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationEventsByEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**ListNotificationEventsResponse**](ListNotificationEventsResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationMessages

> ListNotificationMessagesResponse ListNotificationMessages(ctx, orgId).Limit(limit).Offset(offset).Execute()

list notification messages



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
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationAPI.ListNotificationMessages(context.Background(), orgId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.ListNotificationMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationMessages`: ListNotificationMessagesResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationAPI.ListNotificationMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**ListNotificationMessagesResponse**](ListNotificationMessagesResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

