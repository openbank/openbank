package storage

import (
	"context"

	accountspb "github.com/openbank/gunk/gunk/v1/accounts"
	transactionspb "github.com/openbank/gunk/gunk/v1/transactions"
)

// AccountStore provides storage operations for account resource.
type AccountStore interface {
	GetAccount(ctx context.Context, id string) (*accountspb.Account, error)
	GetAccounts(ctx context.Context, nextID string) ([]*accountspb.Account, bool, error)
	CreateAccount(ctx context.Context, account *accountspb.Account) error
	UpdateAccount(ctx context.Context, account *accountspb.Account) error
	DeleteAccount(ctx context.Context, id string) error
}

// TransactionStore provides storage operations for transaction resource.
type TransactionStore interface {
	GetTransaction(ctx context.Context, id string) (*transactionspb.Transaction, error)
	GetTransactions(ctx context.Context, nextID string) ([]*transactionspb.Transaction, bool, error)
	CreateTransaction(ctx context.Context, req *transactionspb.Transaction) error
}
