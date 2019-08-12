// This is temporary solution to use golang const for authentication scope list.
// TODO : Find a way to generate it with gunk if possible.
package auth

const (
	// TranscationCreate scope for the transaction read operation.
	Scope_TransactionCreate = "https://bank.to/transaction.create"

	// TransactionRead scope for the transaction read operation.
	Scope_TransactionRead = "https://bank.to/transaction.read"

	// TransactionUpdate scope for the transaction update operation.
	Scope_TransactionUpdate = "https://bank.to/transaction.update"

	// TransactionDelete scope for the transaction delete operation.
	Scope_TransactionDelete = "https://bank.to/transaction.delete"

	// AccountsCreate scope fot the account crate operation.
	Scope_AccountsCreate = "https://bank.to/account.create"

	// AccountsRead scope for the account read account operation.
	Scope_AccountsRead = "https://bank.to/account.read"

	// AccountsUpdate scope for the account update operation.
	Scope_AccountsUpdate = "https://bank.to/account.update"

	// AccountsDelete scope for the account delete operation.
	Scope_AccountsDelete = "https://bank.to/account.delete"

	// CardsCreate scope for the card create operation.
	Scope_CardsCreate = "https://bank.to/card.create"

	// CardsRead scope for the card read operation.
	Scope_CardsRead = "https://bank.to/card.read"

	// CardsUpdate scope for the card update operation.
	Scope_CardsUpdate = "https://bank.to/card.update"

	// CardsDelete scope for the card delete operation.
	Scope_CardsDelete = "https://bank.to/card.delete"
)
