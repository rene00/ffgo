# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountsAC**](AutocompleteApi.md#GetAccountsAC) | **Get** /api/v1/autocomplete/accounts | Returns all accounts of the user returned in a basic auto-complete array.
[**GetBillsAC**](AutocompleteApi.md#GetBillsAC) | **Get** /api/v1/autocomplete/bills | Returns all bills of the user returned in a basic auto-complete array.
[**GetBudgetsAC**](AutocompleteApi.md#GetBudgetsAC) | **Get** /api/v1/autocomplete/budgets | Returns all budgets of the user returned in a basic auto-complete array.
[**GetCategoriesAC**](AutocompleteApi.md#GetCategoriesAC) | **Get** /api/v1/autocomplete/categories | Returns all categories of the user returned in a basic auto-complete array.
[**GetCurrenciesAC**](AutocompleteApi.md#GetCurrenciesAC) | **Get** /api/v1/autocomplete/currencies | Returns all currencies of the user returned in a basic auto-complete array.
[**GetCurrenciesCodeAC**](AutocompleteApi.md#GetCurrenciesCodeAC) | **Get** /api/v1/autocomplete/currencies-with-code | Returns all currencies of the user returned in a basic auto-complete array. This endpoint is DEPRECATED and I suggest you DO NOT use it.
[**GetObjectGroupsAC**](AutocompleteApi.md#GetObjectGroupsAC) | **Get** /api/v1/autocomplete/object-groups | Returns all object groups of the user returned in a basic auto-complete array.
[**GetPiggiesAC**](AutocompleteApi.md#GetPiggiesAC) | **Get** /api/v1/autocomplete/piggy-banks | Returns all piggy banks of the user returned in a basic auto-complete array.
[**GetPiggiesBalanceAC**](AutocompleteApi.md#GetPiggiesBalanceAC) | **Get** /api/v1/autocomplete/piggy-banks-with-balance | Returns all piggy banks of the user returned in a basic auto-complete array complemented with balance information.
[**GetRecurringAC**](AutocompleteApi.md#GetRecurringAC) | **Get** /api/v1/autocomplete/recurring | Returns all recurring transactions of the user returned in a basic auto-complete array.
[**GetRuleGroupsAC**](AutocompleteApi.md#GetRuleGroupsAC) | **Get** /api/v1/autocomplete/rule-groups | Returns all rule groups of the user returned in a basic auto-complete array.
[**GetRulesAC**](AutocompleteApi.md#GetRulesAC) | **Get** /api/v1/autocomplete/rules | Returns all rules of the user returned in a basic auto-complete array.
[**GetTagAC**](AutocompleteApi.md#GetTagAC) | **Get** /api/v1/autocomplete/tags | Returns all tags of the user returned in a basic auto-complete array.
[**GetTransactionTypesAC**](AutocompleteApi.md#GetTransactionTypesAC) | **Get** /api/v1/autocomplete/transaction-types | Returns all transaction types returned in a basic auto-complete array. English only.
[**GetTransactionsAC**](AutocompleteApi.md#GetTransactionsAC) | **Get** /api/v1/autocomplete/transactions | Returns all transaction descriptions of the user returned in a basic auto-complete array.
[**GetTransactionsIDAC**](AutocompleteApi.md#GetTransactionsIDAC) | **Get** /api/v1/autocomplete/transactions-with-id | Returns all transactions, complemented with their ID, of the user returned in a basic auto-complete array. This endpoint is DEPRECATED and I suggest you DO NOT use it.

# **GetAccountsAC**
> []AutocompleteAccount GetAccountsAC(ctx, optional)
Returns all accounts of the user returned in a basic auto-complete array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetAccountsACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetAccountsACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query for accounts. | 
 **limit** | **optional.Int32**| The number of items returned. | 
 **date** | **optional.String**| If the account is an asset account or a liability, the autocomplete will also return the balance of the account on this date. | 
 **type_** | [**optional.Interface of AccountTypeFilter**](.md)| Optional filter on the account type(s) used in the autocomplete. | 

### Return type

[**[]AutocompleteAccount**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBillsAC**
> []AutocompleteBill GetBillsAC(ctx, optional)
Returns all bills of the user returned in a basic auto-complete array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetBillsACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetBillsACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query for bills. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompleteBill**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBudgetsAC**
> []AutocompleteBudget GetBudgetsAC(ctx, optional)
Returns all budgets of the user returned in a basic auto-complete array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetBudgetsACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetBudgetsACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned | 

### Return type

[**[]AutocompleteBudget**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategoriesAC**
> []AutocompleteCategory GetCategoriesAC(ctx, optional)
Returns all categories of the user returned in a basic auto-complete array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetCategoriesACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetCategoriesACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompleteCategory**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrenciesAC**
> []AutocompleteCurrency GetCurrenciesAC(ctx, optional)
Returns all currencies of the user returned in a basic auto-complete array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetCurrenciesACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetCurrenciesACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompleteCurrency**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrenciesCodeAC**
> []AutocompleteCurrencyCode GetCurrenciesCodeAC(ctx, optional)
Returns all currencies of the user returned in a basic auto-complete array. This endpoint is DEPRECATED and I suggest you DO NOT use it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetCurrenciesCodeACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetCurrenciesCodeACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompleteCurrencyCode**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectGroupsAC**
> []AutocompleteObjectGroup GetObjectGroupsAC(ctx, optional)
Returns all object groups of the user returned in a basic auto-complete array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetObjectGroupsACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetObjectGroupsACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompleteObjectGroup**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPiggiesAC**
> []AutocompletePiggy GetPiggiesAC(ctx, optional)
Returns all piggy banks of the user returned in a basic auto-complete array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetPiggiesACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetPiggiesACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompletePiggy**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPiggiesBalanceAC**
> []AutocompletePiggyBalance GetPiggiesBalanceAC(ctx, optional)
Returns all piggy banks of the user returned in a basic auto-complete array complemented with balance information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetPiggiesBalanceACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetPiggiesBalanceACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompletePiggyBalance**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecurringAC**
> []AutocompleteRecurrence GetRecurringAC(ctx, optional)
Returns all recurring transactions of the user returned in a basic auto-complete array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetRecurringACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetRecurringACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompleteRecurrence**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuleGroupsAC**
> []AutocompleteRuleGroup GetRuleGroupsAC(ctx, optional)
Returns all rule groups of the user returned in a basic auto-complete array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetRuleGroupsACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetRuleGroupsACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompleteRuleGroup**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRulesAC**
> []AutocompleteRule GetRulesAC(ctx, optional)
Returns all rules of the user returned in a basic auto-complete array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetRulesACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetRulesACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompleteRule**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagAC**
> []AutocompleteTag GetTagAC(ctx, optional)
Returns all tags of the user returned in a basic auto-complete array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetTagACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetTagACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompleteTag**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionTypesAC**
> []AutocompleteTransactionType GetTransactionTypesAC(ctx, optional)
Returns all transaction types returned in a basic auto-complete array. English only.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetTransactionTypesACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetTransactionTypesACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompleteTransactionType**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionsAC**
> []AutocompleteTransaction GetTransactionsAC(ctx, optional)
Returns all transaction descriptions of the user returned in a basic auto-complete array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetTransactionsACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetTransactionsACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompleteTransaction**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionsIDAC**
> []AutocompleteTransactionId GetTransactionsIDAC(ctx, optional)
Returns all transactions, complemented with their ID, of the user returned in a basic auto-complete array. This endpoint is DEPRECATED and I suggest you DO NOT use it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutocompleteApiGetTransactionsIDACOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutocompleteApiGetTransactionsIDACOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The autocomplete search query. | 
 **limit** | **optional.Int32**| The number of items returned. | 

### Return type

[**[]AutocompleteTransactionId**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

