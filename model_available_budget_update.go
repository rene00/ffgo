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

type AvailableBudgetUpdate struct {
	// Use either currency_id or currency_code.
	CurrencyId string `json:"currency_id,omitempty"`
	// Use either currency_id or currency_code.
	CurrencyCode string `json:"currency_code,omitempty"`
	Amount string `json:"amount,omitempty"`
	// Start date of the available budget.
	Start string `json:"start,omitempty"`
	// End date of the available budget.
	End string `json:"end,omitempty"`
}