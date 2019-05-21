package netmux

import (
	"net"
)

// Listener
type Listener struct {
	nm *Netmux
}

// Accept satisfies the net.Listener interface.
func (l *Listener) Accept() (net.Conn, error) {
	return nil, nil
}

// Close satisfies the net.Listener interface.
func (l *Listener) Close() error {
	return nil
}

// Addr satisfies the net.Listener interface.
func (l *Listener) Addr() net.Addr {
	return nil
}
