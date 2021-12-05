# Rule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Title** | **string** |  | [default to null]
**Description** | **string** |  | [optional] [default to null]
**RuleGroupId** | **string** | ID of the rule group under which the rule must be stored. Either this field or rule_group_title is mandatory. | [default to null]
**RuleGroupTitle** | **string** | Title of the rule group under which the rule must be stored. Either this field or rule_group_id is mandatory. | [optional] [default to null]
**Order** | **int32** |  | [optional] [default to null]
**Trigger** | **string** | Which action is necessary for the rule to fire? Use either store-journal or update-journal. | [default to null]
**Active** | **bool** | Whether or not the rule is even active. Default is true. | [optional] [default to null]
**Strict** | **bool** | If the rule is set to be strict, ALL triggers must hit in order for the rule to fire. Otherwise, just one is enough. Default value is true. | [optional] [default to null]
**StopProcessing** | **bool** | If this value is true and the rule is triggered, other rules  after this one in the group will be skipped. Default value is false. | [optional] [default to null]
**Triggers** | [**[]RuleTrigger**](RuleTrigger.md) |  | [default to null]
**Actions** | [**[]RuleAction**](RuleAction.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

