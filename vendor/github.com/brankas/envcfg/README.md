# About envcfg

`envcfg` is a package to retrieve configuration settings from the environment,
in adherence with [12-factor][12-factor] app principles.

## Installing

Install in the usual [Go][go-project] fashion:

```sh
$ go get -u github.com/brankas/envcfg
```

## Using

`envcfg` can be used similarly to the following:

```go
// examples/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brankas/envcfg"
)

func main() {
	// load config from the default APP_CONFIG environment variable
	config, err := envcfg.New()
	if err != nil {
		log.Fatal(err)
	}

	// load additional config from SOME_OTHER_VAR environment variable
	config2, err := envcfg.New(
		envcfg.VarName("SOME_OTHER_VAR"),
	)
	if err != nil {
		log.Fatal(err)
	}
	config2 = config2

	// read a config key
	val := config.GetKey("mysection.mykeyname")
	log.Printf("> val: %s", val)

	// create a http.Server with a host, port, and TLS based on config pulled
	// from environment
	s := &http.Server{
		Addr:      fmt.Sprintf("%s:%d", config.Host(), config.Port()),
		TLSConfig: config.TLS(nil),
	}
	log.Fatal(s.ListenAndServe())
}
```

Please see the [GoDoc][godoc-api] listing for a full API listing.

### Environment Variables

By default, `envcfg` loads configuration data from the base64-encoded
environment variable `APP_CONFIG`, and expects a [git config][git-config]
style configuration file of sections, subsections, and keys.

### Config files and key data

Configuration data is stored in [git config][git-config] style configuration
files, and uses the [github.com/knq/ini][knq-ini] package for parsing
configuration data.


The below is the [`examples/sample.config`](examples/sample.config):

```ini
# examples/sample.config
[runtime]
environment="$ENV||production"                   ; production / development / etc.

[server]
host="$HOST||example.com"                        ; hostname
port="$PORT||8443"                               ; port
certs="$CERTS||./env/certs"                      ; certificate directory cache
certProvider="dns:godo:<domain>:<email>:<token>" ; certificate provider

[google]
creds="$GOOGLECREDS||env/gsa.json||file"         ; "file" encoded gsa credentials loaded from disk,
                                                 ; relative to the current working directory

[service "something"]
key="$NAME||subdir/myfile||relfile"              ; "relfile" encoded file will be loaded relative to the original config file
                                                 ; for example, if the config file was on disk at /etc/myapp then
                                                 ; the /etc/myapp/subdir/myfile value would be loaded

[example]
b64value="$B64VALUE||e30K||base64"               ; "base64" encoded value
```

### Certificate Providers

The following certificate providers are supported for `server.certProvider`:

| Name   | Type                   | Example                                                                                    |
|--------|------------------------|--------------------------------------------------------------------------------------------|
| `auto` | LetsEncrypt TLS-SNI-01 | `auto`                                                                                     |
| `disk` | Certificate on disk    | `disk:/path/to/cert.pem:/path/to/key.pem`                                                  |
| `dns`  | LetsEncrypt DNS-01     | `dns:clouddns:mydomain.com:ken@example.com:my-managed-zone-name:/path/to/credentials.json` |

#### LetsEncrypt DNS-01 Provider

DNS certificate provider strings have a general form of `dns:<type>:<domain>:<email>:...`.

| Type       | DNS Provider     | Parameters                                                              |
|------------|------------------|-------------------------------------------------------------------------|
| `godo`     | DigitalOcean DNS | `<type>:<domain>:<email>:<token>`                                       |
| `clouddns` | Google Cloud DNS | `<type>:<domain>:<email>:<managed-zone-name>:/path/to/credentials.json` |

[12-factor]: https://12factor.net
[go-project]: https://golang.org/project
[godoc-api]: https://godoc.org/github.com/brankas/envcfg
[git-config]: https://git-scm.com/docs/git-config
[knq-ini]: https://github.com/knq/ini
