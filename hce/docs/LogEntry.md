# LogEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | E.g. &#x60;error&#x60;, &#x60;warn&#x60; etc. These should become an enum.\n | [optional] [default to null]
**Message** | **string** | The log message.\n | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The timestamp, in &#x60;2016-04-17T11:40:43.448Z&#x60; format.\n | [optional] [default to null]
**Items** | [**[]ErrorItem**](ErrorItem.md) | Individual error items.\n | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


