# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTransaction**](TransactionsApi.md#DeleteTransaction) | **Delete** /api/v1/transactions/{id} | Delete a transaction.
[**DeleteTransactionJournal**](TransactionsApi.md#DeleteTransactionJournal) | **Delete** /api/v1/transaction-journals/{id} | Delete split from transaction
[**GetTransaction**](TransactionsApi.md#GetTransaction) | **Get** /api/v1/transactions/{id} | Get a single transaction.
[**GetTransactionByJournal**](TransactionsApi.md#GetTransactionByJournal) | **Get** /api/v1/transaction-journals/{id} | Get a single transaction, based on one of the underlying transaction journals (transaction splits).
[**ListAttachmentByTransaction**](TransactionsApi.md#ListAttachmentByTransaction) | **Get** /api/v1/transactions/{id}/attachments | Lists all attachments.
[**ListEventByTransaction**](TransactionsApi.md#ListEventByTransaction) | **Get** /api/v1/transactions/{id}/piggy_bank_events | Lists all piggy bank events.
[**ListLinksByJournal**](TransactionsApi.md#ListLinksByJournal) | **Get** /api/v1/transaction-journals/{id}/links | Lists all the transaction links for an individual journal (individual split).
[**ListTransaction**](TransactionsApi.md#ListTransaction) | **Get** /api/v1/transactions | List all the user&#x27;s transactions. 
[**StoreTransaction**](TransactionsApi.md#StoreTransaction) | **Post** /api/v1/transactions | Store a new transaction
[**UpdateTransaction**](TransactionsApi.md#UpdateTransaction) | **Put** /api/v1/transactions/{id} | Update existing transaction.

# **DeleteTransaction**
> DeleteTransaction(ctx, id)
Delete a transaction.

Delete a transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the transaction. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTransactionJournal**
> DeleteTransactionJournal(ctx, id)
Delete split from transaction

Delete an individual journal (split) from a transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the transaction journal (the split) you wish to delete. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransaction**
> TransactionSingle GetTransaction(ctx, id)
Get a single transaction.

Get a single transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the transaction. | 

### Return type

[**TransactionSingle**](TransactionSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionByJournal**
> TransactionSingle GetTransactionByJournal(ctx, id)
Get a single transaction, based on one of the underlying transaction journals (transaction splits).

Get a single transaction by underlying journal (split).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the transaction journal (split). | 

### Return type

[**TransactionSingle**](TransactionSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAttachmentByTransaction**
> AttachmentArray ListAttachmentByTransaction(ctx, id, optional)
Lists all attachments.

Lists all attachments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the transaction. | 
 **optional** | ***TransactionsApiListAttachmentByTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionsApiListAttachmentByTransactionOpts struct
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

# **ListEventByTransaction**
> PiggyBankEventArray ListEventByTransaction(ctx, id, optional)
Lists all piggy bank events.

Lists all piggy bank events.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the transaction. | 
 **optional** | ***TransactionsApiListEventByTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionsApiListEventByTransactionOpts struct
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

# **ListLinksByJournal**
> TransactionLinkArray ListLinksByJournal(ctx, id, optional)
Lists all the transaction links for an individual journal (individual split).

Lists all the transaction links for an individual journal (a split). Don't use the group ID, you need the actual underlying journal (the split).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the transaction journal / the split. | 
 **optional** | ***TransactionsApiListLinksByJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionsApiListLinksByJournalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is 50. | 

### Return type

[**TransactionLinkArray**](TransactionLinkArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransaction**
> TransactionArray ListTransaction(ctx, optional)
List all the user's transactions. 

List all the user's transactions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransactionsApiListTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionsApiListTransactionOpts struct
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

# **StoreTransaction**
> TransactionSingle StoreTransaction(ctx, body, errorIfDuplicateHash, applyRules, fireWebhooks, groupTitle, transactions)
Store a new transaction

Creates a new transaction. The data required can be submitted as a JSON body or as a list of parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionStore**](TransactionStore.md)| JSON array or key&#x3D;value pairs with the necessary transaction information. See the model for the exact specifications. | 
  **errorIfDuplicateHash** | **bool**|  | 
  **applyRules** | **bool**|  | 
  **fireWebhooks** | **bool**|  | 
  **groupTitle** | **string**|  | 
  **transactions** | [**[]TransactionSplitStore**](TransactionSplitStore.md)|  | 

### Return type

[**TransactionSingle**](TransactionSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTransaction**
> TransactionSingle UpdateTransaction(ctx, body, applyRules, fireWebhooks, groupTitle, transactions, id)
Update existing transaction.

Update an existing transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionUpdate**](TransactionUpdate.md)| JSON array with updated transaction information. See the model for the exact specifications. | 
  **applyRules** | **bool**|  | 
  **fireWebhooks** | **bool**|  | 
  **groupTitle** | **string**|  | 
  **transactions** | [**[]TransactionSplitUpdate**](TransactionSplitUpdate.md)|  | 
  **id** | **int32**| The ID of the transaction. | 

### Return type

[**TransactionSingle**](TransactionSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

