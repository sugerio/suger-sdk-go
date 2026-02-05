# \CRMEnrichmentAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgOrgIdCrmEnrichmentOnDemandPost**](CRMEnrichmentAPI.md#OrgOrgIdCrmEnrichmentOnDemandPost) | **Post** /org/{orgId}/crm/enrichment/on-demand | Trigger on-demand CRM record enrichment
[**OrgOrgIdIntegrationPartnerServiceEnrichmentProgressGet**](CRMEnrichmentAPI.md#OrgOrgIdIntegrationPartnerServiceEnrichmentProgressGet) | **Get** /org/{orgId}/integration/{partner}/{service}/enrichment/progress | Get CRM enrichment progress
[**OrgOrgIdIntegrationPartnerServiceEnrichmentValidateQueryPost**](CRMEnrichmentAPI.md#OrgOrgIdIntegrationPartnerServiceEnrichmentValidateQueryPost) | **Post** /org/{orgId}/integration/{partner}/{service}/enrichment/validate-query | Validate CRM query



## OrgOrgIdCrmEnrichmentOnDemandPost

> TriggerOnDemandEnrichmentResponse OrgOrgIdCrmEnrichmentOnDemandPost(ctx, orgId).CompanyDomain(companyDomain).CrmRecordId(crmRecordId).CrmRecordType(crmRecordType).CompanyName(companyName).Execute()

Trigger on-demand CRM record enrichment



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
	companyDomain := "companyDomain_example" // string | Company domain (e.g., acme.com)
	crmRecordId := "crmRecordId_example" // string | CRM record ID
	crmRecordType := "crmRecordType_example" // string | CRM record type (e.g., Opportunity, deals)
	companyName := "companyName_example" // string | Company name (fallback if domain lookup fails) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CRMEnrichmentAPI.OrgOrgIdCrmEnrichmentOnDemandPost(context.Background(), orgId).CompanyDomain(companyDomain).CrmRecordId(crmRecordId).CrmRecordType(crmRecordType).CompanyName(companyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CRMEnrichmentAPI.OrgOrgIdCrmEnrichmentOnDemandPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrgOrgIdCrmEnrichmentOnDemandPost`: TriggerOnDemandEnrichmentResponse
	fmt.Fprintf(os.Stdout, "Response from `CRMEnrichmentAPI.OrgOrgIdCrmEnrichmentOnDemandPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgOrgIdCrmEnrichmentOnDemandPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyDomain** | **string** | Company domain (e.g., acme.com) | 
 **crmRecordId** | **string** | CRM record ID | 
 **crmRecordType** | **string** | CRM record type (e.g., Opportunity, deals) | 
 **companyName** | **string** | Company name (fallback if domain lookup fails) | 

### Return type

[**TriggerOnDemandEnrichmentResponse**](TriggerOnDemandEnrichmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgOrgIdIntegrationPartnerServiceEnrichmentProgressGet

> PkgHandlerGetEnrichmentProgressResponse OrgOrgIdIntegrationPartnerServiceEnrichmentProgressGet(ctx, orgId, partner, service).RecordType(recordType).Execute()

Get CRM enrichment progress



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
	partner := "partner_example" // string | Partner (SALESFORCE or HUBSPOT)
	service := "service_example" // string | Service (CRM)
	recordType := "recordType_example" // string | Record type (e.g., Opportunity, Account, deals, companies)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CRMEnrichmentAPI.OrgOrgIdIntegrationPartnerServiceEnrichmentProgressGet(context.Background(), orgId, partner, service).RecordType(recordType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CRMEnrichmentAPI.OrgOrgIdIntegrationPartnerServiceEnrichmentProgressGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrgOrgIdIntegrationPartnerServiceEnrichmentProgressGet`: PkgHandlerGetEnrichmentProgressResponse
	fmt.Fprintf(os.Stdout, "Response from `CRMEnrichmentAPI.OrgOrgIdIntegrationPartnerServiceEnrichmentProgressGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**partner** | **string** | Partner (SALESFORCE or HUBSPOT) | 
**service** | **string** | Service (CRM) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgOrgIdIntegrationPartnerServiceEnrichmentProgressGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **recordType** | **string** | Record type (e.g., Opportunity, Account, deals, companies) | 

### Return type

[**PkgHandlerGetEnrichmentProgressResponse**](PkgHandlerGetEnrichmentProgressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgOrgIdIntegrationPartnerServiceEnrichmentValidateQueryPost

> PkgHandlerValidateQueryResponse OrgOrgIdIntegrationPartnerServiceEnrichmentValidateQueryPost(ctx, orgId, partner, service).Request(request).Execute()

Validate CRM query



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
	partner := "partner_example" // string | Partner (SALESFORCE or HUBSPOT)
	service := "service_example" // string | Service (CRM)
	request := *openapiclient.NewPkgHandlerValidateQueryRequest("RecordQuery_example", "RecordType_example") // PkgHandlerValidateQueryRequest | Query validation request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CRMEnrichmentAPI.OrgOrgIdIntegrationPartnerServiceEnrichmentValidateQueryPost(context.Background(), orgId, partner, service).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CRMEnrichmentAPI.OrgOrgIdIntegrationPartnerServiceEnrichmentValidateQueryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrgOrgIdIntegrationPartnerServiceEnrichmentValidateQueryPost`: PkgHandlerValidateQueryResponse
	fmt.Fprintf(os.Stdout, "Response from `CRMEnrichmentAPI.OrgOrgIdIntegrationPartnerServiceEnrichmentValidateQueryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**partner** | **string** | Partner (SALESFORCE or HUBSPOT) | 
**service** | **string** | Service (CRM) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgOrgIdIntegrationPartnerServiceEnrichmentValidateQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**PkgHandlerValidateQueryRequest**](PkgHandlerValidateQueryRequest.md) | Query validation request | 

### Return type

[**PkgHandlerValidateQueryResponse**](PkgHandlerValidateQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

