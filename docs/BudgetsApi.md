# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBudget**](BudgetsApi.md#DeleteBudget) | **Delete** /api/v1/budgets/{id} | Delete a budget.
[**DeleteBudgetLimit**](BudgetsApi.md#DeleteBudgetLimit) | **Delete** /api/v1/budgets/{id}/limits/{limitId} | Delete a budget limit.
[**GetBudget**](BudgetsApi.md#GetBudget) | **Get** /api/v1/budgets/{id} | Get a single budget.
[**GetBudgetLimit**](BudgetsApi.md#GetBudgetLimit) | **Get** /api/v1/budgets/{id}/limits/{limitId} | Get single budget limit.
[**ListAttachmentByBudget**](BudgetsApi.md#ListAttachmentByBudget) | **Get** /api/v1/budgets/{id}/attachments | Lists all attachments of a budget.
[**ListBudget**](BudgetsApi.md#ListBudget) | **Get** /api/v1/budgets | List all budgets.
[**ListBudgetLimit**](BudgetsApi.md#ListBudgetLimit) | **Get** /api/v1/budget-limits | Get list of budget limits by date
[**ListBudgetLimitByBudget**](BudgetsApi.md#ListBudgetLimitByBudget) | **Get** /api/v1/budgets/{id}/limits | Get all limits for a budget.
[**ListTransactionByBudget**](BudgetsApi.md#ListTransactionByBudget) | **Get** /api/v1/budgets/{id}/transactions | All transactions to a budget.
[**ListTransactionByBudgetLimit**](BudgetsApi.md#ListTransactionByBudgetLimit) | **Get** /api/v1/budgets/{id}/limits/{limitId}/transactions | List all transactions by a budget limit ID.
[**StoreBudget**](BudgetsApi.md#StoreBudget) | **Post** /api/v1/budgets | Store a new budget
[**StoreBudgetLimit**](BudgetsApi.md#StoreBudgetLimit) | **Post** /api/v1/budgets/{id}/limits | Store new budget limit.
[**UpdateBudget**](BudgetsApi.md#UpdateBudget) | **Put** /api/v1/budgets/{id} | Update existing budget.
[**UpdateBudgetLimit**](BudgetsApi.md#UpdateBudgetLimit) | **Put** /api/v1/budgets/{id}/limits/{limitId} | Update existing budget limit.

# **DeleteBudget**
> DeleteBudget(ctx, id)
Delete a budget.

Delete a budget. Transactions will not be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the budget. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBudgetLimit**
> DeleteBudgetLimit(ctx, id, limitId)
Delete a budget limit.

Delete a budget limit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the budget. The budget limit MUST be associated to the budget ID. | 
  **limitId** | **int32**| The ID of the budget limit. The budget limit MUST be associated to the budget ID. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBudget**
> BudgetSingle GetBudget(ctx, id, optional)
Get a single budget.

Get a single budget. If the start date and end date are submitted as well, the \"spent\" array will be updated accordingly.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the requested budget. | 
 **optional** | ***BudgetsApiGetBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetsApiGetBudgetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| A date formatted YYYY-MM-DD, to get info on how much the user has spent.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD, to get info on how much the user has spent.  | 

### Return type

[**BudgetSingle**](BudgetSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBudgetLimit**
> BudgetLimitSingle GetBudgetLimit(ctx, id, limitId)
Get single budget limit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the budget. The budget limit MUST be associated to the budget ID. | 
  **limitId** | **int32**| The ID of the budget limit. The budget limit MUST be associated to the budget ID. | 

### Return type

[**BudgetLimitSingle**](BudgetLimitSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAttachmentByBudget**
> AttachmentArray ListAttachmentByBudget(ctx, id, optional)
Lists all attachments of a budget.

Lists all attachments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the budget. | 
 **optional** | ***BudgetsApiListAttachmentByBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetsApiListAttachmentByBudgetOpts struct
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

# **ListBudget**
> BudgetArray ListBudget(ctx, optional)
List all budgets.

List all the budgets the user has made. If the start date and end date are submitted as well, the \"spent\" array will be updated accordingly.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BudgetsApiListBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetsApiListBudgetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. The default pagination is 50. | 
 **start** | **optional.String**| A date formatted YYYY-MM-DD, to get info on how much the user has spent. You must submit both start and end.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD, to get info on how much the user has spent. You must submit both start and end.  | 

### Return type

[**BudgetArray**](BudgetArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBudgetLimit**
> BudgetLimitArray ListBudgetLimit(ctx, start, end)
Get list of budget limits by date

Get all budget limits for for this date range. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 

### Return type

[**BudgetLimitArray**](BudgetLimitArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBudgetLimitByBudget**
> BudgetLimitArray ListBudgetLimitByBudget(ctx, id, optional)
Get all limits for a budget.

Get all budget limits for this budget and the money spent, and money left. You can limit the list by submitting a date range as well. The \"spent\" array for each budget limit is NOT influenced by the start and end date of your query, but by the start and end date of the budget limit itself. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the requested budget. | 
 **optional** | ***BudgetsApiListBudgetLimitByBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetsApiListBudgetLimitByBudgetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| A date formatted YYYY-MM-DD.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD.  | 

### Return type

[**BudgetLimitArray**](BudgetLimitArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransactionByBudget**
> TransactionArray ListTransactionByBudget(ctx, id, optional)
All transactions to a budget.

Get all transactions linked to a budget, possibly limited by start and end

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the budget. | 
 **optional** | ***BudgetsApiListTransactionByBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetsApiListTransactionByBudgetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Limits the number of results on one page. | 
 **page** | **optional.Int32**| Page number. The default pagination is 50. | 
 **start** | **optional.String**| A date formatted YYYY-MM-DD.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD.  | 
 **type_** | [**optional.Interface of TransactionTypeFilter**](.md)| Optional filter on the transaction type(s) returned | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransactionByBudgetLimit**
> TransactionArray ListTransactionByBudgetLimit(ctx, id, limitId, optional)
List all transactions by a budget limit ID.

List all the transactions within one budget limit. The start and end date are dictated by the budget limit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the budget. The budget limit MUST be associated to the budget ID. | 
  **limitId** | **int32**| The ID of the budget limit. The budget limit MUST be associated to the budget ID. | 
 **optional** | ***BudgetsApiListTransactionByBudgetLimitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetsApiListTransactionByBudgetLimitOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Page number. The default pagination is 50. | 
 **type_** | [**optional.Interface of TransactionTypeFilter**](.md)| Optional filter on the transaction type(s) returned | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreBudget**
> BudgetSingle StoreBudget(ctx, body, name, active, order, autoBudgetType, autoBudgetCurrencyId, autoBudgetCurrencyCode, autoBudgetAmount, autoBudgetPeriod)
Store a new budget

Creates a new budget. The data required can be submitted as a JSON body or as a list of parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BudgetStore**](BudgetStore.md)| JSON array or key&#x3D;value pairs with the necessary budget information. See the model for the exact specifications. | 
  **name** | **string**|  | 
  **active** | **bool**|  | 
  **order** | **int32**|  | 
  **autoBudgetType** | **string**|  | 
  **autoBudgetCurrencyId** | **string**|  | 
  **autoBudgetCurrencyCode** | **string**|  | 
  **autoBudgetAmount** | **string**|  | 
  **autoBudgetPeriod** | **string**|  | 

### Return type

[**BudgetSingle**](BudgetSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreBudgetLimit**
> BudgetLimitSingle StoreBudgetLimit(ctx, body, currencyId, currencyCode, budgetId, start, period, end, amount, id)
Store new budget limit.

Store a new budget limit under this budget.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BudgetLimitStore**](BudgetLimitStore.md)| JSON array or key&#x3D;value pairs with the necessary budget information. See the model for the exact specifications. | 
  **currencyId** | **string**|  | 
  **currencyCode** | **string**|  | 
  **budgetId** | **string**|  | 
  **start** | **string**|  | 
  **period** | **string**|  | 
  **end** | **string**|  | 
  **amount** | **string**|  | 
  **id** | **int32**| The ID of the budget. | 

### Return type

[**BudgetLimitSingle**](BudgetLimitSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBudget**
> BudgetSingle UpdateBudget(ctx, body, name, active, order, autoBudgetType, autoBudgetCurrencyId, autoBudgetCurrencyCode, autoBudgetAmount, autoBudgetPeriod, id)
Update existing budget.

Update existing budget. This endpoint cannot be used to set budget amount limits.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BudgetUpdate**](BudgetUpdate.md)| JSON array with updated budget information. See the model for the exact specifications. | 
  **name** | **string**|  | 
  **active** | **bool**|  | 
  **order** | **int32**|  | 
  **autoBudgetType** | **string**|  | 
  **autoBudgetCurrencyId** | **string**|  | 
  **autoBudgetCurrencyCode** | **string**|  | 
  **autoBudgetAmount** | **string**|  | 
  **autoBudgetPeriod** | **string**|  | 
  **id** | **int32**| The ID of the budget. | 

### Return type

[**BudgetSingle**](BudgetSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBudgetLimit**
> BudgetLimitSingle UpdateBudgetLimit(ctx, body, createdAt, updatedAt, start, end, currencyId, currencyCode, currencyName, currencySymbol, currencyDecimalPlaces, budgetId, period, amount, spent, id, limitId)
Update existing budget limit.

Update existing budget limit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BudgetLimit**](BudgetLimit.md)| JSON array with updated budget limit information. See the model for the exact specifications. | 
  **createdAt** | **time.Time**|  | 
  **updatedAt** | **time.Time**|  | 
  **start** | **time.Time**|  | 
  **end** | **time.Time**|  | 
  **currencyId** | **string**|  | 
  **currencyCode** | **string**|  | 
  **currencyName** | **string**|  | 
  **currencySymbol** | **string**|  | 
  **currencyDecimalPlaces** | **int32**|  | 
  **budgetId** | **string**|  | 
  **period** | **string**|  | 
  **amount** | **string**|  | 
  **spent** | **string**|  | 
  **id** | **int32**| The ID of the budget. The budget limit MUST be associated to the budget ID. | 
  **limitId** | **int32**| The ID of the budget limit. The budget limit MUST be associated to the budget ID. | 

### Return type

[**BudgetLimitSingle**](BudgetLimitSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

