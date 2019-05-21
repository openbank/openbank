package godop

import (
	"bytes"
	"context"
	"io/ioutil"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

// Option is the Client option type.
type Option func(c *Client) error

// Domain is a Client option to set the domain.
func Domain(domain string) Option {
	return func(c *Client) error {
		c.domain = domain
		return nil
	}
}

// GodoClient is a Client option to pass an already created godo client.
func GodoClient(client *godo.Client) Option {
	return func(c *Client) error {
		c.client = client
		return nil
	}
}

// GodoClientToken is a Client option to pass only the godo client token, and a
// new godo client will be created.
func GodoClientToken(ctxt context.Context, token string) Option {
	return func(c *Client) error {
		return GodoClient(godo.NewClient(oauth2.NewClient(
			ctxt,
			oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: token,
				},
			),
		)))(c)
	}
}

// GodoClientTokenFile is a Client option to create a new godo client using a
// token stored in a file on disk.
func GodoClientTokenFile(ctxt context.Context, filename string) Option {
	return func(c *Client) error {
		tok, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		return GodoClientToken(ctxt, string(bytes.TrimSpace(tok)))(c)
	}
}

// Logf is a Client option to specify the logging function used.
func Logf(f func(string, ...interface{})) Option {
	return func(c *Client) error {
		c.logf = f
		return nil
	}
}

// Errorf is a Client option to specify the error logging function used.
func Errorf(f func(string, ...interface{})) Option {
	return func(c *Client) error {
		c.errf = f
		return nil
	}
}
