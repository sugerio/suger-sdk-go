# \OfferAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOffer**](OfferAPI.md#CancelOffer) | **Post** /org/{orgId}/offer/{offerId}/cancel | cancel offer
[**CreateOffer**](OfferAPI.md#CreateOffer) | **Post** /org/{orgId}/offer | create offer
[**CreateOrUpdateDraftOffer**](OfferAPI.md#CreateOrUpdateDraftOffer) | **Post** /org/{orgId}/draftOffer | create or update draft offer
[**DeleteOffer**](OfferAPI.md#DeleteOffer) | **Delete** /org/{orgId}/offer/{offerId} | delete offer
[**ExtendPrivateOfferExpiryDate**](OfferAPI.md#ExtendPrivateOfferExpiryDate) | **Post** /org/{orgId}/offer/{offerId}/extendExpiryDate | extend offer expiry date
[**GetOffer**](OfferAPI.md#GetOffer) | **Get** /org/{orgId}/offer/{offerId} | get offer
[**GetOfferEula**](OfferAPI.md#GetOfferEula) | **Get** /org/{orgId}/offer/{offerId}/eula | get offer EULA
[**ListOffersByContact**](OfferAPI.md#ListOffersByContact) | **Get** /org/{orgId}/contact/{contactId}/offer | list offers by contact
[**ListOffersByOrganization**](OfferAPI.md#ListOffersByOrganization) | **Get** /org/{orgId}/offer | list offers by organization
[**ListOffersByPartner**](OfferAPI.md#ListOffersByPartner) | **Get** /org/{orgId}/partner/{partner}/offer | list offers by partner
[**ListOffersByProduct**](OfferAPI.md#ListOffersByProduct) | **Get** /org/{orgId}/product/{productId}/offer | list offers by product
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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOffersByContact

> []WorkloadOffer ListOffersByContact(ctx, orgId, contactId).Execute()

list offers by contact



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
    contactId := "contactId_example" // string | Contact ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OfferAPI.ListOffersByContact(context.Background(), orgId, contactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.ListOffersByContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOffersByContact`: []WorkloadOffer
    fmt.Fprintf(os.Stdout, "Response from `OfferAPI.ListOffersByContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**contactId** | **string** | Contact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOffersByContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WorkloadOffer**](WorkloadOffer.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOffersByOrganization

> []WorkloadOffer ListOffersByOrganization(ctx, orgId).Execute()

list offers by organization



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OfferAPI.ListOffersByOrganization(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.ListOffersByOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOffersByOrganization`: []WorkloadOffer
    fmt.Fprintf(os.Stdout, "Response from `OfferAPI.ListOffersByOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOffersByOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]WorkloadOffer**](WorkloadOffer.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOffersByPartner

> []WorkloadOffer ListOffersByPartner(ctx, orgId, partner).Execute()

list offers by partner



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
    partner := "partner_example" // string | Cloud Partner

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OfferAPI.ListOffersByPartner(context.Background(), orgId, partner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.ListOffersByPartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOffersByPartner`: []WorkloadOffer
    fmt.Fprintf(os.Stdout, "Response from `OfferAPI.ListOffersByPartner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**partner** | **string** | Cloud Partner | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOffersByPartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WorkloadOffer**](WorkloadOffer.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOffersByProduct

> []WorkloadOffer ListOffersByProduct(ctx, orgId, productId).Execute()

list offers by product



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
    productId := "productId_example" // string | Product ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OfferAPI.ListOffersByProduct(context.Background(), orgId, productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfferAPI.ListOffersByProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOffersByProduct`: []WorkloadOffer
    fmt.Fprintf(os.Stdout, "Response from `OfferAPI.ListOffersByProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOffersByProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WorkloadOffer**](WorkloadOffer.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

