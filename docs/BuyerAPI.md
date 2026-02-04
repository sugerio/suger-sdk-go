# \BuyerAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseCreditWallet**](BuyerAPI.md#CloseCreditWallet) | **Patch** /org/{orgId}/buyer/{buyerId}/wallet/{walletId}/close | close credit wallet
[**CreateBuyer**](BuyerAPI.md#CreateBuyer) | **Post** /org/{orgId}/buyer | create buyer
[**CreateCreditWallet**](BuyerAPI.md#CreateCreditWallet) | **Post** /org/{orgId}/buyer/{buyerId}/wallet | create credit wallet
[**DeleteBuyerWallet**](BuyerAPI.md#DeleteBuyerWallet) | **Delete** /org/{orgId}/buyer/{buyerId}/wallet/{walletId} | delete buyer wallet
[**GetBuyer**](BuyerAPI.md#GetBuyer) | **Get** /org/{orgId}/buyer/{buyerId} | get buyer
[**ListBuyerWallets**](BuyerAPI.md#ListBuyerWallets) | **Get** /org/{orgId}/buyer/{buyerId}/wallet | list buyer&#39;s wallets
[**ListBuyers**](BuyerAPI.md#ListBuyers) | **Get** /org/{orgId}/buyer | list buyers
[**QueryBuyers**](BuyerAPI.md#QueryBuyers) | **Get** /org/{orgId}/buyer/query | query buyers
[**SetBuyerDefaultWallet**](BuyerAPI.md#SetBuyerDefaultWallet) | **Patch** /org/{orgId}/buyer/{buyerId}/wallet/{walletId}/default | set buyer default wallet
[**UpdateBuyer**](BuyerAPI.md#UpdateBuyer) | **Patch** /org/{orgId}/buyer/{buyerId} | update buyer
[**UpdateCreditWallet**](BuyerAPI.md#UpdateCreditWallet) | **Patch** /org/{orgId}/buyer/{buyerId}/wallet/{walletId} | update credit wallet



## CloseCreditWallet

> BillingWallet CloseCreditWallet(ctx, orgId, buyerId, walletId).Execute()

close credit wallet



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
	buyerId := "buyerId_example" // string | Buyer ID
	walletId := "walletId_example" // string | Wallet ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerAPI.CloseCreditWallet(context.Background(), orgId, buyerId, walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.CloseCreditWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloseCreditWallet`: BillingWallet
	fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.CloseCreditWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**buyerId** | **string** | Buyer ID | 
**walletId** | **string** | Wallet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseCreditWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BillingWallet**](BillingWallet.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBuyer

> IdentityBuyer CreateBuyer(ctx, orgId).Data(data).Execute()

create buyer



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
	data := *openapiclient.NewCreateBuyerParams() // CreateBuyerParams | CreateBuyerParams

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerAPI.CreateBuyer(context.Background(), orgId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.CreateBuyer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBuyer`: IdentityBuyer
	fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.CreateBuyer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**CreateBuyerParams**](CreateBuyerParams.md) | CreateBuyerParams | 

### Return type

[**IdentityBuyer**](IdentityBuyer.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCreditWallet

> BillingWallet CreateCreditWallet(ctx, orgId, buyerId).Execute()

create credit wallet



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
	buyerId := "buyerId_example" // string | Buyer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerAPI.CreateCreditWallet(context.Background(), orgId, buyerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.CreateCreditWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCreditWallet`: BillingWallet
	fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.CreateCreditWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**buyerId** | **string** | Buyer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreditWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BillingWallet**](BillingWallet.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBuyerWallet

> string DeleteBuyerWallet(ctx, orgId, buyerId, walletId).Execute()

delete buyer wallet



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
	buyerId := "buyerId_example" // string | Buyer ID
	walletId := "walletId_example" // string | Wallet ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerAPI.DeleteBuyerWallet(context.Background(), orgId, buyerId, walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.DeleteBuyerWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBuyerWallet`: string
	fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.DeleteBuyerWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**buyerId** | **string** | Buyer ID | 
**walletId** | **string** | Wallet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerWalletRequest struct via the builder pattern


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
	openapiclient "github.com/sugerio/suger-sdk-go"
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

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuyerWallets

> []BillingWallet ListBuyerWallets(ctx, orgId, buyerId).Execute()

list buyer's wallets



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
	buyerId := "buyerId_example" // string | Buyer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerAPI.ListBuyerWallets(context.Background(), orgId, buyerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.ListBuyerWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBuyerWallets`: []BillingWallet
	fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.ListBuyerWallets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**buyerId** | **string** | Buyer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerWalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]BillingWallet**](BillingWallet.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuyers

> []IdentityBuyer ListBuyers(ctx, orgId).Partner(partner).EmailDomain(emailDomain).AwsAccountId(awsAccountId).Limit(limit).Offset(offset).Execute()

list buyers



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
	emailDomain := "emailDomain_example" // string | filter by email domain (optional)
	awsAccountId := "awsAccountId_example" // string | filter by awsAccountId (optional)
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerAPI.ListBuyers(context.Background(), orgId).Partner(partner).EmailDomain(emailDomain).AwsAccountId(awsAccountId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.ListBuyers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBuyers`: []IdentityBuyer
	fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.ListBuyers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partner** | **string** | filter by partner | 
 **emailDomain** | **string** | filter by email domain | 
 **awsAccountId** | **string** | filter by awsAccountId | 
 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**[]IdentityBuyer**](IdentityBuyer.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryBuyers

> GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmIdentityBuyer QueryBuyers(ctx, orgId).PageSize(pageSize).PageNumber(pageNumber).Q(q).S(s).Execute()

query buyers



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
	q := "q_example" // string | LISP-style filter expression (e.g., '(= partner \\ (optional)
	s := "s_example" // string | Sort fields: 'field:asc,field2:desc' or '-field,field2' format (e.g., 'creation_time:desc,name:asc' or '-creation_time,name') (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerAPI.QueryBuyers(context.Background(), orgId).PageSize(pageSize).PageNumber(pageNumber).Q(q).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.QueryBuyers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryBuyers`: GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmIdentityBuyer
	fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.QueryBuyers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryBuyersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | Number of items per page (default 20, max 1000) | 
 **pageNumber** | **int32** | Page number (default 1) | 
 **q** | **string** | LISP-style filter expression (e.g., &#39;(&#x3D; partner \\ | 
 **s** | **string** | Sort fields: &#39;field:asc,field2:desc&#39; or &#39;-field,field2&#39; format (e.g., &#39;creation_time:desc,name:asc&#39; or &#39;-creation_time,name&#39;) | 

### Return type

[**GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmIdentityBuyer**](GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmIdentityBuyer.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBuyerDefaultWallet

> IdentityBuyer SetBuyerDefaultWallet(ctx, orgId, buyerId, walletId).Execute()

set buyer default wallet



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
	buyerId := "buyerId_example" // string | Buyer ID
	walletId := "walletId_example" // string | Wallet ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerAPI.SetBuyerDefaultWallet(context.Background(), orgId, buyerId, walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.SetBuyerDefaultWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetBuyerDefaultWallet`: IdentityBuyer
	fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.SetBuyerDefaultWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**buyerId** | **string** | Buyer ID | 
**walletId** | **string** | Wallet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetBuyerDefaultWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IdentityBuyer**](IdentityBuyer.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

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
	openapiclient "github.com/sugerio/suger-sdk-go"
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

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCreditWallet

> BillingWallet UpdateCreditWallet(ctx, orgId, buyerId, walletId).Execute()

update credit wallet



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
	buyerId := "buyerId_example" // string | Buyer ID
	walletId := "walletId_example" // string | Wallet ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerAPI.UpdateCreditWallet(context.Background(), orgId, buyerId, walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerAPI.UpdateCreditWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCreditWallet`: BillingWallet
	fmt.Fprintf(os.Stdout, "Response from `BuyerAPI.UpdateCreditWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**buyerId** | **string** | Buyer ID | 
**walletId** | **string** | Wallet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCreditWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BillingWallet**](BillingWallet.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

