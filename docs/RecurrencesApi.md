# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecurrence**](RecurrencesApi.md#DeleteRecurrence) | **Delete** /api/v1/recurrences/{id} | Delete a recurring transaction.
[**GetRecurrence**](RecurrencesApi.md#GetRecurrence) | **Get** /api/v1/recurrences/{id} | Get a single recurring transaction.
[**ListRecurrence**](RecurrencesApi.md#ListRecurrence) | **Get** /api/v1/recurrences | List all recurring transactions.
[**ListTransactionByRecurrence**](RecurrencesApi.md#ListTransactionByRecurrence) | **Get** /api/v1/recurrences/{id}/transactions | List all transactions created by a recurring transaction.
[**StoreRecurrence**](RecurrencesApi.md#StoreRecurrence) | **Post** /api/v1/recurrences | Store a new recurring transaction
[**UpdateRecurrence**](RecurrencesApi.md#UpdateRecurrence) | **Put** /api/v1/recurrences/{id} | Update existing recurring transaction.

# **DeleteRecurrence**
> DeleteRecurrence(ctx, id)
Delete a recurring transaction.

Delete a recurring transaction. Transactions created by the recurring transaction will not be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the recurring transaction. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecurrence**
> RecurrenceSingle GetRecurrence(ctx, id)
Get a single recurring transaction.

Get a single recurring transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the recurring transaction. | 

### Return type

[**RecurrenceSingle**](RecurrenceSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRecurrence**
> RecurrenceArray ListRecurrence(ctx, optional)
List all recurring transactions.

List all recurring transactions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecurrencesApiListRecurrenceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecurrencesApiListRecurrenceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. The default pagination is 50. | 

### Return type

[**RecurrenceArray**](RecurrenceArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransactionByRecurrence**
> TransactionArray ListTransactionByRecurrence(ctx, id, optional)
List all transactions created by a recurring transaction.

List all transactions created by a recurring transaction, optionally limited to the date ranges specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the recurring transaction. | 
 **optional** | ***RecurrencesApiListTransactionByRecurrenceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecurrencesApiListTransactionByRecurrenceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is 50. | 
 **start** | **optional.String**| A date formatted YYYY-MM-DD. Both the start and end date must be present.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD. Both the start and end date must be present.  | 
 **type_** | [**optional.Interface of TransactionTypeFilter**](.md)| Optional filter on the transaction type(s) returned | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreRecurrence**
> RecurrenceSingle StoreRecurrence(ctx, body, type_, title, description, firstDate, repeatUntil, nrOfRepetitions, applyRules, active, notes, repetitions, transactions)
Store a new recurring transaction

Creates a new recurring transaction. The data required can be submitted as a JSON body or as a list of parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RecurrenceStore**](RecurrenceStore.md)| JSON array or key&#x3D;value pairs with the necessary recurring transaction information. See the model for the exact specifications. | 
  **type_** | **string**|  | 
  **title** | **string**|  | 
  **description** | **string**|  | 
  **firstDate** | **string**|  | 
  **repeatUntil** | **string**|  | 
  **nrOfRepetitions** | **int32**|  | 
  **applyRules** | **bool**|  | 
  **active** | **bool**|  | 
  **notes** | **string**|  | 
  **repetitions** | [**[]RecurrenceRepetitionStore**](RecurrenceRepetitionStore.md)|  | 
  **transactions** | [**[]RecurrenceTransactionStore**](RecurrenceTransactionStore.md)|  | 

### Return type

[**RecurrenceSingle**](RecurrenceSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRecurrence**
> RecurrenceSingle UpdateRecurrence(ctx, body, title, description, firstDate, repeatUntil, nrOfRepetitions, applyRules, active, notes, repetitions, transactions, id)
Update existing recurring transaction.

Update existing recurring transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RecurrenceUpdate**](RecurrenceUpdate.md)| JSON array with updated recurring transaction information. See the model for the exact specifications. | 
  **title** | **string**|  | 
  **description** | **string**|  | 
  **firstDate** | **string**|  | 
  **repeatUntil** | **string**|  | 
  **nrOfRepetitions** | **int32**|  | 
  **applyRules** | **bool**|  | 
  **active** | **bool**|  | 
  **notes** | **string**|  | 
  **repetitions** | [**[]RecurrenceRepetitionUpdate**](RecurrenceRepetitionUpdate.md)|  | 
  **transactions** | [**[]RecurrenceTransactionUpdate**](RecurrenceTransactionUpdate.md)|  | 
  **id** | **int32**| The ID of the recurring transaction. | 

### Return type

[**RecurrenceSingle**](RecurrenceSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

