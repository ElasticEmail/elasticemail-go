# \ListsApi

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListsByNameContactsPost**](ListsApi.md#ListsByNameContactsPost) | **Post** /lists/{name}/contacts | Add Contacts to List
[**ListsByNameContactsRemovePost**](ListsApi.md#ListsByNameContactsRemovePost) | **Post** /lists/{name}/contacts/remove | Remove Contacts from List
[**ListsByNameDelete**](ListsApi.md#ListsByNameDelete) | **Delete** /lists/{name} | Delete List
[**ListsByNameGet**](ListsApi.md#ListsByNameGet) | **Get** /lists/{name} | Load List
[**ListsByNamePut**](ListsApi.md#ListsByNamePut) | **Put** /lists/{name} | Update List
[**ListsGet**](ListsApi.md#ListsGet) | **Get** /lists | Load Lists
[**ListsPost**](ListsApi.md#ListsPost) | **Post** /lists | Add List



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
    resp, r, err := apiClient.ListsApi.ListsByNameContactsPost(context.Background(), name).EmailsPayload(emailsPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.ListsByNameContactsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsByNameContactsPost`: ContactsList
    fmt.Fprintf(os.Stdout, "Response from `ListsApi.ListsByNameContactsPost`: %v\n", resp)
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
    r, err := apiClient.ListsApi.ListsByNameContactsRemovePost(context.Background(), name).EmailsPayload(emailsPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.ListsByNameContactsRemovePost``: %v\n", err)
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
    r, err := apiClient.ListsApi.ListsByNameDelete(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.ListsByNameDelete``: %v\n", err)
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
    resp, r, err := apiClient.ListsApi.ListsByNameGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.ListsByNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsByNameGet`: ContactsList
    fmt.Fprintf(os.Stdout, "Response from `ListsApi.ListsByNameGet`: %v\n", resp)
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
    resp, r, err := apiClient.ListsApi.ListsByNamePut(context.Background(), name).ListUpdatePayload(listUpdatePayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.ListsByNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsByNamePut`: ContactsList
    fmt.Fprintf(os.Stdout, "Response from `ListsApi.ListsByNamePut`: %v\n", resp)
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
    resp, r, err := apiClient.ListsApi.ListsGet(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.ListsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsGet`: []ContactsList
    fmt.Fprintf(os.Stdout, "Response from `ListsApi.ListsGet`: %v\n", resp)
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
    resp, r, err := apiClient.ListsApi.ListsPost(context.Background()).ListPayload(listPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.ListsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsPost`: ContactsList
    fmt.Fprintf(os.Stdout, "Response from `ListsApi.ListsPost`: %v\n", resp)
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

