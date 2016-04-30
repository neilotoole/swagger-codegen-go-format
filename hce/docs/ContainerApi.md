# \ContainerApi

All URIs are relative to *https://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddImage**](ContainerApi.md#AddImage) | **Post** /containers/images | Add a Image instance.
[**AddImageRegistry**](ContainerApi.md#AddImageRegistry) | **Post** /containers/images/registries | Add a ImageRegistry instance.
[**GetImage**](ContainerApi.md#GetImage) | **Get** /containers/images/{image_id} | Get the Image specified in the request.
[**GetImageRegistries**](ContainerApi.md#GetImageRegistries) | **Get** /containers/images/registries | List ImageRegistry instances.
[**GetImageRegistry**](ContainerApi.md#GetImageRegistry) | **Get** /containers/images/registries/{registry_id} | Get the ImageRegistry specified in the request.
[**GetImageTypes**](ContainerApi.md#GetImageTypes) | **Get** /containers/images/types | Enumeration of Container image types, e.g. &#x60;DOCKER&#x60;, &#x60;ROCKET&#x60;, etc.\n
[**GetImages**](ContainerApi.md#GetImages) | **Get** /containers/images | List Image instances.
[**RemoveImage**](ContainerApi.md#RemoveImage) | **Delete** /containers/images/{image_id} | Remove (unregister) the specified Image.
[**RemoveImageRegistry**](ContainerApi.md#RemoveImageRegistry) | **Delete** /containers/images/registries/{registry_id} | Remove (unregister) the specified ImageRegistry.
[**UpdateImage**](ContainerApi.md#UpdateImage) | **Put** /containers/images/{image_id} | Update the specified container image.
[**UpdateImageRegistry**](ContainerApi.md#UpdateImageRegistry) | **Put** /containers/images/registries/{registry_id} | Update the specified container registry.


# **AddImage**
> Image AddImage($image)

Add a Image instance.

Add (register) a Image instance, such as a Docker image.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image** | [**Image**](Image.md)| Image to add. | 

### Return type

[**Image**](Image.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddImageRegistry**
> ImageRegistry AddImageRegistry($registry)

Add a ImageRegistry instance.

Add (register) a ImageRegistry instance, such as DockerHub, or a local registry.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registry** | [**ImageRegistry**](ImageRegistry.md)| ImageRegistry to add. | 

### Return type

[**ImageRegistry**](ImageRegistry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImage**
> Image GetImage($imageId)

Get the Image specified in the request.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageId** | **int32**| The (HCE) image id. | 

### Return type

[**Image**](Image.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImageRegistries**
> []ImageRegistry GetImageRegistries()

List ImageRegistry instances.

List the registered ImageRegistry instances.\n


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]ImageRegistry**](ImageRegistry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImageRegistry**
> ImageRegistry GetImageRegistry($registryId)

Get the ImageRegistry specified in the request.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registryId** | **int32**| The (HCE) item id. | 

### Return type

[**ImageRegistry**](ImageRegistry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImageTypes**
> []string GetImageTypes()

Enumeration of Container image types, e.g. `DOCKER`, `ROCKET`, etc.\n


### Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImages**
> []Image GetImages()

List Image instances.

List the registered Image instances.\n


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]Image**](Image.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveImage**
> RemoveImage($imageId)

Remove (unregister) the specified Image.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageId** | **int32**| The (HCE) Image id. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveImageRegistry**
> RemoveImageRegistry($registryId)

Remove (unregister) the specified ImageRegistry.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registryId** | **int32**| The (HCE) ImageRegistry id. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateImage**
> Image UpdateImage($imageId, $user)

Update the specified container image.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageId** | **int32**| Container image id. | 
 **user** | [**Image**](Image.md)| Image object (containing values to be updated). | 

### Return type

[**Image**](Image.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateImageRegistry**
> ImageRegistry UpdateImageRegistry($registryId, $user)

Update the specified container registry.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registryId** | **int32**| Registry id. | 
 **user** | [**ImageRegistry**](ImageRegistry.md)| Registry object (containing values to be updated). | 

### Return type

[**ImageRegistry**](ImageRegistry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

