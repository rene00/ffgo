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

type TransactionSplitStore struct {
	// Type of transaction.
	Type_ string `json:"type"`
	// Date of the transaction
	Date time.Time `json:"date"`
	// Amount of the transaction.
	Amount string `json:"amount"`
	// Description of the transaction.
	Description string `json:"description"`
	// Order of this entry in the list of transactions.
	Order int32 `json:"order,omitempty"`
	// Currency ID. Default is the source account's currency, or the user's default currency. The value you submit may be overruled by the source or destination account.
	CurrencyId string `json:"currency_id,omitempty"`
	// Currency code. Default is the source account's currency, or the user's default currency. The value you submit may be overruled by the source or destination account.
	CurrencyCode string `json:"currency_code,omitempty"`
	// The amount in a foreign currency.
	ForeignAmount string `json:"foreign_amount,omitempty"`
	// Currency ID of the foreign currency. Default is null. Is required when you submit a foreign amount.
	ForeignCurrencyId string `json:"foreign_currency_id,omitempty"`
	// Currency code of the foreign currency. Default is NULL. Can be used instead of the foreign_currency_id, but this or the ID is required when submitting a foreign amount.
	ForeignCurrencyCode string `json:"foreign_currency_code,omitempty"`
	// The budget ID for this transaction.
	BudgetId string `json:"budget_id,omitempty"`
	// The name of the budget to be used. If the budget name is unknown, the ID will be used or the value will be ignored.
	BudgetName string `json:"budget_name,omitempty"`
	// The category ID for this transaction.
	CategoryId string `json:"category_id,omitempty"`
	// The name of the category to be used. If the category is unknown, it will be created. If the ID and the name point to different categories, the ID overrules the name.
	CategoryName string `json:"category_name,omitempty"`
	// ID of the source account. For a withdrawal or a transfer, this must always be an asset account. For deposits, this must be a revenue account.
	SourceId string `json:"source_id"`
	// Name of the source account. For a withdrawal or a transfer, this must always be an asset account. For deposits, this must be a revenue account. Can be used instead of the source_id. If the transaction is a deposit, the source_name can be filled in freely: the account will be created based on the name.
	SourceName string `json:"source_name,omitempty"`
	// ID of the destination account. For a deposit or a transfer, this must always be an asset account. For withdrawals this must be an expense account.
	DestinationId string `json:"destination_id"`
	// Name of the destination account. You can submit the name instead of the ID. For everything except transfers, the account will be auto-generated if unknown, so submitting a name is enough.
	DestinationName string `json:"destination_name,omitempty"`
	// If the transaction has been reconciled already. When you set this, the amount can no longer be edited by the user.
	Reconciled bool `json:"reconciled,omitempty"`
	// Optional. Use either this or the piggy_bank_name
	PiggyBankId int32 `json:"piggy_bank_id,omitempty"`
	// Optional. Use either this or the piggy_bank_id
	PiggyBankName string `json:"piggy_bank_name,omitempty"`
	// Optional. Use either this or the bill_name
	BillId string `json:"bill_id,omitempty"`
	// Optional. Use either this or the bill_id
	BillName string `json:"bill_name,omitempty"`
	// Array of tags.
	Tags []string `json:"tags,omitempty"`
	Notes string `json:"notes,omitempty"`
	// Reference to internal reference of other systems.
	InternalReference string `json:"internal_reference,omitempty"`
	// Reference to external ID in other systems.
	ExternalId string `json:"external_id,omitempty"`
	// Internal ID of bunq transaction. Field is no longer used but still works.
	BunqPaymentId string `json:"bunq_payment_id,omitempty"`
	// SEPA Clearing Code
	SepaCc string `json:"sepa_cc,omitempty"`
	// SEPA Opposing Account Identifier
	SepaCtOp string `json:"sepa_ct_op,omitempty"`
	// SEPA end-to-end Identifier
	SepaCtId string `json:"sepa_ct_id,omitempty"`
	// SEPA mandate identifier
	SepaDb string `json:"sepa_db,omitempty"`
	// SEPA Country
	SepaCountry string `json:"sepa_country,omitempty"`
	// SEPA External Purpose indicator
	SepaEp string `json:"sepa_ep,omitempty"`
	// SEPA Creditor Identifier
	SepaCi string `json:"sepa_ci,omitempty"`
	// SEPA Batch ID
	SepaBatchId string `json:"sepa_batch_id,omitempty"`
	InterestDate time.Time `json:"interest_date,omitempty"`
	BookDate time.Time `json:"book_date,omitempty"`
	ProcessDate time.Time `json:"process_date,omitempty"`
	DueDate time.Time `json:"due_date,omitempty"`
	PaymentDate time.Time `json:"payment_date,omitempty"`
	InvoiceDate time.Time `json:"invoice_date,omitempty"`
}