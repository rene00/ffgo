# BasicSummaryEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | This is a reference to the type of info shared, not influenced by translations or user preferences. The EUR value is a reference to the currency code. Possibilities are: balance-in-ABC, spent-in-ABC, earned-in-ABC, bills-paid-in-ABC, bills-unpaid-in-ABC, left-to-spend-in-ABC and net-worth-in-ABC. | [optional] [default to null]
**Title** | **string** | A translated title for the information shared. | [optional] [default to null]
**MonetaryValue** | **float64** | The amount as a float. | [optional] [default to null]
**CurrencyId** | **string** | The currency ID of the associated currency. | [optional] [default to null]
**CurrencyCode** | **string** |  | [optional] [default to null]
**CurrencySymbol** | **string** |  | [optional] [default to null]
**CurrencyDecimalPlaces** | **int32** | Number of decimals for the associated currency. | [optional] [default to null]
**ValueParsed** | **string** | The amount formatted according to the users locale | [optional] [default to null]
**LocalIcon** | **string** | Reference to a font-awesome icon without the fa- part. | [optional] [default to null]
**SubTitle** | **string** | A short explanation of the amounts origin. Already formatted according to the locale of the user or translated, if relevant. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

