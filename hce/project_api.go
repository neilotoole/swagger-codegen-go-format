package hce

import (
    "strings"
    "fmt"
    "errors"
    "encoding/json"
)

type ProjectApi struct {
    Configuration Configuration
}

func NewProjectApi() *ProjectApi{
    configuration := NewConfiguration()
    return &ProjectApi {
        Configuration: *configuration,
    }
}

func NewProjectApiWithBasePath(basePath string) *ProjectApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &ProjectApi {
        Configuration: *configuration,
    }
}

/**
 * Add a member (user) to a project.
 * 
 * @param projectId The project id.
 * @param userId The user id.
 * @return void
 */
func (a ProjectApi) AddMember (projectId int32, userId int32) (APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/projects/{project_id}/members"
 path = strings.Replace(path, "{" + "project_id" + "}", fmt.Sprintf("%v", projectId), -1)

  // verify the required parameter 'projectId' is set
  if &projectId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'projectId' when calling ProjectApi->AddMember")
  }
  // verify the required parameter 'userId' is set
  if &userId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'userId' when calling ProjectApi->AddMember")
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
  postBody = &userId


  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *NewAPIResponse(httpResponse.RawResponse), err
  }

  return *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Add a pipeline task for a specific project.
 * 
 * @param projectId Project id to filter by.
 * @param pipelineTask PipelineTask object.
 * @return PipelineTask
 */
