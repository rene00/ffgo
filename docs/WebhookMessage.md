# WebhookMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Sent** | **bool** | If this message is sent yet. | [optional] [default to null]
**Errored** | **bool** | If this message has errored out. | [optional] [default to null]
**WebhookId** | **string** | The ID of the webhook this message belongs to. | [optional] [default to null]
**Uuid** | **string** | Long UUID string for identification of this webhook message. | [optional] [default to null]
**String_** | **string** | The actual message that is sent or will be sent as JSON string. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

