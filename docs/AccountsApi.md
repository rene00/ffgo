# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccount**](AccountsApi.md#DeleteAccount) | **Delete** /api/v1/accounts/{id} | Permanently delete account.
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /api/v1/accounts/{id} | Get single account.
[**ListAccount**](AccountsApi.md#ListAccount) | **Get** /api/v1/accounts | List all accounts.
[**ListAttachmentByAccount**](AccountsApi.md#ListAttachmentByAccount) | **Get** /api/v1/accounts/{id}/attachments | Lists all attachments.
[**ListPiggyBankByAccount**](AccountsApi.md#ListPiggyBankByAccount) | **Get** /api/v1/accounts/{id}/piggy_banks | List all piggy banks related to the account.
[**ListTransactionByAccount**](AccountsApi.md#ListTransactionByAccount) | **Get** /api/v1/accounts/{id}/transactions | List all transactions related to the account.
[**StoreAccount**](AccountsApi.md#StoreAccount) | **Post** /api/v1/accounts | Create new account.
[**UpdateAccount**](AccountsApi.md#UpdateAccount) | **Put** /api/v1/accounts/{id} | Update existing account.

# **DeleteAccount**
> DeleteAccount(ctx, id)
Permanently delete account.

Will permanently delete an account. Any associated transactions and piggy banks are ALSO deleted. Cannot be recovered from. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the account. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccount**
> AccountSingle GetAccount(ctx, id, optional)
Get single account.

Returns a single account by its ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the account. | 
 **optional** | ***AccountsApiGetAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsApiGetAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **optional.String**| A date formatted YYYY-MM-DD. When added to the request, Firefly III will show the account&#x27;s balance on that day.  | 

### Return type

[**AccountSingle**](AccountSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccount**
> AccountArray ListAccount(ctx, optional)
List all accounts.

This endpoint returns a list of all the accounts owned by the authenticated user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsApiListAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsApiListAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. The default pagination is per 50 items. | 
 **date** | **optional.String**| A date formatted YYYY-MM-DD. When added to the request, Firefly III will show the account&#x27;s balance on that day.  | 
 **type_** | [**optional.Interface of AccountTypeFilter**](.md)| Optional filter on the account type(s) returned | 

### Return type

[**AccountArray**](AccountArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAttachmentByAccount**
> AttachmentArray ListAttachmentByAccount(ctx, id, optional)
Lists all attachments.

Lists all attachments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the account. | 
 **optional** | ***AccountsApiListAttachmentByAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsApiListAttachmentByAccountOpts struct
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

# **ListPiggyBankByAccount**
> PiggyBankArray ListPiggyBankByAccount(ctx, id, optional)
List all piggy banks related to the account.

This endpoint returns a list of all the piggy banks connected to the account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the account. | 
 **optional** | ***AccountsApiListPiggyBankByAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsApiListPiggyBankByAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is per 50 items. | 

### Return type

[**PiggyBankArray**](PiggyBankArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransactionByAccount**
> TransactionArray ListTransactionByAccount(ctx, id, optional)
List all transactions related to the account.

This endpoint returns a list of all the transactions connected to the account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the account. | 
 **optional** | ***AccountsApiListTransactionByAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsApiListTransactionByAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is per 50 items. | 
 **limit** | **optional.Int32**| Limits the number of results on one page. | 
 **start** | **optional.String**| A date formatted YYYY-MM-DD.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD.  | 
 **type_** | [**optional.Interface of TransactionTypeFilter**](.md)| Optional filter on the transaction type(s) returned. | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreAccount**
> AccountSingle StoreAccount(ctx, body, name, type_, iban, bic, accountNumber, openingBalance, openingBalanceDate, virtualBalance, currencyId, currencyCode, active, order, includeNetWorth, accountRole, creditCardType, monthlyPaymentDate, liabilityType, liabilityDirection, interest, interestPeriod, notes, latitude, longitude, zoomLevel)
Create new account.

Creates a new account. The data required can be submitted as a JSON body or as a list of parameters (in key=value pairs, like a webform).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountStore**](AccountStore.md)| JSON array with the necessary account information or key&#x3D;value pairs. See the model for the exact specifications. | 
  **name** | **string**|  | 
  **type_** | **string**|  | 
  **iban** | **string**|  | 
  **bic** | **string**|  | 
  **accountNumber** | **string**|  | 
  **openingBalance** | **string**|  | 
  **openingBalanceDate** | **string**|  | 
  **virtualBalance** | **string**|  | 
  **currencyId** | **string**|  | 
  **currencyCode** | **string**|  | 
  **active** | **bool**|  | 
  **order** | **int32**|  | 
  **includeNetWorth** | **bool**|  | 
  **accountRole** | **string**|  | 
  **creditCardType** | **string**|  | 
  **monthlyPaymentDate** | **string**|  | 
  **liabilityType** | **string**|  | 
  **liabilityDirection** | **string**|  | 
  **interest** | **string**|  | 
  **interestPeriod** | **string**|  | 
  **notes** | **string**|  | 
  **latitude** | **float64**|  | 
  **longitude** | **float64**|  | 
  **zoomLevel** | **int32**|  | 

### Return type

[**AccountSingle**](AccountSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccount**
> AccountSingle UpdateAccount(ctx, body, name, iban, bic, accountNumber, openingBalance, openingBalanceDate, virtualBalance, currencyId, currencyCode, active, order, includeNetWorth, accountRole, creditCardType, monthlyPaymentDate, liabilityType, interest, interestPeriod, notes, latitude, longitude, zoomLevel, id)
Update existing account.

Used to update a single account. All fields that are not submitted will be cleared (set to NULL). The model will tell you which fields are mandatory. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountUpdate**](AccountUpdate.md)| JSON array or formdata with updated account information. See the model for the exact specifications. | 
  **name** | **string**|  | 
  **iban** | **string**|  | 
  **bic** | **string**|  | 
  **accountNumber** | **string**|  | 
  **openingBalance** | **string**|  | 
  **openingBalanceDate** | **string**|  | 
  **virtualBalance** | **string**|  | 
  **currencyId** | **string**|  | 
  **currencyCode** | **string**|  | 
  **active** | **bool**|  | 
  **order** | **int32**|  | 
  **includeNetWorth** | **bool**|  | 
  **accountRole** | **string**|  | 
  **creditCardType** | **string**|  | 
  **monthlyPaymentDate** | **string**|  | 
  **liabilityType** | **string**|  | 
  **interest** | **string**|  | 
  **interestPeriod** | **string**|  | 
  **notes** | **string**|  | 
  **latitude** | **float64**|  | 
  **longitude** | **float64**|  | 
  **zoomLevel** | **int32**|  | 
  **id** | **int32**| The ID of the account. | 

### Return type

[**AccountSingle**](AccountSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

