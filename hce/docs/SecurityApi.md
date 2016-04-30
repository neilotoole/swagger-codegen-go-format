# \SecurityApi

All URIs are relative to *https://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ForgetCredential**](SecurityApi.md#ForgetCredential) | **Delete** /auth/credentials/{credential_id} | Remove the specified credential from the authstore.
[**ListCredentials**](SecurityApi.md#ListCredentials) | **Get** /auth/credentials | List credentials for the specified user.
[**StoreCredential**](SecurityApi.md#StoreCredential) | **Post** /auth/credentials | Store auth Credential.
[**UpdateCredential**](SecurityApi.md#UpdateCredential) | **Put** /auth/credentials/{credential_id} | Update the specified credential.\n


# **ForgetCredential**
> ForgetCredential($credentialId)

Remove the specified credential from the authstore.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialId** | **int32**| The credential id. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCredentials**
> []CredentialView ListCredentials($userId)

List credentials for the specified user.

Get a view of the credentials for the specified user. The returned data does not include secret fields.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| List the credentials owned by this user. | 

### Return type

[**[]CredentialView**](CredentialView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreCredential**
> CredentialView StoreCredential($credential)

Store auth Credential.

Add security credential to the authstore. After this operation, the Credential is still\nunlinked to any entity.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credential** | [**CredentialDetail**](CredentialDetail.md)| The credential details. | 

### Return type

[**CredentialView**](CredentialView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCredential**
> CredentialView UpdateCredential($credentialId, $credentialUpdate)

Update the specified credential.\n

Update the specified Credential and return the updated `Credential` instance.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialId** | **int32**| The id of the &#x60;Credential&#x60; instance.\n | 
 **credentialUpdate** | [**CredentialUpdate**](CredentialUpdate.md)| The &#x60;CredentialUpdate&#x60; instance (containing values to be updated).\n | 

### Return type

[**CredentialView**](CredentialView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

