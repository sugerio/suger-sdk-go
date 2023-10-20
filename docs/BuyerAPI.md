# \BuyerAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBuyer**](BuyerAPI.md#GetBuyer) | **Get** /org/{orgId}/buyer/{buyerId} | get buyer
[**ListBuyersByContact**](BuyerAPI.md#ListBuyersByContact) | **Get** /org/{orgId}/contact/{contactId}/buyer | list buyers by contact
[**ListBuyersByOrganization**](BuyerAPI.md#ListBuyersByOrganization) | **Get** /org/{orgId}/buyer | list buyers by organization
[**ListBuyersByPartner**](BuyerAPI.md#ListBuyersByPartner) | **Get** /org/{orgId}/partner/{partner}/buyer | list buyers by partner
[**UpdateBuyer**](BuyerAPI.md#UpdateBuyer) | **Patch** /org/{orgId}/buyer/{buyerId} | update buyer



## GetBuyer

> IdentityBuyer GetBuyer(ctx, orgId, buyerId).Execute()

get buyer



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
    resp, r, err := apiClient.BuyerAPI.GetBuyer(context.Background(), orgId, buyerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.GetBuyer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyer`: IdentityBuyer
    fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.GetBuyer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**buyerId** | **string** | Buyer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdentityBuyer**](IdentityBuyer.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuyersByContact

> []IdentityBuyer ListBuyersByContact(ctx, orgId, contactId).Execute()

list buyers by contact



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
    resp, r, err := apiClient.BuyerAPI.ListBuyersByContact(context.Background(), orgId, contactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.ListBuyersByContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyersByContact`: []IdentityBuyer
    fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.ListBuyersByContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**contactId** | **string** | Contact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyersByContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]IdentityBuyer**](IdentityBuyer.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuyersByOrganization

> []IdentityBuyer ListBuyersByOrganization(ctx, orgId).Execute()

list buyers by organization



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
    resp, r, err := apiClient.BuyerAPI.ListBuyersByOrganization(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.ListBuyersByOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyersByOrganization`: []IdentityBuyer
    fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.ListBuyersByOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyersByOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IdentityBuyer**](IdentityBuyer.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuyersByPartner

> []IdentityBuyer ListBuyersByPartner(ctx, orgId, partner).Execute()

list buyers by partner



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
    resp, r, err := apiClient.BuyerAPI.ListBuyersByPartner(context.Background(), orgId, partner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.ListBuyersByPartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyersByPartner`: []IdentityBuyer
    fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.ListBuyersByPartner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**partner** | **string** | Cloud Partner | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyersByPartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]IdentityBuyer**](IdentityBuyer.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBuyer

> IdentityBuyer UpdateBuyer(ctx, orgId, buyerId).Data(data).Execute()

update buyer



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
    data := *openapiclient.NewUpdateBuyerParams() // UpdateBuyerParams | UpdateBuyerParams

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerAPI.UpdateBuyer(context.Background(), orgId, buyerId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.UpdateBuyer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyer`: IdentityBuyer
    fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.UpdateBuyer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**buyerId** | **string** | Buyer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**UpdateBuyerParams**](UpdateBuyerParams.md) | UpdateBuyerParams | 

### Return type

[**IdentityBuyer**](IdentityBuyer.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

