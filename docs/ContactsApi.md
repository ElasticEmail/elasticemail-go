# \ContactsApi

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactsByEmailDelete**](ContactsApi.md#ContactsByEmailDelete) | **Delete** /contacts/{email} | Delete Contact
[**ContactsByEmailGet**](ContactsApi.md#ContactsByEmailGet) | **Get** /contacts/{email} | Load Contact
[**ContactsByEmailPut**](ContactsApi.md#ContactsByEmailPut) | **Put** /contacts/{email} | Update Contact
[**ContactsDeletePost**](ContactsApi.md#ContactsDeletePost) | **Post** /contacts/delete | Delete Contacts Bulk
[**ContactsExportByIdStatusGet**](ContactsApi.md#ContactsExportByIdStatusGet) | **Get** /contacts/export/{id}/status | Check Export Status
[**ContactsExportPost**](ContactsApi.md#ContactsExportPost) | **Post** /contacts/export | Export Contacts
[**ContactsGet**](ContactsApi.md#ContactsGet) | **Get** /contacts | Load Contacts
[**ContactsImportPost**](ContactsApi.md#ContactsImportPost) | **Post** /contacts/import | Upload Contacts
[**ContactsPost**](ContactsApi.md#ContactsPost) | **Post** /contacts | Add Contact



## ContactsByEmailDelete

> ContactsByEmailDelete(ctx, email).Execute()

Delete Contact



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
    email := "mail@example.com" // string | Proper email address.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContactsApi.ContactsByEmailDelete(context.Background(), email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.ContactsByEmailDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiContactsByEmailDeleteRequest struct via the builder pattern


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


## ContactsByEmailGet

> Contact ContactsByEmailGet(ctx, email).Execute()

Load Contact



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
    email := "mail@example.com" // string | Proper email address.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.ContactsByEmailGet(context.Background(), email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.ContactsByEmailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsByEmailGet`: Contact
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.ContactsByEmailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Proper email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsByEmailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Contact**](Contact.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsByEmailPut

> Contact ContactsByEmailPut(ctx, email).ContactUpdatePayload(contactUpdatePayload).Execute()

Update Contact



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
    email := "mail@example.com" // string | Proper email address.
    contactUpdatePayload := *openapiclient.NewContactUpdatePayload() // ContactUpdatePayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.ContactsByEmailPut(context.Background(), email).ContactUpdatePayload(contactUpdatePayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.ContactsByEmailPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsByEmailPut`: Contact
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.ContactsByEmailPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Proper email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsByEmailPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contactUpdatePayload** | [**ContactUpdatePayload**](ContactUpdatePayload.md) |  | 

### Return type

[**Contact**](Contact.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsDeletePost

> ContactsDeletePost(ctx).EmailsPayload(emailsPayload).Execute()

Delete Contacts Bulk



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
    emailsPayload := *openapiclient.NewEmailsPayload() // EmailsPayload | Provide either rule or a list of emails, not both.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContactsApi.ContactsDeletePost(context.Background()).EmailsPayload(emailsPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.ContactsDeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsDeletePostRequest struct via the builder pattern


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


## ContactsExportByIdStatusGet

> ExportStatus ContactsExportByIdStatusGet(ctx, id).Execute()

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
    resp, r, err := apiClient.ContactsApi.ContactsExportByIdStatusGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.ContactsExportByIdStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsExportByIdStatusGet`: ExportStatus
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.ContactsExportByIdStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the exported file | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsExportByIdStatusGetRequest struct via the builder pattern


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


## ContactsExportPost

> ExportLink ContactsExportPost(ctx).FileFormat(fileFormat).Rule(rule).Emails(emails).CompressionFormat(compressionFormat).FileName(fileName).Execute()

Export Contacts



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
    fileFormat := openapiclient.ExportFileFormats("Csv") // ExportFileFormats | Format of the exported file (optional) (default to "Csv")
    rule := "Status%20=%20Engaged" // string | Query used for filtering. (optional)
    emails := []string{"Inner_example"} // []string | Comma delimited list of contact emails (optional)
    compressionFormat := openapiclient.CompressionFormat("None") // CompressionFormat | FileResponse compression format. None or Zip. (optional) (default to "None")
    fileName := "filename.txt" // string | Name of your file including extension. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.ContactsExportPost(context.Background()).FileFormat(fileFormat).Rule(rule).Emails(emails).CompressionFormat(compressionFormat).FileName(fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.ContactsExportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsExportPost`: ExportLink
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.ContactsExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileFormat** | [**ExportFileFormats**](ExportFileFormats.md) | Format of the exported file | [default to &quot;Csv&quot;]
 **rule** | **string** | Query used for filtering. | 
 **emails** | **[]string** | Comma delimited list of contact emails | 
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


## ContactsGet

> []Contact ContactsGet(ctx).Limit(limit).Offset(offset).Execute()

Load Contacts



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
    resp, r, err := apiClient.ContactsApi.ContactsGet(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.ContactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsGet`: []Contact
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.ContactsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsGetRequest struct via the builder pattern


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


## ContactsImportPost

> ContactsImportPost(ctx).ListName(listName).EncodingName(encodingName).FileUrl(fileUrl).File(file).Execute()

Upload Contacts



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
    listName := "listName_example" // string | Name of an existing list to add these contacts to (optional)
    encodingName := "encodingName_example" // string | In what encoding the file is uploaded (optional)
    fileUrl := "fileUrl_example" // string | Optional url of csv to import (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContactsApi.ContactsImportPost(context.Background()).ListName(listName).EncodingName(encodingName).FileUrl(fileUrl).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.ContactsImportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listName** | **string** | Name of an existing list to add these contacts to | 
 **encodingName** | **string** | In what encoding the file is uploaded | 
 **fileUrl** | **string** | Optional url of csv to import | 
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


## ContactsPost

> []Contact ContactsPost(ctx).ContactPayload(contactPayload).Listnames(listnames).Execute()

Add Contact



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
    contactPayload := []openapiclient.ContactPayload{*openapiclient.NewContactPayload("mail@example.com")} // []ContactPayload | 
    listnames := []string{"Inner_example"} // []string | Names of lists to which the uploaded contacts should be added to (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.ContactsPost(context.Background()).ContactPayload(contactPayload).Listnames(listnames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.ContactsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsPost`: []Contact
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.ContactsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactPayload** | [**[]ContactPayload**](ContactPayload.md) |  | 
 **listnames** | **[]string** | Names of lists to which the uploaded contacts should be added to | 

### Return type

[**[]Contact**](Contact.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

