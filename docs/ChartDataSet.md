# ChartDataSet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | This is the title of the current set. It can refer to an account, a budget or another object (by name). | [optional] [default to null]
**CurrencyId** | **string** | The currency ID of the currency associated to the data in the entries. | [optional] [default to null]
**CurrencyCode** | **string** |  | [optional] [default to null]
**CurrencySymbol** | **string** |  | [optional] [default to null]
**CurrencyDecimalPlaces** | **int32** | Number of decimals for the currency associated to the data in the entries. | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Type_** | **string** | Indicated the type of chart that is expected to be rendered. You can safely ignore this if you want. | [optional] [default to null]
**YAxisID** | **int32** | Used to indicate the Y axis for this data set. Is usually between 0 and 1 (left and right side of the chart). | [optional] [default to null]
**Entries** | [**[]ChartDataPoint**](ChartDataPoint.md) | The actual entries for this data set. They &#x27;key&#x27; value is the label for the data point. The value is the actual (numerical) value. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

