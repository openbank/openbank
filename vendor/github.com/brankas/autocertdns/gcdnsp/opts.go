package gcdnsp

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"cloud.google.com/go/compute/metadata"
	"github.com/knq/jwt/gserviceaccount"

	dns "google.golang.org/api/dns/v2beta1"
)

// Option represents a Google Cloud DNS Client option.
type Option func(*Client) error

// ProjectID is an option that sets the project ID.
func ProjectID(projectID string) Option {
	return func(c *Client) error {
		c.projectID = projectID
		return nil
	}
}

// ManagedZone is a Client option to set the managed zone.
func ManagedZone(managedZone string) Option {
	return func(c *Client) error {
		c.managedZone = managedZone
		return nil
	}
}

// Domain is a Client option to set the domain.
func Domain(domain string) Option {
	return func(c *Client) error {
		c.domain = domain
		return nil
	}
}

// Nameservers is a Client option to set the nameservers to check.
func Nameservers(nameservers ...string) Option {
	return func(c *Client) error {
		c.nameservers = nameservers
		return nil
	}
}

// PropagationWait is a Client option to set the propagation wait timeout.
func PropagationWait(d time.Duration) Option {
	return func(c *Client) error {
		c.propagationWait = d
		return nil
	}
}

// CheckDelay is a Client option to set the delay between DNS name propagation
// checks.
func CheckDelay(d time.Duration) Option {
	return func(c *Client) error {
		c.checkDelay = d
		return nil
	}
}

// ProvisionDelay is a Client option to set the delay after a successful
// provision and name propagation.
func ProvisionDelay(d time.Duration) Option {
	return func(c *Client) error {
		c.provisionDelay = d
		return nil
	}
}

// DNSService is an option that sets the Google Cloud DNS service to use.
func DNSService(dnsService *dns.Service) Option {
	return func(c *Client) error {
		c.dnsService = dnsService
		return nil
	}
}

// HTTPClient is a Client option that sets the http.Client used.
func HTTPClient(client *http.Client) Option {
	return func(c *Client) error {
		// create dns service
		l, err := dns.New(client)
		if err != nil {
			return err
		}
		return DNSService(l)(c)
	}
}

// requiredScopes are the oauth2 scopes required for cloud dns.
var requiredScopes = []string{
	dns.CloudPlatformScope,
	dns.NdevClouddnsReadwriteScope,
}

// GoogleServiceAccountCredentialsJSON is an option that creates the Google
// Cloud DNS service using the supplied Google service account credentials.
//
// Google Service Account credentials can be downloaded from the Google Cloud
// console: https://console.cloud.google.com/iam-admin/serviceaccounts/
func GoogleServiceAccountCredentialsJSON(buf []byte) Option {
	return func(c *Client) error {
		var err error

		// load credentials
		gsa, err := gserviceaccount.FromJSON(buf)
		if err != nil {
			return err
		}

		// check project id
		if gsa.ProjectID == "" {
			return errors.New("google service account credentials missing project_id")
		}

		// set project id
		err = ProjectID(gsa.ProjectID)(c)
		if err != nil {
			return err
		}

		// create token source
		ts, err := gsa.TokenSource(nil, requiredScopes...)
		if err != nil {
			return err
		}

		// set client
		return HTTPClient(&http.Client{
			Transport: &oauth2.Transport{
				Source: oauth2.ReuseTokenSource(nil, ts),
			},
		})(c)
	}
}

// GoogleServiceAccountCredentialsFile is an option that loads Google Service
// Account credentials for use with the Client from the specified
// file.
//
// Google Service Account credentials can be downloaded from the Google Cloud
// console: https://console.cloud.google.com/iam-admin/serviceaccounts/
func GoogleServiceAccountCredentialsFile(path string) Option {
	return func(c *Client) error {
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		return GoogleServiceAccountCredentialsJSON(buf)(c)
	}
}

// GoogleComputeCredentials is an option that loads the Google Service Account
// credentials from the GCE metadata associated with the GCE compute instance.
//
// If serviceAccount is empty, then the default service account credentials
// associated with the GCE instance will be used.
func GoogleComputeCredentials(serviceAccount string) Option {
	return func(c *Client) error {
		var err error

		// get compute metadata scopes associated with the service account
		scopes, err := metadata.Scopes(serviceAccount)
		if err != nil {
			return err
		}

		// check if all the necessary scopes are provided
		for _, s := range requiredScopes {
			if !sliceContains(scopes, s) {
				// NOTE: if you are seeing this error, you probably need to
				// recreate your compute instance with the correct scope
				//
				// as of August 2016, there is not a way to add a scope to an
				// existing compute instance
				return fmt.Errorf("missing required scope %s in compute metadata", s)
			}
		}

		return HTTPClient(&http.Client{
			Transport: &oauth2.Transport{
				Source: google.ComputeTokenSource(serviceAccount),
			},
		})(c)
	}
}

// sliceContains returns true if haystack contains needle.
func sliceContains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
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

// IgnorePropagationErrors is a Client option to ignore propagation errors.
func IgnorePropagationErrors(c *Client) error {
	c.ignorePropagationErrors = true
	return nil
}
