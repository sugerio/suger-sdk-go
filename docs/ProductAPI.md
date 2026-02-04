# \ProductAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateDraftProduct**](ProductAPI.md#CreateOrUpdateDraftProduct) | **Post** /org/{orgId}/draftProduct | create or update draft product
[**CreateProduct**](ProductAPI.md#CreateProduct) | **Post** /org/{orgId}/product | create product
[**DeleteProduct**](ProductAPI.md#DeleteProduct) | **Delete** /org/{orgId}/product/{productId} | delete product
[**GetProduct**](ProductAPI.md#GetProduct) | **Get** /org/{orgId}/product/{productId} | get product
[**ListProductMeteringDimensions**](ProductAPI.md#ListProductMeteringDimensions) | **Get** /org/{orgId}/product/{productId}/dimension | list metering dimensions of product
[**ListProducts**](ProductAPI.md#ListProducts) | **Get** /org/{orgId}/product | list products
[**ListProductsByPartner**](ProductAPI.md#ListProductsByPartner) | **Get** /org/{orgId}/partner/{partner}/product | list products by partner
[**PublishProduct**](ProductAPI.md#PublishProduct) | **Patch** /org/{orgId}/product/{productId}/publish | publish product
[**QueryProducts**](ProductAPI.md#QueryProducts) | **Get** /org/{orgId}/product/query | query products
[**UpdateProduct**](ProductAPI.md#UpdateProduct) | **Patch** /org/{orgId}/product/{productId} | update product
[**UpdateProductFulfillmentUrl**](ProductAPI.md#UpdateProductFulfillmentUrl) | **Patch** /org/{orgId}/product/{productId}/fulfillmentUrl | update product fulfillment url
[**UpdateProductMetaInfo**](ProductAPI.md#UpdateProductMetaInfo) | **Patch** /org/{orgId}/product/{productId}/metaInfo | update product meta info



## CreateOrUpdateDraftProduct

> WorkloadProduct CreateOrUpdateDraftProduct(ctx, orgId).Data(data).Execute()

create or update draft product



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
	data := *openapiclient.NewWorkloadProduct() // WorkloadProduct | the draft product to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.CreateOrUpdateDraftProduct(context.Background(), orgId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.CreateOrUpdateDraftProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrUpdateDraftProduct`: WorkloadProduct
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.CreateOrUpdateDraftProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateDraftProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WorkloadProduct**](WorkloadProduct.md) | the draft product to create | 

### Return type

[**WorkloadProduct**](WorkloadProduct.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProduct

> WorkloadProduct CreateProduct(ctx, orgId).Data(data).Execute()

create product



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
	data := *openapiclient.NewWorkloadProduct() // WorkloadProduct | the product to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.CreateProduct(context.Background(), orgId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.CreateProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProduct`: WorkloadProduct
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.CreateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WorkloadProduct**](WorkloadProduct.md) | the product to create | 

### Return type

[**WorkloadProduct**](WorkloadProduct.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProduct

> string DeleteProduct(ctx, orgId, productId).Execute()

delete product



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
	productId := "productId_example" // string | Product ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.DeleteProduct(context.Background(), orgId, productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.DeleteProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProduct`: string
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.DeleteProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductRequest struct via the builder pattern


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


## GetProduct

> WorkloadProduct GetProduct(ctx, orgId, productId).Execute()

get product



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
	productId := "productId_example" // string | Product ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.GetProduct(context.Background(), orgId, productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProduct`: WorkloadProduct
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkloadProduct**](WorkloadProduct.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductMeteringDimensions

> []MeteringDimension ListProductMeteringDimensions(ctx, orgId, productId).Execute()

list metering dimensions of product



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
	productId := "productId_example" // string | Product ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ListProductMeteringDimensions(context.Background(), orgId, productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ListProductMeteringDimensions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductMeteringDimensions`: []MeteringDimension
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ListProductMeteringDimensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProductMeteringDimensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]MeteringDimension**](MeteringDimension.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProducts

> []WorkloadProduct ListProducts(ctx, orgId).Partner(partner).Limit(limit).Offset(offset).View(view).ForCreateOffer(forCreateOffer).SalesforceOpportunityId(salesforceOpportunityId).SalesforceQuoteId(salesforceQuoteId).HubspotDealId(hubspotDealId).Execute()

list products



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
	partner := "partner_example" // string | Filter by partner (optional)
	limit := int32(56) // int32 | List pagination size, default 100, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)
	view := "view_example" // string | Response format: 'lite' or 'full' (default: full). 'lite' excludes info/metaInfo fields. (optional)
	forCreateOffer := true // bool | Enable product filtering for create offer purpose (optional)
	salesforceOpportunityId := "salesforceOpportunityId_example" // string | Salesforce Opportunity ID for product filtering (used with forCreateOffer=true) (optional)
	salesforceQuoteId := "salesforceQuoteId_example" // string | Salesforce Quote ID for product filtering (used with forCreateOffer=true) (optional)
	hubspotDealId := "hubspotDealId_example" // string | HubSpot Deal ID for product filtering (used with forCreateOffer=true) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ListProducts(context.Background(), orgId).Partner(partner).Limit(limit).Offset(offset).View(view).ForCreateOffer(forCreateOffer).SalesforceOpportunityId(salesforceOpportunityId).SalesforceQuoteId(salesforceQuoteId).HubspotDealId(hubspotDealId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ListProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProducts`: []WorkloadProduct
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ListProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partner** | **string** | Filter by partner | 
 **limit** | **int32** | List pagination size, default 100, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 
 **view** | **string** | Response format: &#39;lite&#39; or &#39;full&#39; (default: full). &#39;lite&#39; excludes info/metaInfo fields. | 
 **forCreateOffer** | **bool** | Enable product filtering for create offer purpose | 
 **salesforceOpportunityId** | **string** | Salesforce Opportunity ID for product filtering (used with forCreateOffer&#x3D;true) | 
 **salesforceQuoteId** | **string** | Salesforce Quote ID for product filtering (used with forCreateOffer&#x3D;true) | 
 **hubspotDealId** | **string** | HubSpot Deal ID for product filtering (used with forCreateOffer&#x3D;true) | 

### Return type

[**[]WorkloadProduct**](WorkloadProduct.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductsByPartner

> []WorkloadProduct ListProductsByPartner(ctx, orgId, partner).Execute()

list products by partner



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
	partner := "partner_example" // string | Cloud Partner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ListProductsByPartner(context.Background(), orgId, partner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ListProductsByPartner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductsByPartner`: []WorkloadProduct
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ListProductsByPartner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**partner** | **string** | Cloud Partner | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProductsByPartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WorkloadProduct**](WorkloadProduct.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishProduct

> WorkloadProduct PublishProduct(ctx, orgId, productId).Data(data).Execute()

publish product



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
	productId := "productId_example" // string | Product ID
	data := *openapiclient.NewWorkloadProduct() // WorkloadProduct | the product to publish

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.PublishProduct(context.Background(), orgId, productId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.PublishProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishProduct`: WorkloadProduct
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.PublishProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**WorkloadProduct**](WorkloadProduct.md) | the product to publish | 

### Return type

[**WorkloadProduct**](WorkloadProduct.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryProducts

> GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmProduct QueryProducts(ctx, orgId).PageSize(pageSize).PageNumber(pageNumber).Fields(fields).Q(q).S(s).Execute()

query products



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
	pageSize := int32(56) // int32 | Number of items per page (default 100, max 200) (optional)
	pageNumber := int32(56) // int32 | Page number (default 1) (optional)
	fields := "fields_example" // string | Comma-separated fields to return (e.g., 'id,name,status'). Dot notation supported. (optional)
	q := "q_example" // string | LISP-style filter expression (e.g., '(= partner \\ (optional)
	s := "s_example" // string | Sort fields: 'field:asc,field2:desc' or '-field,field2' format (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.QueryProducts(context.Background(), orgId).PageSize(pageSize).PageNumber(pageNumber).Fields(fields).Q(q).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.QueryProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryProducts`: GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmProduct
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.QueryProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | Number of items per page (default 100, max 200) | 
 **pageNumber** | **int32** | Page number (default 1) | 
 **fields** | **string** | Comma-separated fields to return (e.g., &#39;id,name,status&#39;). Dot notation supported. | 
 **q** | **string** | LISP-style filter expression (e.g., &#39;(&#x3D; partner \\ | 
 **s** | **string** | Sort fields: &#39;field:asc,field2:desc&#39; or &#39;-field,field2&#39; format | 

### Return type

[**GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmProduct**](GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmProduct.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProduct

> WorkloadProduct UpdateProduct(ctx, orgId, productId).Data(data).Execute()

update product



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
	productId := "productId_example" // string | Product ID
	data := *openapiclient.NewWorkloadProduct() // WorkloadProduct | the product to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.UpdateProduct(context.Background(), orgId, productId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdateProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProduct`: WorkloadProduct
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**WorkloadProduct**](WorkloadProduct.md) | the product to update | 

### Return type

[**WorkloadProduct**](WorkloadProduct.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProductFulfillmentUrl

> WorkloadProduct UpdateProductFulfillmentUrl(ctx, orgId, productId).Data(data).Execute()

update product fulfillment url



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
	productId := "productId_example" // string | Product ID
	data := *openapiclient.NewUpdateProductParams("FulfillmentUrl_example", "Id_example", "OrganizationID_example") // UpdateProductParams | Update Product Params

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.UpdateProductFulfillmentUrl(context.Background(), orgId, productId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdateProductFulfillmentUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProductFulfillmentUrl`: WorkloadProduct
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdateProductFulfillmentUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductFulfillmentUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**UpdateProductParams**](UpdateProductParams.md) | Update Product Params | 

### Return type

[**WorkloadProduct**](WorkloadProduct.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProductMetaInfo

> WorkloadMetaInfo UpdateProductMetaInfo(ctx, orgId, productId).Data(data).Execute()

update product meta info



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
	productId := "productId_example" // string | Product ID
	data := *openapiclient.NewWorkloadMetaInfo() // WorkloadMetaInfo | Product meta info to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.UpdateProductMetaInfo(context.Background(), orgId, productId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdateProductMetaInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProductMetaInfo`: WorkloadMetaInfo
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdateProductMetaInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductMetaInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**WorkloadMetaInfo**](WorkloadMetaInfo.md) | Product meta info to update | 

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

