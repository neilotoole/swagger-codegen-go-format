# \ProjectApi

All URIs are relative to *https://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMember**](ProjectApi.md#AddMember) | **Post** /projects/{project_id}/members | Add a member (user) to a project.
[**AddPipelineTask**](ProjectApi.md#AddPipelineTask) | **Post** /pipelines/tasks | Add a pipeline task for a specific project.
[**CreateProject**](ProjectApi.md#CreateProject) | **Post** /projects | Create a new project.
[**DeleteProject**](ProjectApi.md#DeleteProject) | **Delete** /projects/{project_id} | Delete the specified &#x60;Project&#x60; instance.\n
[**GetPipelineTask**](ProjectApi.md#GetPipelineTask) | **Get** /pipelines/tasks/{task_id} | Get the specified pipeline task.
[**GetPipelineTasks**](ProjectApi.md#GetPipelineTasks) | **Get** /pipelines/tasks | List the pipeline tasks, optionally filtering by project.
[**GetProject**](ProjectApi.md#GetProject) | **Get** /projects/{project_id} | Get the specified instance.\n
[**GetProjectMembers**](ProjectApi.md#GetProjectMembers) | **Get** /projects/{project_id}/members | Get the project members.
[**GetProjects**](ProjectApi.md#GetProjects) | **Get** /projects | List projects, optionally filtering.
[**RemoveMember**](ProjectApi.md#RemoveMember) | **Delete** /projects/{project_id}/members | Remove a member (user) from a project.
[**RemovePipelineTask**](ProjectApi.md#RemovePipelineTask) | **Delete** /pipelines/tasks/{task_id} | Remove the specified pipeline task.
[**UpdateProject**](ProjectApi.md#UpdateProject) | **Put** /projects/{project_id} | Update the specified project.\n


# **AddMember**
> AddMember($projectId, $userId)

Add a member (user) to a project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32**| The project id. | 
 **userId** | **int32**| The user id. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPipelineTask**
> PipelineTask AddPipelineTask($projectId, $pipelineTask)

Add a pipeline task for a specific project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32**| Project id to filter by. | 
 **pipelineTask** | [**PipelineTask**](PipelineTask.md)| PipelineTask object. | 

### Return type

[**PipelineTask**](PipelineTask.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProject**
> Project CreateProject($project)

Create a new project.

Create a new `Project`. The newly created `Project` will initially not have any\nmembers or owners.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**Project**](Project.md)| Project to create. | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> DeleteProject($projectId)

Delete the specified `Project` instance.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32**| The id of the &#x60;Project&#x60; instance to delete.\n | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineTask**
> PipelineTask GetPipelineTask($taskId)

Get the specified pipeline task.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **int32**| PipelineTask id. | 

### Return type

[**PipelineTask**](PipelineTask.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineTasks**
> []PipelineTask GetPipelineTasks($projectId)

List the pipeline tasks, optionally filtering by project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32**| Project id to filter by. | 

### Return type

[**[]PipelineTask**](PipelineTask.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProject**
> Project GetProject($userId, $projectId)

Get the specified instance.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| Filter the list of projects by whether this user is a member of that project. | 
 **projectId** | **int32**| HCE id of the &#x60;Project&#x60; to get. | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectMembers**
> []User GetProjectMembers($projectId)

Get the project members.

Get all project members (including project owners).


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32**| The instance id. | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjects**
> []Project GetProjects($userId, $deploymentTargetId)

List projects, optionally filtering.

Get the list of projects, optionally filtering by user membership, and/or deployment targets that the project is associated with.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| Filter the list of projects by whether this user is a member of that project. | 
 **deploymentTargetId** | **int32**| Filter the list of projects by whether the project has this environment as a target. | [optional] 

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveMember**
> RemoveMember($projectId, $userId)

Remove a member (user) from a project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32**| The project id. | 
 **userId** | **int32**| The user id. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemovePipelineTask**
> RemovePipelineTask($taskId)

Remove the specified pipeline task.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **int32**| The PipelineTask id to remove. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProject**
> Project UpdateProject($projectId, $project)

Update the specified project.\n

Update the specified `Project` and return the updated `Project` instance.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32**| The id of the &#x60;Project&#x60; instance.\n | 
 **project** | [**Project**](Project.md)| The &#x60;Project&#x60; instance (containing values to be updated).\n | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

