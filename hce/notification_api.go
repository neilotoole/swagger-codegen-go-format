package hce

import (
    "strings"
    "fmt"
    "errors"
    "encoding/json"
)

type NotificationApi struct {
    Configuration Configuration
}

func NewNotificationApi() *NotificationApi{
    configuration := NewConfiguration()
    return &NotificationApi {
        Configuration: *configuration,
    }
}

func NewNotificationApiWithBasePath(basePath string) *NotificationApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &NotificationApi {
        Configuration: *configuration,
    }
}

/**
 * Add a NotificationTarget (for a project).
 * 
 * @param projectId The HCE Project id.
 * @param notificationTarget NotificationTarget object.
 * @return NotificationTarget
 */
func (a NotificationApi) AddNotificationTarget (projectId int32, notificationTarget NotificationTarget) (NotificationTarget, APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/notifications/targets"

  // verify the required parameter 'projectId' is set
  if &projectId == nil {
      return *new(NotificationTarget), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'projectId' when calling NotificationApi->AddNotificationTarget")
  }
  // verify the required parameter 'notificationTarget' is set
  if &notificationTarget == nil {
      return *new(NotificationTarget), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'notificationTarget' when calling NotificationApi->AddNotificationTarget")
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
  postBody = &notificationTarget

  var successPayload = new(NotificationTarget)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Get the specified notification target.
 * 
 * @param notificationTargetId NotificationTarget id.
 * @return NotificationTarget
 */
func (a NotificationApi) GetNotificationTarget (notificationTargetId int32) (NotificationTarget, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/notifications/targets/{notification_target_id}"
 path = strings.Replace(path, "{" + "notification_target_id" + "}", fmt.Sprintf("%v", notificationTargetId), -1)

  // verify the required parameter 'notificationTargetId' is set
  if &notificationTargetId == nil {
      return *new(NotificationTarget), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'notificationTargetId' when calling NotificationApi->GetNotificationTarget")
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


  var successPayload = new(NotificationTarget)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * List notifiction targets, optionally filtering.
 * 
 * @param projectId The HCE Project id.
 * @return []NotificationTarget
 */
func (a NotificationApi) GetNotificationTargets (projectId int32) ([]NotificationTarget, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/notifications/targets"

  // verify the required parameter 'projectId' is set
  if &projectId == nil {
      return *new([]NotificationTarget), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'projectId' when calling NotificationApi->GetNotificationTargets")
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


  var successPayload = new([]NotificationTarget)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Remove the specified notification target.
 * 
 * @param notificationTargetId NotificationTarget id.
 * @return void
 */
func (a NotificationApi) RemoveNotificationTarget (notificationTargetId int32) (APIResponse, error) {

  var httpMethod = "Delete"
 // create path and map variables
  path := a.Configuration.BasePath + "/notifications/targets/{notification_target_id}"
 path = strings.Replace(path, "{" + "notification_target_id" + "}", fmt.Sprintf("%v", notificationTargetId), -1)

  // verify the required parameter 'notificationTargetId' is set
  if &notificationTargetId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'notificationTargetId' when calling NotificationApi->RemoveNotificationTarget")
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
 * Update the specified notification target.
 * 
 * @param notificationTargetId NotificationTarget id.
 * @param user NotificationTarget object (containing values to be updated).
 * @return NotificationTarget
 */
func (a NotificationApi) UpdateNotificationTarget (notificationTargetId int32, user NotificationTarget) (NotificationTarget, APIResponse, error) {

  var httpMethod = "Put"
 // create path and map variables
  path := a.Configuration.BasePath + "/notifications/targets/{notification_target_id}"
 path = strings.Replace(path, "{" + "notification_target_id" + "}", fmt.Sprintf("%v", notificationTargetId), -1)

  // verify the required parameter 'notificationTargetId' is set
  if &notificationTargetId == nil {
      return *new(NotificationTarget), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'notificationTargetId' when calling NotificationApi->UpdateNotificationTarget")
  }
  // verify the required parameter 'user' is set
  if &user == nil {
      return *new(NotificationTarget), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'user' when calling NotificationApi->UpdateNotificationTarget")
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
  postBody = &user

  var successPayload = new(NotificationTarget)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
