# \EmailsApi

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailsByMsgidViewGet**](EmailsApi.md#EmailsByMsgidViewGet) | **Get** /emails/{msgid}/view | View Email
[**EmailsMergefilePost**](EmailsApi.md#EmailsMergefilePost) | **Post** /emails/mergefile | Send Bulk Emails CSV
[**EmailsPost**](EmailsApi.md#EmailsPost) | **Post** /emails | Send Bulk Emails
[**EmailsTransactionalPost**](EmailsApi.md#EmailsTransactionalPost) | **Post** /emails/transactional | Send Transactional Email



## EmailsByMsgidViewGet

> EmailData EmailsByMsgidViewGet(ctx, msgid).Execute()

View Email



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
    msgid := "msgid_example" // string | Message identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailsApi.EmailsByMsgidViewGet(context.Background(), msgid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailsApi.EmailsByMsgidViewGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailsByMsgidViewGet`: EmailData
    fmt.Fprintf(os.Stdout, "Response from `EmailsApi.EmailsByMsgidViewGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgid** | **string** | Message identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailsByMsgidViewGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailData**](EmailData.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailsMergefilePost

> EmailSend EmailsMergefilePost(ctx).MergeEmailPayload(mergeEmailPayload).Execute()

Send Bulk Emails CSV



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
    mergeEmailPayload := *openapiclient.NewMergeEmailPayload(*openapiclient.NewMessageAttachment(string(123), "Name_example")) // MergeEmailPayload | Email data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailsApi.EmailsMergefilePost(context.Background()).MergeEmailPayload(mergeEmailPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailsApi.EmailsMergefilePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailsMergefilePost`: EmailSend
    fmt.Fprintf(os.Stdout, "Response from `EmailsApi.EmailsMergefilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailsMergefilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mergeEmailPayload** | [**MergeEmailPayload**](MergeEmailPayload.md) | Email data | 

### Return type

[**EmailSend**](EmailSend.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailsPost

> EmailSend EmailsPost(ctx).EmailMessageData(emailMessageData).Execute()

Send Bulk Emails



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
    emailMessageData := *openapiclient.NewEmailMessageData([]openapiclient.EmailRecipient{*openapiclient.NewEmailRecipient("mail@example.com")}) // EmailMessageData | Email data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailsApi.EmailsPost(context.Background()).EmailMessageData(emailMessageData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailsApi.EmailsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailsPost`: EmailSend
    fmt.Fprintf(os.Stdout, "Response from `EmailsApi.EmailsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailMessageData** | [**EmailMessageData**](EmailMessageData.md) | Email data | 

### Return type

[**EmailSend**](EmailSend.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailsTransactionalPost

> EmailSend EmailsTransactionalPost(ctx).EmailTransactionalMessageData(emailTransactionalMessageData).Execute()

Send Transactional Email



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
    emailTransactionalMessageData := *openapiclient.NewEmailTransactionalMessageData(*openapiclient.NewTransactionalRecipient([]string{"To_example"})) // EmailTransactionalMessageData | Email data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailsApi.EmailsTransactionalPost(context.Background()).EmailTransactionalMessageData(emailTransactionalMessageData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailsApi.EmailsTransactionalPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailsTransactionalPost`: EmailSend
    fmt.Fprintf(os.Stdout, "Response from `EmailsApi.EmailsTransactionalPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailsTransactionalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailTransactionalMessageData** | [**EmailTransactionalMessageData**](EmailTransactionalMessageData.md) | Email data | 

### Return type

[**EmailSend**](EmailSend.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

