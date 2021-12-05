# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAvailableBudget**](AvailableBudgetsApi.md#DeleteAvailableBudget) | **Delete** /api/v1/available_budgets/{id} | Delete an available budget.
[**GetAvailableBudget**](AvailableBudgetsApi.md#GetAvailableBudget) | **Get** /api/v1/available_budgets/{id} | Get a single available budget.
[**ListAvailableBudget**](AvailableBudgetsApi.md#ListAvailableBudget) | **Get** /api/v1/available_budgets | List all available budget amounts.
[**StoreAvailableBudget**](AvailableBudgetsApi.md#StoreAvailableBudget) | **Post** /api/v1/available_budgets | Store a new available budget
[**UpdateAvailableBudget**](AvailableBudgetsApi.md#UpdateAvailableBudget) | **Put** /api/v1/available_budgets/{id} | Update existing available budget, to change for example the date range of the amount or the amount itself.

# **DeleteAvailableBudget**
> DeleteAvailableBudget(ctx, id)
Delete an available budget.

Delete an available budget. Not much more to say.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the available budget. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvailableBudget**
> AvailableBudgetSingle GetAvailableBudget(ctx, id)
Get a single available budget.

Get a single available budget, by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the available budget. | 

### Return type

[**AvailableBudgetSingle**](AvailableBudgetSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAvailableBudget**
> AvailableBudgetArray ListAvailableBudget(ctx, optional)
List all available budget amounts.

Firefly III allows users to set the amount that is available to be budgeted in so-called \"available budgets\". For example, the user could have 1200,- available to be divided during the coming month. This amount is used on the /budgets page. This endpoint returns all of these amounts and the periods for which they are set. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AvailableBudgetsApiListAvailableBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AvailableBudgetsApiListAvailableBudgetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. The default pagination is 50. | 
 **start** | **optional.String**| A date formatted YYYY-MM-DD.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD.  | 

### Return type

[**AvailableBudgetArray**](AvailableBudgetArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreAvailableBudget**
> AvailableBudgetSingle StoreAvailableBudget(ctx, body, currencyId, currencyCode, amount, start, end)
Store a new available budget

Creates a new available budget for a specified period. The data required can be submitted as a JSON body or as a list of parameters. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AvailableBudgetStore**](AvailableBudgetStore.md)| JSON array or key&#x3D;value pairs with the necessary available budget information. See the model for the exact specifications. | 
  **currencyId** | **string**|  | 
  **currencyCode** | **string**|  | 
  **amount** | **string**|  | 
  **start** | **string**|  | 
  **end** | **string**|  | 

### Return type

[**AvailableBudgetSingle**](AvailableBudgetSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAvailableBudget**
> AvailableBudgetSingle UpdateAvailableBudget(ctx, body, currencyId, currencyCode, amount, start, end, id)
Update existing available budget, to change for example the date range of the amount or the amount itself.

Update existing available budget.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AvailableBudgetUpdate**](AvailableBudgetUpdate.md)| JSON array or form value with updated available budget information. See the model for the exact specifications. | 
  **currencyId** | **string**|  | 
  **currencyCode** | **string**|  | 
  **amount** | **string**|  | 
  **start** | **string**|  | 
  **end** | **string**|  | 
  **id** | **int32**| The ID of the object.X | 

### Return type

[**AvailableBudgetSingle**](AvailableBudgetSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

