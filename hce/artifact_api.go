package hce

import (
    "strings"
    "fmt"
    "errors"
    "os"
    "io/ioutil"
    "encoding/json"
)

type ArtifactApi struct {
    Configuration Configuration
}

func NewArtifactApi() *ArtifactApi{
    configuration := NewConfiguration()
    return &ArtifactApi {
        Configuration: *configuration,
    }
}

func NewArtifactApiWithBasePath(basePath string) *ArtifactApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &ArtifactApi {
        Configuration: *configuration,
    }
}

/**
 * Deletes the artifact specified in the request params.
 * Deletes the artifact specified in the request params.\nIf artifact with specified id not found, returns 404.\n
 * @param artifactId The id of the artifact to delete.
 * @return void
 */
func (a ArtifactApi) DeleteArtifact (artifactId int32) (APIResponse, error) {

  var httpMethod = "Delete"
 // create path and map variables
  path := a.Configuration.BasePath + "/artifacts/{artifact_id}"
 path = strings.Replace(path, "{" + "artifact_id" + "}", fmt.Sprintf("%v", artifactId), -1)

  // verify the required parameter 'artifactId' is set
  if &artifactId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'artifactId' when calling ArtifactApi->DeleteArtifact")
  }

  headerParams := make(map[string]string)
  queryParams := make(map[string]string)
  formParams := make(map[string]string)
  var postBody interface{}
  var fileName string
  var fileBytes []byte

  
  // add default headers if any
  for key := range a.Configuration.DefaultHeader {
      headerParams[key] = a.Configuration.DefaultHeader[key]
  }
  

  // to determine the Content-Type header
  localVarHttpContentTypes := []string {
      "application/json", 
  }
  //set Content-Type header
  localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
  if localVarHttpContentType != "" {    
      headerParams["Content-Type"] = localVarHttpContentType
  }
  // to determine the Accept header
  localVarHttpHeaderAccepts := []string {
      "application/json", 
  }
  //set Accept header
  localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
  if localVarHttpHeaderAccept != "" {  
      headerParams["Accept"] = localVarHttpHeaderAccept
  }



  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *NewAPIResponse(httpResponse.RawResponse), err
  }

  return *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Download the file specified by the artifact id. This operation may result in a redirect.
 * 
 * @param artifactId The id of the artifact to download.
 * @return *os.File
 */
