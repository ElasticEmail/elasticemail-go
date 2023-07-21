# \SegmentsApi

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SegmentsByNameDelete**](SegmentsApi.md#SegmentsByNameDelete) | **Delete** /segments/{name} | Delete Segment
[**SegmentsByNameGet**](SegmentsApi.md#SegmentsByNameGet) | **Get** /segments/{name} | Load Segment
[**SegmentsByNamePut**](SegmentsApi.md#SegmentsByNamePut) | **Put** /segments/{name} | Update Segment
[**SegmentsGet**](SegmentsApi.md#SegmentsGet) | **Get** /segments | Load Segments
[**SegmentsPost**](SegmentsApi.md#SegmentsPost) | **Post** /segments | Add Segment



## SegmentsByNameDelete

> SegmentsByNameDelete(ctx, name).Execute()

Delete Segment



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
    name := "My Segment 1" // string | Name of your segment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SegmentsApi.SegmentsByNameDelete(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.SegmentsByNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of your segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSegmentsByNameDeleteRequest struct via the builder pattern


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


## SegmentsByNameGet

> Segment SegmentsByNameGet(ctx, name).Execute()

Load Segment



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
    name := "name_example" // string | Name of the segment you want to load. Will load all contacts if the 'All Contacts' name has been provided

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsApi.SegmentsByNameGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.SegmentsByNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SegmentsByNameGet`: Segment
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.SegmentsByNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the segment you want to load. Will load all contacts if the &#39;All Contacts&#39; name has been provided | 

### Other Parameters

Other parameters are passed through a pointer to a apiSegmentsByNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Segment**](Segment.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SegmentsByNamePut

> Segment SegmentsByNamePut(ctx, name).SegmentPayload(segmentPayload).Execute()

Update Segment



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
    name := "My Segment 1" // string | Name of your segment.
    segmentPayload := *openapiclient.NewSegmentPayload("Name_example", "Rule_example") // SegmentPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsApi.SegmentsByNamePut(context.Background(), name).SegmentPayload(segmentPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.SegmentsByNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SegmentsByNamePut`: Segment
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.SegmentsByNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of your segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSegmentsByNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **segmentPayload** | [**SegmentPayload**](SegmentPayload.md) |  | 

### Return type

[**Segment**](Segment.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SegmentsGet

> []Segment SegmentsGet(ctx).Limit(limit).Offset(offset).Execute()

Load Segments



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
    limit := int32(100) // int32 | Maximum number of returned items. (optional)
    offset := int32(20) // int32 | How many items should be returned ahead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsApi.SegmentsGet(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.SegmentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SegmentsGet`: []Segment
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.SegmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSegmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]Segment**](Segment.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SegmentsPost

> Segment SegmentsPost(ctx).SegmentPayload(segmentPayload).Execute()

Add Segment



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
    segmentPayload := *openapiclient.NewSegmentPayload("Name_example", "Rule_example") // SegmentPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsApi.SegmentsPost(context.Background()).SegmentPayload(segmentPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.SegmentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SegmentsPost`: Segment
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.SegmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSegmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **segmentPayload** | [**SegmentPayload**](SegmentPayload.md) |  | 

### Return type

[**Segment**](Segment.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

