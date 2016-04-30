# Project

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The internal project id. | [optional] [default to null]
**Name** | **string** | The user-provided label for the project. | [default to null]
**Type_** | **string** | The platform type of the project, e.g. &#39;nodejs&#39; or &#39;java&#39;, etc. | [default to null]
**UserId** | **int32** | The internal user id of the project owner. | [optional] [default to null]
**JoinCode** | **string** | The (semi-secret) code that an invited user users to join the project. | [optional] [default to null]
**Token** | **string** | The GitHub token to use (presumably???) | [optional] [default to null]
**BranchRefName** | **string** | The branch to use (FIXME: This is a duplicate of what&#39;s in \&quot;Repo\&quot;)\n | [optional] [default to null]
**Repo** | [**Repo**](Repo.md) |  | [optional] [default to null]
**Targets** | [**[]DeploymentTarget**](DeploymentTarget.md) | An array of target deployment environments (such as a CloudFoundry or AWS instance) that the project is deployed to. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


