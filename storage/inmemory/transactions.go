package inmemory

import (
	"context"
	"fmt"
	"sync"

	transactionspb "github.com/openbank/gunk/gunk/v1/transactions"
)

// TransactionStore ...
type TransactionStore struct {
	data sync.Map
	keys []string
}

// NewTransactionStore creates a new inmemory storage
func NewTransactionStore() *TransactionStore {
	return &TransactionStore{
		keys: []string{},
	}
}

// GetTransaction returns a transaction.
func (s *TransactionStore) GetTransaction(ctx context.Context, id string) (*transactionspb.Transaction, error) {
	if a, ok := s.data.Load(id); ok {
		return a.(*transactionspb.Transaction), nil
	}
	return nil, fmt.Errorf("not found")
}

// GetTransactions returns a list of transactions.
func (s *TransactionStore) GetTransactions(ctx context.Context, nextID string) ([]*transactionspb.Transaction, bool, error) {
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
	var res []*transactionspb.Transaction
	for _, k := range sel {
		if a, ok := s.data.Load(k); ok {
			res = append(res, a.(*transactionspb.Transaction))
		}
	}
	return res, hasMore, nil
}

// CreateTransaction stores a transaction.
func (s *TransactionStore) CreateTransaction(ctx context.Context, transaction *transactionspb.Transaction) error {
	// LoadOrStore tries to load or store the key/value.
	// loaded is true if key is already present in the map.
	if _, loaded := s.data.LoadOrStore(transaction.TransactionID, transaction); loaded {
		return fmt.Errorf("transaction %s already exists", transaction.TransactionID)
	}
	s.keys = append(s.keys, transaction.TransactionID)
	return nil
}
