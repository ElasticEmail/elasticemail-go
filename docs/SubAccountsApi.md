# \SubAccountsApi

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubaccountsByEmailCreditsPatch**](SubAccountsApi.md#SubaccountsByEmailCreditsPatch) | **Patch** /subaccounts/{email}/credits | Add, Subtract Email Credits
[**SubaccountsByEmailDelete**](SubAccountsApi.md#SubaccountsByEmailDelete) | **Delete** /subaccounts/{email} | Delete SubAccount
[**SubaccountsByEmailGet**](SubAccountsApi.md#SubaccountsByEmailGet) | **Get** /subaccounts/{email} | Load SubAccount
[**SubaccountsByEmailSettingsEmailPut**](SubAccountsApi.md#SubaccountsByEmailSettingsEmailPut) | **Put** /subaccounts/{email}/settings/email | Update SubAccount Email Settings
[**SubaccountsGet**](SubAccountsApi.md#SubaccountsGet) | **Get** /subaccounts | Load SubAccounts
[**SubaccountsPost**](SubAccountsApi.md#SubaccountsPost) | **Post** /subaccounts | Add SubAccount



## SubaccountsByEmailCreditsPatch

> SubaccountsByEmailCreditsPatch(ctx, email).SubaccountEmailCreditsPayload(subaccountEmailCreditsPayload).Execute()

Add, Subtract Email Credits



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
    email := "mail@example.com" // string | Email address of Sub-Account
    subaccountEmailCreditsPayload := *openapiclient.NewSubaccountEmailCreditsPayload(int32(123)) // SubaccountEmailCreditsPayload | Amount of email credits to add or subtract from the current SubAccount email credits pool (positive or negative value)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubAccountsApi.SubaccountsByEmailCreditsPatch(context.Background(), email).SubaccountEmailCreditsPayload(subaccountEmailCreditsPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountsApi.SubaccountsByEmailCreditsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Email address of Sub-Account | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountsByEmailCreditsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subaccountEmailCreditsPayload** | [**SubaccountEmailCreditsPayload**](SubaccountEmailCreditsPayload.md) | Amount of email credits to add or subtract from the current SubAccount email credits pool (positive or negative value) | 

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


## SubaccountsByEmailDelete

> SubaccountsByEmailDelete(ctx, email).Execute()

Delete SubAccount



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
    email := "mail@example.com" // string | Email address of Sub-Account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubAccountsApi.SubaccountsByEmailDelete(context.Background(), email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountsApi.SubaccountsByEmailDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Email address of Sub-Account | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountsByEmailDeleteRequest struct via the builder pattern


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


## SubaccountsByEmailGet

> SubAccountInfo SubaccountsByEmailGet(ctx, email).Execute()

Load SubAccount



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
    email := "mail@example.com" // string | Email address of Sub-Account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubAccountsApi.SubaccountsByEmailGet(context.Background(), email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountsApi.SubaccountsByEmailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountsByEmailGet`: SubAccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SubAccountsApi.SubaccountsByEmailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Email address of Sub-Account | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountsByEmailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubAccountInfo**](SubAccountInfo.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountsByEmailSettingsEmailPut

> SubaccountEmailSettings SubaccountsByEmailSettingsEmailPut(ctx, email).SubaccountEmailSettings(subaccountEmailSettings).Execute()

Update SubAccount Email Settings



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
    email := "email_example" // string | 
    subaccountEmailSettings := *openapiclient.NewSubaccountEmailSettings() // SubaccountEmailSettings | Updated Email Settings

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubAccountsApi.SubaccountsByEmailSettingsEmailPut(context.Background(), email).SubaccountEmailSettings(subaccountEmailSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountsApi.SubaccountsByEmailSettingsEmailPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountsByEmailSettingsEmailPut`: SubaccountEmailSettings
    fmt.Fprintf(os.Stdout, "Response from `SubAccountsApi.SubaccountsByEmailSettingsEmailPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountsByEmailSettingsEmailPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subaccountEmailSettings** | [**SubaccountEmailSettings**](SubaccountEmailSettings.md) | Updated Email Settings | 

### Return type

[**SubaccountEmailSettings**](SubaccountEmailSettings.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountsGet

> []SubAccountInfo SubaccountsGet(ctx).Limit(limit).Offset(offset).Execute()

Load SubAccounts



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
    resp, r, err := apiClient.SubAccountsApi.SubaccountsGet(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountsApi.SubaccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountsGet`: []SubAccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SubAccountsApi.SubaccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]SubAccountInfo**](SubAccountInfo.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountsPost

> SubAccountInfo SubaccountsPost(ctx).SubaccountPayload(subaccountPayload).Execute()

Add SubAccount



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
    subaccountPayload := *openapiclient.NewSubaccountPayload("mail@example.com", "********") // SubaccountPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubAccountsApi.SubaccountsPost(context.Background()).SubaccountPayload(subaccountPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountsApi.SubaccountsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountsPost`: SubAccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SubAccountsApi.SubaccountsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subaccountPayload** | [**SubaccountPayload**](SubaccountPayload.md) |  | 

### Return type

[**SubAccountInfo**](SubAccountInfo.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

