# \ReportAPI

All URIs are relative to *http://https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRevenueReport**](ReportAPI.md#GetRevenueReport) | **Post** /org/{orgId}/revenueReport | get revenue report
[**ListDailyRevenueRecords**](ReportAPI.md#ListDailyRevenueRecords) | **Get** /org/{orgId}/dailyRevenueRecord | list daily revenue records
[**ListRevenueRecordDetails**](ReportAPI.md#ListRevenueRecordDetails) | **Get** /org/{orgId}/partner/{partner}/revenueRecordDetail | list revenue record details
[**ListRevenueRecords**](ReportAPI.md#ListRevenueRecords) | **Get** /org/{orgId}/partner/{partner}/revenueRecord | list revenue records
[**ListUsageMeteringDailyRecords**](ReportAPI.md#ListUsageMeteringDailyRecords) | **Get** /org/{orgId}/partner/{partner}/usageMeteringDailyRecord | list usage metering daily records



## GetRevenueReport

> RevenueReport GetRevenueReport(ctx, orgId).Data(data).Execute()

get revenue report



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
	data := *openapiclient.NewGetRevenueReportParams("OrganizationID_example", "Partner_example", openapiclient.RevenueReportType("GrossAmount"), "Service_example") // GetRevenueReportParams | Get Revenue Report Params

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportAPI.GetRevenueReport(context.Background(), orgId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportAPI.GetRevenueReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevenueReport`: RevenueReport
	fmt.Fprintf(os.Stdout, "Response from `ReportAPI.GetRevenueReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevenueReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**GetRevenueReportParams**](GetRevenueReportParams.md) | Get Revenue Report Params | 

### Return type

[**RevenueReport**](RevenueReport.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDailyRevenueRecords

> []RevenueRecord ListDailyRevenueRecords(ctx, orgId).Partner(partner).ProductId(productId).BuyerId(buyerId).EntitlementId(entitlementId).StartDate(startDate).EndDate(endDate).Execute()

list daily revenue records



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
	partner := "partner_example" // string | Cloud Partner (optional)
	productId := "productId_example" // string | Filter daily revenue records by the given product ID (optional)
	buyerId := "buyerId_example" // string | Filter daily revenue records by the given buyer ID (optional)
	entitlementId := "entitlementId_example" // string | Filter daily revenue records by the given entitlement ID (optional)
	startDate := "startDate_example" // string | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate (optional)
	endDate := "endDate_example" // string | end date (UTC) in YYYY-MM-DD format, default is today (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportAPI.ListDailyRevenueRecords(context.Background(), orgId).Partner(partner).ProductId(productId).BuyerId(buyerId).EntitlementId(entitlementId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportAPI.ListDailyRevenueRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDailyRevenueRecords`: []RevenueRecord
	fmt.Fprintf(os.Stdout, "Response from `ReportAPI.ListDailyRevenueRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDailyRevenueRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partner** | **string** | Cloud Partner | 
 **productId** | **string** | Filter daily revenue records by the given product ID | 
 **buyerId** | **string** | Filter daily revenue records by the given buyer ID | 
 **entitlementId** | **string** | Filter daily revenue records by the given entitlement ID | 
 **startDate** | **string** | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate | 
 **endDate** | **string** | end date (UTC) in YYYY-MM-DD format, default is today | 

### Return type

[**[]RevenueRecord**](RevenueRecord.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRevenueRecordDetails

> ListRevenueRecordDetailsResponse ListRevenueRecordDetails(ctx, orgId, partner).ProductId(productId).BuyerId(buyerId).EntitlementId(entitlementId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()

list revenue record details



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
	productId := "productId_example" // string | Filter revenue record details by the given product ID (optional)
	buyerId := "buyerId_example" // string | Filter revenue record details by the given buyer ID (optional)
	entitlementId := "entitlementId_example" // string | Filter revenue record details by the given entitlement ID (optional)
	startDate := "startDate_example" // string | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate (optional)
	endDate := "endDate_example" // string | end date (UTC) in YYYY-MM-DD format, default is today (optional)
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportAPI.ListRevenueRecordDetails(context.Background(), orgId, partner).ProductId(productId).BuyerId(buyerId).EntitlementId(entitlementId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportAPI.ListRevenueRecordDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRevenueRecordDetails`: ListRevenueRecordDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportAPI.ListRevenueRecordDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**partner** | **string** | Cloud Partner | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRevenueRecordDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productId** | **string** | Filter revenue record details by the given product ID | 
 **buyerId** | **string** | Filter revenue record details by the given buyer ID | 
 **entitlementId** | **string** | Filter revenue record details by the given entitlement ID | 
 **startDate** | **string** | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate | 
 **endDate** | **string** | end date (UTC) in YYYY-MM-DD format, default is today | 
 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**ListRevenueRecordDetailsResponse**](ListRevenueRecordDetailsResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRevenueRecords

> ListRevenueRecordsResponse ListRevenueRecords(ctx, orgId, partner).ProductId(productId).EntitlementId(entitlementId).BuyerId(buyerId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()

list revenue records



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
	productId := "productId_example" // string | Filter revenue records by the given product ID (optional)
	entitlementId := "entitlementId_example" // string | Filter revenue records by the given entitlement ID (optional)
	buyerId := "buyerId_example" // string | Filter revenue records by the given buyer ID (optional)
	startDate := "startDate_example" // string | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate (optional)
	endDate := "endDate_example" // string | end date (UTC) in YYYY-MM-DD format, default is today (optional)
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportAPI.ListRevenueRecords(context.Background(), orgId, partner).ProductId(productId).EntitlementId(entitlementId).BuyerId(buyerId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportAPI.ListRevenueRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRevenueRecords`: ListRevenueRecordsResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportAPI.ListRevenueRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**partner** | **string** | Cloud Partner | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRevenueRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productId** | **string** | Filter revenue records by the given product ID | 
 **entitlementId** | **string** | Filter revenue records by the given entitlement ID | 
 **buyerId** | **string** | Filter revenue records by the given buyer ID | 
 **startDate** | **string** | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate | 
 **endDate** | **string** | end date (UTC) in YYYY-MM-DD format, default is today | 
 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**ListRevenueRecordsResponse**](ListRevenueRecordsResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageMeteringDailyRecords

> ListUsageMeteringDailyRecordsResponse ListUsageMeteringDailyRecords(ctx, orgId, partner).BuyerId(buyerId).EntitlementId(entitlementId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()

list usage metering daily records



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
	partner := "partner_example" // string | Filter by the Cloud Partner
	buyerId := "buyerId_example" // string | Filter usage metering daily records by the given buyer ID (optional)
	entitlementId := "entitlementId_example" // string | Filter usage metering daily records by the given entitlement ID (optional)
	startDate := "startDate_example" // string | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate (optional)
	endDate := "endDate_example" // string | end date (UTC) in YYYY-MM-DD format, default is today (optional)
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportAPI.ListUsageMeteringDailyRecords(context.Background(), orgId, partner).BuyerId(buyerId).EntitlementId(entitlementId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportAPI.ListUsageMeteringDailyRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsageMeteringDailyRecords`: ListUsageMeteringDailyRecordsResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportAPI.ListUsageMeteringDailyRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**partner** | **string** | Filter by the Cloud Partner | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageMeteringDailyRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **buyerId** | **string** | Filter usage metering daily records by the given buyer ID | 
 **entitlementId** | **string** | Filter usage metering daily records by the given entitlement ID | 
 **startDate** | **string** | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate | 
 **endDate** | **string** | end date (UTC) in YYYY-MM-DD format, default is today | 
 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**ListUsageMeteringDailyRecordsResponse**](ListUsageMeteringDailyRecordsResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

