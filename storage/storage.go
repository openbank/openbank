package storage

import (
	"context"

	accountspb "github.com/openbank/gunk/gunk/v1/accounts"
)

// AccountStore provides storage operation for account resource.
type AccountStore interface {
	GetAccount(ctx context.Context, id string) (*accountspb.Account, error)
	GetAccounts(ctx context.Context, nextID string) ([]*accountspb.Account, bool, error)
	CreateAccount(ctx context.Context, account *accountspb.Account) error
	UpdateAccount(ctx context.Context, account *accountspb.Account) error
	DeleteAccount(ctx context.Context, id string) error
}
