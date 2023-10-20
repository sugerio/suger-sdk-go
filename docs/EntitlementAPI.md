# \EntitlementAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEntitlementCredit**](EntitlementAPI.md#AddEntitlementCredit) | **Post** /org/{orgId}/entitlement/{entitlementId}/addCredit | add entitlement credit
[**ApproveEntitlement**](EntitlementAPI.md#ApproveEntitlement) | **Post** /org/{orgId}/entitlement/{entitlementId}/approve | approve entitlement
[**DivideEntitlementCommit**](EntitlementAPI.md#DivideEntitlementCommit) | **Post** /org/{orgId}/entitlement/{entitlementId}/divideCommit | divide entitlement commit
[**GetEntitlement**](EntitlementAPI.md#GetEntitlement) | **Get** /org/{orgId}/entitlement/{entitlementId} | get entitlement
[**ListEntitlements**](EntitlementAPI.md#ListEntitlements) | **Get** /org/{orgId}/entitlement | list entitlements
[**ListEntitlementsByBuyer**](EntitlementAPI.md#ListEntitlementsByBuyer) | **Get** /org/{orgId}/buyer/{buyerId}/entitlement | list entitlements by buyer
[**ListEntitlementsByOffer**](EntitlementAPI.md#ListEntitlementsByOffer) | **Get** /org/{orgId}/offer/{offerId}/entitlement | list entitlements by offer
[**ListEntitlementsByPartner**](EntitlementAPI.md#ListEntitlementsByPartner) | **Get** /org/{orgId}/partner/{partner}/entitlement | list entitlements by partner
[**ListEntitlementsByProduct**](EntitlementAPI.md#ListEntitlementsByProduct) | **Get** /org/{orgId}/product/{productId}/entitlement | list entitlements by product
[**UpdateEntitlementMetaInfo**](EntitlementAPI.md#UpdateEntitlementMetaInfo) | **Patch** /org/{orgId}/entitlement/{entitlementId}/metaInfo | update entitlement meta info
[**UpdateEntitlementName**](EntitlementAPI.md#UpdateEntitlementName) | **Patch** /org/{orgId}/entitlement/{entitlementId}/entitlementName | update entitlement name



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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlements

> []WorkloadEntitlement ListEntitlements(ctx, orgId).Execute()

list entitlements



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
    resp, r, err := apiClient.EntitlementAPI.ListEntitlements(context.Background(), orgId).Execute()
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


### Return type

[**[]WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlementsByBuyer

> []WorkloadEntitlement ListEntitlementsByBuyer(ctx, orgId, buyerId).Execute()

list entitlements by buyer



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
    buyerId := "buyerId_example" // string | Buyer ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementAPI.ListEntitlementsByBuyer(context.Background(), orgId, buyerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.ListEntitlementsByBuyer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntitlementsByBuyer`: []WorkloadEntitlement
    fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.ListEntitlementsByBuyer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**buyerId** | **string** | Buyer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementsByBuyerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlementsByOffer

> []WorkloadEntitlement ListEntitlementsByOffer(ctx, orgId, offerId).Execute()

list entitlements by offer



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
    resp, r, err := apiClient.EntitlementAPI.ListEntitlementsByOffer(context.Background(), orgId, offerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.ListEntitlementsByOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntitlementsByOffer`: []WorkloadEntitlement
    fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.ListEntitlementsByOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**offerId** | **string** | Offer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementsByOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlementsByPartner

> []WorkloadEntitlement ListEntitlementsByPartner(ctx, orgId, partner).Execute()

list entitlements by partner



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
    resp, r, err := apiClient.EntitlementAPI.ListEntitlementsByPartner(context.Background(), orgId, partner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.ListEntitlementsByPartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntitlementsByPartner`: []WorkloadEntitlement
    fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.ListEntitlementsByPartner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**partner** | **string** | Cloud Partner | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementsByPartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlementsByProduct

> []WorkloadEntitlement ListEntitlementsByProduct(ctx, orgId, productId).Execute()

list entitlements by product



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
    resp, r, err := apiClient.EntitlementAPI.ListEntitlementsByProduct(context.Background(), orgId, productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.ListEntitlementsByProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntitlementsByProduct`: []WorkloadEntitlement
    fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.ListEntitlementsByProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementsByProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    orgId := "orgId_example" // string | Organization ID
    entitlementId := "entitlementId_example" // string | Entitlement ID
    data := *openapiclient.NewGithubComSugerioMarketplaceServiceRdsDbLibUpdateEntitlementNameParams() // GithubComSugerioMarketplaceServiceRdsDbLibUpdateEntitlementNameParams | UpdateEntitlementNameParams

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


 **data** | [**GithubComSugerioMarketplaceServiceRdsDbLibUpdateEntitlementNameParams**](GithubComSugerioMarketplaceServiceRdsDbLibUpdateEntitlementNameParams.md) | UpdateEntitlementNameParams | 

### Return type

[**WorkloadEntitlement**](WorkloadEntitlement.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

