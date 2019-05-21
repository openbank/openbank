package stringid

import (
	"github.com/google/uuid"
)

// NewUUIDGenerator creates a ID generator that generates UUIDv4 style IDs.
func NewUUIDGenerator() Generator {
	return GeneratorFunc(func() string {
		return uuid.New().String()
	})
}
