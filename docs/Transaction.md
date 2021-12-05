# Transaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**User** | **string** | User ID | [optional] [default to null]
**GroupTitle** | **string** | Title of the transaction if it has been split in more than one piece. Empty otherwise. | [optional] [default to null]
**Transactions** | [**[]TransactionSplit**](TransactionSplit.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

