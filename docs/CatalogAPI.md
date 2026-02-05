# \CatalogAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVendorDetails**](CatalogAPI.md#GetVendorDetails) | **Get** /org/{orgId}/catalog/vendor/{vendorId} | get vendor details
[**ListVendors**](CatalogAPI.md#ListVendors) | **Get** /org/{orgId}/catalog/vendor | list vendors



## GetVendorDetails

> GithubComSugerioMarketplaceServicePkgOrmGlobalCompany GetVendorDetails(ctx, orgId, vendorId).Execute()

get vendor details



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
	vendorId := "vendorId_example" // string | Vendor ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPI.GetVendorDetails(context.Background(), orgId, vendorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.GetVendorDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVendorDetails`: GithubComSugerioMarketplaceServicePkgOrmGlobalCompany
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.GetVendorDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**vendorId** | **string** | Vendor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GithubComSugerioMarketplaceServicePkgOrmGlobalCompany**](GithubComSugerioMarketplaceServicePkgOrmGlobalCompany.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVendors

> GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmGlobalCompany ListVendors(ctx, orgId).PageSize(pageSize).PageNumber(pageNumber).Q(q).S(s).Execute()

list vendors



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
	pageSize := int32(56) // int32 | Number of items per page (default 20, max 1000) (optional)
	pageNumber := int32(56) // int32 | Page number (default 1) (optional)
	q := "q_example" // string | LISP-style filter expression (e.g., '(and (= is_vendor true) (ilike name \\ (optional)
	s := "s_example" // string | Sort fields: 'field:asc,field2:desc' or '-field,field2' format (e.g., 'name:asc' or '-name') (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPI.ListVendors(context.Background(), orgId).PageSize(pageSize).PageNumber(pageNumber).Q(q).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.ListVendors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVendors`: GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmGlobalCompany
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.ListVendors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVendorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | Number of items per page (default 20, max 1000) | 
 **pageNumber** | **int32** | Page number (default 1) | 
 **q** | **string** | LISP-style filter expression (e.g., &#39;(and (&#x3D; is_vendor true) (ilike name \\ | 
 **s** | **string** | Sort fields: &#39;field:asc,field2:desc&#39; or &#39;-field,field2&#39; format (e.g., &#39;name:asc&#39; or &#39;-name&#39;) | 

### Return type

[**GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmGlobalCompany**](GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmGlobalCompany.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

