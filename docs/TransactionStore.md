# TransactionStore

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorIfDuplicateHash** | **bool** | Break if the submitted transaction exists already. | [optional] [default to null]
**ApplyRules** | **bool** | Whether or not to apply rules when submitting transaction. | [optional] [default to null]
**FireWebhooks** | **bool** | Whether or not to fire the webhooks that are related to this event. | [optional] [default to true]
**GroupTitle** | **string** | Title of the transaction if it has been split in more than one piece. Empty otherwise. | [optional] [default to null]
**Transactions** | [**[]TransactionSplitStore**](TransactionSplitStore.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

