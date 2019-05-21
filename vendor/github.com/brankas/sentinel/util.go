package sentinel

import (
	"context"
)

// Error is a sentinel error.
type Error string

// Error satisfies the error interface.
func (err Error) Error() string {
	return string(err)
}

// Error values.
const (
	// ErrAlreadyStarted is the already started error.
	ErrAlreadyStarted Error = "already started"

	// ErrInvalidType is the invalid type error.
	ErrInvalidType Error = "invalid type"
)

// convertAndAppendContextFuncs converts and appends funcs in v to o.
func convertAndAppendContextFuncs(o []func(context.Context) error, v ...interface{}) ([]func(context.Context) error, error) {
	for _, z := range v {
		var t func(context.Context) error
		switch f := z.(type) {
		case func(context.Context) error:
			t = f

		case func():
			t = func(context.Context) error {
				f()
				return nil
			}

		case func() error:
			t = func(context.Context) error {
				return f()
			}
		}

		if t == nil {
			return nil, ErrInvalidType
		}

		o = append(o, t)
	}
	return o, nil
}
