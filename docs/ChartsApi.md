# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChartAccountOverview**](ChartsApi.md#GetChartAccountOverview) | **Get** /api/v1/chart/account/overview | Dashboard chart with asset account balance information.

# **GetChartAccountOverview**
> []ChartDataSet GetChartAccountOverview(ctx, start, end)
Dashboard chart with asset account balance information.

This endpoint returns the data required to generate a chart with basic asset account balance information. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 

### Return type

[**[]ChartDataSet**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

