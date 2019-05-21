// Package netmux provides a simple way to multiplex (mux) network connections
// based on content.
package netmux

import (
	"crypto/tls"
	"net"
)

// Netmux is a network connection multiplexor.
type Netmux struct {
	l net.Listener

	// proxies are the allowed proxy networks.
	proxies []*net.IPNet

	// Default is the default listener.
	Default *Listener
}

// New creates a connection multiplexor for the supplied listener.
func New(l net.Listener, opts ...Option) (*Netmux, error) {
	var err error

	nm := &Netmux{
		l: l,
	}

	// apply opts
	for _, o := range opts {
		if err = o(nm); err != nil {
			return nil, err
		}
	}

	// set default listener
	if nm.Default == nil {
		nm.Default = nm.Listen(Any())
	}

	return nm, nil
}

// Listen is a utility func wrapping a call to net.Listen and passing the
// returned net.Listener and provided options to New.
func Listen(network, address string, opts ...Option) (*Netmux, error) {
	l, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return New(l, opts...)
}

// ListenTLS is a utility func wrapping a call to tls.Listen and passing the
// returned net.Listener and provided options to New.
func ListenTLS(network, laddr string, config *tls.Config, opts ...Option) (*Netmux, error) {
	l, err := tls.Listen(network, laddr, config)
	if err != nil {
		return nil, err
	}
	return New(l, opts...)
}

// Listen creates a listener matching any supplied matchers.
func (nm *Netmux) Listen(matchers ...Matcher) *Listener {
	l := &Listener{
		nm: nm,
	}
	return l
}
