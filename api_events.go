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
	"time"
	"reflect"
)


// EventsApiService EventsApi service
type EventsApiService service

type ApiEventsByTransactionidGetRequest struct {
	ctx context.Context
	ApiService *EventsApiService
	transactionid string
	from *time.Time
	to *time.Time
	orderBy *EventsOrderBy
	limit *int32
	offset *int32
}

// Starting date for search in YYYY-MM-DDThh:mm:ss format.
func (r ApiEventsByTransactionidGetRequest) From(from time.Time) ApiEventsByTransactionidGetRequest {
	r.from = &from
	return r
}

// Ending date for search in YYYY-MM-DDThh:mm:ss format.
func (r ApiEventsByTransactionidGetRequest) To(to time.Time) ApiEventsByTransactionidGetRequest {
	r.to = &to
	return r
}

func (r ApiEventsByTransactionidGetRequest) OrderBy(orderBy EventsOrderBy) ApiEventsByTransactionidGetRequest {
	r.orderBy = &orderBy
	return r
}

// Maximum number of returned items.
func (r ApiEventsByTransactionidGetRequest) Limit(limit int32) ApiEventsByTransactionidGetRequest {
	r.limit = &limit
	return r
}

// How many items should be returned ahead.
func (r ApiEventsByTransactionidGetRequest) Offset(offset int32) ApiEventsByTransactionidGetRequest {
	r.offset = &offset
	return r
}

func (r ApiEventsByTransactionidGetRequest) Execute() ([]RecipientEvent, *http.Response, error) {
	return r.ApiService.EventsByTransactionidGetExecute(r)
}

/*
EventsByTransactionidGet Load Email Events

Returns a log of delivery events for the specific transaction ID. Required Access Level: ViewReports

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param transactionid ID number of transaction
 @return ApiEventsByTransactionidGetRequest
*/
func (a *EventsApiService) EventsByTransactionidGet(ctx context.Context, transactionid string) ApiEventsByTransactionidGetRequest {
	return ApiEventsByTransactionidGetRequest{
		ApiService: a,
		ctx: ctx,
		transactionid: transactionid,
	}
}

