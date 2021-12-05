# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTag**](TagsApi.md#DeleteTag) | **Delete** /api/v1/tags/{tag} | Delete an tag.
[**GetTag**](TagsApi.md#GetTag) | **Get** /api/v1/tags/{tag} | Get a single tag.
[**ListAttachmentByTag**](TagsApi.md#ListAttachmentByTag) | **Get** /api/v1/tags/{tag}/attachments | Lists all attachments.
[**ListTag**](TagsApi.md#ListTag) | **Get** /api/v1/tags | List all tags.
[**ListTransactionByTag**](TagsApi.md#ListTransactionByTag) | **Get** /api/v1/tags/{tag}/transactions | List all transactions with this tag.
[**StoreTag**](TagsApi.md#StoreTag) | **Post** /api/v1/tags | Store a new tag
[**UpdateTag**](TagsApi.md#UpdateTag) | **Put** /api/v1/tags/{tag} | Update existing tag.

# **DeleteTag**
> DeleteTag(ctx, tag)
Delete an tag.

Delete an tag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tag** | **string**| Either the tag itself or the tag ID. If you use the tag itself, and it contains international (non-ASCII) characters, your milage may vary. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTag**
> TagSingle GetTag(ctx, tag, optional)
Get a single tag.

Get a single tag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tag** | **string**| Either the tag itself or the tag ID. If you use the tag itself, and it contains international (non-ASCII) characters, your milage may vary. | 
 **optional** | ***TagsApiGetTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagsApiGetTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number | 

### Return type

[**TagSingle**](TagSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAttachmentByTag**
> AttachmentArray ListAttachmentByTag(ctx, tag, optional)
Lists all attachments.

Lists all attachments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tag** | **string**| Either the tag itself or the tag ID. | 
 **optional** | ***TagsApiListAttachmentByTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagsApiListAttachmentByTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is 50. | 

### Return type

[**AttachmentArray**](AttachmentArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTag**
> TagArray ListTag(ctx, optional)
List all tags.

List all of the user's tags.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TagsApiListTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagsApiListTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number | 

### Return type

[**TagArray**](TagArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransactionByTag**
> TransactionArray ListTransactionByTag(ctx, tag, optional)
List all transactions with this tag.

List all transactions with this tag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tag** | **string**| Either the tag itself or the tag ID. | 
 **optional** | ***TagsApiListTransactionByTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagsApiListTransactionByTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is 50. | 
 **start** | **optional.String**| A date formatted YYYY-MM-DD. This is the start date of the selected range (inclusive).  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD. This is the end date of the selected range (inclusive).  | 
 **type_** | [**optional.Interface of TransactionTypeFilter**](.md)| Optional filter on the transaction type(s) returned. | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreTag**
> TagSingle StoreTag(ctx, body, tag, date, description, latitude, longitude, zoomLevel)
Store a new tag

Creates a new tag. The data required can be submitted as a JSON body or as a list of parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagModelStore**](TagModelStore.md)| JSON array or key&#x3D;value pairs with the necessary tag information. See the model for the exact specifications. | 
  **tag** | **string**|  | 
  **date** | **string**|  | 
  **description** | **string**|  | 
  **latitude** | **float64**|  | 
  **longitude** | **float64**|  | 
  **zoomLevel** | **int32**|  | 

### Return type

[**TagSingle**](TagSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTag**
> TagSingle UpdateTag(ctx, body, tag, date, description, latitude, longitude, zoomLevel, tag)
Update existing tag.

Update existing tag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagModelUpdate**](TagModelUpdate.md)| JSON array with updated tag information. See the model for the exact specifications. | 
  **tag** | **string**|  | 
  **date** | **string**|  | 
  **description** | **string**|  | 
  **latitude** | **float64**|  | 
  **longitude** | **float64**|  | 
  **zoomLevel** | **int32**|  | 
  **tag** | **string**| Either the tag itself or the tag ID. If you use the tag itself, and it contains international (non-ASCII) characters, your milage may vary. | 

### Return type

[**TagSingle**](TagSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

