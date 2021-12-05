# RuleTrigger

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Type_** | **string** | The type of thing this trigger responds to. A limited set is possible | [default to null]
**Value** | **string** | The accompanying value the trigger responds to. This value is often mandatory, but this depends on the trigger. | [default to null]
**Order** | **int32** | Order of the trigger | [optional] [default to null]
**Active** | **bool** | If the trigger is active. Defaults to true. | [optional] [default to null]
**StopProcessing** | **bool** | When true, other triggers will not be checked if this trigger was triggered. Defaults to false. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

