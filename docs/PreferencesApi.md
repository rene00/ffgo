# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPreference**](PreferencesApi.md#GetPreference) | **Get** /api/v1/preferences/{name} | Return a single preference.
[**ListPreference**](PreferencesApi.md#ListPreference) | **Get** /api/v1/preferences | List all users preferences.
[**StorePreference**](PreferencesApi.md#StorePreference) | **Post** /api/v1/preferences | Store a new preference for this user.
[**UpdatePreference**](PreferencesApi.md#UpdatePreference) | **Put** /api/v1/preferences/{name} | Update preference

# **GetPreference**
> PreferenceSingle GetPreference(ctx, name)
Return a single preference.

Return a single preference and the value.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the preference. | 

### Return type

[**PreferenceSingle**](PreferenceSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPreference**
> PreferenceArray ListPreference(ctx, optional)
List all users preferences.

List all of the preferences of the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PreferencesApiListPreferenceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreferencesApiListPreferenceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. The default pagination is 50. | 

### Return type

[**PreferenceArray**](PreferenceArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StorePreference**
> PreferenceSingle StorePreference(ctx, body, createdAt, updatedAt, name, data)
Store a new preference for this user.

This endpoint creates a new preference. The name and data are free-format, and entirely up to you. If the preference is not used in Firefly III itself it may not be configurable through the user interface, but you can use this endpoint to persist custom data for your own app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Preference**](Preference.md)| JSON array with the necessary preference information or key&#x3D;value pairs. See the model for the exact specifications. | 
  **createdAt** | **time.Time**|  | 
  **updatedAt** | **time.Time**|  | 
  **name** | **string**|  | 
  **data** | [**Object**](.md)|  | 

### Return type

[**PreferenceSingle**](PreferenceSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePreference**
> PreferenceSingle UpdatePreference(ctx, body, data, name)
Update preference

Update a user's preference.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PreferenceUpdate**](PreferenceUpdate.md)| JSON array or key&#x3D;value pairs with the necessary preference information. See the model for the exact specifications. | 
  **data** | [**Object**](.md)|  | 
  **name** | **string**| The name of the preference. Will always overwrite. Will be created if it does not exist. | 

### Return type

[**PreferenceSingle**](PreferenceSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

