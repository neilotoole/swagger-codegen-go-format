package hce

import (
  "net/http"
)


type APIResponse struct {
  *http.Response
  
  Message  string  `json:"message,omitempty"`
}

func NewAPIResponse(r *http.Response) *APIResponse {
  response := &APIResponse{Response: r}

  return response
}

func NewAPIResponseWithError(errorMessage string) *APIResponse {
  response := &APIResponse{Message: errorMessage}

  return response
}