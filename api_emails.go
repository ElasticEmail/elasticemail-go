/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://app.elasticemail.com/marketing/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// EmailsApiService EmailsApi service
type EmailsApiService service

type ApiEmailsByMsgidViewGetRequest struct {
	ctx context.Context
	ApiService *EmailsApiService
	msgid string
}

func (r ApiEmailsByMsgidViewGetRequest) Execute() (*EmailData, *http.Response, error) {
	return r.ApiService.EmailsByMsgidViewGetExecute(r)
}

/*
EmailsByMsgidViewGet View Email

Returns email details for viewing or rendering. Required Access Level: None

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param msgid Message identifier
 @return ApiEmailsByMsgidViewGetRequest
*/
func (a *EmailsApiService) EmailsByMsgidViewGet(ctx context.Context, msgid string) ApiEmailsByMsgidViewGetRequest {
	return ApiEmailsByMsgidViewGetRequest{
		ApiService: a,
		ctx: ctx,
		msgid: msgid,
	}
}

// Execute executes the request
//  @return EmailData
func (a *EmailsApiService) EmailsByMsgidViewGetExecute(r ApiEmailsByMsgidViewGetRequest) (*EmailData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmailData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailsApiService.EmailsByMsgidViewGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/emails/{msgid}/view"
	localVarPath = strings.Replace(localVarPath, "{"+"msgid"+"}", url.PathEscape(parameterValueToString(r.msgid, "msgid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-ElasticEmail-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEmailsMergefilePostRequest struct {
	ctx context.Context
	ApiService *EmailsApiService
	mergeEmailPayload *MergeEmailPayload
}

// Email data
func (r ApiEmailsMergefilePostRequest) MergeEmailPayload(mergeEmailPayload MergeEmailPayload) ApiEmailsMergefilePostRequest {
	r.mergeEmailPayload = &mergeEmailPayload
	return r
}

func (r ApiEmailsMergefilePostRequest) Execute() (*EmailSend, *http.Response, error) {
	return r.ApiService.EmailsMergefilePostExecute(r)
}

/*
EmailsMergefilePost Send Bulk Emails CSV

Send bulk merge email. Required Access Level: SendHttp

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEmailsMergefilePostRequest
*/
func (a *EmailsApiService) EmailsMergefilePost(ctx context.Context) ApiEmailsMergefilePostRequest {
	return ApiEmailsMergefilePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmailSend
func (a *EmailsApiService) EmailsMergefilePostExecute(r ApiEmailsMergefilePostRequest) (*EmailSend, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmailSend
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailsApiService.EmailsMergefilePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/emails/mergefile"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mergeEmailPayload == nil {
		return localVarReturnValue, nil, reportError("mergeEmailPayload is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.mergeEmailPayload
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-ElasticEmail-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEmailsPostRequest struct {
	ctx context.Context
	ApiService *EmailsApiService
	emailMessageData *EmailMessageData
}

// Email data
func (r ApiEmailsPostRequest) EmailMessageData(emailMessageData EmailMessageData) ApiEmailsPostRequest {
	r.emailMessageData = &emailMessageData
	return r
}

func (r ApiEmailsPostRequest) Execute() (*EmailSend, *http.Response, error) {
	return r.ApiService.EmailsPostExecute(r)
}

/*
EmailsPost Send Bulk Emails

Send bulk merge email. Required Access Level: SendHttp

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEmailsPostRequest
*/
func (a *EmailsApiService) EmailsPost(ctx context.Context) ApiEmailsPostRequest {
	return ApiEmailsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmailSend
func (a *EmailsApiService) EmailsPostExecute(r ApiEmailsPostRequest) (*EmailSend, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmailSend
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailsApiService.EmailsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/emails"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.emailMessageData == nil {
		return localVarReturnValue, nil, reportError("emailMessageData is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.emailMessageData
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-ElasticEmail-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEmailsTransactionalPostRequest struct {
	ctx context.Context
	ApiService *EmailsApiService
	emailTransactionalMessageData *EmailTransactionalMessageData
}

// Email data
func (r ApiEmailsTransactionalPostRequest) EmailTransactionalMessageData(emailTransactionalMessageData EmailTransactionalMessageData) ApiEmailsTransactionalPostRequest {
	r.emailTransactionalMessageData = &emailTransactionalMessageData
	return r
}

func (r ApiEmailsTransactionalPostRequest) Execute() (*EmailSend, *http.Response, error) {
	return r.ApiService.EmailsTransactionalPostExecute(r)
}

/*
EmailsTransactionalPost Send Transactional Email

Send transactional emails (recipients will be known to each other). Required Access Level: SendHttp

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEmailsTransactionalPostRequest
*/
func (a *EmailsApiService) EmailsTransactionalPost(ctx context.Context) ApiEmailsTransactionalPostRequest {
	return ApiEmailsTransactionalPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmailSend
func (a *EmailsApiService) EmailsTransactionalPostExecute(r ApiEmailsTransactionalPostRequest) (*EmailSend, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmailSend
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailsApiService.EmailsTransactionalPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/emails/transactional"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.emailTransactionalMessageData == nil {
		return localVarReturnValue, nil, reportError("emailTransactionalMessageData is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.emailTransactionalMessageData
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-ElasticEmail-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
