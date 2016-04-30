# \PipelineApi

All URIs are relative to *https://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePipelineExecution**](PipelineApi.md#DeletePipelineExecution) | **Delete** /pipelines/executions/{execution_id} | Delete the specified build.
[**GetPipelineEvent**](PipelineApi.md#GetPipelineEvent) | **Get** /pipelines/events/{event_id} | Get the specified pipeline event.
[**GetPipelineEvents**](PipelineApi.md#GetPipelineEvents) | **Get** /pipelines/events | List pipeline events, optionally filtering by Build id.
[**GetPipelineExecution**](PipelineApi.md#GetPipelineExecution) | **Get** /pipelines/executions/{execution_id} | Gets the specified build.
[**GetPipelineExecutions**](PipelineApi.md#GetPipelineExecutions) | **Get** /pipelines/executions | List executions, optionally filtering by project_id.
[**PipelineEventOccurred**](PipelineApi.md#PipelineEventOccurred) | **Post** /pipelines/events | Record a PipelineEvent.
[**TriggerPipelineExecution**](PipelineApi.md#TriggerPipelineExecution) | **Post** /pipelines/triggers | Trigger execution of a pipeline(s).


# **DeletePipelineExecution**
> DeletePipelineExecution($executionId)

Delete the specified build.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionId** | **int32**| Build id. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineEvent**
> PipelineEvent GetPipelineEvent($eventId)

Get the specified pipeline event.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventId** | **int32**| PipelineEvent id. | 

### Return type

[**PipelineEvent**](PipelineEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineEvents**
> []PipelineEvent GetPipelineEvents($executionId)

List pipeline events, optionally filtering by Build id.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionId** | **int32**| execution id. | 

### Return type

[**[]PipelineEvent**](PipelineEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineExecution**
> PipelineExecution GetPipelineExecution($executionId)

Gets the specified build.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionId** | **int32**| Build id. | 

### Return type

[**PipelineExecution**](PipelineExecution.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineExecutions**
> []PipelineExecution GetPipelineExecutions($projectId, $limit, $offset)

List executions, optionally filtering by project_id.

Return an array of builds, typically filtered by project_id, and often paged.\nFor example `/pipelines/executions?project_id=2&limit=25&offset=50` return an\narray of builds for project with id `2`, will return a maximum of `25` items,\nstarting with item `50`.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32**| Project id. | 
 **limit** | **int32**| Maximum number of items to return. | 
 **offset** | **int32**| Returned array of items should begin at this offset. | 

### Return type

[**[]PipelineExecution**](PipelineExecution.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PipelineEventOccurred**
> PipelineEvent PipelineEventOccurred($executionId, $pipelineEvent)

Record a PipelineEvent.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionId** | **int32**| execution id. | 
 **pipelineEvent** | [**PipelineEvent**](PipelineEvent.md)| Pipeline event. | 

### Return type

[**PipelineEvent**](PipelineEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerPipelineExecution**
> PipelineTrigger TriggerPipelineExecution($buildTrigger)

Trigger execution of a pipeline(s).


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildTrigger** | [**PipelineTrigger**](PipelineTrigger.md)| PipelineTrigger event. | 

### Return type

[**PipelineTrigger**](PipelineTrigger.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

