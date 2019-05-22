package inmemory

const pageSize = 10

// Storage wraps the different storages.
type Storage struct {
	AccountStore     *AccountStore
	TransactionStore *TransactionStore
}

// NewStorage creates a new storage.
func NewStorage() *Storage {
	return &Storage{
		AccountStore:     NewAccountStore(),
		TransactionStore: NewTransactionStore(),
	}
}
