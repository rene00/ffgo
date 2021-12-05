# RecurrenceUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | [optional] [default to null]
**Description** | **string** | Not to be confused with the description of the actual transaction(s) being created. | [optional] [default to null]
**FirstDate** | **string** | First time the recurring transaction will fire. | [optional] [default to null]
**RepeatUntil** | **string** | Date until the recurring transaction can fire. After that date, it&#x27;s basically inactive. Use either this field or repetitions. | [optional] [default to null]
**NrOfRepetitions** | **int32** | Max number of created transactions. Use either this field or repeat_until. | [optional] [default to null]
**ApplyRules** | **bool** | Whether or not to fire the rules after the creation of a transaction. | [optional] [default to null]
**Active** | **bool** | If the recurrence is even active. | [optional] [default to null]
**Notes** | **string** |  | [optional] [default to null]
**Repetitions** | [**[]RecurrenceRepetitionUpdate**](RecurrenceRepetitionUpdate.md) |  | [optional] [default to null]
**Transactions** | [**[]RecurrenceTransactionUpdate**](RecurrenceTransactionUpdate.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

