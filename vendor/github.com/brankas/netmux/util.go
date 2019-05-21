package netmux

// Error is a network connection mux error.
type Error string

// Error satisfies the error interface.
func (err Error) Error() string {
	return string(err)
}

// Error values.
const (
	// ErrListenerClosed is the listener closed error.
	ErrListenerClosed Error = "listener closed"
)
