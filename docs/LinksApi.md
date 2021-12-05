# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLinkType**](LinksApi.md#DeleteLinkType) | **Delete** /api/v1/link_types/{id} | Permanently delete link type.
[**DeleteTransactionLink**](LinksApi.md#DeleteTransactionLink) | **Delete** /api/v1/transaction_links/{id} | Permanently delete link between transactions.
[**GetLinkType**](LinksApi.md#GetLinkType) | **Get** /api/v1/link_types/{id} | Get single a link type.
[**GetTransactionLink**](LinksApi.md#GetTransactionLink) | **Get** /api/v1/transaction_links/{id} | Get a single link.
[**ListLinkType**](LinksApi.md#ListLinkType) | **Get** /api/v1/link_types | List all types of links.
[**ListTransactionByLinkType**](LinksApi.md#ListTransactionByLinkType) | **Get** /api/v1/link_types/{id}/transactions | List all transactions under this link type.
[**ListTransactionLink**](LinksApi.md#ListTransactionLink) | **Get** /api/v1/transaction_links | List all transaction links.
[**StoreLinkType**](LinksApi.md#StoreLinkType) | **Post** /api/v1/link_types | Create a new link type
[**StoreTransactionLink**](LinksApi.md#StoreTransactionLink) | **Post** /api/v1/transaction_links | Create a new link between transactions
[**UpdateLinkType**](LinksApi.md#UpdateLinkType) | **Put** /api/v1/link_types/{id} | Update existing link type.
[**UpdateTransactionLink**](LinksApi.md#UpdateTransactionLink) | **Put** /api/v1/transaction_links/{id} | Update an existing link between transactions.

# **DeleteLinkType**
> DeleteLinkType(ctx, id)
Permanently delete link type.

Will permanently delete a link type. The links between transactions will be removed. The transactions themselves remain. You cannot delete some of the system provided link types, indicated by the editable=false flag when you list it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the link type. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTransactionLink**
> DeleteTransactionLink(ctx, id)
Permanently delete link between transactions.

Will permanently delete link. Transactions remain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the transaction link. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLinkType**
> LinkTypeSingle GetLinkType(ctx, id)
Get single a link type.

Returns a single link type by its ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the link type. | 

### Return type

[**LinkTypeSingle**](LinkTypeSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionLink**
> TransactionLinkSingle GetTransactionLink(ctx, id)
Get a single link.

Returns a single link by its ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the transaction link. | 

### Return type

[**TransactionLinkSingle**](TransactionLinkSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLinkType**
> LinkTypeArray ListLinkType(ctx, optional)
List all types of links.

List all the link types the system has. These include the default ones as well as any new ones. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LinksApiListLinkTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LinksApiListLinkTypeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. The default pagination is 50 items. | 

### Return type

[**LinkTypeArray**](LinkTypeArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransactionByLinkType**
> TransactionArray ListTransactionByLinkType(ctx, id, optional)
List all transactions under this link type.

List all transactions under this link type, both the inward and outward transactions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the link type. | 
 **optional** | ***LinksApiListTransactionByLinkTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LinksApiListTransactionByLinkTypeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is per 50 items. | 
 **start** | **optional.String**| A date formatted YYYY-MM-DD, to limit the results.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD, to limit the results.  | 
 **type_** | [**optional.Interface of TransactionTypeFilter**](.md)| Optional filter on the transaction type(s) returned. | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransactionLink**
> TransactionLinkArray ListTransactionLink(ctx, optional)
List all transaction links.

List all the transaction links. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LinksApiListTransactionLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LinksApiListTransactionLinkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. The default pagination is per 50 items. | 

### Return type

[**TransactionLinkArray**](TransactionLinkArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreLinkType**
> LinkTypeSingle StoreLinkType(ctx, body, name, inward, outward, editable)
Create a new link type

Creates a new link type. The data required can be submitted as a JSON body or as a list of parameters (in key=value pairs, like a webform).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LinkType**](LinkType.md)| JSON array with the necessary link type information or key&#x3D;value pairs. See the model for the exact specifications. | 
  **name** | **string**|  | 
  **inward** | **string**|  | 
  **outward** | **string**|  | 
  **editable** | **bool**|  | 

### Return type

[**LinkTypeSingle**](LinkTypeSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreTransactionLink**
> TransactionLinkSingle StoreTransactionLink(ctx, body, linkTypeId, linkTypeName, inwardId, outwardId, notes)
Create a new link between transactions

Store a new link between two transactions. For this end point you need the journal_id from a transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionLinkStore**](TransactionLinkStore.md)| JSON array with the necessary link type information or key&#x3D;value pairs. See the model for the exact specifications. | 
  **linkTypeId** | **string**|  | 
  **linkTypeName** | **string**|  | 
  **inwardId** | **string**|  | 
  **outwardId** | **string**|  | 
  **notes** | **string**|  | 

### Return type

[**TransactionLinkSingle**](TransactionLinkSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLinkType**
> LinkTypeSingle UpdateLinkType(ctx, body, name, inward, outward, id)
Update existing link type.

Used to update a single link type. All fields that are not submitted will be cleared (set to NULL). The model will tell you which fields are mandatory. You cannot update some of the system provided link types, indicated by the editable=false flag when you list it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LinkTypeUpdate**](LinkTypeUpdate.md)| JSON array or formdata with updated link type information. See the model for the exact specifications. | 
  **name** | **string**|  | 
  **inward** | **string**|  | 
  **outward** | **string**|  | 
  **id** | **int32**| The ID of the link type. | 

### Return type

[**LinkTypeSingle**](LinkTypeSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTransactionLink**
> TransactionLinkSingle UpdateTransactionLink(ctx, body, linkTypeId, linkTypeName, inwardId, outwardId, notes, id)
Update an existing link between transactions.

Used to update a single existing link. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionLinkUpdate**](TransactionLinkUpdate.md)| JSON array or formdata with updated link type information. See the model for the exact specifications. | 
  **linkTypeId** | **string**|  | 
  **linkTypeName** | **string**|  | 
  **inwardId** | **string**|  | 
  **outwardId** | **string**|  | 
  **notes** | **string**|  | 
  **id** | **int32**| The ID of the transaction link. | 

### Return type

[**TransactionLinkSingle**](TransactionLinkSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

