# \DeploymentApi

All URIs are relative to *https://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDeploymentTarget**](DeploymentApi.md#AddDeploymentTarget) | **Post** /deployments/targets | Add a deployment target.
[**DeploymentOccurred**](DeploymentApi.md#DeploymentOccurred) | **Post** /deployments | Record a deployment.
[**GetDeployment**](DeploymentApi.md#GetDeployment) | **Get** /deployments/{deployment_id} | Get the specified deployment.
[**GetDeploymentTarget**](DeploymentApi.md#GetDeploymentTarget) | **Get** /deployments/targets/{target_id} | Get the specified environment.
[**GetDeploymentTargets**](DeploymentApi.md#GetDeploymentTargets) | **Get** /deployments/targets | List the registered deployment targets, optionally filtering.
[**GetDeployments**](DeploymentApi.md#GetDeployments) | **Get** /deployments | List deployments, optionally filtering by Project id or Build id.
[**RemoveDeployment**](DeploymentApi.md#RemoveDeployment) | **Delete** /deployments/{deployment_id} | Remove the specified deployment.
[**RemoveDeploymentTarget**](DeploymentApi.md#RemoveDeploymentTarget) | **Delete** /deployments/targets/{target_id} | Remove (unregister) the specified deployment target.
[**UpdateTarget**](DeploymentApi.md#UpdateTarget) | **Put** /deployments/targets/{target_id} | Update the specified DeploymentTarget.


# **AddDeploymentTarget**
> DeploymentTarget AddDeploymentTarget($deploymentTarget)

Add a deployment target.

Add (register) a deployment environment (target), such as a CloudFoundry instance.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentTarget** | [**DeploymentTarget**](DeploymentTarget.md)| DeploymentTarget object. | 

### Return type

[**DeploymentTarget**](DeploymentTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeploymentOccurred**
> Deployment DeploymentOccurred($deployment)

Record a deployment.

Record a deploymennt. FIXME: Does invoking this method REGISTER a deployment that has occurred, or does it initiate a deployment?\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment** | [**Deployment**](Deployment.md)| Deployment object. | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployment**
> Deployment GetDeployment($deploymentId)

Get the specified deployment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentId** | **int32**| Deployment id. | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentTarget**
> DeploymentTarget GetDeploymentTarget($targetId)

Get the specified environment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetId** | **int32**| DeploymentTarget id. | 

### Return type

[**DeploymentTarget**](DeploymentTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentTargets**
> []DeploymentTarget GetDeploymentTargets($userId)

List the registered deployment targets, optionally filtering.

List the registered deployment targets, optionally filtering by owner user_id.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| User id. | 

### Return type

[**[]DeploymentTarget**](DeploymentTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployments**
> []Deployment GetDeployments($projectId, $executionId)

List deployments, optionally filtering by Project id or Build id.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32**| Project id. | 
 **executionId** | **int32**| Build id. | [optional] 

### Return type

[**[]Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveDeployment**
> RemoveDeployment($deploymentId)

Remove the specified deployment.

Remove the specified deployment. FIXME: Does invoking this operation indicate that a deployment has been deleted, or does it initiate deletion of a deployment?\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentId** | **int32**| Deployment id. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveDeploymentTarget**
> RemoveDeploymentTarget($targetId)

Remove (unregister) the specified deployment target.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetId** | **int32**| DeploymentTarget id. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTarget**
> DeploymentTarget UpdateTarget($targetId, $deploymentTarget)

Update the specified DeploymentTarget.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetId** | **int32**| DeploymentTarget id. | 
 **deploymentTarget** | [**DeploymentTarget**](DeploymentTarget.md)| DeploymentTarget object. | 

### Return type

[**DeploymentTarget**](DeploymentTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

