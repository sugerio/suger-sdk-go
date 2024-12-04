# \MeteringAPI

All URIs are relative to *http://https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchReportUsageRecordGroups**](MeteringAPI.md#BatchReportUsageRecordGroups) | **Post** /org/{orgId}/batchCreateUsageRecordGroups | batch report usageRecordGroups
[**BatchValidateUsageRecordGroups**](MeteringAPI.md#BatchValidateUsageRecordGroups) | **Post** /org/{orgId}/batchValidateUsageRecordGroups | batch validate usageRecordGroups
[**CreateBillableMetric**](MeteringAPI.md#CreateBillableMetric) | **Post** /org/{orgId}/billableMetric | create billable metric
[**DeleteUsageRecordGroup**](MeteringAPI.md#DeleteUsageRecordGroup) | **Delete** /org/{orgId}/usageRecordGroup/{usageRecordGroupId} | delete usageRecordGroup
[**GetBillableMetric**](MeteringAPI.md#GetBillableMetric) | **Get** /org/{orgId}/billableMetric/{billableMetricId} | get billable metric
[**GetUsageMeteringConfigInfo**](MeteringAPI.md#GetUsageMeteringConfigInfo) | **Get** /org/{orgId}/usageMeteringConfigInfo | get usage metering config info
[**ListBillableMetrics**](MeteringAPI.md#ListBillableMetrics) | **Get** /org/{orgId}/billableMetric | list billable metrics
[**ListUsageRecordGroups**](MeteringAPI.md#ListUsageRecordGroups) | **Get** /org/{orgId}/usageRecordGroup | list usageRecordGroups
[**ListUsageRecordReports**](MeteringAPI.md#ListUsageRecordReports) | **Get** /org/{orgId}/usageRecordReport | list usageRecordReports
[**ReportUsageRecordGroup**](MeteringAPI.md#ReportUsageRecordGroup) | **Post** /org/{orgId}/entitlement/{entitlementId}/usageRecordGroup | report usageRecordGroup
[**RetryUsageRecordGroup**](MeteringAPI.md#RetryUsageRecordGroup) | **Post** /org/{orgId}/usageRecordGroup/{usageRecordGroupId}/retry | retry usageRecordGroup
[**UpdateBillableMetric**](MeteringAPI.md#UpdateBillableMetric) | **Patch** /org/{orgId}/billableMetric/{billableMetricId} | update billable metric
[**UpdateUsageMeteringConfigInfo**](MeteringAPI.md#UpdateUsageMeteringConfigInfo) | **Patch** /org/{orgId}/usageMeteringConfigInfo | update usage metering config info



## BatchReportUsageRecordGroups

> []MeteringUsageRecordGroup BatchReportUsageRecordGroups(ctx, orgId).UsageRecordGroups(usageRecordGroups).Execute()

batch report usageRecordGroups



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
	usageRecordGroups := []openapiclient.NewUsageRecordGroup{*openapiclient.NewNewUsageRecordGroup("EntitlementID_example", map[string]float32{"key": float32(123)})} // []NewUsageRecordGroup | Array of new usage record groups to report

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeteringAPI.BatchReportUsageRecordGroups(context.Background(), orgId).UsageRecordGroups(usageRecordGroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.BatchReportUsageRecordGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchReportUsageRecordGroups`: []MeteringUsageRecordGroup
	fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.BatchReportUsageRecordGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchReportUsageRecordGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usageRecordGroups** | [**[]NewUsageRecordGroup**](NewUsageRecordGroup.md) | Array of new usage record groups to report | 

### Return type

[**[]MeteringUsageRecordGroup**](MeteringUsageRecordGroup.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchValidateUsageRecordGroups

> string BatchValidateUsageRecordGroups(ctx, orgId).Data(data).Execute()

batch validate usageRecordGroups



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
	data := []openapiclient.NewUsageRecordGroup{*openapiclient.NewNewUsageRecordGroup("EntitlementID_example", map[string]float32{"key": float32(123)})} // []NewUsageRecordGroup | Array of usage record groups to be validated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeteringAPI.BatchValidateUsageRecordGroups(context.Background(), orgId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.BatchValidateUsageRecordGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchValidateUsageRecordGroups`: string
	fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.BatchValidateUsageRecordGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchValidateUsageRecordGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**[]NewUsageRecordGroup**](NewUsageRecordGroup.md) | Array of usage record groups to be validated | 

### Return type

**string**

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBillableMetric

> BillableMetric CreateBillableMetric(ctx, orgId).Data(data).Execute()

create billable metric



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
	data := *openapiclient.NewBillableMetric() // BillableMetric | RequestBody

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeteringAPI.CreateBillableMetric(context.Background(), orgId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.CreateBillableMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBillableMetric`: BillableMetric
	fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.CreateBillableMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillableMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**BillableMetric**](BillableMetric.md) | RequestBody | 

### Return type

[**BillableMetric**](BillableMetric.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsageRecordGroup

> MeteringUsageRecordGroup DeleteUsageRecordGroup(ctx, orgId, usageRecordGroupId).Execute()

delete usageRecordGroup



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
	usageRecordGroupId := "usageRecordGroupId_example" // string | UsageRecordGroup ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeteringAPI.DeleteUsageRecordGroup(context.Background(), orgId, usageRecordGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.DeleteUsageRecordGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUsageRecordGroup`: MeteringUsageRecordGroup
	fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.DeleteUsageRecordGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**usageRecordGroupId** | **string** | UsageRecordGroup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsageRecordGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MeteringUsageRecordGroup**](MeteringUsageRecordGroup.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillableMetric

> BillableMetric GetBillableMetric(ctx, orgId, billableMetricId).Execute()

get billable metric



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
	billableMetricId := "billableMetricId_example" // string | Billable Metric ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeteringAPI.GetBillableMetric(context.Background(), orgId, billableMetricId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.GetBillableMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillableMetric`: BillableMetric
	fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.GetBillableMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**billableMetricId** | **string** | Billable Metric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillableMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BillableMetric**](BillableMetric.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageMeteringConfigInfo

> UsageMeteringConfigInfo GetUsageMeteringConfigInfo(ctx, orgId).Execute()

get usage metering config info



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
	resp, r, err := apiClient.MeteringAPI.GetUsageMeteringConfigInfo(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.GetUsageMeteringConfigInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageMeteringConfigInfo`: UsageMeteringConfigInfo
	fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.GetUsageMeteringConfigInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageMeteringConfigInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsageMeteringConfigInfo**](UsageMeteringConfigInfo.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBillableMetrics

> []BillableMetric ListBillableMetrics(ctx, orgId).Status(status).Execute()

list billable metrics



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
	status := "status_example" // string | Status of the billable metric (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeteringAPI.ListBillableMetrics(context.Background(), orgId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.ListBillableMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBillableMetrics`: []BillableMetric
	fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.ListBillableMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBillableMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Status of the billable metric | 

### Return type

[**[]BillableMetric**](BillableMetric.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordGroups

> ListUsageRecordGroupsResponse ListUsageRecordGroups(ctx, orgId).Partner(partner).BuyerId(buyerId).EntitlementId(entitlementId).Status(status).Source(source).MetaInfo(metaInfo).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()

list usageRecordGroups



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
	buyerId := "buyerId_example" // string | filter by buyer ID, default no filter by buyerId if not provided (optional)
	entitlementId := "entitlementId_example" // string | filter by entitlement ID, default no filter by entitlementId if not provided (optional)
	status := "status_example" // string | The status of the usage record group, default no filter by status if not provided (optional)
	source := "source_example" // string | The source of the usage record group, default no filter by source if not provided (optional)
	metaInfo := "metaInfo_example" // string | metaInfo filter (optional)
	startDate := "startDate_example" // string | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate (optional)
	endDate := "endDate_example" // string | end date (UTC) in YYYY-MM-DD format, default is today (optional)
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeteringAPI.ListUsageRecordGroups(context.Background(), orgId).Partner(partner).BuyerId(buyerId).EntitlementId(entitlementId).Status(status).Source(source).MetaInfo(metaInfo).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.ListUsageRecordGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsageRecordGroups`: ListUsageRecordGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.ListUsageRecordGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRecordGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partner** | **string** | Cloud Partner | 
 **buyerId** | **string** | filter by buyer ID, default no filter by buyerId if not provided | 
 **entitlementId** | **string** | filter by entitlement ID, default no filter by entitlementId if not provided | 
 **status** | **string** | The status of the usage record group, default no filter by status if not provided | 
 **source** | **string** | The source of the usage record group, default no filter by source if not provided | 
 **metaInfo** | **string** | metaInfo filter | 
 **startDate** | **string** | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate | 
 **endDate** | **string** | end date (UTC) in YYYY-MM-DD format, default is today | 
 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**ListUsageRecordGroupsResponse**](ListUsageRecordGroupsResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordReports

> ListUsageRecordReportsResponse ListUsageRecordReports(ctx, orgId).Partner(partner).BuyerId(buyerId).EntitlementId(entitlementId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()

list usageRecordReports



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
	buyerId := "buyerId_example" // string | buyer ID (optional)
	entitlementId := "entitlementId_example" // string | entitlement ID (optional)
	startDate := "startDate_example" // string | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate (optional)
	endDate := "endDate_example" // string | end date (UTC) in YYYY-MM-DD format, default is today (optional)
	limit := int32(56) // int32 | List pagination size, default 1000, max value is 1000 (optional)
	offset := int32(56) // int32 | List pagination offset, default 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeteringAPI.ListUsageRecordReports(context.Background(), orgId).Partner(partner).BuyerId(buyerId).EntitlementId(entitlementId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.ListUsageRecordReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsageRecordReports`: ListUsageRecordReportsResponse
	fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.ListUsageRecordReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRecordReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partner** | **string** | Cloud Partner | 
 **buyerId** | **string** | buyer ID | 
 **entitlementId** | **string** | entitlement ID | 
 **startDate** | **string** | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate | 
 **endDate** | **string** | end date (UTC) in YYYY-MM-DD format, default is today | 
 **limit** | **int32** | List pagination size, default 1000, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**ListUsageRecordReportsResponse**](ListUsageRecordReportsResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportUsageRecordGroup

> MeteringUsageRecordGroup ReportUsageRecordGroup(ctx, orgId, entitlementId).Data(data).Execute()

report usageRecordGroup



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
	entitlementId := "entitlementId_example" // string | Entitlement ID
	data := *openapiclient.NewCreateUsageRecordGroupParams("EntitlementID_example", "OrganizationID_example", map[string]float32{"key": float32(123)}) // CreateUsageRecordGroupParams | RequestBody

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeteringAPI.ReportUsageRecordGroup(context.Background(), orgId, entitlementId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.ReportUsageRecordGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportUsageRecordGroup`: MeteringUsageRecordGroup
	fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.ReportUsageRecordGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportUsageRecordGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**CreateUsageRecordGroupParams**](CreateUsageRecordGroupParams.md) | RequestBody | 

### Return type

[**MeteringUsageRecordGroup**](MeteringUsageRecordGroup.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryUsageRecordGroup

> MeteringUsageRecordGroup RetryUsageRecordGroup(ctx, orgId, usageRecordGroupId).Execute()

retry usageRecordGroup



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
	usageRecordGroupId := "usageRecordGroupId_example" // string | UsageRecordGroup ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeteringAPI.RetryUsageRecordGroup(context.Background(), orgId, usageRecordGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.RetryUsageRecordGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryUsageRecordGroup`: MeteringUsageRecordGroup
	fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.RetryUsageRecordGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**usageRecordGroupId** | **string** | UsageRecordGroup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryUsageRecordGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MeteringUsageRecordGroup**](MeteringUsageRecordGroup.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBillableMetric

> BillableMetric UpdateBillableMetric(ctx, orgId, billableMetricId).Data(data).Execute()

update billable metric



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
	billableMetricId := "billableMetricId_example" // string | Billable Metric ID
	data := *openapiclient.NewUpdateBillableMetricParams() // UpdateBillableMetricParams | RequestBody

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeteringAPI.UpdateBillableMetric(context.Background(), orgId, billableMetricId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.UpdateBillableMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBillableMetric`: BillableMetric
	fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.UpdateBillableMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**billableMetricId** | **string** | Billable Metric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBillableMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**UpdateBillableMetricParams**](UpdateBillableMetricParams.md) | RequestBody | 

### Return type

[**BillableMetric**](BillableMetric.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsageMeteringConfigInfo

> UsageMeteringConfigInfo UpdateUsageMeteringConfigInfo(ctx, orgId).Data(data).Execute()

update usage metering config info



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
	data := *openapiclient.NewUsageMeteringConfigInfo() // UsageMeteringConfigInfo | The usage metering config info to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeteringAPI.UpdateUsageMeteringConfigInfo(context.Background(), orgId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.UpdateUsageMeteringConfigInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUsageMeteringConfigInfo`: UsageMeteringConfigInfo
	fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.UpdateUsageMeteringConfigInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUsageMeteringConfigInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**UsageMeteringConfigInfo**](UsageMeteringConfigInfo.md) | The usage metering config info to be updated | 

### Return type

[**UsageMeteringConfigInfo**](UsageMeteringConfigInfo.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

