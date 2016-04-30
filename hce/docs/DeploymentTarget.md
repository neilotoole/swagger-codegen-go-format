# DeploymentTarget

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The internal id of this target instance. | [optional] [default to null]
**UserId** | **int32** | The owner&#39;s user id. | [default to null]
**Type_** | **string** | The type of target environment, e.g. &#x60;cloudfoundry&#x60;, &#x60;aws&#x60;, etc.. | [default to null]
**Url** | **string** | The URL endpoint that the target is accessible at. | [default to null]
**UserName** | **string** | The username to authenticate to the target with. | [default to null]
**Password** | **string** | The password to authenticate to the target with. | [default to null]
**Name** | **string** | The user-provided label for this target. | [default to null]
**Organization** | **string** | The organization under which a project will be deployed to this target. | [default to null]
**Space** | **string** | The space within an organization that a project will be deployed under on this target. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


