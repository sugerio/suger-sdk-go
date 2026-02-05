# \VendorAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVendorProduct**](VendorAPI.md#GetVendorProduct) | **Get** /org/{orgId}/vendor/{vendorId}/product/{productId} | get vendor product
[**ListVendorProducts**](VendorAPI.md#ListVendorProducts) | **Get** /org/{orgId}/vendor/{vendorId}/product | list vendor products



## GetVendorProduct

> GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing GetVendorProduct(ctx, orgId, vendorId, productId).Execute()

get vendor product



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
	vendorId := "vendorId_example" // string | Vendor ID (Company ID where is_vendor=true)
	productId := "productId_example" // string | Product ID (MarketplaceListing internal UUID)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VendorAPI.GetVendorProduct(context.Background(), orgId, vendorId, productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorAPI.GetVendorProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVendorProduct`: GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing
	fmt.Fprintf(os.Stdout, "Response from `VendorAPI.GetVendorProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**vendorId** | **string** | Vendor ID (Company ID where is_vendor&#x3D;true) | 
**productId** | **string** | Product ID (MarketplaceListing internal UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing**](GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVendorProducts

> GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmMarketplaceListing ListVendorProducts(ctx, orgId, vendorId).PageSize(pageSize).PageNumber(pageNumber).Q(q).S(s).Execute()

list vendor products



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
	vendorId := "vendorId_example" // string | Vendor ID (Company ID where is_vendor=true)
	pageSize := int32(56) // int32 | Number of items per page (default 20, max 1000) (optional)
	pageNumber := int32(56) // int32 | Page number (default 1) (optional)
	q := "q_example" // string | LISP-style filter expression (e.g., '(ilike name \\ (optional)
	s := "s_example" // string | Sort fields: 'field:asc,field2:desc' or '-field,field2' format (e.g., 'name:asc' or '-name') (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VendorAPI.ListVendorProducts(context.Background(), orgId, vendorId).PageSize(pageSize).PageNumber(pageNumber).Q(q).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorAPI.ListVendorProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVendorProducts`: GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmMarketplaceListing
	fmt.Fprintf(os.Stdout, "Response from `VendorAPI.ListVendorProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**vendorId** | **string** | Vendor ID (Company ID where is_vendor&#x3D;true) | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVendorProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | Number of items per page (default 20, max 1000) | 
 **pageNumber** | **int32** | Page number (default 1) | 
 **q** | **string** | LISP-style filter expression (e.g., &#39;(ilike name \\ | 
 **s** | **string** | Sort fields: &#39;field:asc,field2:desc&#39; or &#39;-field,field2&#39; format (e.g., &#39;name:asc&#39; or &#39;-name&#39;) | 

### Return type

[**GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmMarketplaceListing**](GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmMarketplaceListing.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

