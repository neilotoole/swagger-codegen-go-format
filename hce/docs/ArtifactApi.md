# \ArtifactApi

All URIs are relative to *https://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteArtifact**](ArtifactApi.md#DeleteArtifact) | **Delete** /artifacts/{artifact_id} | Deletes the artifact specified in the request params.
[**DownloadArtifact**](ArtifactApi.md#DownloadArtifact) | **Get** /artifacts/{artifact_id}/download | Download the file specified by the artifact id. This operation may result in a redirect.
[**GetArtifact**](ArtifactApi.md#GetArtifact) | **Get** /artifacts/{artifact_id} | Get the artifact specified in request.
[**GetArtifacts**](ArtifactApi.md#GetArtifacts) | **Get** /artifacts | List the list of artifacts associated with the specified PipelineExecution.
[**UploadArtifact**](ArtifactApi.md#UploadArtifact) | **Post** /artifacts | Upload an artifact for the specified PipelineExecution.


# **DeleteArtifact**
> DeleteArtifact($artifactId)

Deletes the artifact specified in the request params.

Deletes the artifact specified in the request params.\nIf artifact with specified id not found, returns 404.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifactId** | **int32**| The id of the artifact to delete. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadArtifact**
> *os.File DownloadArtifact($artifactId)

Download the file specified by the artifact id. This operation may result in a redirect.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifactId** | **int32**| The id of the artifact to download. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifact**
> Artifact GetArtifact($artifactId)

Get the artifact specified in request.

Get the artifact specified in request. Note that this operation returns an\nArtifact object (describing the artifact). Use `download_artifact` to fetch the Artifact file itself.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifactId** | **int32**| The id of the artifact. | 

### Return type

[**Artifact**](Artifact.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifacts**
> []Artifact GetArtifacts($executionId)

List the list of artifacts associated with the specified PipelineExecution.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionId** | **int32**| The PipelineExecution id to list artifacts for. | 

### Return type

[**[]Artifact**](Artifact.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadArtifact**
> Artifact UploadArtifact($executionId, $artifactType, $file)

Upload an artifact for the specified PipelineExecution.

Upload an artifact (file) for the specified PipelineExecution, returning the Artifact (descriptor)\nobject for the recently uploaded file.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionId** | **int32**| The id of the PipelineExecution the artifact is associated with. | 
 **artifactType** | **string**| The artifact type, e.g. \&quot;BUILD_LOG\&quot;. | 
 **file** | ***os.File**| Artifact file. | 

### Return type

[**Artifact**](Artifact.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-gzip
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

