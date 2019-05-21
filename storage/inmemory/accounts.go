package inmemory

import (
	"context"
	"fmt"
	"sync"

	accountspb "github.com/openbank/gunk/gunk/v1/accounts"
)

const pageSize = 10

type AccountStore struct {
	data sync.Map
	keys []string
}

// NewAccountStore creates a new inmemory storage
func NewAccountStore() *AccountStore {
	return &AccountStore{
		keys: []string{},
	}
}

// GetAccount returns an account.
func (s *AccountStore) GetAccount(ctx context.Context, id string) (*accountspb.Account, error) {
	if a, ok := s.data.Load(id); ok {
		return a.(*accountspb.Account), nil
	}
	return nil, fmt.Errorf("not found")
}

// GetAccounts returns a list of accounts.
func (s *AccountStore) GetAccounts(ctx context.Context, nextID string) ([]*accountspb.Account, bool, error) {
	hasMore := false
	var sel []string
	for i, k := range s.keys {
		if k == nextID {
			high := i + pageSize
			if high > len(s.keys) {
				high = len(s.keys)
				hasMore = false
			} else {
				hasMore = true
			}
			sel = s.keys[i:high]
		}
	}
	var res []*accountspb.Account
	for _, k := range sel {
		if a, ok := s.data.Load(k); ok {
			res = append(res, a.(*accountspb.Account))
		}
	}
	return res, hasMore, nil
}

// CreateAccount stores an account.
func (s *AccountStore) CreateAccount(ctx context.Context, account *accountspb.Account) error {
	// LoadOrStore tries to load or store the key/value.
	// loaded is true if key is already present in the map.
	if _, loaded := s.data.LoadOrStore(account.AccountID, account); loaded {
		return fmt.Errorf("account %s already exists", account.AccountID)
	}
	s.keys = append(s.keys, account.AccountID)
	return nil
}

// UpdateAccount updates an account.
func (s *AccountStore) UpdateAccount(ctx context.Context, account *accountspb.Account) error {
	if _, ok := s.data.Load(account.AccountID); ok {
		s.data.Store(account.AccountID, account)
		return nil
	}
	return fmt.Errorf("not found")
}

// DeleteAccount removes an account from the storage.
func (s *AccountStore) DeleteAccount(ctx context.Context, id string) error {
	if _, ok := s.data.Load(id); ok {
		s.data.Delete(id)
		for i, k := range s.keys {
			if k == id {
				s.keys = append(s.keys[:i], s.keys[i+1:]...)
			}
		}
		return nil
	}
	return fmt.Errorf("not found")
}
