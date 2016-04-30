# CredentialUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialId** | **int32** | The id of the credential to update. | [default to null]
**CredentialKey** | **string** | The credential key, e.g. the username, or AWS Access Key ID, etc. May be null. | [optional] [default to null]
**CredentialValue** | **string** | The credential value, e.g. the password, or AWS Secret Key. | [default to null]
**CredentialExtra** | **string** | An optional additional (potentially secret) field, e.g. the user email in the case of a DockerHub authentication. | [optional] [default to null]
**Label** | **string** | An optional label for the credential, e.g. \&quot;Alice&#39;s GitHub OAuth token\&quot;. May be null. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


