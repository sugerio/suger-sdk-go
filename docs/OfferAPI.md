# \OfferAPI

All URIs are relative to *http://https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOffer**](OfferAPI.md#CancelOffer) | **Post** /org/{orgId}/offer/{offerId}/cancel | cancel offer
[**CreateOffer**](OfferAPI.md#CreateOffer) | **Post** /org/{orgId}/offer | create offer
[**CreateOrUpdateDraftOffer**](OfferAPI.md#CreateOrUpdateDraftOffer) | **Post** /org/{orgId}/draftOffer | create or update draft offer
[**DeleteOffer**](OfferAPI.md#DeleteOffer) | **Delete** /org/{orgId}/offer/{offerId} | delete offer
[**ExtendPrivateOfferExpiryDate**](OfferAPI.md#ExtendPrivateOfferExpiryDate) | **Post** /org/{orgId}/offer/{offerId}/extendExpiryDate | extend offer expiry date
[**GetOffer**](OfferAPI.md#GetOffer) | **Get** /org/{orgId}/offer/{offerId} | get offer
[**GetOfferByExternalId**](OfferAPI.md#GetOfferByExternalId) | **Get** /org/{orgId}/offerExternalId/{offerExternalId} | get offer by external ID
[**GetOfferEula**](OfferAPI.md#GetOfferEula) | **Get** /org/{orgId}/offer/{offerId}/eula | get offer EULA
[**GetOfferResellerEula**](OfferAPI.md#GetOfferResellerEula) | **Get** /org/{orgId}/offer/{offerId}/resellerEula | get offer reseller EULA
[**ListOffers**](OfferAPI.md#ListOffers) | **Get** /org/{orgId}/offer | list offers
[**SendOfferNotifications**](OfferAPI.md#SendOfferNotifications) | **Post** /org/{orgId}/offer/{offerId}/notifyContacts | notify offer contacts
[**UpdateOfferMetaInfo**](OfferAPI.md#UpdateOfferMetaInfo) | **Patch** /org/{orgId}/offer/{offerId}/metaInfo | update offer meta info



## CancelOffer

> string CancelOffer(ctx, orgId, offerId).Execute()

cancel offer



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
	offerId := "offerId_example" // string | Offer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OfferAPI.CancelOffer(context.Background(), orgId, offerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.CancelOffer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelOffer`: string
	fmt.Fprintf(os.Stdout, "Response from `OfferAPI.CancelOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**offerId** | **string** | Offer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOfferRequest struct via the builder pattern


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


## CreateOffer

> WorkloadOffer CreateOffer(ctx, orgId).Data(data).Execute()

create offer



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
	data := *openapiclient.NewWorkloadOffer() // WorkloadOffer | Offer to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OfferAPI.CreateOffer(context.Background(), orgId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.CreateOffer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOffer`: WorkloadOffer
	fmt.Fprintf(os.Stdout, "Response from `OfferAPI.CreateOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WorkloadOffer**](WorkloadOffer.md) | Offer to create | 

### Return type

[**WorkloadOffer**](WorkloadOffer.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateDraftOffer

> WorkloadOffer CreateOrUpdateDraftOffer(ctx, orgId).Data(data).Execute()

create or update draft offer



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
	data := *openapiclient.NewWorkloadOffer() // WorkloadOffer | the draft offer to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OfferAPI.CreateOrUpdateDraftOffer(context.Background(), orgId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.CreateOrUpdateDraftOffer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrUpdateDraftOffer`: WorkloadOffer
	fmt.Fprintf(os.Stdout, "Response from `OfferAPI.CreateOrUpdateDraftOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateDraftOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WorkloadOffer**](WorkloadOffer.md) | the draft offer to create | 

### Return type

[**WorkloadOffer**](WorkloadOffer.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOffer

> string DeleteOffer(ctx, orgId, offerId).Execute()

delete offer



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
	offerId := "offerId_example" // string | Offer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OfferAPI.DeleteOffer(context.Background(), orgId, offerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.DeleteOffer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOffer`: string
	fmt.Fprintf(os.Stdout, "Response from `OfferAPI.DeleteOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**offerId** | **string** | Offer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOfferRequest struct via the builder pattern


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


## ExtendPrivateOfferExpiryDate

> string ExtendPrivateOfferExpiryDate(ctx, orgId, offerId).NewExpiryDate(newExpiryDate).Execute()

extend offer expiry date



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
	offerId := "offerId_example" // string | Offer ID
	newExpiryDate := "newExpiryDate_example" // string | new expiry date in YYYY-MM-DD format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OfferAPI.ExtendPrivateOfferExpiryDate(context.Background(), orgId, offerId).NewExpiryDate(newExpiryDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.ExtendPrivateOfferExpiryDate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExtendPrivateOfferExpiryDate`: string
	fmt.Fprintf(os.Stdout, "Response from `OfferAPI.ExtendPrivateOfferExpiryDate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**offerId** | **string** | Offer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendPrivateOfferExpiryDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newExpiryDate** | **string** | new expiry date in YYYY-MM-DD format | 

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


## GetOffer

> WorkloadOffer GetOffer(ctx, orgId, offerId).Execute()

get offer



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
	offerId := "offerId_example" // string | Offer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OfferAPI.GetOffer(context.Background(), orgId, offerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.GetOffer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOffer`: WorkloadOffer
	fmt.Fprintf(os.Stdout, "Response from `OfferAPI.GetOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**offerId** | **string** | Offer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkloadOffer**](WorkloadOffer.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOfferByExternalId

> WorkloadOffer GetOfferByExternalId(ctx, orgId, offerExternalId).Execute()

get offer by external ID



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
	offerExternalId := "offerExternalId_example" // string | Offer External ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OfferAPI.GetOfferByExternalId(context.Background(), orgId, offerExternalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.GetOfferByExternalId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOfferByExternalId`: WorkloadOffer
	fmt.Fprintf(os.Stdout, "Response from `OfferAPI.GetOfferByExternalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**offerExternalId** | **string** | Offer External ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferByExternalIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkloadOffer**](WorkloadOffer.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOfferEula

> string GetOfferEula(ctx, orgId, offerId).Execute()

get offer EULA



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
	offerId := "offerId_example" // string | Offer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OfferAPI.GetOfferEula(context.Background(), orgId, offerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.GetOfferEula``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOfferEula`: string
	fmt.Fprintf(os.Stdout, "Response from `OfferAPI.GetOfferEula`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**offerId** | **string** | Offer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferEulaRequest struct via the builder pattern


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


## GetOfferResellerEula

> string GetOfferResellerEula(ctx, orgId, offerId).Execute()

get offer reseller EULA



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
	offerId := "offerId_example" // string | Offer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OfferAPI.GetOfferResellerEula(context.Background(), orgId, offerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.GetOfferResellerEula``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOfferResellerEula`: string
	fmt.Fprintf(os.Stdout, "Response from `OfferAPI.GetOfferResellerEula`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**offerId** | **string** | Offer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferResellerEulaRequest struct via the builder pattern


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


## ListOffers

> []WorkloadOffer ListOffers(ctx, orgId).Status(status).Partner(partner).OfferType(offerType).ProductId(productId).BuyerId(buyerId).HubspotDealId(hubspotDealId).ContactId(contactId).Limit(limit).Offset(offset).Execute()

list offers



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
	status := "status_example" // string | filter by offer status (optional)
	partner := "partner_example" // string | filter by partner (optional)
	offerType := "offerType_example" // string | filter by offerType (optional)
	productId := "productId_example" // string | filter by productId (optional)
	buyerId := "buyerId_example" // string | filter by buyerId (optional)
	hubspotDealId := "hubspotDealId_example" // string | filter by hubspotDealId (optional)
	contactId := "contactId_example" // string | filter by contactId (optional)
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OfferAPI.ListOffers(context.Background(), orgId).Status(status).Partner(partner).OfferType(offerType).ProductId(productId).BuyerId(buyerId).HubspotDealId(hubspotDealId).ContactId(contactId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.ListOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOffers`: []WorkloadOffer
	fmt.Fprintf(os.Stdout, "Response from `OfferAPI.ListOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | filter by offer status | 
 **partner** | **string** | filter by partner | 
 **offerType** | **string** | filter by offerType | 
 **productId** | **string** | filter by productId | 
 **buyerId** | **string** | filter by buyerId | 
 **hubspotDealId** | **string** | filter by hubspotDealId | 
 **contactId** | **string** | filter by contactId | 
 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**[]WorkloadOffer**](WorkloadOffer.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendOfferNotifications

> NotificationEvent SendOfferNotifications(ctx, orgId, offerId).ContactIds(contactIds).Execute()

notify offer contacts



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
	offerId := "offerId_example" // string | Offer ID
	contactIds := []string{"Property_example"} // []string | List of Contact IDs, if emoty or nil, send notifications to all contacts of the offer (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OfferAPI.SendOfferNotifications(context.Background(), orgId, offerId).ContactIds(contactIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.SendOfferNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendOfferNotifications`: NotificationEvent
	fmt.Fprintf(os.Stdout, "Response from `OfferAPI.SendOfferNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**offerId** | **string** | Offer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendOfferNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contactIds** | **[]string** | List of Contact IDs, if emoty or nil, send notifications to all contacts of the offer | 

### Return type

[**NotificationEvent**](NotificationEvent.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOfferMetaInfo

> WorkloadMetaInfo UpdateOfferMetaInfo(ctx, orgId, offerId).Data(data).Execute()

update offer meta info



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
	offerId := "offerId_example" // string | Offer ID
	data := *openapiclient.NewWorkloadMetaInfo() // WorkloadMetaInfo | Offer meta info to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OfferAPI.UpdateOfferMetaInfo(context.Background(), orgId, offerId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.UpdateOfferMetaInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOfferMetaInfo`: WorkloadMetaInfo
	fmt.Fprintf(os.Stdout, "Response from `OfferAPI.UpdateOfferMetaInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**offerId** | **string** | Offer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOfferMetaInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**WorkloadMetaInfo**](WorkloadMetaInfo.md) | Offer meta info to update | 

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

