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

type RecurrenceStore struct {
	Type_ string `json:"type"`
	Title string `json:"title"`
	// Not to be confused with the description of the actual transaction(s) being created.
	Description string `json:"description,omitempty"`
	// First time the recurring transaction will fire. Must be after today.
	FirstDate string `json:"first_date"`
	// Date until the recurring transaction can fire. Use either this field or repetitions.
	RepeatUntil string `json:"repeat_until"`
	// Max number of created transactions. Use either this field or repeat_until.
	NrOfRepetitions int32 `json:"nr_of_repetitions,omitempty"`
	// Whether or not to fire the rules after the creation of a transaction.
	ApplyRules bool `json:"apply_rules,omitempty"`
	// If the recurrence is even active.
	Active bool `json:"active,omitempty"`
	Notes string `json:"notes,omitempty"`
	Repetitions []RecurrenceRepetitionStore `json:"repetitions"`
	Transactions []RecurrenceTransactionStore `json:"transactions"`
}