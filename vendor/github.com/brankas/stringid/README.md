# About

Package `stringid` provides string based ID generators, and accompanying
middleware for assigning a request ID to standard HTTP request contexts.

## Installing

Install in the usual [Go][go-project] fashion:

```sh
$ go get -u github.com/brankas/stringid
```

## Using

`stringid` can be used similarly to the following:

```go
// examples/goji/main.go
package main

import (
	"fmt"
	"net/http"

	goji "goji.io"
	"goji.io/pat"

	"github.com/brankas/stringid"
)

func main() {
	mux := goji.NewMux()
	mux.Use(stringid.Middleware())
	mux.HandleFunc(pat.New("/"), func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "request id: %s\n", stringid.FromContext(req.Context()))
	})

	http.ListenAndServe(":3000", mux)
}
```

Please see the [GoDoc listing][godoc] for the full API listing.

[go-project]: https://golang.org/project
[godoc]: https://godoc.org/github.com/brankas/stringid
