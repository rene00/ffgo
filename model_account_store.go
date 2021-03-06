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

type AccountStore struct {
	Name string `json:"name"`
	// Can only be one one these account types. import, initial-balance and reconciliation cannot be set manually.
	Type_ string `json:"type"`
	Iban string `json:"iban,omitempty"`
	Bic string `json:"bic,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
	// Represents the opening balance, the initial amount this account holds.
	OpeningBalance string `json:"opening_balance,omitempty"`
	// Represents the date of the opening balance.
	OpeningBalanceDate string `json:"opening_balance_date,omitempty"`
	VirtualBalance string `json:"virtual_balance,omitempty"`
	// Use either currency_id or currency_code. Defaults to the user's default currency.
	CurrencyId string `json:"currency_id,omitempty"`
	// Use either currency_id or currency_code. Defaults to the user's default currency.
	CurrencyCode string `json:"currency_code,omitempty"`
	// If omitted, defaults to true.
	Active bool `json:"active,omitempty"`
	// Order of the account
	Order int32 `json:"order,omitempty"`
	// If omitted, defaults to true.
	IncludeNetWorth bool `json:"include_net_worth,omitempty"`
	// Is only mandatory when the type is asset.
	AccountRole string `json:"account_role,omitempty"`
	// Mandatory when the account_role is ccAsset. Can only be monthlyFull or null.
	CreditCardType string `json:"credit_card_type,omitempty"`
	// Mandatory when the account_role is ccAsset. Moment at which CC payment installments are asked for by the bank.
	MonthlyPaymentDate string `json:"monthly_payment_date,omitempty"`
	// Mandatory when type is liability. Specifies the exact type.
	LiabilityType string `json:"liability_type,omitempty"`
	// 'credit' indicates somebody owes you the liability. 'debit' Indicates you owe this debt yourself. Works only for liabiltiies.
	LiabilityDirection string `json:"liability_direction,omitempty"`
	// Mandatory when type is liability. Interest percentage.
	Interest string `json:"interest,omitempty"`
	// Mandatory when type is liability. Period over which the interest is calculated.
	InterestPeriod string `json:"interest_period,omitempty"`
	Notes string `json:"notes,omitempty"`
	// Latitude of the accounts's location, if applicable. Can be used to draw a map.
	Latitude float64 `json:"latitude,omitempty"`
	// Latitude of the accounts's location, if applicable. Can be used to draw a map.
	Longitude float64 `json:"longitude,omitempty"`
	// Zoom level for the map, if drawn. This to set the box right. Unfortunately this is a proprietary value because each map provider has different zoom levels.
	ZoomLevel int32 `json:"zoom_level,omitempty"`
}
