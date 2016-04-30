package hce

import (
    "strings"
    "fmt"
    "errors"
    "encoding/json"
)

type SecurityApi struct {
    Configuration Configuration
}

func NewSecurityApi() *SecurityApi{
    configuration := NewConfiguration()
    return &SecurityApi {
        Configuration: *configuration,
    }
}

func NewSecurityApiWithBasePath(basePath string) *SecurityApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &SecurityApi {
        Configuration: *configuration,
    }
}

/**
 * Remove the specified credential from the authstore.
 * 
 * @param credentialId The credential id.
 * @return void
 */
func (a SecurityApi) ForgetCredential (credentialId int32) (APIResponse, error) {

  var httpMethod = "Delete"
 // create path and map variables
  path := a.Configuration.BasePath + "/auth/credentials/{credential_id}"
 path = strings.Replace(path, "{" + "credential_id" + "}", fmt.Sprintf("%v", credentialId), -1)

  // verify the required parameter 'credentialId' is set
  if &credentialId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'credentialId' when calling SecurityApi->ForgetCredential")
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
 * List credentials for the specified user.
 * Get a view of the credentials for the specified user. The returned data does not include secret fields.
 * @param userId List the credentials owned by this user.
 * @return []CredentialView
 */
func (a SecurityApi) ListCredentials (userId int32) ([]CredentialView, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/auth/credentials"

  // verify the required parameter 'userId' is set
  if &userId == nil {
      return *new([]CredentialView), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'userId' when calling SecurityApi->ListCredentials")
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


  var successPayload = new([]CredentialView)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Store auth Credential.
 * Add security credential to the authstore. After this operation, the Credential is still\nunlinked to any entity.\n
 * @param credential The credential details.
 * @return CredentialView
 */
func (a SecurityApi) StoreCredential (credential CredentialDetail) (CredentialView, APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/auth/credentials"

  // verify the required parameter 'credential' is set
  if &credential == nil {
      return *new(CredentialView), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'credential' when calling SecurityApi->StoreCredential")
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
  postBody = &credential

  var successPayload = new(CredentialView)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Update the specified credential.\n
 * Update the specified Credential and return the updated &#x60;Credential&#x60; instance.\n
 * @param credentialId The id of the &#x60;Credential&#x60; instance.\n
 * @param credentialUpdate The &#x60;CredentialUpdate&#x60; instance (containing values to be updated).\n
 * @return CredentialView
 */
func (a SecurityApi) UpdateCredential (credentialId int32, credentialUpdate CredentialUpdate) (CredentialView, APIResponse, error) {

  var httpMethod = "Put"
 // create path and map variables
  path := a.Configuration.BasePath + "/auth/credentials/{credential_id}"
 path = strings.Replace(path, "{" + "credential_id" + "}", fmt.Sprintf("%v", credentialId), -1)

  // verify the required parameter 'credentialId' is set
  if &credentialId == nil {
      return *new(CredentialView), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'credentialId' when calling SecurityApi->UpdateCredential")
  }
  // verify the required parameter 'credentialUpdate' is set
  if &credentialUpdate == nil {
      return *new(CredentialView), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'credentialUpdate' when calling SecurityApi->UpdateCredential")
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
  postBody = &credentialUpdate

  var successPayload = new(CredentialView)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
