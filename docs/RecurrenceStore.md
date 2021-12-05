# RecurrenceStore

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Title** | **string** |  | [default to null]
**Description** | **string** | Not to be confused with the description of the actual transaction(s) being created. | [optional] [default to null]
**FirstDate** | **string** | First time the recurring transaction will fire. Must be after today. | [default to null]
**RepeatUntil** | **string** | Date until the recurring transaction can fire. Use either this field or repetitions. | [default to null]
**NrOfRepetitions** | **int32** | Max number of created transactions. Use either this field or repeat_until. | [optional] [default to null]
**ApplyRules** | **bool** | Whether or not to fire the rules after the creation of a transaction. | [optional] [default to null]
**Active** | **bool** | If the recurrence is even active. | [optional] [default to null]
**Notes** | **string** |  | [optional] [default to null]
**Repetitions** | [**[]RecurrenceRepetitionStore**](RecurrenceRepetitionStore.md) |  | [default to null]
**Transactions** | [**[]RecurrenceTransactionStore**](RecurrenceTransactionStore.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

