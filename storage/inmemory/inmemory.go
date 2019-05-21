package inmemory

// Storage wraps the different storages.
type Storage struct {
	AccountStore *AccountStore
}

// NewStorage creates a new storage.
func NewStorage() *Storage {
	return &Storage{
		AccountStore: NewAccountStore(),
	}
}
