# \WebhookAPI

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookByPublicidDelete**](WebhookAPI.md#WebhookByPublicidDelete) | **Delete** /webhook/{publicid} | Delete Webhook
[**WebhookByPublicidGet**](WebhookAPI.md#WebhookByPublicidGet) | **Get** /webhook/{publicid} | Load Webhook
[**WebhookByPublicidPut**](WebhookAPI.md#WebhookByPublicidPut) | **Put** /webhook/{publicid} | Update Webhook
[**WebhookGet**](WebhookAPI.md#WebhookGet) | **Get** /webhook | Load Webhooks
[**WebhookPost**](WebhookAPI.md#WebhookPost) | **Post** /webhook | Add Webhook



## WebhookByPublicidDelete

> WebhookByPublicidDelete(ctx, publicid).Execute()

Delete Webhook



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
	publicid := "publicid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookAPI.WebhookByPublicidDelete(context.Background(), publicid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookByPublicidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookByPublicidDeleteRequest struct via the builder pattern


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


## WebhookByPublicidGet

> Webhook WebhookByPublicidGet(ctx, publicid).Execute()

Load Webhook



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
	publicid := "publicid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookByPublicidGet(context.Background(), publicid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookByPublicidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookByPublicidGet`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookByPublicidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookByPublicidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookByPublicidPut

> Webhook WebhookByPublicidPut(ctx, publicid).WebhookUpdatePayload(webhookUpdatePayload).Execute()

Update Webhook



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
	publicid := "publicid_example" // string | 
	webhookUpdatePayload := *openapiclient.NewWebhookUpdatePayload() // WebhookUpdatePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookByPublicidPut(context.Background(), publicid).WebhookUpdatePayload(webhookUpdatePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookByPublicidPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookByPublicidPut`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookByPublicidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookByPublicidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookUpdatePayload** | [**WebhookUpdatePayload**](WebhookUpdatePayload.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookGet

> []Webhook WebhookGet(ctx).Limit(limit).Offset(offset).Execute()

Load Webhooks



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
	resp, r, err := apiClient.WebhookAPI.WebhookGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookGet`: []Webhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]Webhook**](Webhook.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookPost

> Webhook WebhookPost(ctx).WebhookCreatePayload(webhookCreatePayload).Execute()

Add Webhook



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
	webhookCreatePayload := *openapiclient.NewWebhookCreatePayload("attachment.txt", "http://address.for.notification.com") // WebhookCreatePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookPost(context.Background()).WebhookCreatePayload(webhookCreatePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookPost`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookCreatePayload** | [**WebhookCreatePayload**](WebhookCreatePayload.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

