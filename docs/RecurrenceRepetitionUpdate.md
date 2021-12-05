# RecurrenceRepetitionUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | The type of the repetition. ndom means: the n-th weekday of the month, where you can also specify which day of the week. | [optional] [default to null]
**Moment** | **string** | Information that defined the type of repetition. - For &#x27;daily&#x27;, this is empty. - For &#x27;weekly&#x27;, it is day of the week between 1 and 7 (Monday - Sunday). - For &#x27;ndom&#x27;, it is &#x27;1,2&#x27; or &#x27;4,5&#x27; or something else, where the first number is the week in the month, and the second number is the day in the week (between 1 and 7). &#x27;2,3&#x27; means: the 2nd Wednesday of the month - For &#x27;monthly&#x27; it is the day of the month (1 - 31) - For yearly, it is a full date, ie &#x27;2018-09-17&#x27;. The year you use does not matter.  | [optional] [default to null]
**Skip** | **int32** | How many occurrences to skip. 0 means skip nothing. 1 means every other. | [optional] [default to null]
**Weekend** | **int32** | How to respond when the recurring transaction falls in the weekend. Possible values: 1. Do nothing, just create it 2. Create no transaction. 3. Skip to the previous Friday. 4. Skip to the next Monday.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

