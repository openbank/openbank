// Package godop provides a godo (DigitalOcean API) compatible
// autocertdns.Provisioner.
package godop

import (
	"context"
	"errors"
	"strings"

	"github.com/digitalocean/godo"
)

const (
	// allowedRecordType is the allowed record provisioning type.
	allowedRecordType = "TXT"
)

// Client wraps a DigitalOcean godo.Client.
type Client struct {
	client *godo.Client
	domain string
	logf   func(string, ...interface{})
	errf   func(string, ...interface{})
}

// New wraps a godo.Client with a Client that can also handle DNS provisioning
// requests for use with the autocertdns.Manager.
func New(opts ...Option) (*Client, error) {
	var err error

	c := &Client{
		logf: func(string, ...interface{}) {},
	}

	// apply opts
	for _, o := range opts {
		err = o(c)
		if err != nil {
			return nil, err
		}
	}

	// ensure errf is set
	if c.errf == nil {
		c.errf = func(s string, v ...interface{}) {
			c.logf("ERROR: "+s, v...)
		}
	}

	if c.domain == "" || c.client == nil {
		return nil, errors.New("godop missing domain or godo client")
	}

	return c, nil
}

// Provision creates a DNS record of typ, for the specified domain name and
// with the value in token.
func (c *Client) Provision(ctxt context.Context, typ, name, token string) error {
	if typ != allowedRecordType {
		return errors.New("only TXT records are supported")
	}

	// check name
	if !strings.HasSuffix(name, "."+c.domain) {
		return errors.New("invalid domain")
	}
	name = strings.TrimSuffix(name, "."+c.domain)
	if name == "" {
		return errors.New("invalid name")
	}

	// create dns record
	c.logf("provisioning (type: %s, name: %s, token: %s)", typ, name, token)
	_, _, err := c.client.Domains.CreateRecord(ctxt, c.domain, &godo.DomainRecordEditRequest{
		Type: allowedRecordType,
		Name: name,
		Data: token,
	})
	if err != nil {
		c.errf("unable to provision (type: %s, name: %s, token: %s): %v", typ, name, token, err)
	} /*else {
		c.logf("successfully provisioned (type: %s, name: %s, token: %s)", typ, name, token)
	}*/

	return err
}

// Unprovision deletes the DNS record of typ, for the specified domain name,
// and for the record with the specified token as the value.
func (c *Client) Unprovision(ctxt context.Context, typ, name, token string) error {
	var err error

	if typ != allowedRecordType {
		return errors.New("only TXT records are supported")
	}

	// check name
	if !strings.HasSuffix(name, "."+c.domain) {
		return errors.New("invalid domain")
	}
	name = strings.TrimSuffix(name, "."+c.domain)
	if name == "" {
		return errors.New("invalid name")
	}

	// get current records
	//c.logf("retrieving records (type: %s, name: %s, token: %s)", typ, name, token)
	records, _, err := c.client.Domains.Records(ctxt, c.domain, &godo.ListOptions{PerPage: 10000})
	if err != nil {
		c.errf("could not retrieve records (type: %s, name: %s, token: %s): %v", typ, name, token, err)
		return err
	}
	//c.logf("retrieved %d records (type: %s, name: %s, token: %s)", len(records), typ, name, token)

	// find record and delete if TXT record and token matches
	for _, record := range records {
		if record.Name != name || record.Type != allowedRecordType || record.Data != token {
			continue
		}

		c.logf("unprovisioning (type: %s, name: %s, token: %s)", typ, name, token)
		_, err = c.client.Domains.DeleteRecord(ctxt, c.domain, record.ID)
		if err != nil {
			c.errf("unable to unprovision (type: %s, name: %s, token: %s): %v", typ, name, token, err)
		} /*else {
			c.logf("successfully unprovisioned (type: %s, name: %s, token: %s)", typ, name, token)
		}*/
		return err
	}

	c.errf("could not find record (type: %s, name: %s, token: %s)", typ, name, token)

	return errors.New("record not deleted")
}
