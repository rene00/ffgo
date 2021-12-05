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

type TransactionLinkStore struct {
	// The link type ID to use. You can also use the link_type_name field.
	LinkTypeId string `json:"link_type_id"`
	// The link type name to use. You can also use the link_type_id field.
	LinkTypeName string `json:"link_type_name,omitempty"`
	// The inward transaction transaction_journal_id for the link. This becomes the 'is paid by' transaction of the set.
	InwardId string `json:"inward_id"`
	// The outward transaction transaction_journal_id for the link. This becomes the 'pays for' transaction of the set.
	OutwardId string `json:"outward_id"`
	// Optional. Some notes.
	Notes string `json:"notes,omitempty"`
}
