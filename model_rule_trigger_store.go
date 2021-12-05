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

type RuleTriggerStore struct {
	// The type of thing this trigger responds to. A limited set is possible
	Type_ string `json:"type"`
	// The accompanying value the trigger responds to. This value is often mandatory, but this depends on the trigger.
	Value string `json:"value"`
	// Order of the trigger
	Order int32 `json:"order,omitempty"`
	// If the trigger is active. Defaults to true.
	Active bool `json:"active,omitempty"`
	// When true, other triggers will not be checked if this trigger was triggered. Defaults to false.
	StopProcessing bool `json:"stop_processing,omitempty"`
}