func (a ProjectApi) AddPipelineTask (projectId int32, pipelineTask PipelineTask) (PipelineTask, APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/pipelines/tasks"

  // verify the required parameter 'projectId' is set
  if &projectId == nil {
      return *new(PipelineTask), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'projectId' when calling ProjectApi->AddPipelineTask")
  }
  // verify the required parameter 'pipelineTask' is set
  if &pipelineTask == nil {
      return *new(PipelineTask), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'pipelineTask' when calling ProjectApi->AddPipelineTask")
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
  postBody = &pipelineTask

  var successPayload = new(PipelineTask)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Create a new project.
 * Create a new &#x60;Project&#x60;. The newly created &#x60;Project&#x60; will initially not have any\nmembers or owners.\n
 * @param project Project to create.
 * @return Project
 */
func (a ProjectApi) CreateProject (project Project) (Project, APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/projects"

  // verify the required parameter 'project' is set
  if &project == nil {
      return *new(Project), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'project' when calling ProjectApi->CreateProject")
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
  postBody = &project

  var successPayload = new(Project)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Delete the specified &#x60;Project&#x60; instance.\n
 * 
 * @param projectId The id of the &#x60;Project&#x60; instance to delete.\n
 * @return void
 */
func (a ProjectApi) DeleteProject (projectId int32) (APIResponse, error) {

  var httpMethod = "Delete"
 // create path and map variables
  path := a.Configuration.BasePath + "/projects/{project_id}"
 path = strings.Replace(path, "{" + "project_id" + "}", fmt.Sprintf("%v", projectId), -1)

  // verify the required parameter 'projectId' is set
  if &projectId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'projectId' when calling ProjectApi->DeleteProject")
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
 * Get the specified pipeline task.
 * 
 * @param taskId PipelineTask id.
 * @return PipelineTask
 */
func (a ProjectApi) GetPipelineTask (taskId int32) (PipelineTask, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/pipelines/tasks/{task_id}"
 path = strings.Replace(path, "{" + "task_id" + "}", fmt.Sprintf("%v", taskId), -1)

  // verify the required parameter 'taskId' is set
  if &taskId == nil {
      return *new(PipelineTask), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'taskId' when calling ProjectApi->GetPipelineTask")
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


  var successPayload = new(PipelineTask)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * List the pipeline tasks, optionally filtering by project.
 * 
 * @param projectId Project id to filter by.
 * @return []PipelineTask
 */
func (a ProjectApi) GetPipelineTasks (projectId int32) ([]PipelineTask, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/pipelines/tasks"

  // verify the required parameter 'projectId' is set
  if &projectId == nil {
      return *new([]PipelineTask), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'projectId' when calling ProjectApi->GetPipelineTasks")
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


  var successPayload = new([]PipelineTask)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Get the specified instance.\n
 * 
 * @param userId Filter the list of projects by whether this user is a member of that project.
 * @param projectId HCE id of the &#x60;Project&#x60; to get.
 * @return Project
 */
func (a ProjectApi) GetProject (userId int32, projectId int32) (Project, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/projects/{project_id}"
 path = strings.Replace(path, "{" + "project_id" + "}", fmt.Sprintf("%v", projectId), -1)

  // verify the required parameter 'userId' is set
  if &userId == nil {
      return *new(Project), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'userId' when calling ProjectApi->GetProject")
  }
  // verify the required parameter 'projectId' is set
  if &projectId == nil {
      return *new(Project), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'projectId' when calling ProjectApi->GetProject")
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


  var successPayload = new(Project)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Get the project members.
 * Get all project members (including project owners).
 * @param projectId The instance id.
 * @return []User
 */
func (a ProjectApi) GetProjectMembers (projectId int32) ([]User, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/projects/{project_id}/members"
 path = strings.Replace(path, "{" + "project_id" + "}", fmt.Sprintf("%v", projectId), -1)

  // verify the required parameter 'projectId' is set
  if &projectId == nil {
      return *new([]User), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'projectId' when calling ProjectApi->GetProjectMembers")
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


  var successPayload = new([]User)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * List projects, optionally filtering.
 * Get the list of projects, optionally filtering by user membership, and/or deployment targets that the project is associated with.
 * @param userId Filter the list of projects by whether this user is a member of that project.
 * @param deploymentTargetId Filter the list of projects by whether the project has this environment as a target.
 * @return []Project
 */
func (a ProjectApi) GetProjects (userId int32, deploymentTargetId int32) ([]Project, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/projects"

  // verify the required parameter 'userId' is set
  if &userId == nil {
      return *new([]Project), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'userId' when calling ProjectApi->GetProjects")
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
  queryParams["deploymentTargetId"] = a.Configuration.APIClient.ParameterToString(deploymentTargetId)

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


  var successPayload = new([]Project)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Remove a member (user) from a project.
 * 
 * @param projectId The project id.
 * @param userId The user id.
 * @return void
 */
func (a ProjectApi) RemoveMember (projectId int32, userId int32) (APIResponse, error) {

  var httpMethod = "Delete"
 // create path and map variables
  path := a.Configuration.BasePath + "/projects/{project_id}/members"
 path = strings.Replace(path, "{" + "project_id" + "}", fmt.Sprintf("%v", projectId), -1)

  // verify the required parameter 'projectId' is set
  if &projectId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'projectId' when calling ProjectApi->RemoveMember")
  }
  // verify the required parameter 'userId' is set
  if &userId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'userId' when calling ProjectApi->RemoveMember")
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
  postBody = &userId


  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *NewAPIResponse(httpResponse.RawResponse), err
  }

  return *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Remove the specified pipeline task.
 * 
 * @param taskId The PipelineTask id to remove.
 * @return void
 */
func (a ProjectApi) RemovePipelineTask (taskId int32) (APIResponse, error) {

  var httpMethod = "Delete"
 // create path and map variables
  path := a.Configuration.BasePath + "/pipelines/tasks/{task_id}"
 path = strings.Replace(path, "{" + "task_id" + "}", fmt.Sprintf("%v", taskId), -1)

  // verify the required parameter 'taskId' is set
  if &taskId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'taskId' when calling ProjectApi->RemovePipelineTask")
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
 * Update the specified project.\n
 * Update the specified &#x60;Project&#x60; and return the updated &#x60;Project&#x60; instance.\n
 * @param projectId The id of the &#x60;Project&#x60; instance.\n
 * @param project The &#x60;Project&#x60; instance (containing values to be updated).\n
 * @return Project
 */
func (a ProjectApi) UpdateProject (projectId int32, project Project) (Project, APIResponse, error) {

  var httpMethod = "Put"
 // create path and map variables
  path := a.Configuration.BasePath + "/projects/{project_id}"
 path = strings.Replace(path, "{" + "project_id" + "}", fmt.Sprintf("%v", projectId), -1)

  // verify the required parameter 'projectId' is set
  if &projectId == nil {
      return *new(Project), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'projectId' when calling ProjectApi->UpdateProject")
  }
  // verify the required parameter 'project' is set
  if &project == nil {
      return *new(Project), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'project' when calling ProjectApi->UpdateProject")
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
  postBody = &project

  var successPayload = new(Project)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
