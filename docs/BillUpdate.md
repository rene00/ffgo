# BillUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyId** | **string** | Use either currency_id or currency_code | [optional] [default to null]
**CurrencyCode** | **string** | Use either currency_id or currency_code | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**AmountMin** | **string** |  | [optional] [default to null]
**AmountMax** | **string** |  | [optional] [default to null]
**Date** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The date after which this bill is no longer valid or applicable | [optional] [default to null]
**ExtensionDate** | [**time.Time**](time.Time.md) | The date before which the bill must be renewed (or cancelled) | [optional] [default to null]
**RepeatFreq** | **string** | How often the bill must be paid. | [optional] [default to null]
**Skip** | **int32** | How often the bill must be skipped. 1 means a bi-monthly bill. | [optional] [default to null]
**Active** | **bool** | If the bill is active. | [optional] [default to null]
**Notes** | **string** |  | [optional] [default to null]
**ObjectGroupId** | **string** | The group ID of the group this object is part of. NULL if no group. | [optional] [default to null]
**ObjectGroupTitle** | **string** | The name of the group. NULL if no group. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

