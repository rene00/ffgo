# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBill**](BillsApi.md#DeleteBill) | **Delete** /api/v1/bills/{id} | Delete a bill.
[**GetBill**](BillsApi.md#GetBill) | **Get** /api/v1/bills/{id} | Get a single bill.
[**ListAttachmentByBill**](BillsApi.md#ListAttachmentByBill) | **Get** /api/v1/bills/{id}/attachments | List all attachments uploaded to the bill.
[**ListBill**](BillsApi.md#ListBill) | **Get** /api/v1/bills | List all bills.
[**ListRuleByBill**](BillsApi.md#ListRuleByBill) | **Get** /api/v1/bills/{id}/rules | List all rules associated with the bill.
[**ListTransactionByBill**](BillsApi.md#ListTransactionByBill) | **Get** /api/v1/bills/{id}/transactions | List all transactions associated with the  bill.
[**StoreBill**](BillsApi.md#StoreBill) | **Post** /api/v1/bills | Store a new bill
[**UpdateBill**](BillsApi.md#UpdateBill) | **Put** /api/v1/bills/{id} | Update existing bill.

# **DeleteBill**
> DeleteBill(ctx, id)
Delete a bill.

Delete a bill. This will not delete any associated rules. Will not remove associated transactions. WILL remove all associated attachments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the bill. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBill**
> BillSingle GetBill(ctx, id, optional)
Get a single bill.

Get a single bill.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the bill. | 
 **optional** | ***BillsApiGetBillOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillsApiGetBillOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| A date formatted YYYY-MM-DD. If it is are added to the request, Firefly III will calculate the appropriate payment and paid dates.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD. If it is added to the request, Firefly III will calculate the appropriate payment and paid dates.  | 

### Return type

[**BillSingle**](BillSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAttachmentByBill**
> AttachmentArray ListAttachmentByBill(ctx, id, optional)
List all attachments uploaded to the bill.

This endpoint will list all attachments linked to the bill.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the bill. | 
 **optional** | ***BillsApiListAttachmentByBillOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillsApiListAttachmentByBillOpts struct
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

# **ListBill**
> BillArray ListBill(ctx, optional)
List all bills.

This endpoint will list all the user's bills.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BillsApiListBillOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillsApiListBillOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. The default pagination is 50. | 
 **start** | **optional.String**| A date formatted YYYY-MM-DD. If it is are added to the request, Firefly III will calculate the appropriate payment and paid dates.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD. If it is added to the request, Firefly III will calculate the appropriate payment and paid dates.  | 

### Return type

[**BillArray**](BillArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRuleByBill**
> RuleArray ListRuleByBill(ctx, id)
List all rules associated with the bill.

This endpoint will list all rules that have an action to set the bill to this bill.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the bill. | 

### Return type

[**RuleArray**](RuleArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransactionByBill**
> TransactionArray ListTransactionByBill(ctx, id, optional)
List all transactions associated with the  bill.

This endpoint will list all transactions linked to this bill.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the bill. | 
 **optional** | ***BillsApiListTransactionByBillOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillsApiListTransactionByBillOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

# **StoreBill**
> BillSingle StoreBill(ctx, body, currencyId, currencyCode, name, amountMin, amountMax, date, endDate, extensionDate, repeatFreq, skip, active, notes, objectGroupId, objectGroupTitle)
Store a new bill

Creates a new bill. The data required can be submitted as a JSON body or as a list of parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BillStore**](BillStore.md)| JSON array or key&#x3D;value pairs with the necessary bill information. See the model for the exact specifications. | 
  **currencyId** | **string**|  | 
  **currencyCode** | **string**|  | 
  **name** | **string**|  | 
  **amountMin** | **string**|  | 
  **amountMax** | **string**|  | 
  **date** | **time.Time**|  | 
  **endDate** | **time.Time**|  | 
  **extensionDate** | **time.Time**|  | 
  **repeatFreq** | **string**|  | 
  **skip** | **int32**|  | 
  **active** | **bool**|  | 
  **notes** | **string**|  | 
  **objectGroupId** | **string**|  | 
  **objectGroupTitle** | **string**|  | 

### Return type

[**BillSingle**](BillSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBill**
> BillSingle UpdateBill(ctx, body, currencyId, currencyCode, name, amountMin, amountMax, date, endDate, extensionDate, repeatFreq, skip, active, notes, objectGroupId, objectGroupTitle, id)
Update existing bill.

Update existing bill.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BillUpdate**](BillUpdate.md)| JSON array or key&#x3D;value pairs with updated bill information. See the model for the exact specifications. | 
  **currencyId** | **string**|  | 
  **currencyCode** | **string**|  | 
  **name** | **string**|  | 
  **amountMin** | **string**|  | 
  **amountMax** | **string**|  | 
  **date** | **time.Time**|  | 
  **endDate** | **time.Time**|  | 
  **extensionDate** | **time.Time**|  | 
  **repeatFreq** | **string**|  | 
  **skip** | **int32**|  | 
  **active** | **bool**|  | 
  **notes** | **string**|  | 
  **objectGroupId** | **string**|  | 
  **objectGroupTitle** | **string**|  | 
  **id** | **int32**| The ID of the bill. | 

### Return type

[**BillSingle**](BillSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

