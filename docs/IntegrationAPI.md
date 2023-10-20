# \IntegrationAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIntegration**](IntegrationAPI.md#CreateIntegration) | **Post** /org/{orgId}/integration | create integration
[**DeleteIntegration**](IntegrationAPI.md#DeleteIntegration) | **Delete** /org/{orgId}/integration/{partner}/{service} | delete integration
[**GetIntegration**](IntegrationAPI.md#GetIntegration) | **Get** /org/{orgId}/integration/{partner}/{service} | get integration
[**ListIntegrationsByOrganization**](IntegrationAPI.md#ListIntegrationsByOrganization) | **Get** /org/{orgId}/integration | list integrations by organization
[**UpdateIntegration**](IntegrationAPI.md#UpdateIntegration) | **Patch** /org/{orgId}/integration/{partner}/{service} | update integration
[**VerifyIntegration**](IntegrationAPI.md#VerifyIntegration) | **Post** /org/{orgId}/integration/{partner}/{service}/verify | verify integration



## CreateIntegration

> IdentityIntegration CreateIntegration(ctx, orgId).Data(data).Execute()

create integration



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
    data := *openapiclient.NewCreateIntegrationParams(*openapiclient.NewIntegrationInfo(), "OrganizationID_example", openapiclient.Partner(""), openapiclient.PartnerService("DEFAULT")) // CreateIntegrationParams | Create Integration Params

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationAPI.CreateIntegration(context.Background(), orgId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.CreateIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIntegration`: IdentityIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.CreateIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**CreateIntegrationParams**](CreateIntegrationParams.md) | Create Integration Params | 

### Return type

[**IdentityIntegration**](IdentityIntegration.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIntegration

> string DeleteIntegration(ctx, orgId, partner, service).Execute()

delete integration



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
    service := "service_example" // string | Partner Service

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationAPI.DeleteIntegration(context.Background(), orgId, partner, service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.DeleteIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIntegration`: string
    fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.DeleteIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**partner** | **string** | Cloud Partner | 
**service** | **string** | Partner Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationRequest struct via the builder pattern


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


## GetIntegration

> IdentityIntegration GetIntegration(ctx, orgId, partner, service).Execute()

get integration



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
    service := "service_example" // string | Partner Service

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationAPI.GetIntegration(context.Background(), orgId, partner, service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegration`: IdentityIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**partner** | **string** | Cloud Partner | 
**service** | **string** | Partner Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IdentityIntegration**](IdentityIntegration.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationsByOrganization

> []IdentityIntegration ListIntegrationsByOrganization(ctx, orgId).Execute()

list integrations by organization



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
    resp, r, err := apiClient.IntegrationAPI.ListIntegrationsByOrganization(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.ListIntegrationsByOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIntegrationsByOrganization`: []IdentityIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.ListIntegrationsByOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationsByOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IdentityIntegration**](IdentityIntegration.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIntegration

> IdentityIntegration UpdateIntegration(ctx, orgId, partner, service).Data(data).Execute()

update integration



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
    service := "service_example" // string | Partner Service
    data := *openapiclient.NewUpdateIntegrationParams(*openapiclient.NewIntegrationInfo(), "OrganizationID_example", openapiclient.Partner(""), openapiclient.PartnerService("DEFAULT")) // UpdateIntegrationParams | Update Integration Params

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationAPI.UpdateIntegration(context.Background(), orgId, partner, service).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.UpdateIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIntegration`: IdentityIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.UpdateIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**partner** | **string** | Cloud Partner | 
**service** | **string** | Partner Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**UpdateIntegrationParams**](UpdateIntegrationParams.md) | Update Integration Params | 

### Return type

[**IdentityIntegration**](IdentityIntegration.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyIntegration

> bool VerifyIntegration(ctx, orgId, partner, service).Execute()

verify integration



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
    service := "service_example" // string | Partner Service

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationAPI.VerifyIntegration(context.Background(), orgId, partner, service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.VerifyIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyIntegration`: bool
    fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.VerifyIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**partner** | **string** | Cloud Partner | 
**service** | **string** | Partner Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**bool**

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

