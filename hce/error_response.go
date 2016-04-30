package hce

import (

)

// A standard error object returned by all API calls.\n
type ErrorResponse struct {
	// The API version, e.g. `2`.
	ApiVersion int32 `json:"api_version,omitempty"`
	// The HTTP status code.
	Status int32 `json:"status,omitempty"`
	// A short message associated with the error.
	Message string `json:"message,omitempty"`
	// A longer, user-facing message associated with the error.
	Detail string `json:"detail,omitempty"`
	// Any additional error information, such as log entries, stack trace, etc..\nLikely to be turned off in production environments. May be null.\n
	Log []LogEntry `json:"log,omitempty"`
}
