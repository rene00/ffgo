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

type RecurrenceRepetition struct {
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The type of the repetition. ndom means: the n-th weekday of the month, where you can also specify which day of the week.
	Type_ string `json:"type"`
	// Information that defined the type of repetition. - For 'daily', this is empty. - For 'weekly', it is day of the week between 1 and 7 (Monday - Sunday). - For 'ndom', it is '1,2' or '4,5' or something else, where the first number is the week in the month, and the second number is the day in the week (between 1 and 7). '2,3' means: the 2nd Wednesday of the month - For 'monthly' it is the day of the month (1 - 31) - For yearly, it is a full date, ie '2018-09-17'. The year you use does not matter. 
	Moment string `json:"moment"`
	// How many occurrences to skip. 0 means skip nothing. 1 means every other.
	Skip int32 `json:"skip,omitempty"`
	// How to respond when the recurring transaction falls in the weekend. Possible values: 1. Do nothing, just create it 2. Create no transaction. 3. Skip to the previous Friday. 4. Skip to the next Monday. 
	Weekend int32 `json:"weekend,omitempty"`
	// Auto-generated repetition description.
	Description string `json:"description,omitempty"`
	// Array of future dates when the repetition will apply to. Auto generated.
	Occurrences []time.Time `json:"occurrences,omitempty"`
}
