package hce

import (
    "strings"
    "fmt"
    "errors"
    "encoding/json"
)

type ContainerApi struct {
    Configuration Configuration
}

func NewContainerApi() *ContainerApi{
    configuration := NewConfiguration()
    return &ContainerApi {
        Configuration: *configuration,
    }
}

func NewContainerApiWithBasePath(basePath string) *ContainerApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &ContainerApi {
        Configuration: *configuration,
    }
}

/**
 * Add a Image instance.
 * Add (register) a Image instance, such as a Docker image.\n
 * @param image Image to add.
 * @return Image
 */
func (a ContainerApi) AddImage (image Image) (Image, APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/containers/images"

  // verify the required parameter 'image' is set
  if &image == nil {
      return *new(Image), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'image' when calling ContainerApi->AddImage")
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
  postBody = &image

  var successPayload = new(Image)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Add a ImageRegistry instance.
 * Add (register) a ImageRegistry instance, such as DockerHub, or a local registry.\n
 * @param registry ImageRegistry to add.
 * @return ImageRegistry
 */
func (a ContainerApi) AddImageRegistry (registry ImageRegistry) (ImageRegistry, APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/containers/images/registries"

  // verify the required parameter 'registry' is set
  if &registry == nil {
      return *new(ImageRegistry), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'registry' when calling ContainerApi->AddImageRegistry")
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
  postBody = &registry

  var successPayload = new(ImageRegistry)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Get the Image specified in the request.
 * 
 * @param imageId The (HCE) image id.
 * @return Image
 */
func (a ContainerApi) GetImage (imageId int32) (Image, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/containers/images/{image_id}"
 path = strings.Replace(path, "{" + "image_id" + "}", fmt.Sprintf("%v", imageId), -1)

  // verify the required parameter 'imageId' is set
  if &imageId == nil {
      return *new(Image), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'imageId' when calling ContainerApi->GetImage")
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


  var successPayload = new(Image)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * List ImageRegistry instances.
 * List the registered ImageRegistry instances.\n
 * @return []ImageRegistry
 */
func (a ContainerApi) GetImageRegistries () ([]ImageRegistry, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/containers/images/registries"


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


  var successPayload = new([]ImageRegistry)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Get the ImageRegistry specified in the request.
 * 
 * @param registryId The (HCE) item id.
 * @return ImageRegistry
 */
func (a ContainerApi) GetImageRegistry (registryId int32) (ImageRegistry, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/containers/images/registries/{registry_id}"
 path = strings.Replace(path, "{" + "registry_id" + "}", fmt.Sprintf("%v", registryId), -1)

  // verify the required parameter 'registryId' is set
  if &registryId == nil {
      return *new(ImageRegistry), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'registryId' when calling ContainerApi->GetImageRegistry")
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


  var successPayload = new(ImageRegistry)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Enumeration of Container image types, e.g. &#x60;DOCKER&#x60;, &#x60;ROCKET&#x60;, etc.\n
 * 
 * @return []string
 */
func (a ContainerApi) GetImageTypes () ([]string, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/containers/images/types"


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


  var successPayload = new([]string)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * List Image instances.
 * List the registered Image instances.\n
 * @return []Image
 */
func (a ContainerApi) GetImages () ([]Image, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/containers/images"


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


  var successPayload = new([]Image)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Remove (unregister) the specified Image.
 * 
 * @param imageId The (HCE) Image id.
 * @return void
 */
func (a ContainerApi) RemoveImage (imageId int32) (APIResponse, error) {

  var httpMethod = "Delete"
 // create path and map variables
  path := a.Configuration.BasePath + "/containers/images/{image_id}"
 path = strings.Replace(path, "{" + "image_id" + "}", fmt.Sprintf("%v", imageId), -1)

  // verify the required parameter 'imageId' is set
  if &imageId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'imageId' when calling ContainerApi->RemoveImage")
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
 * Remove (unregister) the specified ImageRegistry.
 * 
 * @param registryId The (HCE) ImageRegistry id.
 * @return void
 */
func (a ContainerApi) RemoveImageRegistry (registryId int32) (APIResponse, error) {

  var httpMethod = "Delete"
 // create path and map variables
  path := a.Configuration.BasePath + "/containers/images/registries/{registry_id}"
 path = strings.Replace(path, "{" + "registry_id" + "}", fmt.Sprintf("%v", registryId), -1)

  // verify the required parameter 'registryId' is set
  if &registryId == nil {
      return *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'registryId' when calling ContainerApi->RemoveImageRegistry")
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
 * Update the specified container image.
 * 
 * @param imageId Container image id.
 * @param user Image object (containing values to be updated).
 * @return Image
 */
func (a ContainerApi) UpdateImage (imageId int32, user Image) (Image, APIResponse, error) {

  var httpMethod = "Put"
 // create path and map variables
  path := a.Configuration.BasePath + "/containers/images/{image_id}"
 path = strings.Replace(path, "{" + "image_id" + "}", fmt.Sprintf("%v", imageId), -1)

  // verify the required parameter 'imageId' is set
  if &imageId == nil {
      return *new(Image), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'imageId' when calling ContainerApi->UpdateImage")
  }
  // verify the required parameter 'user' is set
  if &user == nil {
      return *new(Image), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'user' when calling ContainerApi->UpdateImage")
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

  var successPayload = new(Image)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Update the specified container registry.
 * 
 * @param registryId Registry id.
 * @param user Registry object (containing values to be updated).
 * @return ImageRegistry
 */
func (a ContainerApi) UpdateImageRegistry (registryId int32, user ImageRegistry) (ImageRegistry, APIResponse, error) {

  var httpMethod = "Put"
 // create path and map variables
  path := a.Configuration.BasePath + "/containers/images/registries/{registry_id}"
 path = strings.Replace(path, "{" + "registry_id" + "}", fmt.Sprintf("%v", registryId), -1)

  // verify the required parameter 'registryId' is set
  if &registryId == nil {
      return *new(ImageRegistry), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'registryId' when calling ContainerApi->UpdateImageRegistry")
  }
  // verify the required parameter 'user' is set
  if &user == nil {
      return *new(ImageRegistry), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'user' when calling ContainerApi->UpdateImageRegistry")
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

  var successPayload = new(ImageRegistry)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
