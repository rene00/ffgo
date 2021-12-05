# WebhookAttempt

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**WebhookMessageId** | **string** | The ID of the webhook message this attempt belongs to. | [optional] [default to null]
**StatusCode** | **int32** | The HTTP status code of the error, if any. | [optional] [default to null]
**Logs** | **string** | Internal log for this attempt. May contain sensitive user data. | [optional] [default to null]
**Response** | **string** | Webhook receiver response for this attempt, if any. May contain sensitive user data. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

