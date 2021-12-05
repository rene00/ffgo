# Recurrence

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]
**Title** | **string** |  | [optional] [default to null]
**Description** | **string** | Not to be confused with the description of the actual transaction(s) being created. | [optional] [default to null]
**FirstDate** | [**time.Time**](time.Time.md) | First time the recurring transaction will fire. Must be after today. | [optional] [default to null]
**LatestDate** | [**time.Time**](time.Time.md) | Last time the recurring transaction has fired. | [optional] [default to null]
**RepeatUntil** | [**time.Time**](time.Time.md) | Date until the recurring transaction can fire. Use either this field or repetitions. | [optional] [default to null]
**NrOfRepetitions** | **int32** | Max number of created transactions. Use either this field or repeat_until. | [optional] [default to null]
**ApplyRules** | **bool** | Whether or not to fire the rules after the creation of a transaction. | [optional] [default to null]
**Active** | **bool** | If the recurrence is even active. | [optional] [default to null]
**Notes** | **string** |  | [optional] [default to null]
**Repetitions** | [**[]RecurrenceRepetition**](RecurrenceRepetition.md) |  | [optional] [default to null]
**Transactions** | [**[]RecurrenceTransaction**](RecurrenceTransaction.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

