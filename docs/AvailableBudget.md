# AvailableBudget

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**CurrencyId** | **string** | Use either currency_id or currency_code. | [optional] [default to null]
**CurrencyCode** | **string** | Use either currency_id or currency_code. | [optional] [default to null]
**CurrencySymbol** | **string** |  | [optional] [default to null]
**CurrencyDecimalPlaces** | **int32** |  | [optional] [default to null]
**Amount** | **string** |  | [default to null]
**Start** | [**time.Time**](time.Time.md) | Start date of the available budget. | [default to null]
**End** | [**time.Time**](time.Time.md) | End date of the available budget. | [default to null]
**SpentInBudgets** | [**[]BudgetSpent**](BudgetSpent.md) |  | [optional] [default to null]
**SpentOutsideBudget** | [**[]BudgetSpent**](BudgetSpent.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

