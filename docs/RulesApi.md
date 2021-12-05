# {{classname}}

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRule**](RulesApi.md#DeleteRule) | **Delete** /api/v1/rules/{id} | Delete an rule.
[**FireRule**](RulesApi.md#FireRule) | **Post** /api/v1/rules/{id}/trigger | Fire the rule on your transactions.
[**GetRule**](RulesApi.md#GetRule) | **Get** /api/v1/rules/{id} | Get a single rule.
[**ListRule**](RulesApi.md#ListRule) | **Get** /api/v1/rules | List all rules.
[**StoreRule**](RulesApi.md#StoreRule) | **Post** /api/v1/rules | Store a new rule
[**TestRule**](RulesApi.md#TestRule) | **Get** /api/v1/rules/{id}/test | Test which transactions would be hit by the rule. No changes will be made.
[**UpdateRule**](RulesApi.md#UpdateRule) | **Put** /api/v1/rules/{id} | Update existing rule.

# **DeleteRule**
> DeleteRule(ctx, id)
Delete an rule.

Delete an rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the rule. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FireRule**
> FireRule(ctx, id, optional)
Fire the rule on your transactions.

Fire the rule group on your transactions. Changes will be made by the rules in the group! Limit the result if you want to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the rule. | 
 **optional** | ***RulesApiFireRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RulesApiFireRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| A date formatted YYYY-MM-DD, to limit the transactions the actions will be applied to. If the start date is not present, it will be set to one year ago. If you use this field, both the start date and the end date must be present.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD, to limit the transactions the actions will be applied to. If the end date is not present, it will be set to today. If you use this field, both the start date and the end date must be present.  | 
 **accounts** | [**optional.Interface of []int64**](int64.md)| Limit the triggering of the rule to these asset accounts or liabilities. Only asset accounts and liabilities will be accepted. Other types will be silently dropped.  | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRule**
> RuleSingle GetRule(ctx, id)
Get a single rule.

Get a single rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the object.X | 

### Return type

[**RuleSingle**](RuleSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRule**
> RuleArray ListRule(ctx, optional)
List all rules.

List all rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RulesApiListRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RulesApiListRuleOpts struct
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

# **StoreRule**
> RuleSingle StoreRule(ctx, body, title, description, ruleGroupId, ruleGroupTitle, order, trigger, active, strict, stopProcessing, triggers, actions)
Store a new rule

Creates a new rule. The data required can be submitted as a JSON body or as a list of parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RuleStore**](RuleStore.md)| JSON array or key&#x3D;value pairs with the necessary rule information. See the model for the exact specifications. | 
  **title** | **string**|  | 
  **description** | **string**|  | 
  **ruleGroupId** | **string**|  | 
  **ruleGroupTitle** | **string**|  | 
  **order** | **int32**|  | 
  **trigger** | **string**|  | 
  **active** | **bool**|  | 
  **strict** | **bool**|  | 
  **stopProcessing** | **bool**|  | 
  **triggers** | [**[]RuleTriggerStore**](RuleTriggerStore.md)|  | 
  **actions** | [**[]RuleActionStore**](RuleActionStore.md)|  | 

### Return type

[**RuleSingle**](RuleSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestRule**
> TransactionArray TestRule(ctx, id, optional)
Test which transactions would be hit by the rule. No changes will be made.

Test which transactions would be hit by the rule. No changes will be made. Limit the result if you want to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the rule. | 
 **optional** | ***RulesApiTestRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RulesApiTestRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| A date formatted YYYY-MM-DD, to limit the transactions the test will be applied to. Both the start date and the end date must be present.  | 
 **end** | **optional.String**| A date formatted YYYY-MM-DD, to limit the transactions the test will be applied to. Both the start date and the end date must be present.  | 
 **accounts** | [**optional.Interface of []int64**](int64.md)| Limit the testing of the rule to these asset accounts or liabilities. Only asset accounts and liabilities will be accepted. Other types will be silently dropped.  | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRule**
> RuleSingle UpdateRule(ctx, body, title, description, ruleGroupId, order, trigger, active, strict, stopProcessing, triggers, actions, id)
Update existing rule.

Update existing rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RuleUpdate**](RuleUpdate.md)| JSON array with updated rule information. See the model for the exact specifications. | 
  **title** | **string**|  | 
  **description** | **string**|  | 
  **ruleGroupId** | **string**|  | 
  **order** | **int32**|  | 
  **trigger** | **string**|  | 
  **active** | **bool**|  | 
  **strict** | **bool**|  | 
  **stopProcessing** | **bool**|  | 
  **triggers** | [**[]RuleTriggerUpdate**](RuleTriggerUpdate.md)|  | 
  **actions** | [**[]RuleActionUpdate**](RuleActionUpdate.md)|  | 
  **id** | **int32**| The ID of the object.X | 

### Return type

[**RuleSingle**](RuleSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

