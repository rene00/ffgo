# RecurrenceTransaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | [default to null]
**Amount** | **string** | Amount of the transaction. | [default to null]
**ForeignAmount** | **string** | Foreign amount of the transaction. | [optional] [default to null]
**CurrencyId** | **string** | Submit either a currency_id or a currency_code. | [optional] [default to null]
**CurrencyCode** | **string** | Submit either a currency_id or a currency_code. | [optional] [default to null]
**CurrencySymbol** | **string** |  | [optional] [default to null]
**CurrencyDecimalPlaces** | **int32** | Number of decimals in the currency | [optional] [default to null]
**ForeignCurrencyId** | **string** | Submit either a foreign_currency_id or a foreign_currency_code, or neither. | [optional] [default to null]
**ForeignCurrencyCode** | **string** | Submit either a foreign_currency_id or a foreign_currency_code, or neither. | [optional] [default to null]
**ForeignCurrencySymbol** | **string** |  | [optional] [default to null]
**ForeignCurrencyDecimalPlaces** | **int32** | Number of decimals in the currency | [optional] [default to null]
**BudgetId** | **string** | The budget ID for this transaction. | [optional] [default to null]
**BudgetName** | **string** | The name of the budget to be used. If the budget name is unknown, the ID will be used or the value will be ignored. | [optional] [default to null]
**CategoryId** | **string** | Category ID for this transaction. | [optional] [default to null]
**CategoryName** | **string** | Category name for this transaction. | [optional] [default to null]
**SourceId** | **string** | ID of the source account. Submit either this or source_name. | [optional] [default to null]
**SourceName** | **string** | Name of the source account. Submit either this or source_id. | [optional] [default to null]
**SourceIban** | **string** |  | [optional] [default to null]
**SourceType** | [***AccountTypeProperty**](AccountTypeProperty.md) |  | [optional] [default to null]
**DestinationId** | **string** | ID of the destination account. Submit either this or destination_name. | [optional] [default to null]
**DestinationName** | **string** | Name of the destination account. Submit either this or destination_id. | [optional] [default to null]
**DestinationIban** | **string** |  | [optional] [default to null]
**DestinationType** | [***AccountTypeProperty**](AccountTypeProperty.md) |  | [optional] [default to null]
**Tags** | **[]string** | Array of tags. | [optional] [default to null]
**PiggyBankId** | **string** | Optional. Use either this or the piggy_bank_name | [optional] [default to null]
**PiggyBankName** | **string** | Optional. Use either this or the piggy_bank_id | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

