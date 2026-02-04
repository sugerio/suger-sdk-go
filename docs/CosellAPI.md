# \CosellAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgOrgIdCosellPartnerAWSSearchPartnerPost**](CosellAPI.md#OrgOrgIdCosellPartnerAWSSearchPartnerPost) | **Post** /org/{orgId}/cosell/partner/AWS/searchPartner | Search partner connections



## OrgOrgIdCosellPartnerAWSSearchPartnerPost

> []GithubComSugerioMarketplaceServicePkgStructsPartnerConnectionSearchResult OrgOrgIdCosellPartnerAWSSearchPartnerPost(ctx, orgId).Query(query).Limit(limit).Execute()

Search partner connections



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
	query := "query_example" // string | Search query (partner name)
	limit := int32(56) // int32 | Maximum results (default 20, max 100) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CosellAPI.OrgOrgIdCosellPartnerAWSSearchPartnerPost(context.Background(), orgId).Query(query).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CosellAPI.OrgOrgIdCosellPartnerAWSSearchPartnerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrgOrgIdCosellPartnerAWSSearchPartnerPost`: []GithubComSugerioMarketplaceServicePkgStructsPartnerConnectionSearchResult
	fmt.Fprintf(os.Stdout, "Response from `CosellAPI.OrgOrgIdCosellPartnerAWSSearchPartnerPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgOrgIdCosellPartnerAWSSearchPartnerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | Search query (partner name) | 
 **limit** | **int32** | Maximum results (default 20, max 100) | 

### Return type

[**[]GithubComSugerioMarketplaceServicePkgStructsPartnerConnectionSearchResult**](GithubComSugerioMarketplaceServicePkgStructsPartnerConnectionSearchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

