# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InsightExpenseAsset**](InsightApi.md#InsightExpenseAsset) | **Get** /api/v1/insight/expense/asset | Insight into expenses, grouped by asset account.
[**InsightExpenseBill**](InsightApi.md#InsightExpenseBill) | **Get** /api/v1/insight/expense/bill | Insight into expenses, grouped by bill.
[**InsightExpenseBudget**](InsightApi.md#InsightExpenseBudget) | **Get** /api/v1/insight/expense/budget | Insight into expenses, grouped by budget.
[**InsightExpenseCategory**](InsightApi.md#InsightExpenseCategory) | **Get** /api/v1/insight/expense/category | Insight into expenses, grouped by category.
[**InsightExpenseExpense**](InsightApi.md#InsightExpenseExpense) | **Get** /api/v1/insight/expense/expense | Insight into expenses, grouped by expense account.
[**InsightExpenseNoBill**](InsightApi.md#InsightExpenseNoBill) | **Get** /api/v1/insight/expense/no-bill | Insight into expenses, without bill.
[**InsightExpenseNoBudget**](InsightApi.md#InsightExpenseNoBudget) | **Get** /api/v1/insight/expense/no-budget | Insight into expenses, without budget.
[**InsightExpenseNoCategory**](InsightApi.md#InsightExpenseNoCategory) | **Get** /api/v1/insight/expense/no-category | Insight into expenses, without category.
[**InsightExpenseNoTag**](InsightApi.md#InsightExpenseNoTag) | **Get** /api/v1/insight/expense/no-tag | Insight into expenses, without tag.
[**InsightExpenseTag**](InsightApi.md#InsightExpenseTag) | **Get** /api/v1/insight/expense/tag | Insight into expenses, grouped by tag.
[**InsightExpenseTotal**](InsightApi.md#InsightExpenseTotal) | **Get** /api/v1/insight/expense/total | Insight into total expenses.
[**InsightIncomeAsset**](InsightApi.md#InsightIncomeAsset) | **Get** /api/v1/insight/income/asset | Insight into income, grouped by asset account.
[**InsightIncomeCategory**](InsightApi.md#InsightIncomeCategory) | **Get** /api/v1/insight/income/category | Insight into income, grouped by category.
[**InsightIncomeNoCategory**](InsightApi.md#InsightIncomeNoCategory) | **Get** /api/v1/insight/income/no-category | Insight into income, without category.
[**InsightIncomeNoTag**](InsightApi.md#InsightIncomeNoTag) | **Get** /api/v1/insight/income/no-tag | Insight into income, without tag.
[**InsightIncomeRevenue**](InsightApi.md#InsightIncomeRevenue) | **Get** /api/v1/insight/income/revenue | Insight into income, grouped by revenue account.
[**InsightIncomeTag**](InsightApi.md#InsightIncomeTag) | **Get** /api/v1/insight/income/tag | Insight into income, grouped by tag.
[**InsightIncomeTotal**](InsightApi.md#InsightIncomeTotal) | **Get** /api/v1/insight/income/total | Insight into total income.
[**InsightTransferCategory**](InsightApi.md#InsightTransferCategory) | **Get** /api/v1/insight/transfer/category | Insight into transfers, grouped by category.
[**InsightTransferNoCategory**](InsightApi.md#InsightTransferNoCategory) | **Get** /api/v1/insight/transfer/no-category | Insight into transfers, without category.
[**InsightTransferNoTag**](InsightApi.md#InsightTransferNoTag) | **Get** /api/v1/insight/transfer/no-tag | Insight into expenses, without tag.
[**InsightTransferTag**](InsightApi.md#InsightTransferTag) | **Get** /api/v1/insight/transfer/tag | Insight into transfers, grouped by tag.
[**InsightTransferTotal**](InsightApi.md#InsightTransferTotal) | **Get** /api/v1/insight/transfer/total | Insight into total transfers.
[**InsightTransfers**](InsightApi.md#InsightTransfers) | **Get** /api/v1/insight/transfer/asset | Insight into transfers, grouped by account.

# **InsightExpenseAsset**
> []InsightGroupEntry InsightExpenseAsset(ctx, start, end, optional)
Insight into expenses, grouped by asset account.

This endpoint gives a summary of the expenses made by the user, grouped by asset account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightExpenseAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightExpenseAssetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightExpenseBill**
> []InsightGroupEntry InsightExpenseBill(ctx, start, end, optional)
Insight into expenses, grouped by bill.

This endpoint gives a summary of the expenses made by the user, grouped by (any) bill. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightExpenseBillOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightExpenseBillOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bills** | [**optional.Interface of []int64**](int64.md)| The bills to be included in the results.  | 
 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightExpenseBudget**
> []InsightGroupEntry InsightExpenseBudget(ctx, start, end, optional)
Insight into expenses, grouped by budget.

This endpoint gives a summary of the expenses made by the user, grouped by (any) budget. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightExpenseBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightExpenseBudgetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **budgets** | [**optional.Interface of []int64**](int64.md)| The budgets to be included in the results.  | 
 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightExpenseCategory**
> []InsightGroupEntry InsightExpenseCategory(ctx, start, end, optional)
Insight into expenses, grouped by category.

This endpoint gives a summary of the expenses made by the user, grouped by (any) category. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightExpenseCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightExpenseCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **categories** | [**optional.Interface of []int64**](int64.md)| The categories to be included in the results.  | 
 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightExpenseExpense**
> []InsightGroupEntry InsightExpenseExpense(ctx, start, end, optional)
Insight into expenses, grouped by expense account.

This endpoint gives a summary of the expenses made by the user, grouped by expense account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightExpenseExpenseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightExpenseExpenseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you add the accounts ID&#x27;s of expense accounts, only those accounts are included in the results. If you include ID&#x27;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. You can combine both asset / liability and expense account ID&#x27;s. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightExpenseNoBill**
> []InsightTotalEntry InsightExpenseNoBill(ctx, start, end, optional)
Insight into expenses, without bill.

This endpoint gives a summary of the expenses made by the user, including only expenses with no bill. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightExpenseNoBillOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightExpenseNoBillOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightExpenseNoBudget**
> []InsightTotalEntry InsightExpenseNoBudget(ctx, start, end, optional)
Insight into expenses, without budget.

This endpoint gives a summary of the expenses made by the user, including only expenses with no budget. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightExpenseNoBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightExpenseNoBudgetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightExpenseNoCategory**
> []InsightTotalEntry InsightExpenseNoCategory(ctx, start, end, optional)
Insight into expenses, without category.

This endpoint gives a summary of the expenses made by the user, including only expenses with no category. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightExpenseNoCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightExpenseNoCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightExpenseNoTag**
> []InsightTotalEntry InsightExpenseNoTag(ctx, start, end, optional)
Insight into expenses, without tag.

This endpoint gives a summary of the expenses made by the user, including only expenses with no tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightExpenseNoTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightExpenseNoTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightExpenseTag**
> []InsightGroupEntry InsightExpenseTag(ctx, start, end, optional)
Insight into expenses, grouped by tag.

This endpoint gives a summary of the expenses made by the user, grouped by (any) tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightExpenseTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightExpenseTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tags** | [**optional.Interface of []int64**](int64.md)| The tags to be included in the results.  | 
 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightExpenseTotal**
> []InsightTotalEntry InsightExpenseTotal(ctx, start, end, optional)
Insight into total expenses.

This endpoint gives a sum of the total expenses made by the user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightExpenseTotalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightExpenseTotalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightIncomeAsset**
> []InsightGroupEntry InsightIncomeAsset(ctx, start, end, optional)
Insight into income, grouped by asset account.

This endpoint gives a summary of the income received by the user, grouped by asset account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightIncomeAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightIncomeAssetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightIncomeCategory**
> []InsightGroupEntry InsightIncomeCategory(ctx, start, end, optional)
Insight into income, grouped by category.

This endpoint gives a summary of the income received by the user, grouped by (any) category. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightIncomeCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightIncomeCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **categories** | [**optional.Interface of []int64**](int64.md)| The categories to be included in the results.  | 
 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightIncomeNoCategory**
> []InsightTotalEntry InsightIncomeNoCategory(ctx, start, end, optional)
Insight into income, without category.

This endpoint gives a summary of the income received by the user, including only income with no category. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightIncomeNoCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightIncomeNoCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightIncomeNoTag**
> []InsightTotalEntry InsightIncomeNoTag(ctx, start, end, optional)
Insight into income, without tag.

This endpoint gives a summary of the income received by the user, including only income with no tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightIncomeNoTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightIncomeNoTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightIncomeRevenue**
> []InsightGroupEntry InsightIncomeRevenue(ctx, start, end, optional)
Insight into income, grouped by revenue account.

This endpoint gives a summary of the income received by the user, grouped by revenue account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightIncomeRevenueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightIncomeRevenueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you add the accounts ID&#x27;s of revenue accounts, only those accounts are included in the results. If you include ID&#x27;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. You can combine both asset / liability and deposit account ID&#x27;s. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightIncomeTag**
> []InsightGroupEntry InsightIncomeTag(ctx, start, end, optional)
Insight into income, grouped by tag.

This endpoint gives a summary of the income received by the user, grouped by (any) tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightIncomeTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightIncomeTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tags** | [**optional.Interface of []int64**](int64.md)| The tags to be included in the results.  | 
 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightIncomeTotal**
> []InsightTotalEntry InsightIncomeTotal(ctx, start, end, optional)
Insight into total income.

This endpoint gives a sum of the total income received by the user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightIncomeTotalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightIncomeTotalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightTransferCategory**
> []InsightGroupEntry InsightTransferCategory(ctx, start, end, optional)
Insight into transfers, grouped by category.

This endpoint gives a summary of the transfers made by the user, grouped by (any) category. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightTransferCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightTransferCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **categories** | [**optional.Interface of []int64**](int64.md)| The categories to be included in the results.  | 
 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightTransferNoCategory**
> []InsightTotalEntry InsightTransferNoCategory(ctx, start, end, optional)
Insight into transfers, without category.

This endpoint gives a summary of the transfers made by the user, including only transfers with no category. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightTransferNoCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightTransferNoCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightTransferNoTag**
> []InsightTotalEntry InsightTransferNoTag(ctx, start, end, optional)
Insight into expenses, without tag.

This endpoint gives a summary of the transfers made by the user, including only transfers with no tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightTransferNoTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightTransferNoTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only transfers from those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightTransferTag**
> []InsightGroupEntry InsightTransferTag(ctx, start, end, optional)
Insight into transfers, grouped by tag.

This endpoint gives a summary of the transfers created by the user, grouped by (any) tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightTransferTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightTransferTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tags** | [**optional.Interface of []int64**](int64.md)| The tags to be included in the results.  | 
 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightTransferTotal**
> []InsightTotalEntry InsightTransferTotal(ctx, start, end, optional)
Insight into total transfers.

This endpoint gives a sum of the total amount transfers made by the user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightTransferTotalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightTransferTotalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsightTransfers**
> []InsightTransferEntry InsightTransfers(ctx, start, end, optional)
Insight into transfers, grouped by account.

This endpoint gives a summary of the transfers made by the user, grouped by asset account or lability. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***InsightApiInsightTransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightApiInsightTransfersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accounts** | [**optional.Interface of []int64**](int64.md)| The accounts to be included in the results. If you include ID&#x27;s of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID&#x27;s will be ignored.  | 

### Return type

[**[]InsightTransferEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

