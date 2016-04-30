package hce

import (
    "strings"
    "fmt"
    "errors"
    "encoding/json"
)

type DeploymentApi struct {
    Configuration Configuration
}

func NewDeploymentApi() *DeploymentApi{
    configuration := NewConfiguration()
    return &DeploymentApi {
        Configuration: *configuration,
    }
}

func NewDeploymentApiWithBasePath(basePath string) *DeploymentApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &DeploymentApi {
        Configuration: *configuration,
    }
}

/**
 * Add a deployment target.
 * Add (register) a deployment environment (target), such as a CloudFoundry instance.\n
 * @param deploymentTarget DeploymentTarget object.
 * @return DeploymentTarget
 */
func (a DeploymentApi) AddDeploymentTarget (deploymentTarget DeploymentTarget) (DeploymentTarget, APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/deployments/targets"

  // verify the required parameter 'deploymentTarget' is set
  if &deploymentTarget == nil {
      return *new(DeploymentTarget), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'deploymentTarget' when calling DeploymentApi->AddDeploymentTarget")
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
  postBody = &deploymentTarget

  var successPayload = new(DeploymentTarget)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Record a deployment.
 * Record a deploymennt. FIXME: Does invoking this method REGISTER a deployment that has occurred, or does it initiate a deployment?\n
 * @param deployment Deployment object.
 * @return Deployment
 */
func (a DeploymentApi) DeploymentOccurred (deployment Deployment) (Deployment, APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/deployments"

  // verify the required parameter 'deployment' is set
  if &deployment == nil {
      return *new(Deployment), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'deployment' when calling DeploymentApi->DeploymentOccurred")
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
  postBody = &deployment

  var successPayload = new(Deployment)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Get the specified deployment.
 * 
 * @param deploymentId Deployment id.
 * @return Deployment
 */
func (a DeploymentApi) GetDeployment (deploymentId int32) (Deployment, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/deployments/{deployment_id}"
 path = strings.Replace(path, "{" + "deployment_id" + "}", fmt.Sprintf("%v", deploymentId), -1)

  // verify the required parameter 'deploymentId' is set
  if &deploymentId == nil {
      return *new(Deployment), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'deploymentId' when calling DeploymentApi->GetDeployment")
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


  var successPayload = new(Deployment)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Get the specified environment.
 * 
 * @param targetId DeploymentTarget id.
 * @return DeploymentTarget
 */
func (a DeploymentApi) GetDeploymentTarget (targetId int32) (DeploymentTarget, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/deployments/targets/{target_id}"
 path = strings.Replace(path, "{" + "target_id" + "}", fmt.Sprintf("%v", targetId), -1)

  // verify the required parameter 'targetId' is set
  if &targetId == nil {
      return *new(DeploymentTarget), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'targetId' when calling DeploymentApi->GetDeploymentTarget")
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


  var successPayload = new(DeploymentTarget)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * List the registered deployment targets, optionally filtering.
 * List the registered deployment targets, optionally filtering by owner user_id.\n
 * @param userId User id.
 * @return []DeploymentTarget
 */
func (a DeploymentApi) GetDeploymentTargets (userId int32) ([]DeploymentTarget, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/deployments/targets"

  // verify the required parameter 'userId' is set
  if &userId == nil {
      return *new([]DeploymentTarget), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'userId' when calling DeploymentApi->GetDeploymentTargets")
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
  
  queryParams["userId"] = a.Configuration.APIClient.ParameterToString(userId)

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


  var successPayload = new([]DeploymentTarget)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * List deployments, optionally filtering by Project id or Build id.
 * 
 * @param projectId Project id.
 * @param executionId Build id.
 * @return []Deployment
 */
func (a DeploymentApi) GetDeployments (projectId int32, executionId int32) ([]Deployment, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/deployments"

  // verify the required parameter 'projectId' is set
  if &projectId == nil {
      return *new([]Deployment), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'projectId' when calling DeploymentApi->GetDeployments")
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


  var successPayload = new([]Deployment)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Remove the specified deployment.
 * Remove the specified deployment. FIXME: Does invoking this operation indicate that a deployment has been deleted, or does it initiate deletion of a deployment?\n
 * @param deploymentId Deployment id.
 * @return void
 */
func (a DeploymentApi) RemoveDeployment (deploymentId int32) (APIResponse, error) {

  var httpMethod = "Delete"
 // create path and map variables
  path := a.Configuration.BasePath + "/deployments/{deployment_id}"
 path = strings.Replace(path, "{" + "deployment_id" + "}", fmt.Sprintf("%v", deploymentId), -1)

  // verify the required parameter 'deploymentId' is set
  if &deploymentId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'deploymentId' when calling DeploymentApi->RemoveDeployment")
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
 * Remove (unregister) the specified deployment target.
 * 
 * @param targetId DeploymentTarget id.
 * @return void
 */
func (a DeploymentApi) RemoveDeploymentTarget (targetId int32) (APIResponse, error) {

  var httpMethod = "Delete"
 // create path and map variables
  path := a.Configuration.BasePath + "/deployments/targets/{target_id}"
 path = strings.Replace(path, "{" + "target_id" + "}", fmt.Sprintf("%v", targetId), -1)

  // verify the required parameter 'targetId' is set
  if &targetId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'targetId' when calling DeploymentApi->RemoveDeploymentTarget")
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
 * Update the specified DeploymentTarget.
 * 
 * @param targetId DeploymentTarget id.
 * @param deploymentTarget DeploymentTarget object.
 * @return DeploymentTarget
 */
func (a DeploymentApi) UpdateTarget (targetId int32, deploymentTarget DeploymentTarget) (DeploymentTarget, APIResponse, error) {

  var httpMethod = "Put"
 // create path and map variables
  path := a.Configuration.BasePath + "/deployments/targets/{target_id}"
 path = strings.Replace(path, "{" + "target_id" + "}", fmt.Sprintf("%v", targetId), -1)

  // verify the required parameter 'targetId' is set
  if &targetId == nil {
      return *new(DeploymentTarget), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'targetId' when calling DeploymentApi->UpdateTarget")
  }
  // verify the required parameter 'deploymentTarget' is set
  if &deploymentTarget == nil {
      return *new(DeploymentTarget), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'deploymentTarget' when calling DeploymentApi->UpdateTarget")
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
  postBody = &deploymentTarget

  var successPayload = new(DeploymentTarget)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
