# ErrorResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **int32** | The API version, e.g. &#x60;2&#x60;. | [default to null]
**Status** | **int32** | The HTTP status code. | [default to null]
**Message** | **string** | A short message associated with the error. | [optional] [default to null]
**Detail** | **string** | A longer, user-facing message associated with the error. | [optional] [default to null]
**Log** | [**[]LogEntry**](LogEntry.md) | Any additional error information, such as log entries, stack trace, etc..\nLikely to be turned off in production environments. May be null.\n | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


