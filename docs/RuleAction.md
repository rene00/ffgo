# RuleAction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Type_** | **string** | The type of thing this action will do. A limited set is possible. | [default to null]
**Value** | **string** | The accompanying value the action will set, change or update. Can be empty, but for some types this value is mandatory. | [default to null]
**Order** | **int32** | Order of the action | [optional] [default to null]
**Active** | **bool** | If the action is active. Defaults to true. | [optional] [default to null]
**StopProcessing** | **bool** | When true, other actions will not be fired after this action has fired. Defaults to false. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

