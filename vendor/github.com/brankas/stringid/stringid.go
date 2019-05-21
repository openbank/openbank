// Package stringid provides string based ID generators, and accompanying
// middleware for assigning a request ID to standard HTTP request contexts.
package stringid

// Generator is the common interface for ID generators.
type Generator interface {
	// Generate generates a string ID.
	Generate() string
}

// GeneratorFunc wraps a standard func as a ID generator.
type GeneratorFunc func() string

// Generate satisfies the Generator interface.
func (f GeneratorFunc) Generate() string {
	return f()
}

// DefaultGenerator is the default ID generator.
var DefaultGenerator Generator

// Generate generates a ID using the DefaultGenerator.
func Generate() string {
	if DefaultGenerator == nil {
		panic("DefaultGenerator cannot be nil")
	}
	return DefaultGenerator.Generate()
}

func init() {
	// create default generator
	DefaultGenerator = NewPushGenerator(nil)
}
