// This is temporary solution to use golang const for authentication scope list.
// TODO : Find a way to generate it with gunk if possible.
package auth

const (
	// TranscationCreate scope for the transaction read operation.
	Scope_TransactionCreate = "https://auth.bnk.to/transaction.create"

	// TransactionRead scope for the transaction read operation.
	Scope_TransactionRead = "https://auth.bnk.to/transaction.read"

	// TransactionUpdate scope for the transaction update operation.
	Scope_TransactionUpdate = "https://auth.bnk.to/transaction.update"

	// TransactionDelete scope for the transaction delete operation.
	Scope_TransactionDelete = "https://auth.bnk.to/transaction.delete"

	// AccountsCreate scope fot the account crate operation.
	Scope_AccountsCreate = "https://auth.bnk.to/account.create"

	// AccountsRead scope for the account read account operation.
	Scope_AccountsRead = "https://auth.bnk.to/account.read"

	// AccountsUpdate scope for the account update operation.
	Scope_AccountsUpdate = "https://auth.bnk.to/account.update"

	// AccountsDelete scope for the account delete operation.
	Scope_AccountsDelete = "https://auth.bnk.to/account.delete"

	// CardsCreate scope for the card create operation.
	Scope_CardsCreate = "https://auth.bnk.to/card.create"

	// CardsRead scope for the card read operation.
	Scope_CardsRead = "https://auth.bnk.to/card.read"

	// CardsUpdate scope for the card update operation.
	Scope_CardsUpdate = "https://auth.bnk.to/card.update"

	// CardsDelete scope for the card delete operation.
	Scope_CardsDelete = "https://auth.bnk.to/card.delete"

	// ProfileCreate scope for the profile create operation.
	Scope_ProfileCreate = "https://auth.bnk.to/profile.create"

	// ProfileRead scope for the profile read operation.
	Scope_ProfileRead = "https://auth.bnk.to/profile.read"

	// ProfileUpdate scope for the profile update operation.
	Scope_ProfileUpdate = "https://auth.bnk.to/profile.update"

	// ProfileDelete scope for the profile delete operation.
	Scope_ProfileDelete = "https://auth.bnk.to/profile.delete"

	// BranchCreate scope for the branch create operation.
	Scoce_BranchCreate = "https://auth.bnk.to/branch.create"

	// BranchRead scope for the branch read operation.
	Scoce_BranchRead = "https://auth.bnk.to/branch.read"

	// BranchUpdate scope for the branch update operation.
	Scoce_BranchUpdate = "https://auth.bnk.to/branch.update"

	// BranchDelete scope fot the branch delete operation.
	Scoce_BranchDelete = "https://auth.bnk.to/branch.delete"
)
