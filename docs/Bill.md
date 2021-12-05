# Bill

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**CurrencyId** | **string** | Use either currency_id or currency_code | [optional] [default to null]
**CurrencyCode** | **string** | Use either currency_id or currency_code | [optional] [default to null]
**CurrencySymbol** | **string** |  | [optional] [default to null]
**CurrencyDecimalPlaces** | **int32** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**AmountMin** | **string** |  | [default to null]
**AmountMax** | **string** |  | [default to null]
**Date** | [**time.Time**](time.Time.md) |  | [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The date after which this bill is no longer valid or applicable | [optional] [default to null]
**ExtensionDate** | [**time.Time**](time.Time.md) | The date before which the bill must be renewed (or cancelled) | [optional] [default to null]
**RepeatFreq** | **string** | How often the bill must be paid. | [default to null]
**Skip** | **int32** | How often the bill must be skipped. 1 means a bi-monthly bill. | [optional] [default to null]
**Active** | **bool** | If the bill is active. | [optional] [default to null]
**Order** | **int32** | Order of the bill. | [optional] [default to null]
**Notes** | **string** |  | [optional] [default to null]
**NextExpectedMatch** | [**time.Time**](time.Time.md) | When the bill is expected to be due. | [optional] [default to null]
**NextExpectedMatchDiff** | **string** | Formatted (locally) when the bill is due. | [optional] [default to null]
**ObjectGroupId** | **string** | The group ID of the group this object is part of. NULL if no group. | [optional] [default to null]
**ObjectGroupOrder** | **int32** | The order of the group. At least 1, for the highest sorting. | [optional] [default to null]
**ObjectGroupTitle** | **string** | The name of the group. NULL if no group. | [optional] [default to null]
**PayDates** | [**[]time.Time**](time.Time.md) | Array of future dates when the bill is expected to be paid. Autogenerated. | [optional] [default to null]
**PaidDates** | [**[]BillPaidDates**](Bill_paid_dates.md) | Array of past transactions when the bill was paid. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
