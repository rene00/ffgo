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

type AutocompleteAccount struct {
	Id string `json:"id"`
	// Name of the account found by an auto-complete search.
	Name string `json:"name"`
	// Asset accounts and liabilities have a second field with the given date's account balance.
	NameWithBalance string `json:"name_with_balance"`
	// Account type of the account found by the auto-complete search.
	Type_ string `json:"type"`
	// ID for the currency used by this account.
	CurrencyId string `json:"currency_id"`
	// Currency name for the currency used by this account.
	CurrencyName string `json:"currency_name"`
	// Currency code for the currency used by this account.
	CurrencyCode string `json:"currency_code"`
	// Currency symbol for the currency used by this account.
	CurrencySymbol string `json:"currency_symbol"`
	// Number of decimal places for the currency used by this account.
	CurrencyDecimalPlaces int32 `json:"currency_decimal_places"`
}
