# \EventsAPI

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsByTransactionidGet**](EventsAPI.md#EventsByTransactionidGet) | **Get** /events/{transactionid} | Load Email Events
[**EventsChannelsByNameExportPost**](EventsAPI.md#EventsChannelsByNameExportPost) | **Post** /events/channels/{name}/export | Export Channel Events
[**EventsChannelsByNameGet**](EventsAPI.md#EventsChannelsByNameGet) | **Get** /events/channels/{name} | Load Channel Events
[**EventsChannelsExportByIdStatusGet**](EventsAPI.md#EventsChannelsExportByIdStatusGet) | **Get** /events/channels/export/{id}/status | Check Channel Export Status
[**EventsExportByIdStatusGet**](EventsAPI.md#EventsExportByIdStatusGet) | **Get** /events/export/{id}/status | Check Export Status
[**EventsExportPost**](EventsAPI.md#EventsExportPost) | **Post** /events/export | Export Events
[**EventsGet**](EventsAPI.md#EventsGet) | **Get** /events | Load Events



## EventsByTransactionidGet

> []RecipientEvent EventsByTransactionidGet(ctx, transactionid).From(from).To(to).OrderBy(orderBy).Limit(limit).Offset(offset).Execute()

Load Email Events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/elasticemail/elasticemail-go"
)

func main() {
	transactionid := "TransactionID" // string | ID number of transaction
	from := time.Now() // time.Time | Starting date for search in YYYY-MM-DDThh:mm:ss format. (optional)
	to := time.Now() // time.Time | Ending date for search in YYYY-MM-DDThh:mm:ss format. (optional)
	orderBy := openapiclient.EventsOrderBy("DateDescending") // EventsOrderBy |  (optional) (default to "DateDescending")
	limit := int32(100) // int32 | Maximum number of returned items. (optional)
	offset := int32(20) // int32 | How many items should be returned ahead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsByTransactionidGet(context.Background(), transactionid).From(from).To(to).OrderBy(orderBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsByTransactionidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsByTransactionidGet`: []RecipientEvent
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsByTransactionidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionid** | **string** | ID number of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsByTransactionidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Starting date for search in YYYY-MM-DDThh:mm:ss format. | 
 **to** | **time.Time** | Ending date for search in YYYY-MM-DDThh:mm:ss format. | 
 **orderBy** | [**EventsOrderBy**](EventsOrderBy.md) |  | [default to &quot;DateDescending&quot;]
 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]RecipientEvent**](RecipientEvent.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsChannelsByNameExportPost

> ExportLink EventsChannelsByNameExportPost(ctx, name).EventTypes(eventTypes).From(from).To(to).FileFormat(fileFormat).CompressionFormat(compressionFormat).FileName(fileName).Execute()

Export Channel Events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/elasticemail/elasticemail-go"
)

func main() {
	name := "Channel01" // string | Name of selected channel.
	eventTypes := []openapiclient.EventType{openapiclient.EventType("Submission")} // []EventType | Types of Events to return (optional)
	from := time.Now() // time.Time | Starting date for search in YYYY-MM-DDThh:mm:ss format. (optional)
	to := time.Now() // time.Time | Ending date for search in YYYY-MM-DDThh:mm:ss format. (optional)
	fileFormat := openapiclient.ExportFileFormats("Csv") // ExportFileFormats | Format of the exported file (optional) (default to "Csv")
	compressionFormat := openapiclient.CompressionFormat("None") // CompressionFormat | FileResponse compression format. None or Zip. (optional) (default to "None")
	fileName := "filename.txt" // string | Name of your file including extension. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsChannelsByNameExportPost(context.Background(), name).EventTypes(eventTypes).From(from).To(to).FileFormat(fileFormat).CompressionFormat(compressionFormat).FileName(fileName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsChannelsByNameExportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsChannelsByNameExportPost`: ExportLink
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsChannelsByNameExportPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of selected channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsChannelsByNameExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventTypes** | [**[]EventType**](EventType.md) | Types of Events to return | 
 **from** | **time.Time** | Starting date for search in YYYY-MM-DDThh:mm:ss format. | 
 **to** | **time.Time** | Ending date for search in YYYY-MM-DDThh:mm:ss format. | 
 **fileFormat** | [**ExportFileFormats**](ExportFileFormats.md) | Format of the exported file | [default to &quot;Csv&quot;]
 **compressionFormat** | [**CompressionFormat**](CompressionFormat.md) | FileResponse compression format. None or Zip. | [default to &quot;None&quot;]
 **fileName** | **string** | Name of your file including extension. | 

### Return type

[**ExportLink**](ExportLink.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsChannelsByNameGet

> []RecipientEvent EventsChannelsByNameGet(ctx, name).EventTypes(eventTypes).From(from).To(to).OrderBy(orderBy).Limit(limit).Offset(offset).Execute()

Load Channel Events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/elasticemail/elasticemail-go"
)

func main() {
	name := "Channel01" // string | Name of selected channel.
	eventTypes := []openapiclient.EventType{openapiclient.EventType("Submission")} // []EventType | Types of Events to return (optional)
	from := time.Now() // time.Time | Starting date for search in YYYY-MM-DDThh:mm:ss format. (optional)
	to := time.Now() // time.Time | Ending date for search in YYYY-MM-DDThh:mm:ss format. (optional)
	orderBy := openapiclient.EventsOrderBy("DateDescending") // EventsOrderBy |  (optional) (default to "DateDescending")
	limit := int32(56) // int32 | How many items to load. Maximum for this request is 1000 items (optional)
	offset := int32(20) // int32 | How many items should be returned ahead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsChannelsByNameGet(context.Background(), name).EventTypes(eventTypes).From(from).To(to).OrderBy(orderBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsChannelsByNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsChannelsByNameGet`: []RecipientEvent
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsChannelsByNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of selected channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsChannelsByNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventTypes** | [**[]EventType**](EventType.md) | Types of Events to return | 
 **from** | **time.Time** | Starting date for search in YYYY-MM-DDThh:mm:ss format. | 
 **to** | **time.Time** | Ending date for search in YYYY-MM-DDThh:mm:ss format. | 
 **orderBy** | [**EventsOrderBy**](EventsOrderBy.md) |  | [default to &quot;DateDescending&quot;]
 **limit** | **int32** | How many items to load. Maximum for this request is 1000 items | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]RecipientEvent**](RecipientEvent.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsChannelsExportByIdStatusGet

> ExportStatus EventsChannelsExportByIdStatusGet(ctx, id).Execute()

Check Channel Export Status



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
	id := "E33EBA7A-C20D-4D3D-8F2F-5EEF42F58E6F" // string | ID of the exported file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsChannelsExportByIdStatusGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsChannelsExportByIdStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsChannelsExportByIdStatusGet`: ExportStatus
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsChannelsExportByIdStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the exported file | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsChannelsExportByIdStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExportStatus**](ExportStatus.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsExportByIdStatusGet

> ExportStatus EventsExportByIdStatusGet(ctx, id).Execute()

Check Export Status



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
	id := "E33EBA7A-C20D-4D3D-8F2F-5EEF42F58E6F" // string | ID of the exported file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsExportByIdStatusGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsExportByIdStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsExportByIdStatusGet`: ExportStatus
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsExportByIdStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the exported file | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsExportByIdStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExportStatus**](ExportStatus.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsExportPost

> ExportLink EventsExportPost(ctx).EventTypes(eventTypes).From(from).To(to).FileFormat(fileFormat).CompressionFormat(compressionFormat).FileName(fileName).Execute()

Export Events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/elasticemail/elasticemail-go"
)

func main() {
	eventTypes := []openapiclient.EventType{openapiclient.EventType("Submission")} // []EventType | Types of Events to return (optional)
	from := time.Now() // time.Time | Starting date for search in YYYY-MM-DDThh:mm:ss format. (optional)
	to := time.Now() // time.Time | Ending date for search in YYYY-MM-DDThh:mm:ss format. (optional)
	fileFormat := openapiclient.ExportFileFormats("Csv") // ExportFileFormats | Format of the exported file (optional) (default to "Csv")
	compressionFormat := openapiclient.CompressionFormat("None") // CompressionFormat | FileResponse compression format. None or Zip. (optional) (default to "None")
	fileName := "filename.txt" // string | Name of your file including extension. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsExportPost(context.Background()).EventTypes(eventTypes).From(from).To(to).FileFormat(fileFormat).CompressionFormat(compressionFormat).FileName(fileName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsExportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsExportPost`: ExportLink
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventTypes** | [**[]EventType**](EventType.md) | Types of Events to return | 
 **from** | **time.Time** | Starting date for search in YYYY-MM-DDThh:mm:ss format. | 
 **to** | **time.Time** | Ending date for search in YYYY-MM-DDThh:mm:ss format. | 
 **fileFormat** | [**ExportFileFormats**](ExportFileFormats.md) | Format of the exported file | [default to &quot;Csv&quot;]
 **compressionFormat** | [**CompressionFormat**](CompressionFormat.md) | FileResponse compression format. None or Zip. | [default to &quot;None&quot;]
 **fileName** | **string** | Name of your file including extension. | 

### Return type

[**ExportLink**](ExportLink.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsGet

> []RecipientEvent EventsGet(ctx).EventTypes(eventTypes).From(from).To(to).OrderBy(orderBy).Limit(limit).Offset(offset).Execute()

Load Events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/elasticemail/elasticemail-go"
)

func main() {
	eventTypes := []openapiclient.EventType{openapiclient.EventType("Submission")} // []EventType | Types of Events to return (optional)
	from := time.Now() // time.Time | Starting date for search in YYYY-MM-DDThh:mm:ss format. (optional)
	to := time.Now() // time.Time | Ending date for search in YYYY-MM-DDThh:mm:ss format. (optional)
	orderBy := openapiclient.EventsOrderBy("DateDescending") // EventsOrderBy |  (optional) (default to "DateDescending")
	limit := int32(56) // int32 | How many items to load. Maximum for this request is 1000 items (optional)
	offset := int32(20) // int32 | How many items should be returned ahead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsGet(context.Background()).EventTypes(eventTypes).From(from).To(to).OrderBy(orderBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsGet`: []RecipientEvent
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventTypes** | [**[]EventType**](EventType.md) | Types of Events to return | 
 **from** | **time.Time** | Starting date for search in YYYY-MM-DDThh:mm:ss format. | 
 **to** | **time.Time** | Ending date for search in YYYY-MM-DDThh:mm:ss format. | 
 **orderBy** | [**EventsOrderBy**](EventsOrderBy.md) |  | [default to &quot;DateDescending&quot;]
 **limit** | **int32** | How many items to load. Maximum for this request is 1000 items | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]RecipientEvent**](RecipientEvent.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

