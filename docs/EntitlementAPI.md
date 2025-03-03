# \EntitlementAPI

All URIs are relative to *http://https://api.suger.cloud*

 Method                                                                                       | HTTP request                                                                            | Description                         
----------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------|-------------------------------------
 [**AddEntitlementCredit**](EntitlementAPI.md#AddEntitlementCredit)                           | **Post** /org/{orgId}/entitlement/{entitlementId}/addCredit                             | add entitlement credit              
 [**ApplyAddonToEntitlement**](EntitlementAPI.md#ApplyAddonToEntitlement)                     | **Post** /org/{orgId}/entitlement/{entitlementId}/addon                                 | apply addon to entitlement          
 [**ApproveEntitlement**](EntitlementAPI.md#ApproveEntitlement)                               | **Post** /org/{orgId}/entitlement/{entitlementId}/approve                               | approve entitlement                 
 [**CancelEntitlement**](EntitlementAPI.md#CancelEntitlement)                                 | **Post** /org/{orgId}/entitlement/{entitlementId}/cancel                                | cancel entitlement                  
 [**CreateEntitlement**](EntitlementAPI.md#CreateEntitlement)                                 | **Post** /org/{orgId}/entitlement                                                       | create entitlement                  
 [**DeleteEntitlementTerm**](EntitlementAPI.md#DeleteEntitlementTerm)                         | **Delete** /org/{orgId}/entitlement/{entitlementId}/entitlementTerm/{entitlementTermId} | delete entitlement term             
 [**DivideEntitlementCommit**](EntitlementAPI.md#DivideEntitlementCommit)                     | **Post** /org/{orgId}/entitlement/{entitlementId}/divideCommit                          | divide entitlement commit           
 [**GetEntitlement**](EntitlementAPI.md#GetEntitlement)                                       | **Get** /org/{orgId}/entitlement/{entitlementId}                                        | get entitlement                     
 [**GetEntitlementTerm**](EntitlementAPI.md#GetEntitlementTerm)                               | **Get** /org/{orgId}/entitlement/{entitlementId}/entitlementTerm/{entitlementTermId}    | get entitlement term                
 [**ListEntitlementTerms**](EntitlementAPI.md#ListEntitlementTerms)                           | **Get** /org/{orgId}/entitlement/{entitlementId}/entitlementTerm                        | list entitlement terms              
 [**ListEntitlements**](EntitlementAPI.md#ListEntitlements)                                   | **Get** /org/{orgId}/entitlement                                                        | list entitlements                   
 [**ScheduleEntitlementCancellation**](EntitlementAPI.md#ScheduleEntitlementCancellation)     | **Post** /org/{orgId}/entitlement/{entitlementId}/scheduleCancellation                  | schedule entitlement cancellation   
 [**UnscheduleEntitlementCancellation**](EntitlementAPI.md#UnscheduleEntitlementCancellation) | **Post** /org/{orgId}/entitlement/{entitlementId}/unscheduleCancellation                | unschedule entitlement cancellation 
 [**UpdateEntitlementMetaInfo**](EntitlementAPI.md#UpdateEntitlementMetaInfo)                 | **Patch** /org/{orgId}/entitlement/{entitlementId}/metaInfo                             | update entitlement meta info        
 [**UpdateEntitlementName**](EntitlementAPI.md#UpdateEntitlementName)                         | **Patch** /org/{orgId}/entitlement/{entitlementId}/entitlementName                      | update entitlement name             
 [**UpdateEntitlementSeat**](EntitlementAPI.md#UpdateEntitlementSeat)                         | **Patch** /org/{orgId}/entitlement/{entitlementId}/seat                                 | update entitlement seat             



## AddEntitlementCredit

> AddEntitlementCreditResponse AddEntitlementCredit(ctx, orgId, entitlementId).Data(data).Execute()

add entitlement credit



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
	entitlementId := "entitlementId_example" // string | Entitlement ID
	data := *openapiclient.NewAddEntitlementCreditParams(float32(123), "EntitlementID_example", "OrganizationID_example") // AddEntitlementCreditParams | RequestBody

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.AddEntitlementCredit(context.Background(), orgId, entitlementId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.AddEntitlementCredit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddEntitlementCredit`: AddEntitlementCreditResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.AddEntitlementCredit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddEntitlementCreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**AddEntitlementCreditParams**](AddEntitlementCreditParams.md) | RequestBody | 

### Return type

[**AddEntitlementCreditResponse**](AddEntitlementCreditResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyAddonToEntitlement

> WorkloadEntitlement ApplyAddonToEntitlement(ctx, orgId, entitlementId).Data(data).Execute()

apply addon to entitlement



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
	entitlementId := "entitlementId_example" // string | Entitlement ID
	data := *openapiclient.NewBillingAddonRecord() // BillingAddonRecord | RequestBody

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.ApplyAddonToEntitlement(context.Background(), orgId, entitlementId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.ApplyAddonToEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyAddonToEntitlement`: WorkloadEntitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.ApplyAddonToEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyAddonToEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**BillingAddonRecord**](BillingAddonRecord.md) | RequestBody | 

### Return type

[**WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApproveEntitlement

> WorkloadEntitlement ApproveEntitlement(ctx, orgId, entitlementId).Execute()

approve entitlement



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
	entitlementId := "entitlementId_example" // string | Entitlement ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.ApproveEntitlement(context.Background(), orgId, entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.ApproveEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveEntitlement`: WorkloadEntitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.ApproveEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelEntitlement

> WorkloadEntitlement CancelEntitlement(ctx, orgId, entitlementId).Execute()

cancel entitlement



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
	entitlementId := "entitlementId_example" // string | Entitlement ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.CancelEntitlement(context.Background(), orgId, entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.CancelEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelEntitlement`: WorkloadEntitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.CancelEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEntitlement

> WorkloadEntitlement CreateEntitlement(ctx, orgId).Data(data).Execute()

create entitlement



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
	data := *openapiclient.NewCreateEntitlementParams() // CreateEntitlementParams | RequestBody

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.CreateEntitlement(context.Background(), orgId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.CreateEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEntitlement`: WorkloadEntitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.CreateEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**CreateEntitlementParams**](CreateEntitlementParams.md) | RequestBody | 

### Return type

[**WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntitlementTerm

> string DeleteEntitlementTerm(ctx, orgId, entitlementId, entitlementTermId).Execute()

delete entitlement term



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
	entitlementId := "entitlementId_example" // string | Entitlement ID
	entitlementTermId := "entitlementTermId_example" // string | Entitlement Term ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.DeleteEntitlementTerm(context.Background(), orgId, entitlementId, entitlementTermId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.DeleteEntitlementTerm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEntitlementTerm`: string
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.DeleteEntitlementTerm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 
**entitlementTermId** | **string** | Entitlement Term ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntitlementTermRequest struct via the builder pattern


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


## DivideEntitlementCommit

> string DivideEntitlementCommit(ctx, orgId, entitlementId).Data(data).Execute()

divide entitlement commit



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
	entitlementId := "entitlementId_example" // string | Entitlement ID
	data := *openapiclient.NewDivideEntitlementCommitParams() // DivideEntitlementCommitParams | RequestBody

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.DivideEntitlementCommit(context.Background(), orgId, entitlementId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.DivideEntitlementCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DivideEntitlementCommit`: string
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.DivideEntitlementCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDivideEntitlementCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**DivideEntitlementCommitParams**](DivideEntitlementCommitParams.md) | RequestBody | 

### Return type

**string**

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlement

> WorkloadEntitlement GetEntitlement(ctx, orgId, entitlementId).Execute()

get entitlement



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
	entitlementId := "entitlementId_example" // string | Entitlement ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.GetEntitlement(context.Background(), orgId, entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.GetEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlement`: WorkloadEntitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.GetEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementTerm

> WorkloadEntitlementTerm GetEntitlementTerm(ctx, orgId, entitlementId, entitlementTermId).Execute()

get entitlement term



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
	entitlementId := "entitlementId_example" // string | Entitlement ID
	entitlementTermId := "entitlementTermId_example" // string | Entitlement Term ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.GetEntitlementTerm(context.Background(), orgId, entitlementId, entitlementTermId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.GetEntitlementTerm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlementTerm`: WorkloadEntitlementTerm
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.GetEntitlementTerm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 
**entitlementTermId** | **string** | Entitlement Term ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementTermRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkloadEntitlementTerm**](WorkloadEntitlementTerm.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlementTerms

> []WorkloadEntitlementTerm ListEntitlementTerms(ctx, orgId, entitlementId).Execute()

list entitlement terms



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
	entitlementId := "entitlementId_example" // string | Entitlement ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.ListEntitlementTerms(context.Background(), orgId, entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.ListEntitlementTerms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntitlementTerms`: []WorkloadEntitlementTerm
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.ListEntitlementTerms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementTermsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WorkloadEntitlementTerm**](WorkloadEntitlementTerm.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlements

> []WorkloadEntitlement ListEntitlements(ctx, orgId).Partner(partner).ProductId(productId).OfferId(offerId).BuyerId(buyerId).Limit(limit).Offset(offset).Execute()

list entitlements



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
	partner := "partner_example" // string | filter by partner (optional)
	productId := "productId_example" // string | filter by productId (optional)
	offerId := "offerId_example" // string | filter by offerId (optional)
	buyerId := "buyerId_example" // string | filter by buyerId (optional)
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.ListEntitlements(context.Background(), orgId).Partner(partner).ProductId(productId).OfferId(offerId).BuyerId(buyerId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.ListEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntitlements`: []WorkloadEntitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.ListEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partner** | **string** | filter by partner | 
 **productId** | **string** | filter by productId | 
 **offerId** | **string** | filter by offerId | 
 **buyerId** | **string** | filter by buyerId | 
 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**[]WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleEntitlementCancellation

> WorkloadEntitlement ScheduleEntitlementCancellation(ctx, orgId, entitlementId).Data(data).Execute()

schedule entitlement cancellation



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
	entitlementId := "entitlementId_example" // string | Entitlement ID
	data := *openapiclient.NewCancellationSchedule() // CancellationSchedule | RequestBody

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.ScheduleEntitlementCancellation(context.Background(), orgId, entitlementId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.ScheduleEntitlementCancellation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleEntitlementCancellation`: WorkloadEntitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.ScheduleEntitlementCancellation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleEntitlementCancellationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**CancellationSchedule**](CancellationSchedule.md) | RequestBody | 

### Return type

[**WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnscheduleEntitlementCancellation

> WorkloadEntitlement UnscheduleEntitlementCancellation(ctx, orgId, entitlementId).Execute()

unschedule entitlement cancellation



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
	entitlementId := "entitlementId_example" // string | Entitlement ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.UnscheduleEntitlementCancellation(context.Background(), orgId, entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.UnscheduleEntitlementCancellation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnscheduleEntitlementCancellation`: WorkloadEntitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.UnscheduleEntitlementCancellation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnscheduleEntitlementCancellationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntitlementMetaInfo

> WorkloadMetaInfo UpdateEntitlementMetaInfo(ctx, orgId, entitlementId).Data(data).Execute()

update entitlement meta info



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
	entitlementId := "entitlementId_example" // string | Entitlement ID
	data := *openapiclient.NewWorkloadMetaInfo() // WorkloadMetaInfo | Entitlement meta info to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.UpdateEntitlementMetaInfo(context.Background(), orgId, entitlementId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.UpdateEntitlementMetaInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEntitlementMetaInfo`: WorkloadMetaInfo
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.UpdateEntitlementMetaInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntitlementMetaInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**WorkloadMetaInfo**](WorkloadMetaInfo.md) | Entitlement meta info to update | 

### Return type

[**WorkloadMetaInfo**](WorkloadMetaInfo.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntitlementName

> WorkloadEntitlement UpdateEntitlementName(ctx, orgId, entitlementId).Data(data).Execute()

update entitlement name



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
	entitlementId := "entitlementId_example" // string | Entitlement ID
	data := *openapiclient.NewGithubComSugerioMarketplaceServicePkgLegacyRdsDbLibUpdateEntitlementNameParams() // GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibUpdateEntitlementNameParams | UpdateEntitlementNameParams

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.UpdateEntitlementName(context.Background(), orgId, entitlementId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.UpdateEntitlementName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEntitlementName`: WorkloadEntitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.UpdateEntitlementName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntitlementNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

**data** | [**GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibUpdateEntitlementNameParams
**](GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibUpdateEntitlementNameParams.md) | UpdateEntitlementNameParams |

### Return type

[**WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntitlementSeat

> WorkloadEntitlement UpdateEntitlementSeat(ctx, orgId, entitlementId).NewSeat(newSeat).Execute()

update entitlement seat



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
	entitlementId := "entitlementId_example" // string | Entitlement ID
	newSeat := int32(56) // int32 | New seat number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.UpdateEntitlementSeat(context.Background(), orgId, entitlementId).NewSeat(newSeat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.UpdateEntitlementSeat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEntitlementSeat`: WorkloadEntitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.UpdateEntitlementSeat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntitlementSeatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newSeat** | **int32** | New seat number | 

### Return type

[**WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

