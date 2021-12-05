# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePiggyBank**](PiggyBanksApi.md#DeletePiggyBank) | **Delete** /api/v1/piggy_banks/{id} | Delete a piggy bank.
[**GetPiggyBank**](PiggyBanksApi.md#GetPiggyBank) | **Get** /api/v1/piggy_banks/{id} | Get a single piggy bank.
[**ListAttachmentByPiggyBank**](PiggyBanksApi.md#ListAttachmentByPiggyBank) | **Get** /api/v1/piggy_banks/{id}/attachments | Lists all attachments.
[**ListEventByPiggyBank**](PiggyBanksApi.md#ListEventByPiggyBank) | **Get** /api/v1/piggy_banks/{id}/events | List all events linked to a piggy bank.
[**ListPiggyBank**](PiggyBanksApi.md#ListPiggyBank) | **Get** /api/v1/piggy_banks | List all piggy banks.
[**StorePiggyBank**](PiggyBanksApi.md#StorePiggyBank) | **Post** /api/v1/piggy_banks | Store a new piggy bank
[**UpdatePiggyBank**](PiggyBanksApi.md#UpdatePiggyBank) | **Put** /api/v1/piggy_banks/{id} | Update existing piggy bank.

# **DeletePiggyBank**
> DeletePiggyBank(ctx, id)
Delete a piggy bank.

Delete a piggy bank.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the piggy bank. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPiggyBank**
> PiggyBankSingle GetPiggyBank(ctx, id)
Get a single piggy bank.

Get a single piggy bank.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the piggy bank. | 

### Return type

[**PiggyBankSingle**](PiggyBankSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAttachmentByPiggyBank**
> AttachmentArray ListAttachmentByPiggyBank(ctx, id, optional)
Lists all attachments.

Lists all attachments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the piggy bank. | 
 **optional** | ***PiggyBanksApiListAttachmentByPiggyBankOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PiggyBanksApiListAttachmentByPiggyBankOpts struct
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

# **ListEventByPiggyBank**
> PiggyBankEventArray ListEventByPiggyBank(ctx, id, optional)
List all events linked to a piggy bank.

List all events linked to a piggy bank (adding and removing money).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the piggy bank | 
 **optional** | ***PiggyBanksApiListEventByPiggyBankOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PiggyBanksApiListEventByPiggyBankOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is 50. | 

### Return type

[**PiggyBankEventArray**](PiggyBankEventArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPiggyBank**
> PiggyBankArray ListPiggyBank(ctx, optional)
List all piggy banks.

List all piggy banks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PiggyBanksApiListPiggyBankOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PiggyBanksApiListPiggyBankOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. The default pagination is 50. | 

### Return type

[**PiggyBankArray**](PiggyBankArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StorePiggyBank**
> PiggyBankSingle StorePiggyBank(ctx, body, name, accountId, targetAmount, currentAmount, startDate, targetDate, order, active, notes, objectGroupId, objectGroupTitle)
Store a new piggy bank

Creates a new piggy bank. The data required can be submitted as a JSON body or as a list of parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PiggyBankStore**](PiggyBankStore.md)| JSON array or key&#x3D;value pairs with the necessary piggy bank information. See the model for the exact specifications. | 
  **name** | **string**|  | 
  **accountId** | **string**|  | 
  **targetAmount** | **string**|  | 
  **currentAmount** | **string**|  | 
  **startDate** | **string**|  | 
  **targetDate** | **string**|  | 
  **order** | **int32**|  | 
  **active** | **bool**|  | 
  **notes** | **string**|  | 
  **objectGroupId** | **string**|  | 
  **objectGroupTitle** | **string**|  | 

### Return type

[**PiggyBankSingle**](PiggyBankSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePiggyBank**
> PiggyBankSingle UpdatePiggyBank(ctx, body, name, accountId, currencyId, currencyCode, targetAmount, currentAmount, startDate, targetDate, order, active, notes, objectGroupId, objectGroupTitle, id)
Update existing piggy bank.

Update existing piggy bank.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PiggyBankUpdate**](PiggyBankUpdate.md)| JSON array with updated piggy bank information. See the model for the exact specifications. | 
  **name** | **string**|  | 
  **accountId** | **string**|  | 
  **currencyId** | **string**|  | 
  **currencyCode** | **string**|  | 
  **targetAmount** | **string**|  | 
  **currentAmount** | **string**|  | 
  **startDate** | **string**|  | 
  **targetDate** | **string**|  | 
  **order** | **int32**|  | 
  **active** | **bool**|  | 
  **notes** | **string**|  | 
  **objectGroupId** | **string**|  | 
  **objectGroupTitle** | **string**|  | 
  **id** | **int32**| The ID of the piggy bank | 

### Return type

[**PiggyBankSingle**](PiggyBankSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

