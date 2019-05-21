package sentinel

import (
	"net/http"
	"os"
)

// Option is a sentinel option.
type Option func(*Sentinel) error

// Register is a sentinel option to register a server, its shutdown listener,
// and ignore error handlers.
//
// Both server and shutdown can have a type of `func()`, `func() error`, or
// `func(context.Context) error`.
func Register(server, shutdown interface{}, ignore ...func(error) bool) Option {
	return func(s *Sentinel) error {
		s.Lock()
		defer s.Unlock()

		if s.started {
			return ErrAlreadyStarted
		}

		return s.Register(server, shutdown, ignore...)
	}
}

// Server is a sentinel option to add server funcs.
//
// Any server can have a type of `func()`, `func() error`, or
// `func(context.Context) error`.
func Server(serverFuncs ...interface{}) Option {
	return func(s *Sentinel) error {
		s.Lock()
		defer s.Unlock()

		if s.started {
			return ErrAlreadyStarted
		}

		var err error
		s.serverFuncs, err = convertAndAppendContextFuncs(s.serverFuncs, serverFuncs...)
		return err
	}
}

// Shutdown is a sentinel option to add shutdown listeners.
//
// Any shutdown listener can have a type of `func()`, `func() error`, or
// `func(context.Context) error`.
func Shutdown(shutdownFuncs ...interface{}) Option {
	return func(s *Sentinel) error {
		s.Lock()
		defer s.Unlock()

		if s.started {
			return ErrAlreadyStarted
		}

		var err error
		s.shutdownFuncs, err = convertAndAppendContextFuncs(s.shutdownFuncs, shutdownFuncs...)
		return err
	}
}

// Ignore is a sentinel option to add ignore error handlers.
func Ignore(ignore ...func(error) bool) Option {
	return func(s *Sentinel) error {
		s.Lock()
		defer s.Unlock()

		if s.started {
			return ErrAlreadyStarted
		}

		s.ignoreErrors = append(s.ignoreErrors, ignore...)
		return nil
	}
}

// Sigs is a sentinel option to set the specified signals for shutdown.
func Sigs(sigs ...os.Signal) Option {
	return func(s *Sentinel) error {
		s.Lock()
		defer s.Unlock()

		if s.started {
			return ErrAlreadyStarted
		}

		s.shutdownSigs = sigs
		return nil
	}
}

// Logf is a sentinel option to set a logger.
func Logf(f func(string, ...interface{})) Option {
	return func(s *Sentinel) error {
		s.logf = f
		return nil
	}
}

// Errorf is a sentinel option to set a error logger.
func Errorf(f func(string, ...interface{})) Option {
	return func(s *Sentinel) error {
		s.errf = f
		return nil
	}
}

// ServerOption is a HTTP server option.
type ServerOption func(*http.Server) error
