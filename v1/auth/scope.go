// This is temporary solution to use golang const for authentication scope list.
// TODO : This is covered by scopegen. After SFF, convert to scopegen.
package auth

const (
	// TransactionRead scope for the transaction read operation.
	Scope_TransactionRead = "https://auth.bnk.to/transaction.read"

	// TransactionWrite scope for the transaction write operation.
	Scope_TransactionWrite = "https://auth.bnk.to/transaction.write"

	// AccountsRead scope for the account read account operation.
	Scope_AccountsRead = "https://auth.bnk.to/account.read"

	// AccountsWrite scope for the account write operation.
	Scope_AccountsWrite = "https://auth.bnk.to/account.write"

	// CardsRead scope for the card read operation.
	Scope_CardsRead = "https://auth.bnk.to/card.read"

	// CardsWrite scope for the card write operation.
	Scope_CardsWrite = "https://auth.bnk.to/card.write"

	// ProfileRead scope for the profile read operation.
	Scope_ProfileRead = "https://auth.bnk.to/profile.read"

	// ProfileWrite scope for the profile write operation.
	Scope_ProfileWrite = "https://auth.bnk.to/profile.write"

	// BranchRead scope for the branch read operation.
	Scope_BranchRead = "https://auth.bnk.to/branch.read"

	// BranchWrite scope for the branch write operation.
	Scope_BranchWrite = "https://auth.bnk.to/branch.write"

	// AccountApplicationRead scope for account application read operation.
	Scope_AccountApplicationRead = "https://auth.bnk.to/account-application.read"

	// AccountApplicationWrite scope for account application write operation.
	Scope_AccountApplicationWrite = "https://auth.bnk.to/account-application.write"

	// BankRead scope for the bank read operation.
	Scope_BankRead = "https://auth.bnk.to/bank.read"

	// BankWrite scope for the bank write operation.
	Scope_BankWrite = "https://auth.bnk.to/bank.write"

	// FXRead scope for forex read operation.
	Scope_FXRead = "https://auth.bnk.to/fx.read"

	// FXWrite scope for forex write operation.
	Scope_FXWrite = "https://auth.bnk.to/fx.write"

	// CounterpartyRead scope for counterparty read operation.
	Scope_CounterpartyRead = "https://auth.bnk.to/counterparty.read"

	// CounterpartyWrite scope for counterparty write operation.
	Scope_CounterpartyWrite = "https://auth.bnk.to/counterparty.write"

	// KYCRead scope for kyc read operation.
	Scope_KYCRead = "https://auth.bnk.to/kyc.read"

	// FXWrite scope for kyc write operation.
	Scope_KYCWrite = "https://auth.bnk.to/kyc.write"

	// CustomerRead scope for customer read operation.
	Scope_CustomerRead = "https://auth.bnk.to/customer.read"

	// CustomerWrite scope for customer write operation.
	Scope_CustomerWrite = "https://auth.bnk.to/customer.write"

	// TransactionMetadataRead scope for transaction metadata read operation.
	Scope_TransactionMetadataRead = "https://auth.bnk.to/transaction-metadata.read"

	// TransactionMetadataWrite scope for transaction metadata write operation.
	Scope_TransactionMetadataWrite = "https://auth.bnk.to/transaction-metadata.write"

	// TransactionRequestRead scope for transaction request read operation.
	Scope_TransactionRequestRead = "https://auth.bnk.to/transaction-request.read"

	// TransactionRequestWrite scope for transaction request write operation.
	Scope_TransactionRequestWrite = "https://auth.bnk.to/transaction-request.write"

	// CounterpartyMetadataRead scope for counterparty metadata read operation.
	Scope_CounterpartyMetadataRead = "https://auth.bnk.to/counterparty-metadata.read"

	// CounterpartyMetadataWrite scope for counterparty metadata write operation.
	Scope_CounterpartyMetadataWrite = "https://auth.bnk.to/counterparty-metadata.write"

	// ConsumerRead scope for consumer read operation.
	Scope_ConsumerRead = "https://auth.bnk.to/consumer.read"

	// ConsumerWrite scope for consumer write operation.
	Scope_ConsumerWrite = "https://auth.bnk.to/consumer.write"

	// AccountPublicRead scope for account public read operation.
	Scope_AccountPublicRead = "https://auth.bnk.to/account-public.read"

	// AccountPublicWrite scope for account public write operation.
	Scope_AccountPublicWrite = "https://auth.bnk.to/account-public.write"

	// ConsentRead scope for consent read operation.
	Scope_ConsentRead = "https://auth.bnk.to/consent.read"

	// ConsentWrite scope for consent write operation.
	Scope_ConsentWrite = "https://auth.bnk.to/consent.write"

	// ProductRead scope for product read operation.
	Scope_ProductRead = "https://auth.bnk.to/product.read"

	// ProductWrite scope for product write operation.
	Scope_ProductWrite = "https://auth.bnk.to/product.write"

	// ProductCollectionRead scope for product collection read operation.
	Scope_ProductCollectionRead = "https://auth.bnk.to/product-collection.read"

	// ProductCollectionWrite scope for product collection write operation.
	Scope_ProductCollectionWrite = "https://auth.bnk.to/product-collection.write"

	// EntitlementRead scope for entitlement read operation.
	Scope_EntitlementRead = "https://auth.bnk.to/entitlement.read"

	// EntitlementWrite scope for entitlement write operation.
	Scope_EntitlementWrite = "https://auth.bnk.to/entitlement.write"

	// Fastcheckout Scopes
	// Scope_FcoWrite scope for create fco transaction request.
	Scope_FcoWrite = "https://auth.bnk.to/fco.write"

	// Scope_FcoRead scope for read fco transactions.
	Scope_FcoRead = "https://auth.bnk.to/fco.read"
)
