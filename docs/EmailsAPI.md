# \EmailsAPI

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailsByMsgidViewGet**](EmailsAPI.md#EmailsByMsgidViewGet) | **Get** /emails/{msgid}/view | View Email
[**EmailsByTransactionidStatusGet**](EmailsAPI.md#EmailsByTransactionidStatusGet) | **Get** /emails/{transactionid}/status | Get Status
[**EmailsMergefilePost**](EmailsAPI.md#EmailsMergefilePost) | **Post** /emails/mergefile | Send Bulk Emails CSV
[**EmailsPost**](EmailsAPI.md#EmailsPost) | **Post** /emails | Send Bulk Emails
[**EmailsTransactionalPost**](EmailsAPI.md#EmailsTransactionalPost) | **Post** /emails/transactional | Send Transactional Email



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
	resp, r, err := apiClient.EmailsAPI.EmailsByMsgidViewGet(context.Background(), msgid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.EmailsByMsgidViewGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailsByMsgidViewGet`: EmailData
	fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.EmailsByMsgidViewGet`: %v\n", resp)
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


## EmailsByTransactionidStatusGet

> EmailJobStatus EmailsByTransactionidStatusGet(ctx, transactionid).ShowFailed(showFailed).ShowSent(showSent).ShowDelivered(showDelivered).ShowPending(showPending).ShowOpened(showOpened).ShowClicked(showClicked).ShowAbuse(showAbuse).ShowUnsubscribed(showUnsubscribed).ShowErrors(showErrors).ShowMessageIDs(showMessageIDs).Execute()

Get Status



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
	transactionid := "transactionid_example" // string | Transaction identifier
	showFailed := true // bool | Include Bounced email addresses. (optional) (default to false)
	showSent := true // bool | Include Sent email addresses. (optional) (default to false)
	showDelivered := true // bool | Include all delivered email addresses. (optional) (default to false)
	showPending := true // bool | Include Ready to send email addresses. (optional) (default to false)
	showOpened := true // bool | Include Opened email addresses. (optional) (default to false)
	showClicked := true // bool | Include Clicked email addresses. (optional) (default to false)
	showAbuse := true // bool | Include Reported as abuse email addresses. (optional) (default to false)
	showUnsubscribed := true // bool | Include Unsubscribed email addresses. (optional) (default to false)
	showErrors := true // bool | Include error messages for bounced emails. (optional) (default to false)
	showMessageIDs := true // bool | Include all MessageIDs for this transaction (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.EmailsByTransactionidStatusGet(context.Background(), transactionid).ShowFailed(showFailed).ShowSent(showSent).ShowDelivered(showDelivered).ShowPending(showPending).ShowOpened(showOpened).ShowClicked(showClicked).ShowAbuse(showAbuse).ShowUnsubscribed(showUnsubscribed).ShowErrors(showErrors).ShowMessageIDs(showMessageIDs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.EmailsByTransactionidStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailsByTransactionidStatusGet`: EmailJobStatus
	fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.EmailsByTransactionidStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionid** | **string** | Transaction identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailsByTransactionidStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showFailed** | **bool** | Include Bounced email addresses. | [default to false]
 **showSent** | **bool** | Include Sent email addresses. | [default to false]
 **showDelivered** | **bool** | Include all delivered email addresses. | [default to false]
 **showPending** | **bool** | Include Ready to send email addresses. | [default to false]
 **showOpened** | **bool** | Include Opened email addresses. | [default to false]
 **showClicked** | **bool** | Include Clicked email addresses. | [default to false]
 **showAbuse** | **bool** | Include Reported as abuse email addresses. | [default to false]
 **showUnsubscribed** | **bool** | Include Unsubscribed email addresses. | [default to false]
 **showErrors** | **bool** | Include error messages for bounced emails. | [default to false]
 **showMessageIDs** | **bool** | Include all MessageIDs for this transaction | [default to false]

### Return type

[**EmailJobStatus**](EmailJobStatus.md)

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
	mergeEmailPayload := *openapiclient.NewMergeEmailPayload(*openapiclient.NewMessageAttachment(string(123), "Name_example"), *openapiclient.NewEmailContent("John Doe <email@domain.com>")) // MergeEmailPayload | Email data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.EmailsMergefilePost(context.Background()).MergeEmailPayload(mergeEmailPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.EmailsMergefilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailsMergefilePost`: EmailSend
	fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.EmailsMergefilePost`: %v\n", resp)
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
	emailMessageData := *openapiclient.NewEmailMessageData([]openapiclient.EmailRecipient{*openapiclient.NewEmailRecipient("mail@example.com")}, *openapiclient.NewEmailContent("John Doe <email@domain.com>")) // EmailMessageData | Email data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.EmailsPost(context.Background()).EmailMessageData(emailMessageData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.EmailsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailsPost`: EmailSend
	fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.EmailsPost`: %v\n", resp)
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
	emailTransactionalMessageData := *openapiclient.NewEmailTransactionalMessageData(*openapiclient.NewTransactionalRecipient([]string{"To_example"}), *openapiclient.NewEmailContent("John Doe <email@domain.com>")) // EmailTransactionalMessageData | Email data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.EmailsTransactionalPost(context.Background()).EmailTransactionalMessageData(emailTransactionalMessageData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.EmailsTransactionalPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailsTransactionalPost`: EmailSend
	fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.EmailsTransactionalPost`: %v\n", resp)
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

