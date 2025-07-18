# \StatisticsAPI

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatisticsCampaignsByNameGet**](StatisticsAPI.md#StatisticsCampaignsByNameGet) | **Get** /statistics/campaigns/{name} | Load Campaign Stats
[**StatisticsCampaignsGet**](StatisticsAPI.md#StatisticsCampaignsGet) | **Get** /statistics/campaigns | Load Campaigns Stats
[**StatisticsChannelsByNameGet**](StatisticsAPI.md#StatisticsChannelsByNameGet) | **Get** /statistics/channels/{name} | Load Channel Stats
[**StatisticsChannelsGet**](StatisticsAPI.md#StatisticsChannelsGet) | **Get** /statistics/channels | Load Channels Stats
[**StatisticsGet**](StatisticsAPI.md#StatisticsGet) | **Get** /statistics | Load Statistics



## StatisticsCampaignsByNameGet

> ChannelLogStatusSummary StatisticsCampaignsByNameGet(ctx, name).Execute()

Load Campaign Stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/radzawee/elasticemail-go"
)

func main() {
	name := "name_example" // string | The name of the campaign to get.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticsAPI.StatisticsCampaignsByNameGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsAPI.StatisticsCampaignsByNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatisticsCampaignsByNameGet`: ChannelLogStatusSummary
	fmt.Fprintf(os.Stdout, "Response from `StatisticsAPI.StatisticsCampaignsByNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the campaign to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatisticsCampaignsByNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChannelLogStatusSummary**](ChannelLogStatusSummary.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticsCampaignsGet

> []ChannelLogStatusSummary StatisticsCampaignsGet(ctx).Limit(limit).Offset(offset).Execute()

Load Campaigns Stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/radzawee/elasticemail-go"
)

func main() {
	limit := int32(100) // int32 | Maximum number of returned items. (optional)
	offset := int32(20) // int32 | How many items should be returned ahead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticsAPI.StatisticsCampaignsGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsAPI.StatisticsCampaignsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatisticsCampaignsGet`: []ChannelLogStatusSummary
	fmt.Fprintf(os.Stdout, "Response from `StatisticsAPI.StatisticsCampaignsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatisticsCampaignsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]ChannelLogStatusSummary**](ChannelLogStatusSummary.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticsChannelsByNameGet

> ChannelLogStatusSummary StatisticsChannelsByNameGet(ctx, name).Execute()

Load Channel Stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/radzawee/elasticemail-go"
)

func main() {
	name := "name_example" // string | The name of the channel to get.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticsAPI.StatisticsChannelsByNameGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsAPI.StatisticsChannelsByNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatisticsChannelsByNameGet`: ChannelLogStatusSummary
	fmt.Fprintf(os.Stdout, "Response from `StatisticsAPI.StatisticsChannelsByNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the channel to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatisticsChannelsByNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChannelLogStatusSummary**](ChannelLogStatusSummary.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticsChannelsGet

> []ChannelLogStatusSummary StatisticsChannelsGet(ctx).Limit(limit).Offset(offset).Execute()

Load Channels Stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/radzawee/elasticemail-go"
)

func main() {
	limit := int32(100) // int32 | Maximum number of returned items. (optional)
	offset := int32(20) // int32 | How many items should be returned ahead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticsAPI.StatisticsChannelsGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsAPI.StatisticsChannelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatisticsChannelsGet`: []ChannelLogStatusSummary
	fmt.Fprintf(os.Stdout, "Response from `StatisticsAPI.StatisticsChannelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatisticsChannelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]ChannelLogStatusSummary**](ChannelLogStatusSummary.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticsGet

> LogStatusSummary StatisticsGet(ctx).From(from).To(to).Execute()

Load Statistics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/radzawee/elasticemail-go"
)

func main() {
	from := time.Now() // time.Time | Starting date for search in YYYY-MM-DDThh:mm:ss format.
	to := time.Now() // time.Time | Ending date for search in YYYY-MM-DDThh:mm:ss format. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticsAPI.StatisticsGet(context.Background()).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsAPI.StatisticsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatisticsGet`: LogStatusSummary
	fmt.Fprintf(os.Stdout, "Response from `StatisticsAPI.StatisticsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatisticsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** | Starting date for search in YYYY-MM-DDThh:mm:ss format. | 
 **to** | **time.Time** | Ending date for search in YYYY-MM-DDThh:mm:ss format. | 

### Return type

[**LogStatusSummary**](LogStatusSummary.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

