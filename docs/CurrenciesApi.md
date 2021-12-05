# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DefaultCurrency**](CurrenciesApi.md#DefaultCurrency) | **Post** /api/v1/currencies/{code}/default | Make currency default currency.
[**DeleteCurrency**](CurrenciesApi.md#DeleteCurrency) | **Delete** /api/v1/currencies/{code} | Delete a currency.
[**DisableCurrency**](CurrenciesApi.md#DisableCurrency) | **Post** /api/v1/currencies/{code}/disable | Disable a currency.
[**EnableCurrency**](CurrenciesApi.md#EnableCurrency) | **Post** /api/v1/currencies/{code}/enable | Enable a single currency.
[**GetCurrency**](CurrenciesApi.md#GetCurrency) | **Get** /api/v1/currencies/{code} | Get a single currency.
[**GetDefaultCurrency**](CurrenciesApi.md#GetDefaultCurrency) | **Get** /api/v1/currencies/default | Get the user&#x27;s default currency.
[**ListAccountByCurrency**](CurrenciesApi.md#ListAccountByCurrency) | **Get** /api/v1/currencies/{code}/accounts | List all accounts with this currency.
[**ListAvailableBudgetByCurrency**](CurrenciesApi.md#ListAvailableBudgetByCurrency) | **Get** /api/v1/currencies/{code}/available_budgets | List all available budgets with this currency.
[**ListBillByCurrency**](CurrenciesApi.md#ListBillByCurrency) | **Get** /api/v1/currencies/{code}/bills | List all bills with this currency.
[**ListBudgetLimitByCurrency**](CurrenciesApi.md#ListBudgetLimitByCurrency) | **Get** /api/v1/currencies/{code}/budget_limits | List all budget limits with this currency
[**ListCurrency**](CurrenciesApi.md#ListCurrency) | **Get** /api/v1/currencies | List all currencies.
[**ListRecurrenceByCurrency**](CurrenciesApi.md#ListRecurrenceByCurrency) | **Get** /api/v1/currencies/{code}/recurrences | List all recurring transactions with this currency.
[**ListRuleByCurrency**](CurrenciesApi.md#ListRuleByCurrency) | **Get** /api/v1/currencies/{code}/rules | List all rules with this currency.
[**ListTransactionByCurrency**](CurrenciesApi.md#ListTransactionByCurrency) | **Get** /api/v1/currencies/{code}/transactions | List all transactions with this currency.
[**StoreCurrency**](CurrenciesApi.md#StoreCurrency) | **Post** /api/v1/currencies | Store a new currency
[**UpdateCurrency**](CurrenciesApi.md#UpdateCurrency) | **Put** /api/v1/currencies/{code} | Update existing currency.

# **DefaultCurrency**
> CurrencySingle DefaultCurrency(ctx, code)
Make currency default currency.

Make this currency the default currency for the user. If the currency is not enabled, it will be enabled as well.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| The currency code. | 

### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCurrency**
> DeleteCurrency(ctx, code)
Delete a currency.

Delete a currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| The currency code. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableCurrency**
> CurrencySingle DisableCurrency(ctx, code)
Disable a currency.

Disable a currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **int32**| The currency code. | 

### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableCurrency**
> CurrencySingle EnableCurrency(ctx, code)
Enable a single currency.

Enable a single currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| The currency code. | 

### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrency**
> CurrencySingle GetCurrency(ctx, code)
Get a single currency.

Get a single currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| The currency code. | 

### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefaultCurrency**
> CurrencySingle GetDefaultCurrency(ctx, )
Get the user's default currency.

Get the user's default currency.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountByCurrency**
> AccountArray ListAccountByCurrency(ctx, code, optional)
List all accounts with this currency.

List all accounts with this currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| The currency code. | 
 **optional** | ***CurrenciesApiListAccountByCurrencyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrenciesApiListAccountByCurrencyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is 50. | 
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

# **ListAvailableBudgetByCurrency**
> AvailableBudgetArray ListAvailableBudgetByCurrency(ctx, code, optional)
List all available budgets with this currency.

List all available budgets with this currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| The currency code. | 
 **optional** | ***CurrenciesApiListAvailableBudgetByCurrencyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrenciesApiListAvailableBudgetByCurrencyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is 50 | 

### Return type

[**AvailableBudgetArray**](AvailableBudgetArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBillByCurrency**
> BillArray ListBillByCurrency(ctx, code, optional)
List all bills with this currency.

List all bills with this currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| The currency code. | 
 **optional** | ***CurrenciesApiListBillByCurrencyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrenciesApiListBillByCurrencyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is 50. | 

### Return type

[**BillArray**](BillArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBudgetLimitByCurrency**
> BudgetLimitArray ListBudgetLimitByCurrency(ctx, code, optional)
List all budget limits with this currency

List all budget limits with this currency

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| The currency code. | 
 **optional** | ***CurrenciesApiListBudgetLimitByCurrencyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrenciesApiListBudgetLimitByCurrencyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is 50. | 
 **start** | **optional.String**| Start date for the budget limit list. | 
 **end** | **optional.String**| End date for the budget limit list. | 

### Return type

[**BudgetLimitArray**](BudgetLimitArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCurrency**
> CurrencyArray ListCurrency(ctx, optional)
List all currencies.

List all currencies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CurrenciesApiListCurrencyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrenciesApiListCurrencyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. The default pagination is 50. | 

### Return type

[**CurrencyArray**](CurrencyArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRecurrenceByCurrency**
> RecurrenceArray ListRecurrenceByCurrency(ctx, code, optional)
List all recurring transactions with this currency.

List all recurring transactions with this currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| The currency code. | 
 **optional** | ***CurrenciesApiListRecurrenceByCurrencyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrenciesApiListRecurrenceByCurrencyOpts struct
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

# **ListRuleByCurrency**
> RuleArray ListRuleByCurrency(ctx, code, optional)
List all rules with this currency.

List all rules with this currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| The currency code. | 
 **optional** | ***CurrenciesApiListRuleByCurrencyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrenciesApiListRuleByCurrencyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination per 50. | 

### Return type

[**RuleArray**](RuleArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransactionByCurrency**
> TransactionArray ListTransactionByCurrency(ctx, code, optional)
List all transactions with this currency.

List all transactions with this currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| The currency code. | 
 **optional** | ***CurrenciesApiListTransactionByCurrencyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrenciesApiListTransactionByCurrencyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is per 50. | 
 **start** | **optional.String**| A date formatted YYYY-MM-DD, to limit the list of transactions.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD, to limit the list of transactions.  | 
 **type_** | [**optional.Interface of TransactionTypeFilter**](.md)| Optional filter on the transaction type(s) returned | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreCurrency**
> CurrencySingle StoreCurrency(ctx, body, enabled, default_, code, name, symbol, decimalPlaces)
Store a new currency

Creates a new currency. The data required can be submitted as a JSON body or as a list of parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CurrencyStore**](CurrencyStore.md)| JSON array or key&#x3D;value pairs with the necessary currency information. See the model for the exact specifications. | 
  **enabled** | **bool**|  | 
  **default_** | **bool**|  | 
  **code** | **string**|  | 
  **name** | **string**|  | 
  **symbol** | **string**|  | 
  **decimalPlaces** | **int32**|  | 

### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCurrency**
> CurrencySingle UpdateCurrency(ctx, body, enabled, default_, code, name, symbol, decimalPlaces, code)
Update existing currency.

Update existing currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CurrencyUpdate**](CurrencyUpdate.md)| JSON array with updated currency information. See the model for the exact specifications. | 
  **enabled** | **bool**|  | 
  **default_** | **bool**|  | 
  **code** | **string**|  | 
  **name** | **string**|  | 
  **symbol** | **string**|  | 
  **decimalPlaces** | **int32**|  | 
  **code** | **string**| The currency code. | 

### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

