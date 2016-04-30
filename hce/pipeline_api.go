package hce

import (
    "strings"
    "fmt"
    "errors"
    "encoding/json"
)

type PipelineApi struct {
    Configuration Configuration
}

func NewPipelineApi() *PipelineApi{
    configuration := NewConfiguration()
    return &PipelineApi {
        Configuration: *configuration,
    }
}

func NewPipelineApiWithBasePath(basePath string) *PipelineApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &PipelineApi {
        Configuration: *configuration,
    }
}

/**
 * Delete the specified build.
 * 
 * @param executionId Build id.
 * @return void
 */
func (a PipelineApi) DeletePipelineExecution (executionId int32) (APIResponse, error) {

  var httpMethod = "Delete"
 // create path and map variables
  path := a.Configuration.BasePath + "/pipelines/executions/{execution_id}"
 path = strings.Replace(path, "{" + "execution_id" + "}", fmt.Sprintf("%v", executionId), -1)

  // verify the required parameter 'executionId' is set
  if &executionId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'executionId' when calling PipelineApi->DeletePipelineExecution")
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
 * Get the specified pipeline event.
 * 
 * @param eventId PipelineEvent id.
 * @return PipelineEvent
 */
func (a PipelineApi) GetPipelineEvent (eventId int32) (PipelineEvent, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/pipelines/events/{event_id}"
 path = strings.Replace(path, "{" + "event_id" + "}", fmt.Sprintf("%v", eventId), -1)

  // verify the required parameter 'eventId' is set
  if &eventId == nil {
      return *new(PipelineEvent), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'eventId' when calling PipelineApi->GetPipelineEvent")
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


  var successPayload = new(PipelineEvent)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * List pipeline events, optionally filtering by Build id.
 * 
 * @param executionId execution id.
 * @return []PipelineEvent
 */
func (a PipelineApi) GetPipelineEvents (executionId int32) ([]PipelineEvent, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/pipelines/events"

  // verify the required parameter 'executionId' is set
  if &executionId == nil {
      return *new([]PipelineEvent), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'executionId' when calling PipelineApi->GetPipelineEvents")
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


  var successPayload = new([]PipelineEvent)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Gets the specified build.
 * 
 * @param executionId Build id.
 * @return PipelineExecution
 */
func (a PipelineApi) GetPipelineExecution (executionId int32) (PipelineExecution, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/pipelines/executions/{execution_id}"
 path = strings.Replace(path, "{" + "execution_id" + "}", fmt.Sprintf("%v", executionId), -1)

  // verify the required parameter 'executionId' is set
  if &executionId == nil {
      return *new(PipelineExecution), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'executionId' when calling PipelineApi->GetPipelineExecution")
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


  var successPayload = new(PipelineExecution)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * List executions, optionally filtering by project_id.
 * Return an array of builds, typically filtered by project_id, and often paged.\nFor example &#x60;/pipelines/executions?project_id&#x3D;2&amp;limit&#x3D;25&amp;offset&#x3D;50&#x60; return an\narray of builds for project with id &#x60;2&#x60;, will return a maximum of &#x60;25&#x60; items,\nstarting with item &#x60;50&#x60;.\n
 * @param projectId Project id.
 * @param limit Maximum number of items to return.
 * @param offset Returned array of items should begin at this offset.
 * @return []PipelineExecution
 */
func (a PipelineApi) GetPipelineExecutions (projectId int32, limit int32, offset int32) ([]PipelineExecution, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/pipelines/executions"

  // verify the required parameter 'projectId' is set
  if &projectId == nil {
      return *new([]PipelineExecution), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'projectId' when calling PipelineApi->GetPipelineExecutions")
  }
  // verify the required parameter 'limit' is set
  if &limit == nil {
      return *new([]PipelineExecution), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'limit' when calling PipelineApi->GetPipelineExecutions")
  }
  // verify the required parameter 'offset' is set
  if &offset == nil {
      return *new([]PipelineExecution), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'offset' when calling PipelineApi->GetPipelineExecutions")
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
  
  queryParams["projectId"] = a.Configuration.APIClient.ParameterToString(projectId)
  queryParams["limit"] = a.Configuration.APIClient.ParameterToString(limit)
  queryParams["offset"] = a.Configuration.APIClient.ParameterToString(offset)

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


  var successPayload = new([]PipelineExecution)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Record a PipelineEvent.
 * 
 * @param executionId execution id.
 * @param pipelineEvent Pipeline event.
 * @return PipelineEvent
 */
func (a PipelineApi) PipelineEventOccurred (executionId int32, pipelineEvent PipelineEvent) (PipelineEvent, APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/pipelines/events"

  // verify the required parameter 'executionId' is set
  if &executionId == nil {
      return *new(PipelineEvent), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'executionId' when calling PipelineApi->PipelineEventOccurred")
  }
  // verify the required parameter 'pipelineEvent' is set
  if &pipelineEvent == nil {
      return *new(PipelineEvent), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'pipelineEvent' when calling PipelineApi->PipelineEventOccurred")
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

  // body params
  postBody = &pipelineEvent

  var successPayload = new(PipelineEvent)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Trigger execution of a pipeline(s).
 * 
 * @param buildTrigger PipelineTrigger event.
 * @return PipelineTrigger
 */
func (a PipelineApi) TriggerPipelineExecution (buildTrigger PipelineTrigger) (PipelineTrigger, APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/pipelines/triggers"

  // verify the required parameter 'buildTrigger' is set
  if &buildTrigger == nil {
      return *new(PipelineTrigger), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'buildTrigger' when calling PipelineApi->TriggerPipelineExecution")
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

  // body params
  postBody = &buildTrigger

  var successPayload = new(PipelineTrigger)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
