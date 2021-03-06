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

type BudgetLimit struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Start date of the budget limit.
	Start time.Time `json:"start"`
	// End date of the budget limit.
	End time.Time `json:"end"`
	// Use either currency_id or currency_code. Defaults to the user's default currency.
	CurrencyId string `json:"currency_id,omitempty"`
	// Use either currency_id or currency_code. Defaults to the user's default currency.
	CurrencyCode string `json:"currency_code,omitempty"`
	CurrencyName string `json:"currency_name,omitempty"`
	CurrencySymbol string `json:"currency_symbol,omitempty"`
	CurrencyDecimalPlaces int32 `json:"currency_decimal_places,omitempty"`
	// The budget ID of the associated budget.
	BudgetId string `json:"budget_id"`
	// Period of the budget limit. Only used when auto-generated by auto-budget.
	Period string `json:"period,omitempty"`
	Amount string `json:"amount"`
	Spent string `json:"spent,omitempty"`
}