// Execute executes the request
//  @return []RecipientEvent
func (a *EventsApiService) EventsByTransactionidGetExecute(r ApiEventsByTransactionidGetRequest) ([]RecipientEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RecipientEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.EventsByTransactionidGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/{transactionid}"
	localVarPath = strings.Replace(localVarPath, "{"+"transactionid"+"}", url.PathEscape(parameterValueToString(r.transactionid, "transactionid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.from != nil {
		parameterAddToQuery(localVarQueryParams, "from", r.from, "")
	}
	if r.to != nil {
		parameterAddToQuery(localVarQueryParams, "to", r.to, "")
	}
	if r.orderBy != nil {
		parameterAddToQuery(localVarQueryParams, "orderBy", r.orderBy, "")
	}
	if r.limit != nil {
		parameterAddToQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToQuery(localVarQueryParams, "offset", r.offset, "")
	}
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

type ApiEventsChannelsByNameExportPostRequest struct {
	ctx context.Context
	ApiService *EventsApiService
	name string
	eventTypes *[]EventType
	from *time.Time
	to *time.Time
	fileFormat *ExportFileFormats
	compressionFormat *CompressionFormat
	fileName *string
}

// Types of Events to return
func (r ApiEventsChannelsByNameExportPostRequest) EventTypes(eventTypes []EventType) ApiEventsChannelsByNameExportPostRequest {
	r.eventTypes = &eventTypes
	return r
}

// Starting date for search in YYYY-MM-DDThh:mm:ss format.
func (r ApiEventsChannelsByNameExportPostRequest) From(from time.Time) ApiEventsChannelsByNameExportPostRequest {
	r.from = &from
	return r
}

// Ending date for search in YYYY-MM-DDThh:mm:ss format.
func (r ApiEventsChannelsByNameExportPostRequest) To(to time.Time) ApiEventsChannelsByNameExportPostRequest {
	r.to = &to
	return r
}

// Format of the exported file
func (r ApiEventsChannelsByNameExportPostRequest) FileFormat(fileFormat ExportFileFormats) ApiEventsChannelsByNameExportPostRequest {
	r.fileFormat = &fileFormat
	return r
}

// FileResponse compression format. None or Zip.
func (r ApiEventsChannelsByNameExportPostRequest) CompressionFormat(compressionFormat CompressionFormat) ApiEventsChannelsByNameExportPostRequest {
	r.compressionFormat = &compressionFormat
	return r
}

// Name of your file including extension.
func (r ApiEventsChannelsByNameExportPostRequest) FileName(fileName string) ApiEventsChannelsByNameExportPostRequest {
	r.fileName = &fileName
	return r
}

func (r ApiEventsChannelsByNameExportPostRequest) Execute() (*ExportLink, *http.Response, error) {
	return r.ApiService.EventsChannelsByNameExportPostExecute(r)
}

/*
EventsChannelsByNameExportPost Export Channel Events

Export delivery events log information to the specified file format. Required Access Level: Export

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of selected channel.
 @return ApiEventsChannelsByNameExportPostRequest
*/
func (a *EventsApiService) EventsChannelsByNameExportPost(ctx context.Context, name string) ApiEventsChannelsByNameExportPostRequest {
	return ApiEventsChannelsByNameExportPostRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return ExportLink
func (a *EventsApiService) EventsChannelsByNameExportPostExecute(r ApiEventsChannelsByNameExportPostRequest) (*ExportLink, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExportLink
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.EventsChannelsByNameExportPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/channels/{name}/export"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.eventTypes != nil {
		t := *r.eventTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToQuery(localVarQueryParams, "eventTypes", s.Index(i), "multi")
			}
		} else {
			parameterAddToQuery(localVarQueryParams, "eventTypes", t, "multi")
		}
	}
	if r.from != nil {
		parameterAddToQuery(localVarQueryParams, "from", r.from, "")
	}
	if r.to != nil {
		parameterAddToQuery(localVarQueryParams, "to", r.to, "")
	}
	if r.fileFormat != nil {
		parameterAddToQuery(localVarQueryParams, "fileFormat", r.fileFormat, "")
	}
	if r.compressionFormat != nil {
		parameterAddToQuery(localVarQueryParams, "compressionFormat", r.compressionFormat, "")
	}
	if r.fileName != nil {
		parameterAddToQuery(localVarQueryParams, "fileName", r.fileName, "")
	}
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

type ApiEventsChannelsByNameGetRequest struct {
	ctx context.Context
	ApiService *EventsApiService
	name string
	eventTypes *[]EventType
	from *time.Time
	to *time.Time
	orderBy *EventsOrderBy
	limit *int32
	offset *int32
}

// Types of Events to return
func (r ApiEventsChannelsByNameGetRequest) EventTypes(eventTypes []EventType) ApiEventsChannelsByNameGetRequest {
	r.eventTypes = &eventTypes
	return r
}

// Starting date for search in YYYY-MM-DDThh:mm:ss format.
func (r ApiEventsChannelsByNameGetRequest) From(from time.Time) ApiEventsChannelsByNameGetRequest {
	r.from = &from
	return r
}

// Ending date for search in YYYY-MM-DDThh:mm:ss format.
func (r ApiEventsChannelsByNameGetRequest) To(to time.Time) ApiEventsChannelsByNameGetRequest {
	r.to = &to
	return r
}

func (r ApiEventsChannelsByNameGetRequest) OrderBy(orderBy EventsOrderBy) ApiEventsChannelsByNameGetRequest {
	r.orderBy = &orderBy
	return r
}

// How many items to load. Maximum for this request is 1000 items
func (r ApiEventsChannelsByNameGetRequest) Limit(limit int32) ApiEventsChannelsByNameGetRequest {
	r.limit = &limit
	return r
}

// How many items should be returned ahead.
func (r ApiEventsChannelsByNameGetRequest) Offset(offset int32) ApiEventsChannelsByNameGetRequest {
	r.offset = &offset
	return r
}

func (r ApiEventsChannelsByNameGetRequest) Execute() ([]RecipientEvent, *http.Response, error) {
	return r.ApiService.EventsChannelsByNameGetExecute(r)
}

/*
EventsChannelsByNameGet Load Channel Events

Returns a log of delivery events filtered by specified parameters. Required Access Level: ViewReports

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of selected channel.
 @return ApiEventsChannelsByNameGetRequest
*/
func (a *EventsApiService) EventsChannelsByNameGet(ctx context.Context, name string) ApiEventsChannelsByNameGetRequest {
	return ApiEventsChannelsByNameGetRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return []RecipientEvent
func (a *EventsApiService) EventsChannelsByNameGetExecute(r ApiEventsChannelsByNameGetRequest) ([]RecipientEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RecipientEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.EventsChannelsByNameGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/channels/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.eventTypes != nil {
		t := *r.eventTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToQuery(localVarQueryParams, "eventTypes", s.Index(i), "multi")
			}
		} else {
			parameterAddToQuery(localVarQueryParams, "eventTypes", t, "multi")
		}
	}
	if r.from != nil {
		parameterAddToQuery(localVarQueryParams, "from", r.from, "")
	}
	if r.to != nil {
		parameterAddToQuery(localVarQueryParams, "to", r.to, "")
	}
	if r.orderBy != nil {
		parameterAddToQuery(localVarQueryParams, "orderBy", r.orderBy, "")
	}
	if r.limit != nil {
		parameterAddToQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToQuery(localVarQueryParams, "offset", r.offset, "")
	}
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

type ApiEventsChannelsExportByIdStatusGetRequest struct {
	ctx context.Context
	ApiService *EventsApiService
	id string
}

func (r ApiEventsChannelsExportByIdStatusGetRequest) Execute() (*ExportStatus, *http.Response, error) {
	return r.ApiService.EventsChannelsExportByIdStatusGetExecute(r)
}

/*
EventsChannelsExportByIdStatusGet Check Channel Export Status

Check the current status of the channel export. Required Access Level: Export

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the exported file
 @return ApiEventsChannelsExportByIdStatusGetRequest
*/
func (a *EventsApiService) EventsChannelsExportByIdStatusGet(ctx context.Context, id string) ApiEventsChannelsExportByIdStatusGetRequest {
	return ApiEventsChannelsExportByIdStatusGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ExportStatus
func (a *EventsApiService) EventsChannelsExportByIdStatusGetExecute(r ApiEventsChannelsExportByIdStatusGetRequest) (*ExportStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExportStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.EventsChannelsExportByIdStatusGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/channels/export/{id}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiEventsExportByIdStatusGetRequest struct {
	ctx context.Context
	ApiService *EventsApiService
	id string
}

func (r ApiEventsExportByIdStatusGetRequest) Execute() (*ExportStatus, *http.Response, error) {
	return r.ApiService.EventsExportByIdStatusGetExecute(r)
}

/*
EventsExportByIdStatusGet Check Export Status

Check the current status of the export. Required Access Level: Export

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the exported file
 @return ApiEventsExportByIdStatusGetRequest
*/
func (a *EventsApiService) EventsExportByIdStatusGet(ctx context.Context, id string) ApiEventsExportByIdStatusGetRequest {
	return ApiEventsExportByIdStatusGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ExportStatus
func (a *EventsApiService) EventsExportByIdStatusGetExecute(r ApiEventsExportByIdStatusGetRequest) (*ExportStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExportStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.EventsExportByIdStatusGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/export/{id}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiEventsExportPostRequest struct {
	ctx context.Context
	ApiService *EventsApiService
	eventTypes *[]EventType
	from *time.Time
	to *time.Time
	fileFormat *ExportFileFormats
	compressionFormat *CompressionFormat
	fileName *string
}

// Types of Events to return
func (r ApiEventsExportPostRequest) EventTypes(eventTypes []EventType) ApiEventsExportPostRequest {
	r.eventTypes = &eventTypes
	return r
}

// Starting date for search in YYYY-MM-DDThh:mm:ss format.
func (r ApiEventsExportPostRequest) From(from time.Time) ApiEventsExportPostRequest {
	r.from = &from
	return r
}

// Ending date for search in YYYY-MM-DDThh:mm:ss format.
func (r ApiEventsExportPostRequest) To(to time.Time) ApiEventsExportPostRequest {
	r.to = &to
	return r
}

// Format of the exported file
func (r ApiEventsExportPostRequest) FileFormat(fileFormat ExportFileFormats) ApiEventsExportPostRequest {
	r.fileFormat = &fileFormat
	return r
}

// FileResponse compression format. None or Zip.
func (r ApiEventsExportPostRequest) CompressionFormat(compressionFormat CompressionFormat) ApiEventsExportPostRequest {
	r.compressionFormat = &compressionFormat
	return r
}

// Name of your file including extension.
func (r ApiEventsExportPostRequest) FileName(fileName string) ApiEventsExportPostRequest {
	r.fileName = &fileName
	return r
}

func (r ApiEventsExportPostRequest) Execute() (*ExportLink, *http.Response, error) {
	return r.ApiService.EventsExportPostExecute(r)
}

/*
EventsExportPost Export Events

Export delivery events log information to the specified file format. Required Access Level: Export

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEventsExportPostRequest
*/
func (a *EventsApiService) EventsExportPost(ctx context.Context) ApiEventsExportPostRequest {
	return ApiEventsExportPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ExportLink
func (a *EventsApiService) EventsExportPostExecute(r ApiEventsExportPostRequest) (*ExportLink, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExportLink
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.EventsExportPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/export"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.eventTypes != nil {
		t := *r.eventTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToQuery(localVarQueryParams, "eventTypes", s.Index(i), "multi")
			}
		} else {
			parameterAddToQuery(localVarQueryParams, "eventTypes", t, "multi")
		}
	}
	if r.from != nil {
		parameterAddToQuery(localVarQueryParams, "from", r.from, "")
	}
	if r.to != nil {
		parameterAddToQuery(localVarQueryParams, "to", r.to, "")
	}
	if r.fileFormat != nil {
		parameterAddToQuery(localVarQueryParams, "fileFormat", r.fileFormat, "")
	}
	if r.compressionFormat != nil {
		parameterAddToQuery(localVarQueryParams, "compressionFormat", r.compressionFormat, "")
	}
	if r.fileName != nil {
		parameterAddToQuery(localVarQueryParams, "fileName", r.fileName, "")
	}
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

type ApiEventsGetRequest struct {
	ctx context.Context
	ApiService *EventsApiService
	eventTypes *[]EventType
	from *time.Time
	to *time.Time
	orderBy *EventsOrderBy
	limit *int32
	offset *int32
}

// Types of Events to return
func (r ApiEventsGetRequest) EventTypes(eventTypes []EventType) ApiEventsGetRequest {
	r.eventTypes = &eventTypes
	return r
}

// Starting date for search in YYYY-MM-DDThh:mm:ss format.
func (r ApiEventsGetRequest) From(from time.Time) ApiEventsGetRequest {
	r.from = &from
	return r
}

// Ending date for search in YYYY-MM-DDThh:mm:ss format.
func (r ApiEventsGetRequest) To(to time.Time) ApiEventsGetRequest {
	r.to = &to
	return r
}

func (r ApiEventsGetRequest) OrderBy(orderBy EventsOrderBy) ApiEventsGetRequest {
	r.orderBy = &orderBy
	return r
}

// How many items to load. Maximum for this request is 1000 items
func (r ApiEventsGetRequest) Limit(limit int32) ApiEventsGetRequest {
	r.limit = &limit
	return r
}

// How many items should be returned ahead.
func (r ApiEventsGetRequest) Offset(offset int32) ApiEventsGetRequest {
	r.offset = &offset
	return r
}

func (r ApiEventsGetRequest) Execute() ([]RecipientEvent, *http.Response, error) {
	return r.ApiService.EventsGetExecute(r)
}

/*
EventsGet Load Events

Returns a log of delivery events filtered by specified parameters. Required Access Level: ViewReports

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEventsGetRequest
*/
func (a *EventsApiService) EventsGet(ctx context.Context) ApiEventsGetRequest {
	return ApiEventsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []RecipientEvent
func (a *EventsApiService) EventsGetExecute(r ApiEventsGetRequest) ([]RecipientEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RecipientEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.EventsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.eventTypes != nil {
		t := *r.eventTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToQuery(localVarQueryParams, "eventTypes", s.Index(i), "multi")
			}
		} else {
			parameterAddToQuery(localVarQueryParams, "eventTypes", t, "multi")
		}
	}
	if r.from != nil {
		parameterAddToQuery(localVarQueryParams, "from", r.from, "")
	}
	if r.to != nil {
		parameterAddToQuery(localVarQueryParams, "to", r.to, "")
	}
	if r.orderBy != nil {
		parameterAddToQuery(localVarQueryParams, "orderBy", r.orderBy, "")
	}
	if r.limit != nil {
		parameterAddToQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToQuery(localVarQueryParams, "offset", r.offset, "")
	}
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
