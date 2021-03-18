// +build tools

package tools

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/gunk/gunk"
	_ "github.com/gunk/gunk/docgen"
	_ "github.com/gunk/gunk/generate"
	_ "github.com/gunk/gunk/httprule"
	_ "github.com/gunk/gunk/scopegen"
	_ "github.com/gunk/opt/http"
	_ "github.com/gunk/opt/openapiv2"
)
