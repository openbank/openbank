package netmux

import (
	"fmt"
	"net"
	"strconv"
)

// Option is a network connection multiplexor option.
type Option func(*Netmux) error

// Proxy is a netmux option that sets the upstream networks or IP addresses
// allowed to send the PROXY protocol header.
//
// The provided network addresses can be any IP or CIDR accepted by net.ParseIP
// or net.ParseCIDR.
//
// When set on a Netmux, any Listener.RemoteAddr() call will return the address
// sent via the PROXY header.
//
// See: http://www.haproxy.org/download/1.8/doc/proxy-protocol.txt
func Proxy(networkAddresses ...string) Option {
	return func(nm *Netmux) error {
		var err error
		for _, s := range networkAddresses {
			n := s
			if z := net.ParseIP(n); z != nil {
				n += "/" + strconv.Itoa(len(z)*8)
			}

			var ipNet *net.IPNet
			_, ipNet, err = net.ParseCIDR(n)
			if err != nil {
				return fmt.Errorf("invalid ip or cidr %q", s)
			}
			nm.proxies = append(nm.proxies, ipNet)
		}
		return nil
	}
}
