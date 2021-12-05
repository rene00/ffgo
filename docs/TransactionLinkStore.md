# TransactionLinkStore

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkTypeId** | **string** | The link type ID to use. You can also use the link_type_name field. | [default to null]
**LinkTypeName** | **string** | The link type name to use. You can also use the link_type_id field. | [optional] [default to null]
**InwardId** | **string** | The inward transaction transaction_journal_id for the link. This becomes the &#x27;is paid by&#x27; transaction of the set. | [default to null]
**OutwardId** | **string** | The outward transaction transaction_journal_id for the link. This becomes the &#x27;pays for&#x27; transaction of the set. | [default to null]
**Notes** | **string** | Optional. Some notes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

