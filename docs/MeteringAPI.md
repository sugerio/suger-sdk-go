# \MeteringAPI

All URIs are relative to *https://api.suger.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchReportUsageRecordGroups**](MeteringAPI.md#BatchReportUsageRecordGroups) | **Post** /org/{orgId}/batchCreateUsageRecordGroups | batch report usageRecordGroups
[**BatchValidateUsageRecordGroups**](MeteringAPI.md#BatchValidateUsageRecordGroups) | **Post** /org/{orgId}/batchValidateUsageRecordGroups | batch validate usageRecordGroups
[**DeleteUsageRecordGroup**](MeteringAPI.md#DeleteUsageRecordGroup) | **Delete** /org/{orgId}/usageRecordGroup/{usageRecordGroupId} | delete usageRecordGroup
[**GetUsageMeteringConfigInfo**](MeteringAPI.md#GetUsageMeteringConfigInfo) | **Get** /org/{orgId}/usageMeteringConfigInfo | get usage metering config info
[**GetUsageRecordGroup**](MeteringAPI.md#GetUsageRecordGroup) | **Get** /org/{orgId}/usageRecordGroup/{usageRecordGroupId} | get usageRecordGroup
[**GetUsageRecordReport**](MeteringAPI.md#GetUsageRecordReport) | **Get** /org/{orgId}/usageRecordReport/{usageRecordReportId} | get usageRecordReport
[**ListUsageRecordGroups**](MeteringAPI.md#ListUsageRecordGroups) | **Get** /org/{orgId}/usageRecordGroup | list usageRecordGroups
[**ListUsageRecordGroupsByEntitlement**](MeteringAPI.md#ListUsageRecordGroupsByEntitlement) | **Get** /org/{orgId}/entitlement/{entitlementId}/usageRecordGroup | list usageRecordGroups by entitlement
[**ListUsageRecordGroupsByProduct**](MeteringAPI.md#ListUsageRecordGroupsByProduct) | **Get** /org/{orgId}/product/{productId}/usageRecordGroup | list usageRecordGroups by product
[**ListUsageRecordReports**](MeteringAPI.md#ListUsageRecordReports) | **Get** /org/{orgId}/usageRecordReport | list usageRecordReports
[**ReportUsageRecordGroup**](MeteringAPI.md#ReportUsageRecordGroup) | **Post** /org/{orgId}/entitlement/{entitlementId}/usageRecordGroup | report usageRecordGroup
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
    openapiclient "github.com/sugerio/suger-sdk-go"
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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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
    openapiclient "github.com/sugerio/suger-sdk-go"
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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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
    openapiclient "github.com/sugerio/suger-sdk-go"
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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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
    openapiclient "github.com/sugerio/suger-sdk-go"
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

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageRecordGroup

> MeteringUsageRecordGroup GetUsageRecordGroup(ctx, orgId, usageRecordGroupId).Execute()

get usageRecordGroup



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
    usageRecordGroupId := "usageRecordGroupId_example" // string | UsageRecordGroup ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MeteringAPI.GetUsageRecordGroup(context.Background(), orgId, usageRecordGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.GetUsageRecordGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageRecordGroup`: MeteringUsageRecordGroup
    fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.GetUsageRecordGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**usageRecordGroupId** | **string** | UsageRecordGroup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageRecordGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MeteringUsageRecordGroup**](MeteringUsageRecordGroup.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageRecordReport

> MeteringUsageRecordReport GetUsageRecordReport(ctx, orgId, usageRecordReportId).Execute()

get usageRecordReport



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
    usageRecordReportId := "usageRecordReportId_example" // string | UsageRecordReport ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MeteringAPI.GetUsageRecordReport(context.Background(), orgId, usageRecordReportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.GetUsageRecordReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageRecordReport`: MeteringUsageRecordReport
    fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.GetUsageRecordReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**usageRecordReportId** | **string** | UsageRecordReport ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageRecordReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MeteringUsageRecordReport**](MeteringUsageRecordReport.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordGroups

> ListUsageRecordGroupsResponse ListUsageRecordGroups(ctx, orgId).Partner(partner).ProductId(productId).EntitlementId(entitlementId).BuyerId(buyerId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()

list usageRecordGroups



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
    partner := "partner_example" // string | Cloud Partner (optional)
    productId := "productId_example" // string | product ID (optional)
    entitlementId := "entitlementId_example" // string | entitlement ID (optional)
    buyerId := "buyerId_example" // string | buyer ID (optional)
    startDate := "startDate_example" // string | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate (optional)
    endDate := "endDate_example" // string | end date (UTC) in YYYY-MM-DD format, default is today (optional)
    limit := int32(56) // int32 | List pagination size, default 20, max value is 1000 (optional)
    offset := int32(56) // int32 | List pagination offset, default 0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MeteringAPI.ListUsageRecordGroups(context.Background(), orgId).Partner(partner).ProductId(productId).EntitlementId(entitlementId).BuyerId(buyerId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()
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
 **productId** | **string** | product ID | 
 **entitlementId** | **string** | entitlement ID | 
 **buyerId** | **string** | buyer ID | 
 **startDate** | **string** | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate | 
 **endDate** | **string** | end date (UTC) in YYYY-MM-DD format, default is today | 
 **limit** | **int32** | List pagination size, default 20, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**ListUsageRecordGroupsResponse**](ListUsageRecordGroupsResponse.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordGroupsByEntitlement

> ListUsageRecordGroupsResponse ListUsageRecordGroupsByEntitlement(ctx, orgId, entitlementId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()

list usageRecordGroups by entitlement



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
    entitlementId := "entitlementId_example" // string | entitlement ID
    startDate := "startDate_example" // string | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate (optional)
    endDate := "endDate_example" // string | end date (UTC) in YYYY-MM-DD format, default is today (optional)
    limit := int32(56) // int32 | List pagination size, default 20, max value is 1000 (optional)
    offset := int32(56) // int32 | List pagination offset, default 0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MeteringAPI.ListUsageRecordGroupsByEntitlement(context.Background(), orgId, entitlementId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.ListUsageRecordGroupsByEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordGroupsByEntitlement`: ListUsageRecordGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.ListUsageRecordGroupsByEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**entitlementId** | **string** | entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRecordGroupsByEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **string** | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate | 
 **endDate** | **string** | end date (UTC) in YYYY-MM-DD format, default is today | 
 **limit** | **int32** | List pagination size, default 20, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**ListUsageRecordGroupsResponse**](ListUsageRecordGroupsResponse.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordGroupsByProduct

> ListUsageRecordGroupsResponse ListUsageRecordGroupsByProduct(ctx, orgId, productId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()

list usageRecordGroups by product



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
    productId := "productId_example" // string | product ID
    startDate := "startDate_example" // string | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate (optional)
    endDate := "endDate_example" // string | end date (UTC) in YYYY-MM-DD format, default is today (optional)
    limit := int32(56) // int32 | List pagination size, default 20, max value is 1000 (optional)
    offset := int32(56) // int32 | List pagination offset, default 0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MeteringAPI.ListUsageRecordGroupsByProduct(context.Background(), orgId, productId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeteringAPI.ListUsageRecordGroupsByProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordGroupsByProduct`: ListUsageRecordGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `MeteringAPI.ListUsageRecordGroupsByProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**productId** | **string** | product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRecordGroupsByProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **string** | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate | 
 **endDate** | **string** | end date (UTC) in YYYY-MM-DD format, default is today | 
 **limit** | **int32** | List pagination size, default 20, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**ListUsageRecordGroupsResponse**](ListUsageRecordGroupsResponse.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordReports

> ListUsageRecordReportsResponse ListUsageRecordReports(ctx, orgId).Partner(partner).ProductId(productId).EntitlementId(entitlementId).BuyerId(buyerId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()

list usageRecordReports



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
    partner := "partner_example" // string | Cloud Partner (optional)
    productId := "productId_example" // string | product ID (optional)
    entitlementId := "entitlementId_example" // string | entitlement ID (optional)
    buyerId := "buyerId_example" // string | buyer ID (optional)
    startDate := "startDate_example" // string | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate (optional)
    endDate := "endDate_example" // string | end date (UTC) in YYYY-MM-DD format, default is today (optional)
    limit := int32(56) // int32 | List pagination size, default 20, max value is 1000 (optional)
    offset := int32(56) // int32 | List pagination offset, default 0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MeteringAPI.ListUsageRecordReports(context.Background(), orgId).Partner(partner).ProductId(productId).EntitlementId(entitlementId).BuyerId(buyerId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()
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
 **productId** | **string** | product ID | 
 **entitlementId** | **string** | entitlement ID | 
 **buyerId** | **string** | buyer ID | 
 **startDate** | **string** | start date (UTC) in YYYY-MM-DD format, default is 30 days before the endDate | 
 **endDate** | **string** | end date (UTC) in YYYY-MM-DD format, default is today | 
 **limit** | **int32** | List pagination size, default 20, max value is 1000 | 
 **offset** | **int32** | List pagination offset, default 0 | 

### Return type

[**ListUsageRecordReportsResponse**](ListUsageRecordReportsResponse.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

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
    openapiclient "github.com/sugerio/suger-sdk-go"
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

[BearerTokenAuth](../README.md#BearerTokenAuth)

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
    openapiclient "github.com/sugerio/suger-sdk-go"
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

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

