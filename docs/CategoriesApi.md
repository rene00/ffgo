# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCategory**](CategoriesApi.md#DeleteCategory) | **Delete** /api/v1/categories/{id} | Delete a category.
[**GetCategory**](CategoriesApi.md#GetCategory) | **Get** /api/v1/categories/{id} | Get a single category.
[**ListAttachmentByCategory**](CategoriesApi.md#ListAttachmentByCategory) | **Get** /api/v1/categories/{id}/attachments | Lists all attachments.
[**ListCategory**](CategoriesApi.md#ListCategory) | **Get** /api/v1/categories | List all categories.
[**ListTransactionByCategory**](CategoriesApi.md#ListTransactionByCategory) | **Get** /api/v1/categories/{id}/transactions | List all transactions in a category.
[**StoreCategory**](CategoriesApi.md#StoreCategory) | **Post** /api/v1/categories | Store a new category
[**UpdateCategory**](CategoriesApi.md#UpdateCategory) | **Put** /api/v1/categories/{id} | Update existing category.

# **DeleteCategory**
> DeleteCategory(ctx, id)
Delete a category.

Delete a category. Transactions will not be removed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the category. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategory**
> CategorySingle GetCategory(ctx, id, optional)
Get a single category.

Get a single category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the category. | 
 **optional** | ***CategoriesApiGetCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoriesApiGetCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| A date formatted YYYY-MM-DD, to show spent and earned info.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD, to show spent and earned info.  | 

### Return type

[**CategorySingle**](CategorySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAttachmentByCategory**
> AttachmentArray ListAttachmentByCategory(ctx, id, optional)
Lists all attachments.

Lists all attachments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the category. | 
 **optional** | ***CategoriesApiListAttachmentByCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoriesApiListAttachmentByCategoryOpts struct
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

# **ListCategory**
> CategoryArray ListCategory(ctx, optional)
List all categories.

List all categories.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CategoriesApiListCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoriesApiListCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. The default pagination is 50. | 

### Return type

[**CategoryArray**](CategoryArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransactionByCategory**
> TransactionArray ListTransactionByCategory(ctx, id, optional)
List all transactions in a category.

List all transactions in a category, optionally limited to the date ranges specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the category. | 
 **optional** | ***CategoriesApiListTransactionByCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoriesApiListTransactionByCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is per 50. | 
 **start** | **optional.String**| A date formatted YYYY-MM-DD, to limit the result list.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD, to limit the result list.  | 
 **type_** | [**optional.Interface of TransactionTypeFilter**](.md)| Optional filter on the transaction type(s) returned | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreCategory**
> CategorySingle StoreCategory(ctx, body, createdAt, updatedAt, name, notes, spent, earned)
Store a new category

Creates a new category. The data required can be submitted as a JSON body or as a list of parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Category**](Category.md)| JSON array or key&#x3D;value pairs with the necessary category information. See the model for the exact specifications. | 
  **createdAt** | **time.Time**|  | 
  **updatedAt** | **time.Time**|  | 
  **name** | **string**|  | 
  **notes** | **string**|  | 
  **spent** | [**[]CategorySpent**](CategorySpent.md)|  | 
  **earned** | [**[]CategoryEarned**](CategoryEarned.md)|  | 

### Return type

[**CategorySingle**](CategorySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCategory**
> CategorySingle UpdateCategory(ctx, body, name, notes, id)
Update existing category.

Update existing category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CategoryUpdate**](CategoryUpdate.md)| JSON array with updated category information. See the model for the exact specifications. | 
  **name** | **string**|  | 
  **notes** | **string**|  | 
  **id** | **int32**| The ID of the category. | 

### Return type

[**CategorySingle**](CategorySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

