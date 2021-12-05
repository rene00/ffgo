# Budget

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Active** | **bool** |  | [optional] [default to null]
**Order** | **int32** |  | [optional] [default to null]
**AutoBudgetType** | **string** | The type of auto-budget that Firefly III must create. | [optional] [default to null]
**AutoBudgetCurrencyId** | **string** | Use either currency_id or currency_code. Defaults to the user&#x27;s default currency. | [optional] [default to null]
**AutoBudgetCurrencyCode** | **string** | Use either currency_id or currency_code. Defaults to the user&#x27;s default currency. | [optional] [default to null]
**AutoBudgetAmount** | **string** |  | [optional] [default to null]
**AutoBudgetPeriod** | **string** | Period for the auto budget | [optional] [default to null]
**Spent** | [**[]BudgetSpent**](BudgetSpent.md) | Information on how much was spent in this budget. Is only filled in when the start and end date are submitted. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

