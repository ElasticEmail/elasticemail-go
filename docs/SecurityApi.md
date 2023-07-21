# \SecurityApi

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityApikeysByNameDelete**](SecurityApi.md#SecurityApikeysByNameDelete) | **Delete** /security/apikeys/{name} | Delete ApiKey
[**SecurityApikeysByNameGet**](SecurityApi.md#SecurityApikeysByNameGet) | **Get** /security/apikeys/{name} | Load ApiKey
[**SecurityApikeysByNamePut**](SecurityApi.md#SecurityApikeysByNamePut) | **Put** /security/apikeys/{name} | Update ApiKey
[**SecurityApikeysGet**](SecurityApi.md#SecurityApikeysGet) | **Get** /security/apikeys | List ApiKeys
[**SecurityApikeysPost**](SecurityApi.md#SecurityApikeysPost) | **Post** /security/apikeys | Add ApiKey
[**SecuritySmtpByNameDelete**](SecurityApi.md#SecuritySmtpByNameDelete) | **Delete** /security/smtp/{name} | Delete SMTP Credential
[**SecuritySmtpByNameGet**](SecurityApi.md#SecuritySmtpByNameGet) | **Get** /security/smtp/{name} | Load SMTP Credential
[**SecuritySmtpByNamePut**](SecurityApi.md#SecuritySmtpByNamePut) | **Put** /security/smtp/{name} | Update SMTP Credential
[**SecuritySmtpGet**](SecurityApi.md#SecuritySmtpGet) | **Get** /security/smtp | List SMTP Credentials
[**SecuritySmtpPost**](SecurityApi.md#SecuritySmtpPost) | **Post** /security/smtp | Add SMTP Credential



## SecurityApikeysByNameDelete

> SecurityApikeysByNameDelete(ctx, name).Subaccount(subaccount).Execute()

Delete ApiKey



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
    name := "name_example" // string | Name of the ApiKey
    subaccount := "subaccount_example" // string | Email of the subaccount of which ApiKey should be deleted (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecurityApi.SecurityApikeysByNameDelete(context.Background(), name).Subaccount(subaccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityApikeysByNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the ApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityApikeysByNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subaccount** | **string** | Email of the subaccount of which ApiKey should be deleted | 

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


## SecurityApikeysByNameGet

> ApiKey SecurityApikeysByNameGet(ctx, name).Subaccount(subaccount).Execute()

Load ApiKey



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
    name := "name_example" // string | Name of the ApiKey
    subaccount := "subaccount_example" // string | Email of the subaccount of which ApiKey should be loaded (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecurityApikeysByNameGet(context.Background(), name).Subaccount(subaccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityApikeysByNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityApikeysByNameGet`: ApiKey
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecurityApikeysByNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the ApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityApikeysByNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subaccount** | **string** | Email of the subaccount of which ApiKey should be loaded | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityApikeysByNamePut

> ApiKey SecurityApikeysByNamePut(ctx, name).ApiKeyPayload(apiKeyPayload).Execute()

Update ApiKey



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
    name := "name_example" // string | Name of the ApiKey
    apiKeyPayload := *openapiclient.NewApiKeyPayload("Name_example", []openapiclient.AccessLevel{openapiclient.AccessLevel("None")}) // ApiKeyPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecurityApikeysByNamePut(context.Background(), name).ApiKeyPayload(apiKeyPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityApikeysByNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityApikeysByNamePut`: ApiKey
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecurityApikeysByNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the ApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityApikeysByNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKeyPayload** | [**ApiKeyPayload**](ApiKeyPayload.md) |  | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityApikeysGet

> []ApiKey SecurityApikeysGet(ctx).Subaccount(subaccount).Execute()

List ApiKeys



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
    subaccount := "subaccount_example" // string | Email of the subaccount of which ApiKeys should be loaded (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecurityApikeysGet(context.Background()).Subaccount(subaccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityApikeysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityApikeysGet`: []ApiKey
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecurityApikeysGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityApikeysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subaccount** | **string** | Email of the subaccount of which ApiKeys should be loaded | 

### Return type

[**[]ApiKey**](ApiKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityApikeysPost

> ApiKeyNew SecurityApikeysPost(ctx).ApiKeyPayload(apiKeyPayload).Execute()

Add ApiKey



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
    apiKeyPayload := *openapiclient.NewApiKeyPayload("Name_example", []openapiclient.AccessLevel{openapiclient.AccessLevel("None")}) // ApiKeyPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecurityApikeysPost(context.Background()).ApiKeyPayload(apiKeyPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityApikeysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityApikeysPost`: ApiKeyNew
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecurityApikeysPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityApikeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKeyPayload** | [**ApiKeyPayload**](ApiKeyPayload.md) |  | 

### Return type

[**ApiKeyNew**](ApiKeyNew.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySmtpByNameDelete

> SecuritySmtpByNameDelete(ctx, name).Subaccount(subaccount).Execute()

Delete SMTP Credential



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
    name := "name_example" // string | Name of the SMTP Credential
    subaccount := "subaccount_example" // string | Email of the subaccount of which credential should be deleted (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecurityApi.SecuritySmtpByNameDelete(context.Background(), name).Subaccount(subaccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecuritySmtpByNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the SMTP Credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySmtpByNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subaccount** | **string** | Email of the subaccount of which credential should be deleted | 

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


## SecuritySmtpByNameGet

> SmtpCredentials SecuritySmtpByNameGet(ctx, name).Subaccount(subaccount).Execute()

Load SMTP Credential



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
    name := "name_example" // string | Name of the SMTP Credential
    subaccount := "subaccount_example" // string | Email of the subaccount of which credential should be loaded (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecuritySmtpByNameGet(context.Background(), name).Subaccount(subaccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecuritySmtpByNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecuritySmtpByNameGet`: SmtpCredentials
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecuritySmtpByNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the SMTP Credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySmtpByNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subaccount** | **string** | Email of the subaccount of which credential should be loaded | 

### Return type

[**SmtpCredentials**](SmtpCredentials.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySmtpByNamePut

> SmtpCredentials SecuritySmtpByNamePut(ctx, name).SmtpCredentialsPayload(smtpCredentialsPayload).Execute()

Update SMTP Credential



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
    name := "name_example" // string | Name of the SMTP Credential
    smtpCredentialsPayload := *openapiclient.NewSmtpCredentialsPayload("Name_example") // SmtpCredentialsPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecuritySmtpByNamePut(context.Background(), name).SmtpCredentialsPayload(smtpCredentialsPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecuritySmtpByNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecuritySmtpByNamePut`: SmtpCredentials
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecuritySmtpByNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the SMTP Credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySmtpByNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smtpCredentialsPayload** | [**SmtpCredentialsPayload**](SmtpCredentialsPayload.md) |  | 

### Return type

[**SmtpCredentials**](SmtpCredentials.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySmtpGet

> []SmtpCredentials SecuritySmtpGet(ctx).Subaccount(subaccount).Execute()

List SMTP Credentials



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
    subaccount := "subaccount_example" // string | Email of the subaccount of which credentials should be listed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecuritySmtpGet(context.Background()).Subaccount(subaccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecuritySmtpGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecuritySmtpGet`: []SmtpCredentials
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecuritySmtpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySmtpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subaccount** | **string** | Email of the subaccount of which credentials should be listed | 

### Return type

[**[]SmtpCredentials**](SmtpCredentials.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySmtpPost

> SmtpCredentialsNew SecuritySmtpPost(ctx).SmtpCredentialsPayload(smtpCredentialsPayload).Execute()

Add SMTP Credential



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
    smtpCredentialsPayload := *openapiclient.NewSmtpCredentialsPayload("Name_example") // SmtpCredentialsPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecuritySmtpPost(context.Background()).SmtpCredentialsPayload(smtpCredentialsPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecuritySmtpPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecuritySmtpPost`: SmtpCredentialsNew
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecuritySmtpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySmtpPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smtpCredentialsPayload** | [**SmtpCredentialsPayload**](SmtpCredentialsPayload.md) |  | 

### Return type

[**SmtpCredentialsNew**](SmtpCredentialsNew.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

