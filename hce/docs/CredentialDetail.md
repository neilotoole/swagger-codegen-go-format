# CredentialDetail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | **int32** | The user id of the owner of this credential. May be null. | [default to null]
**CredentialType** | **string** | The type of credential. | [default to null]
**CredentialKey** | **string** | The credential key, e.g. the username, or AWS Access Key ID, etc. May be null. | [optional] [default to null]
**CredentialValue** | **string** | The credential value, e.g. the password, or AWS Secret Key. | [default to null]
**CredentialExtra** | **string** | An optional additional (secret) field in the case of a three-value credential, e.g. the user email of a DockerHub credential. | [optional] [default to null]
**Label** | **string** | An optional label for the credential, e.g. \&quot;Alice&#39;s GitHub OAuth token\&quot;. May be null. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


