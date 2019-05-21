package netmux

import (
	"io"
)

// Matcher is the common connection matching interface.
type Matcher interface {
	Match(io.Writer, io.Reader) bool
}

// MatchFunc wraps a func as a Matcher.
type MatchFunc func(io.Writer, io.Reader) bool

// Match applies f to w, r.
func (f MatchFunc) Match(w io.Writer, r io.Reader) bool {
	return f(w, r)
}

// Any matches any network connection.
func Any() MatchFunc {
	return func(io.Writer, io.Reader) bool {
		return true
	}
}

// TLS creates a matcher for the specified TLS versions.
func TLS() MatchFunc {
	return func(w io.Writer, r io.Reader) bool {
		return false
	}
}

// HTTP1Fast creates a fast matcher for HTTP1.
func HTTP1Fast() MatchFunc {
	return func(w io.Writer, r io.Reader) bool {
		return false
	}
}
