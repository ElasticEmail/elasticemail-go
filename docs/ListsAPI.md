# \ListsAPI

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListsByListnameContactsGet**](ListsAPI.md#ListsByListnameContactsGet) | **Get** /lists/{listname}/contacts | Load Contacts in List
[**ListsByNameContactsPost**](ListsAPI.md#ListsByNameContactsPost) | **Post** /lists/{name}/contacts | Add Contacts to List
[**ListsByNameContactsRemovePost**](ListsAPI.md#ListsByNameContactsRemovePost) | **Post** /lists/{name}/contacts/remove | Remove Contacts from List
[**ListsByNameDelete**](ListsAPI.md#ListsByNameDelete) | **Delete** /lists/{name} | Delete List
[**ListsByNameGet**](ListsAPI.md#ListsByNameGet) | **Get** /lists/{name} | Load List
[**ListsByNamePut**](ListsAPI.md#ListsByNamePut) | **Put** /lists/{name} | Update List
[**ListsGet**](ListsAPI.md#ListsGet) | **Get** /lists | Load Lists
[**ListsPost**](ListsAPI.md#ListsPost) | **Post** /lists | Add List



## ListsByListnameContactsGet

> []Contact ListsByListnameContactsGet(ctx, listname).Limit(limit).Offset(offset).Execute()

Load Contacts in List



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
	listname := "My List 1" // string | Name of your list.
	limit := int32(100) // int32 | Maximum number of returned items. (optional)
	offset := int32(20) // int32 | How many items should be returned ahead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ListsByListnameContactsGet(context.Background(), listname).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsByListnameContactsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsByListnameContactsGet`: []Contact
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ListsByListnameContactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listname** | **string** | Name of your list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsByListnameContactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]Contact**](Contact.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsByNameContactsPost

> ContactsList ListsByNameContactsPost(ctx, name).EmailsPayload(emailsPayload).Execute()

Add Contacts to List



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
	name := "My List 1" // string | Name of your list.
	emailsPayload := *openapiclient.NewEmailsPayload() // EmailsPayload | Provide either rule or a list of emails, not both.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ListsByNameContactsPost(context.Background(), name).EmailsPayload(emailsPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsByNameContactsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsByNameContactsPost`: ContactsList
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ListsByNameContactsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of your list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsByNameContactsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailsPayload** | [**EmailsPayload**](EmailsPayload.md) | Provide either rule or a list of emails, not both. | 

### Return type

[**ContactsList**](ContactsList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsByNameContactsRemovePost

> ListsByNameContactsRemovePost(ctx, name).EmailsPayload(emailsPayload).Execute()

Remove Contacts from List



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
	name := "My List 1" // string | Name of your list.
	emailsPayload := *openapiclient.NewEmailsPayload() // EmailsPayload | Provide either rule or a list of emails, not both.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.ListsByNameContactsRemovePost(context.Background(), name).EmailsPayload(emailsPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsByNameContactsRemovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of your list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsByNameContactsRemovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailsPayload** | [**EmailsPayload**](EmailsPayload.md) | Provide either rule or a list of emails, not both. | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsByNameDelete

> ListsByNameDelete(ctx, name).Execute()

Delete List



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
	name := "My List 1" // string | Name of your list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.ListsByNameDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsByNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of your list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsByNameDeleteRequest struct via the builder pattern


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


## ListsByNameGet

> ContactsList ListsByNameGet(ctx, name).Execute()

Load List



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
	name := "My List 1" // string | Name of your list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ListsByNameGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsByNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsByNameGet`: ContactsList
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ListsByNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of your list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsByNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContactsList**](ContactsList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsByNamePut

> ContactsList ListsByNamePut(ctx, name).ListUpdatePayload(listUpdatePayload).Execute()

Update List



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
	name := "My List 1" // string | Name of your list.
	listUpdatePayload := *openapiclient.NewListUpdatePayload() // ListUpdatePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ListsByNamePut(context.Background(), name).ListUpdatePayload(listUpdatePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsByNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsByNamePut`: ContactsList
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ListsByNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of your list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsByNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listUpdatePayload** | [**ListUpdatePayload**](ListUpdatePayload.md) |  | 

### Return type

[**ContactsList**](ContactsList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsGet

> []ContactsList ListsGet(ctx).Limit(limit).Offset(offset).Execute()

Load Lists



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
	resp, r, err := apiClient.ListsAPI.ListsGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsGet`: []ContactsList
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ListsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]ContactsList**](ContactsList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsPost

> ContactsList ListsPost(ctx).ListPayload(listPayload).Execute()

Add List



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
	listPayload := *openapiclient.NewListPayload("My List 1") // ListPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ListsPost(context.Background()).ListPayload(listPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsPost`: ContactsList
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ListsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listPayload** | [**ListPayload**](ListPayload.md) |  | 

### Return type

[**ContactsList**](ContactsList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