func (a ArtifactApi) DownloadArtifact (artifactId int32) (*os.File, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/artifacts/{artifact_id}/download"
 path = strings.Replace(path, "{" + "artifact_id" + "}", fmt.Sprintf("%v", artifactId), -1)

  // verify the required parameter 'artifactId' is set
  if &artifactId == nil {
      return *new(*os.File), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'artifactId' when calling ArtifactApi->DownloadArtifact")
  }

  headerParams := make(map[string]string)
  queryParams := make(map[string]string)
  formParams := make(map[string]string)
  var postBody interface{}
  var fileName string
  var fileBytes []byte

  
  // add default headers if any
  for key := range a.Configuration.DefaultHeader {
      headerParams[key] = a.Configuration.DefaultHeader[key]
  }
  

  // to determine the Content-Type header
  localVarHttpContentTypes := []string {
      "application/json", 
  }
  //set Content-Type header
  localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
  if localVarHttpContentType != "" {    
      headerParams["Content-Type"] = localVarHttpContentType
  }
  // to determine the Accept header
  localVarHttpHeaderAccepts := []string {
      "application/json", 
  }
  //set Accept header
  localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
  if localVarHttpHeaderAccept != "" {  
      headerParams["Accept"] = localVarHttpHeaderAccept
  }


  var successPayload = new(*os.File)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Get the artifact specified in request.
 * Get the artifact specified in request. Note that this operation returns an\nArtifact object (describing the artifact). Use &#x60;download_artifact&#x60; to fetch the Artifact file itself.\n
 * @param artifactId The id of the artifact.
 * @return Artifact
 */
func (a ArtifactApi) GetArtifact (artifactId int32) (Artifact, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/artifacts/{artifact_id}"
 path = strings.Replace(path, "{" + "artifact_id" + "}", fmt.Sprintf("%v", artifactId), -1)

  // verify the required parameter 'artifactId' is set
  if &artifactId == nil {
      return *new(Artifact), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'artifactId' when calling ArtifactApi->GetArtifact")
  }

  headerParams := make(map[string]string)
  queryParams := make(map[string]string)
  formParams := make(map[string]string)
  var postBody interface{}
  var fileName string
  var fileBytes []byte

  
  // add default headers if any
  for key := range a.Configuration.DefaultHeader {
      headerParams[key] = a.Configuration.DefaultHeader[key]
  }
  

  // to determine the Content-Type header
  localVarHttpContentTypes := []string {
      "application/json", 
  }
  //set Content-Type header
  localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
  if localVarHttpContentType != "" {    
      headerParams["Content-Type"] = localVarHttpContentType
  }
  // to determine the Accept header
  localVarHttpHeaderAccepts := []string {
      "application/json", 
  }
  //set Accept header
  localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
  if localVarHttpHeaderAccept != "" {  
      headerParams["Accept"] = localVarHttpHeaderAccept
  }


  var successPayload = new(Artifact)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * List the list of artifacts associated with the specified PipelineExecution.
 * 
 * @param executionId The PipelineExecution id to list artifacts for.
 * @return []Artifact
 */
func (a ArtifactApi) GetArtifacts (executionId int32) ([]Artifact, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/artifacts"

  // verify the required parameter 'executionId' is set
  if &executionId == nil {
      return *new([]Artifact), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'executionId' when calling ArtifactApi->GetArtifacts")
  }

  headerParams := make(map[string]string)
  queryParams := make(map[string]string)
  formParams := make(map[string]string)
  var postBody interface{}
  var fileName string
  var fileBytes []byte

  
  // add default headers if any
  for key := range a.Configuration.DefaultHeader {
      headerParams[key] = a.Configuration.DefaultHeader[key]
  }
  
  queryParams["executionId"] = a.Configuration.APIClient.ParameterToString(executionId)

  // to determine the Content-Type header
  localVarHttpContentTypes := []string {
      "application/json", 
  }
  //set Content-Type header
  localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
  if localVarHttpContentType != "" {    
      headerParams["Content-Type"] = localVarHttpContentType
  }
  // to determine the Accept header
  localVarHttpHeaderAccepts := []string {
      "application/json", 
  }
  //set Accept header
  localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
  if localVarHttpHeaderAccept != "" {  
      headerParams["Accept"] = localVarHttpHeaderAccept
  }


  var successPayload = new([]Artifact)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Upload an artifact for the specified PipelineExecution.
 * Upload an artifact (file) for the specified PipelineExecution, returning the Artifact (descriptor)\nobject for the recently uploaded file.\n
 * @param executionId The id of the PipelineExecution the artifact is associated with.
 * @param artifactType The artifact type, e.g. \&quot;BUILD_LOG\&quot;.
 * @param file Artifact file.
 * @return Artifact
 */
func (a ArtifactApi) UploadArtifact (executionId int32, artifactType string, file *os.File) (Artifact, APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/artifacts"

  // verify the required parameter 'executionId' is set
  if &executionId == nil {
      return *new(Artifact), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'executionId' when calling ArtifactApi->UploadArtifact")
  }
  // verify the required parameter 'artifactType' is set
  if &artifactType == nil {
      return *new(Artifact), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'artifactType' when calling ArtifactApi->UploadArtifact")
  }
  // verify the required parameter 'file' is set
  if &file == nil {
      return *new(Artifact), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'file' when calling ArtifactApi->UploadArtifact")
  }

  headerParams := make(map[string]string)
  queryParams := make(map[string]string)
  formParams := make(map[string]string)
  var postBody interface{}
  var fileName string
  var fileBytes []byte

  
  // add default headers if any
  for key := range a.Configuration.DefaultHeader {
      headerParams[key] = a.Configuration.DefaultHeader[key]
  }
  
  queryParams["executionId"] = a.Configuration.APIClient.ParameterToString(executionId)
  queryParams["artifactType"] = a.Configuration.APIClient.ParameterToString(artifactType)

  // to determine the Content-Type header
  localVarHttpContentTypes := []string {
      "application/json", 
      "application/x-gzip", 
  }
  //set Content-Type header
  localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
  if localVarHttpContentType != "" {    
      headerParams["Content-Type"] = localVarHttpContentType
  }
  // to determine the Accept header
  localVarHttpHeaderAccepts := []string {
      "application/json", 
  }
  //set Accept header
  localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
  if localVarHttpHeaderAccept != "" {  
      headerParams["Accept"] = localVarHttpHeaderAccept
  }

  fbs, _ := ioutil.ReadAll(file)
  fileBytes = fbs
  fileName = file.Name()

  var successPayload = new(Artifact)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
