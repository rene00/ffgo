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

type TransactionTypeFilter string

// List of TransactionTypeFilter
const (
	ALL_TransactionTypeFilter TransactionTypeFilter = "all"
	WITHDRAWAL_TransactionTypeFilter TransactionTypeFilter = "withdrawal"
	WITHDRAWALS_TransactionTypeFilter TransactionTypeFilter = "withdrawals"
	EXPENSE_TransactionTypeFilter TransactionTypeFilter = "expense"
	DEPOSIT_TransactionTypeFilter TransactionTypeFilter = "deposit"
	DEPOSITS_TransactionTypeFilter TransactionTypeFilter = "deposits"
	INCOME_TransactionTypeFilter TransactionTypeFilter = "income"
	TRANSFER_TransactionTypeFilter TransactionTypeFilter = "transfer"
	TRANSFERS_TransactionTypeFilter TransactionTypeFilter = "transfers"
	OPENING_BALANCE_TransactionTypeFilter TransactionTypeFilter = "opening_balance"
	RECONCILIATION_TransactionTypeFilter TransactionTypeFilter = "reconciliation"
	SPECIAL_TransactionTypeFilter TransactionTypeFilter = "special"
	SPECIALS_TransactionTypeFilter TransactionTypeFilter = "specials"
	DEFAULT__TransactionTypeFilter TransactionTypeFilter = "default"
)