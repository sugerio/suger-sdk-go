# \BillingAPI

All URIs are relative to *http://https://api.suger.cloud*

 Method                                                                             | HTTP request                                                                           | Description               
------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------|---------------------------
 [**CreateAddon**](BillingAPI.md#CreateAddon)                                       | **Post** /org/{orgId}/addon                                                            | create addon              
 [**CreateRefund**](BillingAPI.md#CreateRefund)                                     | **Post** /org/{orgId}/buyer/{buyerId}/paymentTransaction/{paymentTransactionId}/refund | create refund.            
 [**DeleteAddon**](BillingAPI.md#DeleteAddon)                                       | **Delete** /org/{orgId}/addon/{addonId}                                                | delete addon              
 [**GetAddon**](BillingAPI.md#GetAddon)                                             | **Get** /org/{orgId}/addon/{addonId}                                                   | get addon                 
 [**GetInvoiceV2**](BillingAPI.md#GetInvoiceV2)                                     | **Get** /org/{orgId}/invoice/{invoiceId}                                               | get invoice               
 [**IssueInvoiceV2**](BillingAPI.md#IssueInvoiceV2)                                 | **Patch** /org/{orgId}/invoice/{invoiceId}/issue                                       | issue invoice             
 [**ListAddons**](BillingAPI.md#ListAddons)                                         | **Get** /org/{orgId}/addon                                                             | list addons               
 [**ListInvoices**](BillingAPI.md#ListInvoices)                                     | **Get** /org/{orgId}/invoice                                                           | list invoices             
 [**ListPaymentTransactions**](BillingAPI.md#ListPaymentTransactions)               | **Get** /org/{orgId}/paymentTransaction                                                | list payment transactions 
 [**ListRefundOfPaymentTransaction**](BillingAPI.md#ListRefundOfPaymentTransaction) | **Get** /org/{orgId}/buyer/{buyerId}/paymentTransaction/{paymentTransactionId}/refund  | list refunds.             
 [**PayInvoiceV2**](BillingAPI.md#PayInvoiceV2)                                     | **Patch** /org/{orgId}/invoice/{invoiceId}/pay                                         | pay invoice               
 [**PreviewInvoiceEmail**](BillingAPI.md#PreviewInvoiceEmail)                       | **Get** /org/{orgId}/invoice/{invoiceId}/preview                                       | preview invoice email     
 [**UpdateAddon**](BillingAPI.md#UpdateAddon)                                       | **Patch** /org/{orgId}/addon/{addonId}                                                 | update addon              
 [**UpdateInvoiceInfoV2**](BillingAPI.md#UpdateInvoiceInfoV2)                       | **Patch** /org/{orgId}/invoice/{invoiceId}/info                                        | update invoice info       
 [**VoidInvoiceV2**](BillingAPI.md#VoidInvoiceV2)                                   | **Patch** /org/{orgId}/invoice/{invoiceId}/void                                        | void invoice              



## CreateAddon

> BillingAddon CreateAddon(ctx, orgId).Data(data).Execute()

create addon



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
	data := *openapiclient.NewCreateAndUpdateAddonParams() // CreateAndUpdateAddonParams | CreateAndUpdateAddonParams

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.CreateAddon(context.Background(), orgId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.CreateAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAddon`: BillingAddon
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.CreateAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**CreateAndUpdateAddonParams**](CreateAndUpdateAddonParams.md) | CreateAndUpdateAddonParams | 

### Return type

[**BillingAddon**](BillingAddon.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRefund

> BillingPaymentTransaction CreateRefund(ctx, orgId, buyerId, paymentTransactionId).Amount(amount).Execute()

create refund.



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
	buyerId := "buyerId_example" // string | Buyer ID
	paymentTransactionId := "paymentTransactionId_example" // string | Payment transaction ID
	amount := float32(8.14) // float32 | Refund amount

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.CreateRefund(context.Background(), orgId, buyerId, paymentTransactionId).Amount(amount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.CreateRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRefund`: BillingPaymentTransaction
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.CreateRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**buyerId** | **string** | Buyer ID | 
**paymentTransactionId** | **string** | Payment transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amount** | **float32** | Refund amount | 

### Return type

[**BillingPaymentTransaction**](BillingPaymentTransaction.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAddon

> string DeleteAddon(ctx, orgId, addonId).Execute()

delete addon



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
	addonId := "addonId_example" // string | Addon ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.DeleteAddon(context.Background(), orgId, addonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.DeleteAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAddon`: string
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.DeleteAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**addonId** | **string** | Addon ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddon

> BillingAddon GetAddon(ctx, orgId, addonId).Execute()

get addon



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
	addonId := "addonId_example" // string | Addon ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetAddon(context.Background(), orgId, addonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddon`: BillingAddon
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**addonId** | **string** | Addon ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BillingAddon**](BillingAddon.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceV2

> BillingInvoice GetInvoiceV2(ctx, orgId, invoiceId).Execute()

get invoice

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
	invoiceId := "invoiceId_example" // string | Invoice ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetInvoiceV2(context.Background(), orgId, invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetInvoiceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceV2`: BillingInvoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetInvoiceV2`: %v\n", resp)
}
```

### Path Parameters


 Name          | Type                | Description                                                                 | Notes 
---------------|---------------------|-----------------------------------------------------------------------------|-------
 **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. 
 **orgId**     | **string**          | Organization ID                                                             |
 **invoiceId** | **string**          | Invoice ID                                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceV2Request struct via the builder pattern


 Name | Type | Description | Notes 
------|------|-------------|-------

### Return type

[**BillingInvoice**](BillingInvoice.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueInvoiceV2

> BillingInvoice IssueInvoiceV2(ctx, orgId, invoiceId).ContactIds(contactIds).Execute()

issue invoice

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
	invoiceId := "invoiceId_example" // string | Invoice ID
	contactIds := []string{"Property_example"} // []string | List of Contact IDs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.IssueInvoiceV2(context.Background(), orgId, invoiceId).ContactIds(contactIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.IssueInvoiceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueInvoiceV2`: BillingInvoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.IssueInvoiceV2`: %v\n", resp)
}
```

### Path Parameters


 Name          | Type                | Description                                                                 | Notes 
---------------|---------------------|-----------------------------------------------------------------------------|-------
 **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. 
 **orgId**     | **string**          | Organization ID                                                             |
 **invoiceId** | **string**          | Invoice ID                                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiIssueInvoiceV2Request struct via the builder pattern


 Name | Type | Description | Notes 
------|------|-------------|-------

**contactIds** | **[]string** | List of Contact IDs |

### Return type

[**BillingInvoice**](BillingInvoice.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAddons

> []BillingAddon ListAddons(ctx, orgId).Limit(limit).Offset(offset).Execute()

list addons



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
	resp, r, err := apiClient.BillingAPI.ListAddons(context.Background(), orgId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListAddons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAddons`: []BillingAddon
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListAddons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAddonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**[]BillingAddon**](BillingAddon.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvoices

> []BillingInvoice ListInvoices(ctx, orgId).EntitlementId(entitlementId).BuyerId(buyerId).Status(status).Limit(limit).Offset(offset).Execute()

list invoices



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
	entitlementId := "entitlementId_example" // string | Optional, filter by the entitlement ID (optional)
	buyerId := "buyerId_example" // string | Optional, filter by the given buyer ID (optional)
	status := "status_example" // string | Optional, filter by invoice status as filter, if not provided, all status invoices are returned (optional)
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.ListInvoices(context.Background(), orgId).EntitlementId(entitlementId).BuyerId(buyerId).Status(status).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInvoices`: []BillingInvoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entitlementId** | **string** | Optional, filter by the entitlement ID | 
 **buyerId** | **string** | Optional, filter by the given buyer ID | 
 **status** | **string** | Optional, filter by invoice status as filter, if not provided, all status invoices are returned | 
 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**[]BillingInvoice**](BillingInvoice.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentTransactions

> []BillingPaymentTransaction ListPaymentTransactions(ctx, orgId).BuyerId(buyerId).EntitlementId(entitlementId).InvoiceId(invoiceId).Status(status).Type_(type_).Limit(limit).Offset(offset).Execute()

list payment transactions



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
	buyerId := "buyerId_example" // string | Optional, filter by the given buyer ID (optional)
	entitlementId := "entitlementId_example" // string | Optional, filter by the given entitlement ID (optional)
	invoiceId := "invoiceId_example" // string | Optional, filter by the given invoice ID (optional)
	status := "status_example" // string | Optional, filter by status (optional)
	type_ := "type__example" // string | Optional, filter by transaction type (optional)
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.ListPaymentTransactions(context.Background(), orgId).BuyerId(buyerId).EntitlementId(entitlementId).InvoiceId(invoiceId).Status(status).Type_(type_).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListPaymentTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPaymentTransactions`: []BillingPaymentTransaction
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListPaymentTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buyerId** | **string** | Optional, filter by the given buyer ID | 
 **entitlementId** | **string** | Optional, filter by the given entitlement ID | 
 **invoiceId** | **string** | Optional, filter by the given invoice ID | 
 **status** | **string** | Optional, filter by status | 
 **type_** | **string** | Optional, filter by transaction type | 
 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**[]BillingPaymentTransaction**](BillingPaymentTransaction.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRefundOfPaymentTransaction

> []BillingPaymentTransaction ListRefundOfPaymentTransaction(ctx, orgId, buyerId, paymentTransactionId).Execute()

list refunds.



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
	buyerId := "buyerId_example" // string | Buyer ID
	paymentTransactionId := "paymentTransactionId_example" // string | Payment transaction ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.ListRefundOfPaymentTransaction(context.Background(), orgId, buyerId, paymentTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListRefundOfPaymentTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRefundOfPaymentTransaction`: []BillingPaymentTransaction
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListRefundOfPaymentTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**buyerId** | **string** | Buyer ID | 
**paymentTransactionId** | **string** | Payment transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRefundOfPaymentTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]BillingPaymentTransaction**](BillingPaymentTransaction.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PayInvoiceV2

> BillingInvoice PayInvoiceV2(ctx, orgId, invoiceId).Execute()

pay invoice



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
	invoiceId := "invoiceId_example" // string | Invoice ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.PayInvoiceV2(context.Background(), orgId, invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PayInvoiceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PayInvoiceV2`: BillingInvoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.PayInvoiceV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**invoiceId** | **string** | Invoice ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPayInvoiceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BillingInvoice**](BillingInvoice.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PreviewInvoiceEmail

> string PreviewInvoiceEmail(ctx, orgId, invoiceId).Execute()

preview invoice email

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
	invoiceId := "invoiceId_example" // string | Invoice ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.PreviewInvoiceEmail(context.Background(), orgId, invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PreviewInvoiceEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewInvoiceEmail`: string
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.PreviewInvoiceEmail`: %v\n", resp)
}
```

### Path Parameters


 Name          | Type                | Description                                                                 | Notes 
---------------|---------------------|-----------------------------------------------------------------------------|-------
 **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. 
 **orgId**     | **string**          | Organization ID                                                             |
 **invoiceId** | **string**          | Invoice ID                                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewInvoiceEmailRequest struct via the builder pattern

 Name | Type | Description | Notes 
------|------|-------------|-------

### Return type

**string**

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddon

> BillingAddon UpdateAddon(ctx, orgId, addonId).Data(data).Execute()

update addon



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
	addonId := "addonId_example" // string | Addon ID
	data := *openapiclient.NewCreateAndUpdateAddonParams() // CreateAndUpdateAddonParams | CreateAndUpdateAddonParams

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.UpdateAddon(context.Background(), orgId, addonId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpdateAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAddon`: BillingAddon
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpdateAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**addonId** | **string** | Addon ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**CreateAndUpdateAddonParams**](CreateAndUpdateAddonParams.md) | CreateAndUpdateAddonParams | 

### Return type

[**BillingAddon**](BillingAddon.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoiceInfoV2

> BillingInvoiceInfo UpdateInvoiceInfoV2(ctx, orgId, invoiceId).Data(data).Execute()

update invoice info

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
	invoiceId := "invoiceId_example" // string | Invoice ID
	data := *openapiclient.NewUpdateInvoiceInfoRequest() // UpdateInvoiceInfoRequest | Update Invoice Info Request Params

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.UpdateInvoiceInfoV2(context.Background(), orgId, invoiceId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpdateInvoiceInfoV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoiceInfoV2`: BillingInvoiceInfo
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpdateInvoiceInfoV2`: %v\n", resp)
}
```

### Path Parameters

 Name          | Type                | Description                                                                 | Notes 
---------------|---------------------|-----------------------------------------------------------------------------|-------
 **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. 
 **orgId**     | **string**          | Organization ID                                                             |
 **invoiceId** | **string**          | Invoice ID                                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceInfoV2Request struct via the builder pattern

 Name | Type | Description | Notes 
------|------|-------------|-------

**data** | [**UpdateInvoiceInfoRequest**](UpdateInvoiceInfoRequest.md) | Update Invoice Info Request Params |

### Return type

[**BillingInvoiceInfo**](BillingInvoiceInfo.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidInvoiceV2

> BillingInvoice VoidInvoiceV2(ctx, orgId, invoiceId).Execute()

void invoice

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
	invoiceId := "invoiceId_example" // string | Invoice ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.VoidInvoiceV2(context.Background(), orgId, invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.VoidInvoiceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidInvoiceV2`: BillingInvoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.VoidInvoiceV2`: %v\n", resp)
}
```

### Path Parameters

 Name          | Type                | Description                                                                 | Notes 
---------------|---------------------|-----------------------------------------------------------------------------|-------
 **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. 
 **orgId**     | **string**          | Organization ID                                                             |
 **invoiceId** | **string**          | Invoice ID                                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiVoidInvoiceV2Request struct via the builder pattern

 Name | Type | Description | Notes 
------|------|-------------|-------

### Return type

[**BillingInvoice**](BillingInvoice.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

