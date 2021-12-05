# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration**](ConfigurationApi.md#GetConfiguration) | **Get** /api/v1/configuration | Get Firefly III system configuration values.
[**GetSingleConfiguration**](ConfigurationApi.md#GetSingleConfiguration) | **Get** /api/v1/configuration/{name} | Get a single Firefly III system configuration value
[**SetConfiguration**](ConfigurationApi.md#SetConfiguration) | **Put** /api/v1/configuration/{name} | Update configuration value

# **GetConfiguration**
> []Configuration GetConfiguration(ctx, )
Get Firefly III system configuration values.

Returns all editable and not-editable configuration values for this Firefly III installation

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Configuration**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSingleConfiguration**
> ConfigurationSingle GetSingleConfiguration(ctx, name)
Get a single Firefly III system configuration value

Returns one configuration variable for this Firefly III installation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | [**ConfigValueFilter**](.md)| The name of the configuration value you want to know. | 

### Return type

[**ConfigurationSingle**](ConfigurationSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetConfiguration**
> ConfigurationSingle SetConfiguration(ctx, value, body, name)
Update configuration value

Set a single configuration value. Not all configuration values can be updated so the list of accepted configuration variables is small.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **value** | [**Object**](.md)|  | 
  **body** | [**ConfigurationUpdate**](ConfigurationUpdate.md)| JSON array with the necessary account information or key&#x3D;value pairs. See the model for the exact specifications. | 
  **name** | [**ConfigValueUpdateFilter**](.md)| The name of the configuration value you want to update. | 

### Return type

[**ConfigurationSingle**](ConfigurationSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

