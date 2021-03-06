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

type Bill struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Use either currency_id or currency_code
	CurrencyId string `json:"currency_id,omitempty"`
	// Use either currency_id or currency_code
	CurrencyCode string `json:"currency_code,omitempty"`
	CurrencySymbol string `json:"currency_symbol,omitempty"`
	CurrencyDecimalPlaces int32 `json:"currency_decimal_places,omitempty"`
	Name string `json:"name"`
	AmountMin string `json:"amount_min"`
	AmountMax string `json:"amount_max"`
	Date time.Time `json:"date"`
	// The date after which this bill is no longer valid or applicable
	EndDate time.Time `json:"end_date,omitempty"`
	// The date before which the bill must be renewed (or cancelled)
	ExtensionDate time.Time `json:"extension_date,omitempty"`
	// How often the bill must be paid.
	RepeatFreq string `json:"repeat_freq"`
	// How often the bill must be skipped. 1 means a bi-monthly bill.
	Skip int32 `json:"skip,omitempty"`
	// If the bill is active.
	Active bool `json:"active,omitempty"`
	// Order of the bill.
	Order int32 `json:"order,omitempty"`
	Notes string `json:"notes,omitempty"`
	// When the bill is expected to be due.
	NextExpectedMatch time.Time `json:"next_expected_match,omitempty"`
	// Formatted (locally) when the bill is due.
	NextExpectedMatchDiff string `json:"next_expected_match_diff,omitempty"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupId string `json:"object_group_id,omitempty"`
	// The order of the group. At least 1, for the highest sorting.
	ObjectGroupOrder int32 `json:"object_group_order,omitempty"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle string `json:"object_group_title,omitempty"`
	// Array of future dates when the bill is expected to be paid. Autogenerated.
	PayDates []time.Time `json:"pay_dates,omitempty"`
	// Array of past transactions when the bill was paid.
	PaidDates []BillPaidDates `json:"paid_dates,omitempty"`
}
