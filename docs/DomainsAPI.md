# \DomainsAPI

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsByDomainDelete**](DomainsAPI.md#DomainsByDomainDelete) | **Delete** /domains/{domain} | Delete Domain
[**DomainsByDomainGet**](DomainsAPI.md#DomainsByDomainGet) | **Get** /domains/{domain} | Load Domain
[**DomainsByDomainPut**](DomainsAPI.md#DomainsByDomainPut) | **Put** /domains/{domain} | Update Domain
[**DomainsByDomainRestrictedGet**](DomainsAPI.md#DomainsByDomainRestrictedGet) | **Get** /domains/{domain}/restricted | Check for domain restriction
[**DomainsByDomainVerificationPut**](DomainsAPI.md#DomainsByDomainVerificationPut) | **Put** /domains/{domain}/verification | Verify Domain
[**DomainsByEmailDefaultPatch**](DomainsAPI.md#DomainsByEmailDefaultPatch) | **Patch** /domains/{email}/default | Set Default
[**DomainsGet**](DomainsAPI.md#DomainsGet) | **Get** /domains | Load Domains
[**DomainsPost**](DomainsAPI.md#DomainsPost) | **Post** /domains | Add Domain



## DomainsByDomainDelete

> DomainsByDomainDelete(ctx, domain).Execute()

Delete Domain



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
	domain := "domain_example" // string | Name of the given domain

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DomainsAPI.DomainsByDomainDelete(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsByDomainDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Name of the given domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsByDomainDeleteRequest struct via the builder pattern


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


## DomainsByDomainGet

> DomainData DomainsByDomainGet(ctx, domain).Execute()

Load Domain



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
	domain := "domain_example" // string | Name of the given domain

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsByDomainGet(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsByDomainGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsByDomainGet`: DomainData
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsByDomainGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Name of the given domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsByDomainGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainData**](DomainData.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsByDomainPut

> DomainDetail DomainsByDomainPut(ctx, domain).DomainUpdatePayload(domainUpdatePayload).Execute()

Update Domain



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
	domain := "domain_example" // string | Name of the given domain
	domainUpdatePayload := *openapiclient.NewDomainUpdatePayload() // DomainUpdatePayload | Updated Domain resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsByDomainPut(context.Background(), domain).DomainUpdatePayload(domainUpdatePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsByDomainPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsByDomainPut`: DomainDetail
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsByDomainPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Name of the given domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsByDomainPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainUpdatePayload** | [**DomainUpdatePayload**](DomainUpdatePayload.md) | Updated Domain resource | 

### Return type

[**DomainDetail**](DomainDetail.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsByDomainRestrictedGet

> bool DomainsByDomainRestrictedGet(ctx, domain).Execute()

Check for domain restriction



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
	domain := "domain_example" // string | Name of the given domain

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsByDomainRestrictedGet(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsByDomainRestrictedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsByDomainRestrictedGet`: bool
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsByDomainRestrictedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Name of the given domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsByDomainRestrictedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsByDomainVerificationPut

> DomainData DomainsByDomainVerificationPut(ctx, domain).Body(body).Execute()

Verify Domain



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
	domain := "domain_example" // string | Name of the given domain
	body := "body_example" // string | Tracking type used in the Tracking verification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsByDomainVerificationPut(context.Background(), domain).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsByDomainVerificationPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsByDomainVerificationPut`: DomainData
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsByDomainVerificationPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Name of the given domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsByDomainVerificationPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Tracking type used in the Tracking verification | 

### Return type

[**DomainData**](DomainData.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsByEmailDefaultPatch

> DomainDetail DomainsByEmailDefaultPatch(ctx, email).Execute()

Set Default



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
	email := "email_example" // string | Default email sender, example: mail@yourdomain.com

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsByEmailDefaultPatch(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsByEmailDefaultPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsByEmailDefaultPatch`: DomainDetail
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsByEmailDefaultPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Default email sender, example: mail@yourdomain.com | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsByEmailDefaultPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainDetail**](DomainDetail.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsGet

> []DomainDetail DomainsGet(ctx).Execute()

Load Domains



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsGet`: []DomainDetail
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsGetRequest struct via the builder pattern


### Return type

[**[]DomainDetail**](DomainDetail.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsPost

> DomainDetail DomainsPost(ctx).DomainPayload(domainPayload).Execute()

Add Domain



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
	domainPayload := *openapiclient.NewDomainPayload() // DomainPayload | Domain to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsPost(context.Background()).DomainPayload(domainPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsPost`: DomainDetail
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainPayload** | [**DomainPayload**](DomainPayload.md) | Domain to add | 

### Return type

[**DomainDetail**](DomainDetail.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

