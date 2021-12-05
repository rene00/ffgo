# Account

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Active** | **bool** | If omitted, defaults to true. | [optional] [default to null]
**Order** | **int32** | Order of the account. Is NULL if account is not asset or liability. | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Type_** | **string** | Can only be one one these account types. import, initial-balance and reconciliation cannot be set manually. | [default to null]
**AccountRole** | **string** | Is only mandatory when the type is asset. | [optional] [default to null]
**CurrencyId** | **string** | Use either currency_id or currency_code. Defaults to the user&#x27;s default currency. | [optional] [default to null]
**CurrencyCode** | **string** | Use either currency_id or currency_code. Defaults to the user&#x27;s default currency. | [optional] [default to null]
**CurrencySymbol** | **string** |  | [optional] [default to null]
**CurrencyDecimalPlaces** | **int32** |  | [optional] [default to null]
**CurrentBalance** | **string** |  | [optional] [default to null]
**CurrentBalanceDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Iban** | **string** |  | [optional] [default to null]
**Bic** | **string** |  | [optional] [default to null]
**AccountNumber** | **string** |  | [optional] [default to null]
**OpeningBalance** | **string** | Represents the opening balance, the initial amount this account holds. | [optional] [default to null]
**CurrentDebt** | **string** | Represents the current debt for liabilities. | [optional] [default to null]
**OpeningBalanceDate** | [**time.Time**](time.Time.md) | Represents the date of the opening balance. | [optional] [default to null]
**VirtualBalance** | **string** |  | [optional] [default to null]
**IncludeNetWorth** | **bool** | If omitted, defaults to true. | [optional] [default to null]
**CreditCardType** | **string** | Mandatory when the account_role is ccAsset. Can only be monthlyFull or null. | [optional] [default to null]
**MonthlyPaymentDate** | [**time.Time**](time.Time.md) | Mandatory when the account_role is ccAsset. Moment at which CC payment installments are asked for by the bank. | [optional] [default to null]
**LiabilityType** | **string** | Mandatory when type is liability. Specifies the exact type. | [optional] [default to null]
**LiabilityDirection** | **string** | &#x27;credit&#x27; indicates somebody owes you the liability. &#x27;debit&#x27; Indicates you owe this debt yourself. Works only for liabiltiies. | [optional] [default to null]
**Interest** | **string** | Mandatory when type is liability. Interest percentage. | [optional] [default to null]
**InterestPeriod** | **string** | Mandatory when type is liability. Period over which the interest is calculated. | [optional] [default to null]
**Notes** | **string** |  | [optional] [default to null]
**Latitude** | **float64** | Latitude of the accounts&#x27;s location, if applicable. Can be used to draw a map. | [optional] [default to null]
**Longitude** | **float64** | Latitude of the accounts&#x27;s location, if applicable. Can be used to draw a map. | [optional] [default to null]
**ZoomLevel** | **int32** | Zoom level for the map, if drawn. This to set the box right. Unfortunately this is a proprietary value because each map provider has different zoom levels. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

