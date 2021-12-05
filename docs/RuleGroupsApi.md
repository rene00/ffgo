# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRuleGroup**](RuleGroupsApi.md#DeleteRuleGroup) | **Delete** /api/v1/rule_groups/{id} | Delete a rule group.
[**FireRuleGroup**](RuleGroupsApi.md#FireRuleGroup) | **Post** /api/v1/rule_groups/{id}/trigger | Fire the rule group on your transactions.
[**GetRuleGroup**](RuleGroupsApi.md#GetRuleGroup) | **Get** /api/v1/rule_groups/{id} | Get a single rule group.
[**ListRuleByGroup**](RuleGroupsApi.md#ListRuleByGroup) | **Get** /api/v1/rule_groups/{id}/rules | List rules in this rule group.
[**ListRuleGroup**](RuleGroupsApi.md#ListRuleGroup) | **Get** /api/v1/rule_groups | List all rule groups.
[**StoreRuleGroup**](RuleGroupsApi.md#StoreRuleGroup) | **Post** /api/v1/rule_groups | Store a new rule group.
[**TestRuleGroup**](RuleGroupsApi.md#TestRuleGroup) | **Get** /api/v1/rule_groups/{id}/test | Test which transactions would be hit by the rule group. No changes will be made.
[**UpdateRuleGroup**](RuleGroupsApi.md#UpdateRuleGroup) | **Put** /api/v1/rule_groups/{id} | Update existing rule group.

# **DeleteRuleGroup**
> DeleteRuleGroup(ctx, id)
Delete a rule group.

Delete a rule group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the rule group. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FireRuleGroup**
> FireRuleGroup(ctx, id, optional)
Fire the rule group on your transactions.

Fire the rule group on your transactions. Changes will be made by the rules in the rule group! Limit the result if you want to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the rule group. | 
 **optional** | ***RuleGroupsApiFireRuleGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RuleGroupsApiFireRuleGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| A date formatted YYYY-MM-DD, to limit the transactions the actions will be applied to. Both the start date and the end date must be present.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD, to limit the transactions the actions will be applied to. Both the start date and the end date must be present.  | 
 **accounts** | [**optional.Interface of []int64**](int64.md)| Limit the triggering of the rule group to these asset accounts or liabilities. Only asset accounts and liabilities will be accepted. Other types will be silently dropped.  | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuleGroup**
> RuleGroupSingle GetRuleGroup(ctx, id)
Get a single rule group.

Get a single rule group. This does not include the rules. For that, see below.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the rule group. | 

### Return type

[**RuleGroupSingle**](RuleGroupSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRuleByGroup**
> RuleArray ListRuleByGroup(ctx, id, optional)
List rules in this rule group.

List rules in this rule group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the rule group. | 
 **optional** | ***RuleGroupsApiListRuleByGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RuleGroupsApiListRuleByGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is 50. | 

### Return type

[**RuleArray**](RuleArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRuleGroup**
> RuleGroupArray ListRuleGroup(ctx, optional)
List all rule groups.

List all rule groups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RuleGroupsApiListRuleGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RuleGroupsApiListRuleGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. The default pagination is 50 | 

### Return type

[**RuleGroupArray**](RuleGroupArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreRuleGroup**
> RuleGroupSingle StoreRuleGroup(ctx, body, title, description, order, active)
Store a new rule group.

Creates a new rule group. The data required can be submitted as a JSON body or as a list of parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RuleGroupStore**](RuleGroupStore.md)| JSON array or key&#x3D;value pairs with the necessary rule group information. See the model for the exact specifications. | 
  **title** | **string**|  | 
  **description** | **string**|  | 
  **order** | **int32**|  | 
  **active** | **bool**|  | 

### Return type

[**RuleGroupSingle**](RuleGroupSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestRuleGroup**
> TransactionArray TestRuleGroup(ctx, id, optional)
Test which transactions would be hit by the rule group. No changes will be made.

Test which transactions would be hit by the rule group. No changes will be made. Limit the result if you want to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the rule group. | 
 **optional** | ***RuleGroupsApiTestRuleGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RuleGroupsApiTestRuleGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. The default pagination is 50 items. | 
 **start** | **optional.String**| A date formatted YYYY-MM-DD, to limit the transactions the test will be applied to. Both the start date and the end date must be present.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD, to limit the transactions the test will be applied to. Both the start date and the end date must be present.  | 
 **searchLimit** | **optional.Int32**| Maximum number of transactions Firefly III will try. Don&#x27;t set this too high, or it will take Firefly III very long to run the test. I suggest a max of 200.  | 
 **triggeredLimit** | **optional.Int32**| Maximum number of transactions the rule group can actually trigger on, before Firefly III stops. I would suggest setting this to 10 or 15. Don&#x27;t go above the user&#x27;s page size, because browsing to page 2 or 3 of a test result would fire the test again, making any navigation efforts very slow.  | 
 **accounts** | [**optional.Interface of []int64**](int64.md)| Limit the testing of the rule group to these asset accounts or liabilities. Only asset accounts and liabilities will be accepted. Other types will be silently dropped.  | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRuleGroup**
> RuleGroupSingle UpdateRuleGroup(ctx, body, title, description, order, active, id)
Update existing rule group.

Update existing rule group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RuleGroupUpdate**](RuleGroupUpdate.md)| JSON array with updated rule group information. See the model for the exact specifications. | 
  **title** | **string**|  | 
  **description** | **string**|  | 
  **order** | **int32**|  | 
  **active** | **bool**|  | 
  **id** | **int32**| The ID of the rule group. | 

### Return type

[**RuleGroupSingle**](RuleGroupSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

