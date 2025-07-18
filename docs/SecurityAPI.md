# \SecurityAPI

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityApikeysByNameDelete**](SecurityAPI.md#SecurityApikeysByNameDelete) | **Delete** /security/apikeys/{name} | Delete ApiKey
[**SecurityApikeysByNameGet**](SecurityAPI.md#SecurityApikeysByNameGet) | **Get** /security/apikeys/{name} | Load ApiKey
[**SecurityApikeysByNamePut**](SecurityAPI.md#SecurityApikeysByNamePut) | **Put** /security/apikeys/{name} | Update ApiKey
[**SecurityApikeysGet**](SecurityAPI.md#SecurityApikeysGet) | **Get** /security/apikeys | List ApiKeys
[**SecurityApikeysPost**](SecurityAPI.md#SecurityApikeysPost) | **Post** /security/apikeys | Add ApiKey
[**SecuritySmtpByNameDelete**](SecurityAPI.md#SecuritySmtpByNameDelete) | **Delete** /security/smtp/{name} | Delete SMTP Credential
[**SecuritySmtpByNameGet**](SecurityAPI.md#SecuritySmtpByNameGet) | **Get** /security/smtp/{name} | Load SMTP Credential
[**SecuritySmtpByNamePut**](SecurityAPI.md#SecuritySmtpByNamePut) | **Put** /security/smtp/{name} | Update SMTP Credential
[**SecuritySmtpGet**](SecurityAPI.md#SecuritySmtpGet) | **Get** /security/smtp | List SMTP Credentials
[**SecuritySmtpPost**](SecurityAPI.md#SecuritySmtpPost) | **Post** /security/smtp | Add SMTP Credential



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
	openapiclient "github.com/radzawee/elasticemail-go"
)

func main() {
	name := "name_example" // string | Name of the ApiKey
	subaccount := "subaccount_example" // string | Email of the subaccount of which ApiKey should be deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.SecurityApikeysByNameDelete(context.Background(), name).Subaccount(subaccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.SecurityApikeysByNameDelete``: %v\n", err)
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
	openapiclient "github.com/radzawee/elasticemail-go"
)

func main() {
	name := "name_example" // string | Name of the ApiKey
	subaccount := "subaccount_example" // string | Email of the subaccount of which ApiKey should be loaded (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.SecurityApikeysByNameGet(context.Background(), name).Subaccount(subaccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.SecurityApikeysByNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityApikeysByNameGet`: ApiKey
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.SecurityApikeysByNameGet`: %v\n", resp)
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
	openapiclient "github.com/radzawee/elasticemail-go"
)

func main() {
	name := "name_example" // string | Name of the ApiKey
	apiKeyPayload := *openapiclient.NewApiKeyPayload("Name_example", []openapiclient.AccessLevel{openapiclient.AccessLevel("None")}) // ApiKeyPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.SecurityApikeysByNamePut(context.Background(), name).ApiKeyPayload(apiKeyPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.SecurityApikeysByNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityApikeysByNamePut`: ApiKey
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.SecurityApikeysByNamePut`: %v\n", resp)
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
	openapiclient "github.com/radzawee/elasticemail-go"
)

func main() {
	subaccount := "subaccount_example" // string | Email of the subaccount of which ApiKeys should be loaded (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.SecurityApikeysGet(context.Background()).Subaccount(subaccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.SecurityApikeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityApikeysGet`: []ApiKey
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.SecurityApikeysGet`: %v\n", resp)
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

> NewApiKey SecurityApikeysPost(ctx).ApiKeyPayload(apiKeyPayload).Execute()

Add ApiKey



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
	apiKeyPayload := *openapiclient.NewApiKeyPayload("Name_example", []openapiclient.AccessLevel{openapiclient.AccessLevel("None")}) // ApiKeyPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.SecurityApikeysPost(context.Background()).ApiKeyPayload(apiKeyPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.SecurityApikeysPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityApikeysPost`: NewApiKey
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.SecurityApikeysPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityApikeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKeyPayload** | [**ApiKeyPayload**](ApiKeyPayload.md) |  | 

### Return type

[**NewApiKey**](NewApiKey.md)

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
	openapiclient "github.com/radzawee/elasticemail-go"
)

func main() {
	name := "name_example" // string | Name of the SMTP Credential
	subaccount := "subaccount_example" // string | Email of the subaccount of which credential should be deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.SecuritySmtpByNameDelete(context.Background(), name).Subaccount(subaccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.SecuritySmtpByNameDelete``: %v\n", err)
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
	openapiclient "github.com/radzawee/elasticemail-go"
)

func main() {
	name := "name_example" // string | Name of the SMTP Credential
	subaccount := "subaccount_example" // string | Email of the subaccount of which credential should be loaded (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.SecuritySmtpByNameGet(context.Background(), name).Subaccount(subaccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.SecuritySmtpByNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecuritySmtpByNameGet`: SmtpCredentials
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.SecuritySmtpByNameGet`: %v\n", resp)
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
	openapiclient "github.com/radzawee/elasticemail-go"
)

func main() {
	name := "name_example" // string | Name of the SMTP Credential
	smtpCredentialsPayload := *openapiclient.NewSmtpCredentialsPayload("Name_example") // SmtpCredentialsPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.SecuritySmtpByNamePut(context.Background(), name).SmtpCredentialsPayload(smtpCredentialsPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.SecuritySmtpByNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecuritySmtpByNamePut`: SmtpCredentials
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.SecuritySmtpByNamePut`: %v\n", resp)
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
	openapiclient "github.com/radzawee/elasticemail-go"
)

func main() {
	subaccount := "subaccount_example" // string | Email of the subaccount of which credentials should be listed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.SecuritySmtpGet(context.Background()).Subaccount(subaccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.SecuritySmtpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecuritySmtpGet`: []SmtpCredentials
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.SecuritySmtpGet`: %v\n", resp)
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

> NewSmtpCredentials SecuritySmtpPost(ctx).SmtpCredentialsPayload(smtpCredentialsPayload).Execute()

Add SMTP Credential



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
	smtpCredentialsPayload := *openapiclient.NewSmtpCredentialsPayload("Name_example") // SmtpCredentialsPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.SecuritySmtpPost(context.Background()).SmtpCredentialsPayload(smtpCredentialsPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.SecuritySmtpPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecuritySmtpPost`: NewSmtpCredentials
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.SecuritySmtpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySmtpPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smtpCredentialsPayload** | [**SmtpCredentialsPayload**](SmtpCredentialsPayload.md) |  | 

### Return type

[**NewSmtpCredentials**](NewSmtpCredentials.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

