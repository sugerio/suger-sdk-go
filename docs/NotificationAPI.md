# \NotificationAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotificationMessage**](NotificationAPI.md#GetNotificationMessage) | **Get** /org/{orgId}/notificationMessage/{notificationMessageId} | Get the notification message
[**ListNotificationMessages**](NotificationAPI.md#ListNotificationMessages) | **Get** /org/{orgId}/notificationMessage | List the notification messages



## GetNotificationMessage

> NotificationMessage GetNotificationMessage(ctx, orgId, notificationMessageId).Execute()

Get the notification message



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

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationMessages

> ListNotificationMessagesResponse ListNotificationMessages(ctx, orgId).Limit(limit).Offset(offset).Execute()

List the notification messages



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
    limit := int32(56) // int32 | List pagination size, default 20, max value is 1000 (optional)
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

 **limit** | **int32** | List pagination size, default 20, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**ListNotificationMessagesResponse**](ListNotificationMessagesResponse.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

