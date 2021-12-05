# AccountUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to null]
**Iban** | **string** |  | [optional] [default to null]
**Bic** | **string** |  | [optional] [default to null]
**AccountNumber** | **string** |  | [optional] [default to null]
**OpeningBalance** | **string** |  | [optional] [default to null]
**OpeningBalanceDate** | **string** |  | [optional] [default to null]
**VirtualBalance** | **string** |  | [optional] [default to null]
**CurrencyId** | **string** | Use either currency_id or currency_code. Defaults to the user&#x27;s default currency. | [optional] [default to null]
**CurrencyCode** | **string** | Use either currency_id or currency_code. Defaults to the user&#x27;s default currency. | [optional] [default to null]
**Active** | **bool** | If omitted, defaults to true. | [optional] [default to null]
**Order** | **int32** | Order of the account | [optional] [default to null]
**IncludeNetWorth** | **bool** | If omitted, defaults to true. | [optional] [default to null]
**AccountRole** | **string** | Is only mandatory when the type is asset. | [optional] [default to null]
**CreditCardType** | **string** | Mandatory when the account_role is ccAsset. Can only be monthlyFull or null. | [optional] [default to null]
**MonthlyPaymentDate** | **string** | Mandatory when the account_role is ccAsset. Moment at which CC payment installments are asked for by the bank. | [optional] [default to null]
**LiabilityType** | **string** | Mandatory when type is liability. Specifies the exact type. | [optional] [default to null]
**Interest** | **string** | Mandatory when type is liability. Interest percentage. | [optional] [default to null]
**InterestPeriod** | **string** | Mandatory when type is liability. Period over which the interest is calculated. | [optional] [default to null]
**Notes** | **string** |  | [optional] [default to null]
**Latitude** | **float64** | Latitude of the account&#x27;s location, if applicable. Can be used to draw a map. If omitted, the existing location will be kept. If submitted as NULL, the current location will be removed. | [optional] [default to null]
**Longitude** | **float64** | Latitude of the account&#x27;s location, if applicable. Can be used to draw a map. If omitted, the existing location will be kept. If submitted as NULL, the current location will be removed. | [optional] [default to null]
**ZoomLevel** | **int32** | Zoom level for the map, if drawn. This to set the box right. Unfortunately this is a proprietary value because each map provider has different zoom levels. If omitted, the existing location will be kept. If submitted as NULL, the current location will be removed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

