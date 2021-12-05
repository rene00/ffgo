/*
 * Firefly III API v1.5.4
 *
 * This is the documentation of the Firefly III API. You can find accompanying documentation on the website of Firefly III itself (see below). Please report any bugs or issues. You may use the \"Authorize\" button to try the API below. This file was last generated on 2021-09-25T14:21:28+00:00 
 *
 * API version: 1.5.4
 * Contact: james@firefly-iii.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package ffgo
import (
	"time"
)

type Rule struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Title string `json:"title"`
	Description string `json:"description,omitempty"`
	// ID of the rule group under which the rule must be stored. Either this field or rule_group_title is mandatory.
	RuleGroupId string `json:"rule_group_id"`
	// Title of the rule group under which the rule must be stored. Either this field or rule_group_id is mandatory.
	RuleGroupTitle string `json:"rule_group_title,omitempty"`
	Order int32 `json:"order,omitempty"`
	// Which action is necessary for the rule to fire? Use either store-journal or update-journal.
	Trigger string `json:"trigger"`
	// Whether or not the rule is even active. Default is true.
	Active bool `json:"active,omitempty"`
	// If the rule is set to be strict, ALL triggers must hit in order for the rule to fire. Otherwise, just one is enough. Default value is true.
	Strict bool `json:"strict,omitempty"`
	// If this value is true and the rule is triggered, other rules  after this one in the group will be skipped. Default value is false.
	StopProcessing bool `json:"stop_processing,omitempty"`
	Triggers []RuleTrigger `json:"triggers"`
	Actions []RuleAction `json:"actions"`
}
