# \SupportAPI

All URIs are relative to *http://https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseSupportTicket**](SupportAPI.md#CloseSupportTicket) | **Patch** /org/{orgId}/support/ticket/{ticketId}/close | close support ticket
[**CreateSupportTicket**](SupportAPI.md#CreateSupportTicket) | **Post** /org/{orgId}/support/ticket | create support ticket
[**CreateSupportTicketAttachment**](SupportAPI.md#CreateSupportTicketAttachment) | **Post** /org/{orgId}/support/ticket/{ticketId}/attachment | create support ticket attachment
[**CreateSupportTicketComment**](SupportAPI.md#CreateSupportTicketComment) | **Post** /org/{orgId}/support/ticket/{ticketId}/comment | create support ticket comment
[**GetSupportTicket**](SupportAPI.md#GetSupportTicket) | **Get** /org/{orgId}/support/ticket/{ticketId} | get support ticket
[**ListSupportTickets**](SupportAPI.md#ListSupportTickets) | **Get** /org/{orgId}/support/ticket | list support tickets
[**ReopenSupportTicket**](SupportAPI.md#ReopenSupportTicket) | **Patch** /org/{orgId}/support/ticket/{ticketId}/reopen | reopen support ticket
[**UpdateSupportTicket**](SupportAPI.md#UpdateSupportTicket) | **Patch** /org/{orgId}/support/ticket/{ticketId} | update support ticket



## CloseSupportTicket

> SupportTicket CloseSupportTicket(ctx, orgId, ticketId).Execute()

close support ticket



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
	ticketId := "ticketId_example" // string | Ticket ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.CloseSupportTicket(context.Background(), orgId, ticketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.CloseSupportTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloseSupportTicket`: SupportTicket
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.CloseSupportTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**ticketId** | **string** | Ticket ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseSupportTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SupportTicket**](SupportTicket.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSupportTicket

> SupportTicket CreateSupportTicket(ctx, orgId).Body(body).Execute()

create support ticket



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
	body := *openapiclient.NewSupportTicket() // SupportTicket | Ticket create request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.CreateSupportTicket(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.CreateSupportTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSupportTicket`: SupportTicket
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.CreateSupportTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupportTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SupportTicket**](SupportTicket.md) | Ticket create request | 

### Return type

[**SupportTicket**](SupportTicket.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSupportTicketAttachment

> []SupportTicketAttachment CreateSupportTicketAttachment(ctx, orgId, ticketId).File(file).Execute()

create support ticket attachment



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
	ticketId := "ticketId_example" // string | Ticket ID
	file := os.NewFile(1234, "some_file") // *os.File | File to upload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.CreateSupportTicketAttachment(context.Background(), orgId, ticketId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.CreateSupportTicketAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSupportTicketAttachment`: []SupportTicketAttachment
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.CreateSupportTicketAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**ticketId** | **string** | Ticket ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupportTicketAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | File to upload | 

### Return type

[**[]SupportTicketAttachment**](SupportTicketAttachment.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSupportTicketComment

> SupportTicketComment CreateSupportTicketComment(ctx, orgId, ticketId).Body(body).Execute()

create support ticket comment



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
	ticketId := "ticketId_example" // string | Ticket ID
	body := *openapiclient.NewSupportTicketComment() // SupportTicketComment | Ticket create request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.CreateSupportTicketComment(context.Background(), orgId, ticketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.CreateSupportTicketComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSupportTicketComment`: SupportTicketComment
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.CreateSupportTicketComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**ticketId** | **string** | Ticket ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupportTicketCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SupportTicketComment**](SupportTicketComment.md) | Ticket create request | 

### Return type

[**SupportTicketComment**](SupportTicketComment.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportTicket

> SupportTicket GetSupportTicket(ctx, orgId, ticketId).Execute()

get support ticket



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
	ticketId := "ticketId_example" // string | Ticket ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.GetSupportTicket(context.Background(), orgId, ticketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.GetSupportTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportTicket`: SupportTicket
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.GetSupportTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**ticketId** | **string** | Ticket ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SupportTicket**](SupportTicket.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportTickets

> ListSupportTicketsResponse ListSupportTickets(ctx, orgId).Limit(limit).Offset(offset).Execute()

list support tickets



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
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.ListSupportTickets(context.Background(), orgId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.ListSupportTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSupportTickets`: ListSupportTicketsResponse
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.ListSupportTickets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**ListSupportTicketsResponse**](ListSupportTicketsResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReopenSupportTicket

> SupportTicket ReopenSupportTicket(ctx, orgId, ticketId).Execute()

reopen support ticket



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
	ticketId := "ticketId_example" // string | Ticket ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.ReopenSupportTicket(context.Background(), orgId, ticketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.ReopenSupportTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReopenSupportTicket`: SupportTicket
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.ReopenSupportTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**ticketId** | **string** | Ticket ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReopenSupportTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SupportTicket**](SupportTicket.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSupportTicket

> SupportTicket UpdateSupportTicket(ctx, orgId, ticketId).Body(body).Execute()

update support ticket



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
	ticketId := "ticketId_example" // string | Ticket ID
	body := *openapiclient.NewUpdateSupportTicketRequest(openapiclient.SupportTicketPriority("LOW"), []string{"Watchers_example"}) // UpdateSupportTicketRequest | Ticket create request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.UpdateSupportTicket(context.Background(), orgId, ticketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.UpdateSupportTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSupportTicket`: SupportTicket
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.UpdateSupportTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**ticketId** | **string** | Ticket ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSupportTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateSupportTicketRequest**](UpdateSupportTicketRequest.md) | Ticket create request | 

### Return type

[**SupportTicket**](SupportTicket.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

