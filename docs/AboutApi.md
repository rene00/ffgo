# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAbout**](AboutApi.md#GetAbout) | **Get** /api/v1/about | System information end point.
[**GetCron**](AboutApi.md#GetCron) | **Get** /api/v1/cron/{cliToken} | Cron job endpoint
[**GetCurrentUser**](AboutApi.md#GetCurrentUser) | **Get** /api/v1/about/user | Currently authenticated user endpoint.

# **GetAbout**
> SystemInfo GetAbout(ctx, )
System information end point.

Returns general system information and versions of the (supporting) software. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemInfo**](SystemInfo.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCron**
> CronResult GetCron(ctx, cliToken, optional)
Cron job endpoint

Firefly III has one endpoint for its various cron related tasks. Send a GET to this endpoint to run the cron. The cron requires the CLI token to be present. The cron job will fire for all users. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cliToken** | **string**| The CLI token of any user in Firefly III, required to run the cron job. | 
 **optional** | ***AboutApiGetCronOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AboutApiGetCronOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **optional.String**| A date formatted YYYY-MM-DD. This can be used to make the cron job pretend it&#x27;s running on another day.  | 
 **force** | **optional.Bool**| Forces the cron job to fire, regardless of whether it has fired before. This may result in double transactions or weird budgets, so be careful.  | 

### Return type

[**CronResult**](CronResult.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentUser**
> UserSingle GetCurrentUser(ctx, )
Currently authenticated user endpoint.

Returns the currently authenticated user. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserSingle**](UserSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

