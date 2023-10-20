# \SignupAPI

All URIs are relative to *http://https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClientSignupPageConfigInfo**](SignupAPI.md#GetClientSignupPageConfigInfo) | **Get** /org/{orgId}/clientSignupPageConfigInfo | get client signup page config info
[**UpdateClientSignupPageConfigInfo**](SignupAPI.md#UpdateClientSignupPageConfigInfo) | **Patch** /org/{orgId}/clientSignupPageConfigInfo | update client signup page config info



## GetClientSignupPageConfigInfo

> ClientSignupPageConfigInfo GetClientSignupPageConfigInfo(ctx, orgId).Execute()

get client signup page config info



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
    resp, r, err := apiClient.SignupAPI.GetClientSignupPageConfigInfo(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignupAPI.GetClientSignupPageConfigInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientSignupPageConfigInfo`: ClientSignupPageConfigInfo
    fmt.Fprintf(os.Stdout, "Response from `SignupAPI.GetClientSignupPageConfigInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientSignupPageConfigInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientSignupPageConfigInfo**](ClientSignupPageConfigInfo.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientSignupPageConfigInfo

> ClientSignupPageConfigInfo UpdateClientSignupPageConfigInfo(ctx, orgId).Data(data).Execute()

update client signup page config info



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
    data := *openapiclient.NewClientSignupPageConfigInfo() // ClientSignupPageConfigInfo | The client signup page config info to update with

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignupAPI.UpdateClientSignupPageConfigInfo(context.Background(), orgId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignupAPI.UpdateClientSignupPageConfigInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClientSignupPageConfigInfo`: ClientSignupPageConfigInfo
    fmt.Fprintf(os.Stdout, "Response from `SignupAPI.UpdateClientSignupPageConfigInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientSignupPageConfigInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**ClientSignupPageConfigInfo**](ClientSignupPageConfigInfo.md) | The client signup page config info to update with | 

### Return type

[**ClientSignupPageConfigInfo**](ClientSignupPageConfigInfo.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

