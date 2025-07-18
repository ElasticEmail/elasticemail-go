# \InboundRouteAPI

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InboundrouteByIdDelete**](InboundRouteAPI.md#InboundrouteByIdDelete) | **Delete** /inboundroute/{id} | Delete Route
[**InboundrouteByIdGet**](InboundRouteAPI.md#InboundrouteByIdGet) | **Get** /inboundroute/{id} | Get Route
[**InboundrouteByIdPut**](InboundRouteAPI.md#InboundrouteByIdPut) | **Put** /inboundroute/{id} | Update Route
[**InboundrouteGet**](InboundRouteAPI.md#InboundrouteGet) | **Get** /inboundroute | Get Routes
[**InboundrouteOrderPut**](InboundRouteAPI.md#InboundrouteOrderPut) | **Put** /inboundroute/order | Update Sorting
[**InboundroutePost**](InboundRouteAPI.md#InboundroutePost) | **Post** /inboundroute | Create Route



## InboundrouteByIdDelete

> InboundrouteByIdDelete(ctx, id).Execute()

Delete Route



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InboundRouteAPI.InboundrouteByIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundRouteAPI.InboundrouteByIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInboundrouteByIdDeleteRequest struct via the builder pattern


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


## InboundrouteByIdGet

> InboundRoute InboundrouteByIdGet(ctx, id).Execute()

Get Route



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
	id := "123456" // string | ID number of your attachment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboundRouteAPI.InboundrouteByIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundRouteAPI.InboundrouteByIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InboundrouteByIdGet`: InboundRoute
	fmt.Fprintf(os.Stdout, "Response from `InboundRouteAPI.InboundrouteByIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID number of your attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiInboundrouteByIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InboundRoute**](InboundRoute.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InboundrouteByIdPut

> InboundRoute InboundrouteByIdPut(ctx, id).InboundPayload(inboundPayload).Execute()

Update Route



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
	id := "id_example" // string | 
	inboundPayload := *openapiclient.NewInboundPayload("Filter_example", "Name_example", openapiclient.InboundRouteFilterType("EmailAddress"), openapiclient.InboundRouteActionType("ForwardToEmail")) // InboundPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboundRouteAPI.InboundrouteByIdPut(context.Background(), id).InboundPayload(inboundPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundRouteAPI.InboundrouteByIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InboundrouteByIdPut`: InboundRoute
	fmt.Fprintf(os.Stdout, "Response from `InboundRouteAPI.InboundrouteByIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInboundrouteByIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboundPayload** | [**InboundPayload**](InboundPayload.md) |  | 

### Return type

[**InboundRoute**](InboundRoute.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InboundrouteGet

> []InboundRoute InboundrouteGet(ctx).Execute()

Get Routes



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboundRouteAPI.InboundrouteGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundRouteAPI.InboundrouteGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InboundrouteGet`: []InboundRoute
	fmt.Fprintf(os.Stdout, "Response from `InboundRouteAPI.InboundrouteGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInboundrouteGetRequest struct via the builder pattern


### Return type

[**[]InboundRoute**](InboundRoute.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InboundrouteOrderPut

> []InboundRoute InboundrouteOrderPut(ctx).SortOrderItem(sortOrderItem).Execute()

Update Sorting



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
	sortOrderItem := []openapiclient.SortOrderItem{*openapiclient.NewSortOrderItem("PublicInboundId_example", int32(123))} // []SortOrderItem | Change the ordering of inbound routes for when matching the inbound

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboundRouteAPI.InboundrouteOrderPut(context.Background()).SortOrderItem(sortOrderItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundRouteAPI.InboundrouteOrderPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InboundrouteOrderPut`: []InboundRoute
	fmt.Fprintf(os.Stdout, "Response from `InboundRouteAPI.InboundrouteOrderPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInboundrouteOrderPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortOrderItem** | [**[]SortOrderItem**](SortOrderItem.md) | Change the ordering of inbound routes for when matching the inbound | 

### Return type

[**[]InboundRoute**](InboundRoute.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InboundroutePost

> InboundRoute InboundroutePost(ctx).InboundPayload(inboundPayload).Execute()

Create Route



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
	inboundPayload := *openapiclient.NewInboundPayload("Filter_example", "Name_example", openapiclient.InboundRouteFilterType("EmailAddress"), openapiclient.InboundRouteActionType("ForwardToEmail")) // InboundPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboundRouteAPI.InboundroutePost(context.Background()).InboundPayload(inboundPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundRouteAPI.InboundroutePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InboundroutePost`: InboundRoute
	fmt.Fprintf(os.Stdout, "Response from `InboundRouteAPI.InboundroutePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInboundroutePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboundPayload** | [**InboundPayload**](InboundPayload.md) |  | 

### Return type

[**InboundRoute**](InboundRoute.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

