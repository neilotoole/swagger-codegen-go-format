# \NotificationApi

All URIs are relative to *https://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNotificationTarget**](NotificationApi.md#AddNotificationTarget) | **Post** /notifications/targets | Add a NotificationTarget (for a project).
[**GetNotificationTarget**](NotificationApi.md#GetNotificationTarget) | **Get** /notifications/targets/{notification_target_id} | Get the specified notification target.
[**GetNotificationTargets**](NotificationApi.md#GetNotificationTargets) | **Get** /notifications/targets | List notifiction targets, optionally filtering.
[**RemoveNotificationTarget**](NotificationApi.md#RemoveNotificationTarget) | **Delete** /notifications/targets/{notification_target_id} | Remove the specified notification target.
[**UpdateNotificationTarget**](NotificationApi.md#UpdateNotificationTarget) | **Put** /notifications/targets/{notification_target_id} | Update the specified notification target.


# **AddNotificationTarget**
> NotificationTarget AddNotificationTarget($projectId, $notificationTarget)

Add a NotificationTarget (for a project).




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32**| The HCE Project id. | 
 **notificationTarget** | [**NotificationTarget**](NotificationTarget.md)| NotificationTarget object. | 

### Return type

[**NotificationTarget**](NotificationTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationTarget**
> NotificationTarget GetNotificationTarget($notificationTargetId)

Get the specified notification target.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationTargetId** | **int32**| NotificationTarget id. | 

### Return type

[**NotificationTarget**](NotificationTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationTargets**
> []NotificationTarget GetNotificationTargets($projectId)

List notifiction targets, optionally filtering.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32**| The HCE Project id. | 

### Return type

[**[]NotificationTarget**](NotificationTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveNotificationTarget**
> RemoveNotificationTarget($notificationTargetId)

Remove the specified notification target.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationTargetId** | **int32**| NotificationTarget id. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationTarget**
> NotificationTarget UpdateNotificationTarget($notificationTargetId, $user)

Update the specified notification target.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationTargetId** | **int32**| NotificationTarget id. | 
 **user** | [**NotificationTarget**](NotificationTarget.md)| NotificationTarget object (containing values to be updated). | 

### Return type

[**NotificationTarget**](NotificationTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

