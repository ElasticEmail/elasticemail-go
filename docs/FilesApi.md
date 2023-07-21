# \FilesApi

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesByNameDelete**](FilesApi.md#FilesByNameDelete) | **Delete** /files/{name} | Delete File
[**FilesByNameGet**](FilesApi.md#FilesByNameGet) | **Get** /files/{name} | Download File
[**FilesByNameInfoGet**](FilesApi.md#FilesByNameInfoGet) | **Get** /files/{name}/info | Load File Details
[**FilesGet**](FilesApi.md#FilesGet) | **Get** /files | List Files
[**FilesPost**](FilesApi.md#FilesPost) | **Post** /files | Upload File



## FilesByNameDelete

> FilesByNameDelete(ctx, name).Execute()

Delete File



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
    name := "filename.txt" // string | Name of your file including extension.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FilesApi.FilesByNameDelete(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.FilesByNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of your file including extension. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesByNameDeleteRequest struct via the builder pattern


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


## FilesByNameGet

> *os.File FilesByNameGet(ctx, name).Execute()

Download File



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
    name := "filename.txt" // string | Name of your file including extension.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.FilesByNameGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.FilesByNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesByNameGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.FilesByNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of your file including extension. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesByNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesByNameInfoGet

> FileInfo FilesByNameInfoGet(ctx, name).Execute()

Load File Details



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
    name := "filename.txt" // string | Name of your file including extension.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.FilesByNameInfoGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.FilesByNameInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesByNameInfoGet`: FileInfo
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.FilesByNameInfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of your file including extension. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesByNameInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileInfo**](FileInfo.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesGet

> []FileInfo FilesGet(ctx).Limit(limit).Offset(offset).Execute()

List Files



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
    resp, r, err := apiClient.FilesApi.FilesGet(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.FilesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesGet`: []FileInfo
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.FilesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]FileInfo**](FileInfo.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesPost

> FileInfo FilesPost(ctx).FilePayload(filePayload).ExpiresAfterDays(expiresAfterDays).Execute()

Upload File



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
    filePayload := *openapiclient.NewFilePayload(string(123)) // FilePayload | 
    expiresAfterDays := int32(56) // int32 | After how many days should the file be deleted. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.FilesPost(context.Background()).FilePayload(filePayload).ExpiresAfterDays(expiresAfterDays).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.FilesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesPost`: FileInfo
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.FilesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filePayload** | [**FilePayload**](FilePayload.md) |  | 
 **expiresAfterDays** | **int32** | After how many days should the file be deleted. | 

### Return type

[**FileInfo**](FileInfo.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

