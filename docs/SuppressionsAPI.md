# \SuppressionsAPI

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SuppressionsBouncesGet**](SuppressionsAPI.md#SuppressionsBouncesGet) | **Get** /suppressions/bounces | Get Bounce List
[**SuppressionsBouncesImportPost**](SuppressionsAPI.md#SuppressionsBouncesImportPost) | **Post** /suppressions/bounces/import | Add Bounces Async
[**SuppressionsBouncesPost**](SuppressionsAPI.md#SuppressionsBouncesPost) | **Post** /suppressions/bounces | Add Bounces
[**SuppressionsByEmailDelete**](SuppressionsAPI.md#SuppressionsByEmailDelete) | **Delete** /suppressions/{email} | Delete Suppression
[**SuppressionsByEmailGet**](SuppressionsAPI.md#SuppressionsByEmailGet) | **Get** /suppressions/{email} | Get Suppression
[**SuppressionsComplaintsGet**](SuppressionsAPI.md#SuppressionsComplaintsGet) | **Get** /suppressions/complaints | Get Complaints List
[**SuppressionsComplaintsImportPost**](SuppressionsAPI.md#SuppressionsComplaintsImportPost) | **Post** /suppressions/complaints/import | Add Complaints Async
[**SuppressionsComplaintsPost**](SuppressionsAPI.md#SuppressionsComplaintsPost) | **Post** /suppressions/complaints | Add Complaints
[**SuppressionsGet**](SuppressionsAPI.md#SuppressionsGet) | **Get** /suppressions | Get Suppressions
[**SuppressionsUnsubscribesGet**](SuppressionsAPI.md#SuppressionsUnsubscribesGet) | **Get** /suppressions/unsubscribes | Get Unsubscribes List
[**SuppressionsUnsubscribesImportPost**](SuppressionsAPI.md#SuppressionsUnsubscribesImportPost) | **Post** /suppressions/unsubscribes/import | Add Unsubscribes Async
[**SuppressionsUnsubscribesPost**](SuppressionsAPI.md#SuppressionsUnsubscribesPost) | **Post** /suppressions/unsubscribes | Add Unsubscribes



## SuppressionsBouncesGet

> []Suppression SuppressionsBouncesGet(ctx).Search(search).Limit(limit).Offset(offset).Execute()

Get Bounce List



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
	search := "text" // string | Text fragment used for searching. (optional)
	limit := int32(100) // int32 | Maximum number of returned items. (optional)
	offset := int32(20) // int32 | How many items should be returned ahead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionsAPI.SuppressionsBouncesGet(context.Background()).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.SuppressionsBouncesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppressionsBouncesGet`: []Suppression
	fmt.Fprintf(os.Stdout, "Response from `SuppressionsAPI.SuppressionsBouncesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuppressionsBouncesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Text fragment used for searching. | 
 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]Suppression**](Suppression.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppressionsBouncesImportPost

> SuppressionsBouncesImportPost(ctx).File(file).Execute()

Add Bounces Async



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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SuppressionsAPI.SuppressionsBouncesImportPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.SuppressionsBouncesImportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuppressionsBouncesImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppressionsBouncesPost

> []Suppression SuppressionsBouncesPost(ctx).RequestBody(requestBody).Execute()

Add Bounces



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
	requestBody := []string{"Property_example"} // []string | Emails to add as bounces. Limited to 1000 per request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionsAPI.SuppressionsBouncesPost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.SuppressionsBouncesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppressionsBouncesPost`: []Suppression
	fmt.Fprintf(os.Stdout, "Response from `SuppressionsAPI.SuppressionsBouncesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuppressionsBouncesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | Emails to add as bounces. Limited to 1000 per request | 

### Return type

[**[]Suppression**](Suppression.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppressionsByEmailDelete

> SuppressionsByEmailDelete(ctx, email).Execute()

Delete Suppression



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
	email := "mail@example.com" // string | Proper email address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SuppressionsAPI.SuppressionsByEmailDelete(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.SuppressionsByEmailDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Proper email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppressionsByEmailDeleteRequest struct via the builder pattern


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


## SuppressionsByEmailGet

> Suppression SuppressionsByEmailGet(ctx, email).Execute()

Get Suppression



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
	email := "mail@example.com" // string | Proper email address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionsAPI.SuppressionsByEmailGet(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.SuppressionsByEmailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppressionsByEmailGet`: Suppression
	fmt.Fprintf(os.Stdout, "Response from `SuppressionsAPI.SuppressionsByEmailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Proper email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppressionsByEmailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Suppression**](Suppression.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppressionsComplaintsGet

> []Suppression SuppressionsComplaintsGet(ctx).Search(search).Limit(limit).Offset(offset).Execute()

Get Complaints List



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
	search := "text" // string | Text fragment used for searching. (optional)
	limit := int32(100) // int32 | Maximum number of returned items. (optional)
	offset := int32(20) // int32 | How many items should be returned ahead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionsAPI.SuppressionsComplaintsGet(context.Background()).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.SuppressionsComplaintsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppressionsComplaintsGet`: []Suppression
	fmt.Fprintf(os.Stdout, "Response from `SuppressionsAPI.SuppressionsComplaintsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuppressionsComplaintsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Text fragment used for searching. | 
 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]Suppression**](Suppression.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppressionsComplaintsImportPost

> SuppressionsComplaintsImportPost(ctx).File(file).Execute()

Add Complaints Async



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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SuppressionsAPI.SuppressionsComplaintsImportPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.SuppressionsComplaintsImportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuppressionsComplaintsImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppressionsComplaintsPost

> []Suppression SuppressionsComplaintsPost(ctx).RequestBody(requestBody).Execute()

Add Complaints



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
	requestBody := []string{"Property_example"} // []string | Emails to add as complaints. Limited to 1000 per request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionsAPI.SuppressionsComplaintsPost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.SuppressionsComplaintsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppressionsComplaintsPost`: []Suppression
	fmt.Fprintf(os.Stdout, "Response from `SuppressionsAPI.SuppressionsComplaintsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuppressionsComplaintsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | Emails to add as complaints. Limited to 1000 per request | 

### Return type

[**[]Suppression**](Suppression.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppressionsGet

> []Suppression SuppressionsGet(ctx).Limit(limit).Offset(offset).Execute()

Get Suppressions



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
	resp, r, err := apiClient.SuppressionsAPI.SuppressionsGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.SuppressionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppressionsGet`: []Suppression
	fmt.Fprintf(os.Stdout, "Response from `SuppressionsAPI.SuppressionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuppressionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]Suppression**](Suppression.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppressionsUnsubscribesGet

> []Suppression SuppressionsUnsubscribesGet(ctx).Search(search).Limit(limit).Offset(offset).Execute()

Get Unsubscribes List



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
	search := "text" // string | Text fragment used for searching. (optional)
	limit := int32(100) // int32 | Maximum number of returned items. (optional)
	offset := int32(20) // int32 | How many items should be returned ahead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionsAPI.SuppressionsUnsubscribesGet(context.Background()).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.SuppressionsUnsubscribesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppressionsUnsubscribesGet`: []Suppression
	fmt.Fprintf(os.Stdout, "Response from `SuppressionsAPI.SuppressionsUnsubscribesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuppressionsUnsubscribesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Text fragment used for searching. | 
 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]Suppression**](Suppression.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppressionsUnsubscribesImportPost

> SuppressionsUnsubscribesImportPost(ctx).File(file).Execute()

Add Unsubscribes Async



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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SuppressionsAPI.SuppressionsUnsubscribesImportPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.SuppressionsUnsubscribesImportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuppressionsUnsubscribesImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppressionsUnsubscribesPost

> []Suppression SuppressionsUnsubscribesPost(ctx).RequestBody(requestBody).Execute()

Add Unsubscribes



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
	requestBody := []string{"Property_example"} // []string | Emails to add as unsubscribes. Limited to 1000 per request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionsAPI.SuppressionsUnsubscribesPost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.SuppressionsUnsubscribesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppressionsUnsubscribesPost`: []Suppression
	fmt.Fprintf(os.Stdout, "Response from `SuppressionsAPI.SuppressionsUnsubscribesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuppressionsUnsubscribesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | Emails to add as unsubscribes. Limited to 1000 per request | 

### Return type

[**[]Suppression**](Suppression.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

