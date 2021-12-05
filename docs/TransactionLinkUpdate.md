# TransactionLinkUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkTypeId** | **string** | The link type ID to use. Use this field OR use the link_type_name field. | [optional] [default to null]
**LinkTypeName** | **string** | The link type name to use. Use this field OR use the link_type_id field. | [optional] [default to null]
**InwardId** | **string** | The inward transaction transaction_journal_id for the link. This becomes the &#x27;is paid by&#x27; transaction of the set. | [optional] [default to null]
**OutwardId** | **string** | The outward transaction transaction_journal_id for the link. This becomes the &#x27;pays for&#x27; transaction of the set. | [optional] [default to null]
**Notes** | **string** | Optional. Some notes. If you submit an empty string the current notes will be removed | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

