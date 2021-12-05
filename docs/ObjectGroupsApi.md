# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteObjectGroup**](ObjectGroupsApi.md#DeleteObjectGroup) | **Delete** /api/v1/object_groups/{id} | Delete a object group.
[**GetObjectGroup**](ObjectGroupsApi.md#GetObjectGroup) | **Get** /api/v1/object_groups/{id} | Get a single object group.
[**ListBillByObjectGroup**](ObjectGroupsApi.md#ListBillByObjectGroup) | **Get** /api/v1/object_groups/{id}/bills | List all bills with this object group.
[**ListObjectGroups**](ObjectGroupsApi.md#ListObjectGroups) | **Get** /api/v1/object_groups | List all oject groups.
[**ListPiggyBankByObjectGroup**](ObjectGroupsApi.md#ListPiggyBankByObjectGroup) | **Get** /api/v1/object_groups/{id}/piggy_banks | List all piggy banks related to the object group.
[**UpdateObjectGroup**](ObjectGroupsApi.md#UpdateObjectGroup) | **Put** /api/v1/object_groups/{id} | Update existing object group.

# **DeleteObjectGroup**
> DeleteObjectGroup(ctx, id)
Delete a object group.

Delete a object group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the object group. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectGroup**
> ObjectGroupSingle GetObjectGroup(ctx, id)
Get a single object group.

Get a single object group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the object group. | 

### Return type

[**ObjectGroupSingle**](ObjectGroupSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBillByObjectGroup**
> BillArray ListBillByObjectGroup(ctx, id, optional)
List all bills with this object group.

List all bills with this object group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the account. | 
 **optional** | ***ObjectGroupsApiListBillByObjectGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ObjectGroupsApiListBillByObjectGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is per 50 items. | 

### Return type

[**BillArray**](BillArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListObjectGroups**
> ObjectGroupArray ListObjectGroups(ctx, optional)
List all oject groups.

List all oject groups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ObjectGroupsApiListObjectGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ObjectGroupsApiListObjectGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. The default pagination is 50. | 

### Return type

[**ObjectGroupArray**](ObjectGroupArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPiggyBankByObjectGroup**
> PiggyBankArray ListPiggyBankByObjectGroup(ctx, id, optional)
List all piggy banks related to the object group.

This endpoint returns a list of all the piggy banks connected to the object group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the account. | 
 **optional** | ***ObjectGroupsApiListPiggyBankByObjectGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ObjectGroupsApiListPiggyBankByObjectGroupOpts struct
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

# **UpdateObjectGroup**
> ObjectGroupSingle UpdateObjectGroup(ctx, body, title, order, id)
Update existing object group.

Update existing object group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ObjectGroupUpdate**](ObjectGroupUpdate.md)| JSON array with updated piggy bank information. See the model for the exact specifications. | 
  **title** | **string**|  | 
  **order** | **int32**|  | 
  **id** | **int32**| The ID of the object group | 

### Return type

[**ObjectGroupSingle**](ObjectGroupSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

