# \ContactAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContactToBuyer**](ContactAPI.md#AddContactToBuyer) | **Post** /org/{orgId}/contact/{contactId}/buyer/{buyerId} | add contact to buyer
[**AddContactToOffer**](ContactAPI.md#AddContactToOffer) | **Post** /org/{orgId}/contact/{contactId}/offer/{offerId} | add contact to offer
[**BatchCreateContacts**](ContactAPI.md#BatchCreateContacts) | **Post** /org/{orgId}/contact/batch | batch create contacts
[**CreateContact**](ContactAPI.md#CreateContact) | **Post** /org/{orgId}/contact | create contact
[**GetContact**](ContactAPI.md#GetContact) | **Get** /org/{orgId}/contact/{contactId} | get contact
[**ListContacts**](ContactAPI.md#ListContacts) | **Get** /org/{orgId}/contact | list contacts
[**QueryContacts**](ContactAPI.md#QueryContacts) | **Get** /org/{orgId}/contact/query | query contacts
[**RemoveContactFromBuyer**](ContactAPI.md#RemoveContactFromBuyer) | **Delete** /org/{orgId}/contact/{contactId}/buyer/{buyerId} | remove contact from buyer
[**RemoveContactFromOffer**](ContactAPI.md#RemoveContactFromOffer) | **Delete** /org/{orgId}/contact/{contactId}/offer/{offerId} | remove contact from offer
[**UpdateContact**](ContactAPI.md#UpdateContact) | **Patch** /org/{orgId}/contact/{contactId} | update contact
[**UpdateContactTags**](ContactAPI.md#UpdateContactTags) | **Patch** /org/{orgId}/contact/{contactId}/tag | update contact tags



## AddContactToBuyer

> IdentityBuyer AddContactToBuyer(ctx, orgId, buyerId, contactId).Execute()

add contact to buyer



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
	contactId := "contactId_example" // string | Contact ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.AddContactToBuyer(context.Background(), orgId, buyerId, contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.AddContactToBuyer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddContactToBuyer`: IdentityBuyer
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.AddContactToBuyer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**buyerId** | **string** | Buyer ID | 
**contactId** | **string** | Contact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddContactToBuyerRequest struct via the builder pattern


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


## AddContactToOffer

> string AddContactToOffer(ctx, orgId, contactId, offerId).Execute()

add contact to offer



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
	contactId := "contactId_example" // string | Contact ID
	offerId := "offerId_example" // string | Offer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.AddContactToOffer(context.Background(), orgId, contactId, offerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.AddContactToOffer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddContactToOffer`: string
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.AddContactToOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**contactId** | **string** | Contact ID | 
**offerId** | **string** | Offer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddContactToOfferRequest struct via the builder pattern


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


## BatchCreateContacts

> [][]IdentityContact BatchCreateContacts(ctx, orgId).Data(data).Execute()

batch create contacts



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
	data := []openapiclient.IdentityContact{*openapiclient.NewIdentityContact()} // []IdentityContact | RequestBody

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.BatchCreateContacts(context.Background(), orgId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.BatchCreateContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchCreateContacts`: [][]IdentityContact
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.BatchCreateContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchCreateContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**[]IdentityContact**](IdentityContact.md) | RequestBody | 

### Return type

[**[][]IdentityContact**](array.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContact

> IdentityContact CreateContact(ctx, orgId).Data(data).Execute()

create contact



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
	data := *openapiclient.NewIdentityContact() // IdentityContact | RequestBody

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.CreateContact(context.Background(), orgId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.CreateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContact`: IdentityContact
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.CreateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**IdentityContact**](IdentityContact.md) | RequestBody | 

### Return type

[**IdentityContact**](IdentityContact.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContact

> IdentityContact GetContact(ctx, orgId, contactId).Execute()

get contact



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
	contactId := "contactId_example" // string | Contact ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.GetContact(context.Background(), orgId, contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.GetContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContact`: IdentityContact
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.GetContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**contactId** | **string** | Contact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdentityContact**](IdentityContact.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContacts

> []IdentityContact ListContacts(ctx, orgId).ContactIds(contactIds).EmailDomain(emailDomain).Keyword(keyword).Limit(limit).Offset(offset).Execute()

list contacts



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
	contactIds := "contactIds_example" // string | Comma-separated list of contact IDs to fetch specific contacts, up to 100 IDs (optional)
	emailDomain := "emailDomain_example" // string | Email domain to filter contacts (optional)
	keyword := "keyword_example" // string | the keyword string to search contacts by name or email (uses ORM for better performance) (optional)
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (ignored when contactIds is provided) (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (ignored when contactIds is provided) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.ListContacts(context.Background(), orgId).ContactIds(contactIds).EmailDomain(emailDomain).Keyword(keyword).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.ListContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContacts`: []IdentityContact
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.ListContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contactIds** | **string** | Comma-separated list of contact IDs to fetch specific contacts, up to 100 IDs | 
 **emailDomain** | **string** | Email domain to filter contacts | 
 **keyword** | **string** | the keyword string to search contacts by name or email (uses ORM for better performance) | 
 **limit** | **int32** | List pagination size, default 1000, max value is 1000 (ignored when contactIds is provided) | 
 **offset** | **int32** | List pagination offset, default 0 (ignored when contactIds is provided) | 

### Return type

[**[]IdentityContact**](IdentityContact.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryContacts

> GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmIdentityContact QueryContacts(ctx, orgId).PageSize(pageSize).PageNumber(pageNumber).Q(q).S(s).Execute()

query contacts



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
	q := "q_example" // string | LISP-style filter expression (e.g., '(= name \\ (optional)
	s := "s_example" // string | Sort fields: 'field:asc,field2:desc' or '-field,field2' format (e.g., 'creation_time:desc,name:asc' or '-creation_time,name') (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.QueryContacts(context.Background(), orgId).PageSize(pageSize).PageNumber(pageNumber).Q(q).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.QueryContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryContacts`: GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmIdentityContact
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.QueryContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | Number of items per page (default 20, max 1000) | 
 **pageNumber** | **int32** | Page number (default 1) | 
 **q** | **string** | LISP-style filter expression (e.g., &#39;(&#x3D; name \\ | 
 **s** | **string** | Sort fields: &#39;field:asc,field2:desc&#39; or &#39;-field,field2&#39; format (e.g., &#39;creation_time:desc,name:asc&#39; or &#39;-creation_time,name&#39;) | 

### Return type

[**GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmIdentityContact**](GithubComSugerioMarketplaceServicePkgCrudListBaseResponseGithubComSugerioMarketplaceServicePkgOrmIdentityContact.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveContactFromBuyer

> string RemoveContactFromBuyer(ctx, orgId, buyerId, contactId).Execute()

remove contact from buyer



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
	contactId := "contactId_example" // string | Contact ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.RemoveContactFromBuyer(context.Background(), orgId, buyerId, contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.RemoveContactFromBuyer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveContactFromBuyer`: string
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.RemoveContactFromBuyer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**buyerId** | **string** | Buyer ID | 
**contactId** | **string** | Contact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveContactFromBuyerRequest struct via the builder pattern


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


## RemoveContactFromOffer

> string RemoveContactFromOffer(ctx, orgId, contactId, offerId).Execute()

remove contact from offer



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
	contactId := "contactId_example" // string | Contact ID
	offerId := "offerId_example" // string | Offer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.RemoveContactFromOffer(context.Background(), orgId, contactId, offerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.RemoveContactFromOffer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveContactFromOffer`: string
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.RemoveContactFromOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**contactId** | **string** | Contact ID | 
**offerId** | **string** | Offer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveContactFromOfferRequest struct via the builder pattern


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


## UpdateContact

> IdentityContact UpdateContact(ctx, orgId, contactId).Data(data).Execute()

update contact



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
	contactId := "contactId_example" // string | Contact ID
	data := *openapiclient.NewIdentityContact() // IdentityContact | Request Body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.UpdateContact(context.Background(), orgId, contactId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.UpdateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContact`: IdentityContact
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.UpdateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**contactId** | **string** | Contact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**IdentityContact**](IdentityContact.md) | Request Body | 

### Return type

[**IdentityContact**](IdentityContact.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContactTags

> []string UpdateContactTags(ctx, orgId, contactId).Data(data).Execute()

update contact tags



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
	contactId := "contactId_example" // string | Contact ID
	data := *openapiclient.NewServiceMarketplaceServiceApiUpdateContactTagsRequest() // ServiceMarketplaceServiceApiUpdateContactTagsRequest | Request Body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.UpdateContactTags(context.Background(), orgId, contactId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.UpdateContactTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContactTags`: []string
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.UpdateContactTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**contactId** | **string** | Contact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**ServiceMarketplaceServiceApiUpdateContactTagsRequest**](ServiceMarketplaceServiceApiUpdateContactTagsRequest.md) | Request Body | 

### Return type

**[]string**

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

