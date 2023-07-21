# \CampaignsApi

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CampaignsByNameDelete**](CampaignsApi.md#CampaignsByNameDelete) | **Delete** /campaigns/{name} | Delete Campaign
[**CampaignsByNameGet**](CampaignsApi.md#CampaignsByNameGet) | **Get** /campaigns/{name} | Load Campaign
[**CampaignsByNamePut**](CampaignsApi.md#CampaignsByNamePut) | **Put** /campaigns/{name} | Update Campaign
[**CampaignsGet**](CampaignsApi.md#CampaignsGet) | **Get** /campaigns | Load Campaigns
[**CampaignsPost**](CampaignsApi.md#CampaignsPost) | **Post** /campaigns | Add Campaign



## CampaignsByNameDelete

> CampaignsByNameDelete(ctx, name).Execute()

Delete Campaign



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elasticemail/elasticemail-go"
)

func main() {
    name := "name_example" // string | Name of Campaign to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CampaignsApi.CampaignsByNameDelete(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CampaignsByNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of Campaign to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiCampaignsByNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CampaignsByNameGet

> Campaign CampaignsByNameGet(ctx, name).Execute()

Load Campaign



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elasticemail/elasticemail-go"
)

func main() {
    name := "name_example" // string | Name of Campaign to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.CampaignsByNameGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CampaignsByNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CampaignsByNameGet`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.CampaignsByNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of Campaign to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiCampaignsByNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Campaign**](Campaign.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CampaignsByNamePut

> Campaign CampaignsByNamePut(ctx, name).Campaign(campaign).Execute()

Update Campaign



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elasticemail/elasticemail-go"
)

func main() {
    name := "name_example" // string | Name of Campaign to update
    campaign := *openapiclient.NewCampaign("Name_example", *openapiclient.NewCampaignRecipient()) // Campaign | JSON representation of a campaign

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.CampaignsByNamePut(context.Background(), name).Campaign(campaign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CampaignsByNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CampaignsByNamePut`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.CampaignsByNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of Campaign to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiCampaignsByNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaign** | [**Campaign**](Campaign.md) | JSON representation of a campaign | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CampaignsGet

> []Campaign CampaignsGet(ctx).Search(search).Offset(offset).Limit(limit).Execute()

Load Campaigns



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elasticemail/elasticemail-go"
)

func main() {
    search := "search_example" // string | Text fragment used for searching in Campaign name (using the 'contains' rule) (optional)
    offset := int32(20) // int32 | How many items should be returned ahead. (optional)
    limit := int32(100) // int32 | Maximum number of returned items. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.CampaignsGet(context.Background()).Search(search).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CampaignsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CampaignsGet`: []Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.CampaignsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCampaignsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Text fragment used for searching in Campaign name (using the &#39;contains&#39; rule) | 
 **offset** | **int32** | How many items should be returned ahead. | 
 **limit** | **int32** | Maximum number of returned items. | 

### Return type

[**[]Campaign**](Campaign.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CampaignsPost

> Campaign CampaignsPost(ctx).Campaign(campaign).Execute()

Add Campaign



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elasticemail/elasticemail-go"
)

func main() {
    campaign := *openapiclient.NewCampaign("Name_example", *openapiclient.NewCampaignRecipient()) // Campaign | JSON representation of a campaign

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.CampaignsPost(context.Background()).Campaign(campaign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CampaignsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CampaignsPost`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.CampaignsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCampaignsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaign** | [**Campaign**](Campaign.md) | JSON representation of a campaign | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

