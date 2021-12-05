# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBasicSummary**](SummaryApi.md#GetBasicSummary) | **Get** /api/v1/summary/basic | Returns basic sums of the users data.

# **GetBasicSummary**
> []BasicSummaryEntry GetBasicSummary(ctx, start, end, optional)
Returns basic sums of the users data.

Returns basic sums of the users data, like the net worth, spent and earned amounts. It is multi-currency, and is used in Firefly III to populate the dashboard. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| A date formatted YYYY-MM-DD.  | 
  **end** | **string**| A date formatted YYYY-MM-DD.  | 
 **optional** | ***SummaryApiGetBasicSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SummaryApiGetBasicSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **currencyCode** | **optional.String**| A currency code like EUR or USD, to filter the result.  | 

### Return type

[**[]BasicSummaryEntry**](array.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

