# \UserApi

All URIs are relative to *https://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UserApi.md#CreateUser) | **Post** /users | Create a user.
[**DeleteUser**](UserApi.md#DeleteUser) | **Delete** /users/{user_id} | Delete the specified user.
[**GetUser**](UserApi.md#GetUser) | **Get** /users/{user_id} | Get the specified user.
[**GetUserByGithubId**](UserApi.md#GetUserByGithubId) | **Get** /users/github/{github_id} | Get the specified user.
[**UpdateUser**](UserApi.md#UpdateUser) | **Put** /users/{user_id} | Update the specified user.


# **CreateUser**
> User CreateUser($user)

Create a user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md)| User object. | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser($userId)

Delete the specified user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| User id. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> User GetUser($userId)

Get the specified user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| User id. | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserByGithubId**
> User GetUserByGithubId($githubId)

Get the specified user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **githubId** | **int32**| Users github account id. | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> User UpdateUser($userId, $user)

Update the specified user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| User id. | 
 **user** | [**User**](User.md)| User object (containing values to be updated). | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

