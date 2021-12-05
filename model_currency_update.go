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

type CurrencyUpdate struct {
	// If the currency is enabled
	Enabled bool `json:"enabled,omitempty"`
	// If the currency must be the default for the user. You can only submit TRUE.
	Default_ bool `json:"default,omitempty"`
	// The currency code
	Code string `json:"code,omitempty"`
	// The currency name
	Name string `json:"name,omitempty"`
	// The currency symbol
	Symbol string `json:"symbol,omitempty"`
	// How many decimals to use when displaying this currency. Between 0 and 16.
	DecimalPlaces int32 `json:"decimal_places,omitempty"`
}