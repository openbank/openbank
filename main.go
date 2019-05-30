package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"

	"github.com/brankas/envcfg"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/openbank/openbank/pkg/apierror"
	"github.com/openbank/openbank/services"
	"github.com/openbank/openbank/services/accounts"
	"github.com/openbank/openbank/services/transactions"
	"github.com/openbank/openbank/pkg/logging"
	"github.com/openbank/openbank/storage/inmemory"
)

var (
	logger  *logrus.Entry
	version = "devel"
	service = "openbank"

	config *envcfg.Envcfg
)

func init() {
	logger = logging.NewLogger().WithFields(logrus.Fields{
		"service": service,
		"version": version,
	})
}

func main() {
	logger.Info("starting openbank service")

	var err error
	// Load the configuration
	config, err = envcfg.New()
	if err != nil {
		logging.WithError(err, logger).Fatal("error loading configuration")
	}

	// TLS and grpc credentials
	tlsConfig := config.TLS()
	grpcCreds := credentials.NewTLS(tlsConfig)

	// grpc
	grpcServer, err := setupGRPCServer(grpcCreds)
	if err != nil {
		logging.WithError(err, logger).Fatal("unable to start grpc server")
	}

	// grpc-web handler
	grpcWebServer := grpcweb.WrapServer(grpcServer)

	// grpc-gateway
	grpcGW, err := setupGRPCGateway()
	if err != nil {
		logging.WithError(err, logger).Fatal("failed to setup grpc gateway")
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port()))
	if err != nil {
		logging.WithError(err, logger).WithField("port", config.Port).Fatal("unable to listen on port")
	}
	logger.Infof("listening on %v", l.Addr())

	// cors handler
	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods: []string{
			"HEAD", "GET", "POST",
			"PUT", "PATCH", "DELETE",
		},
		AllowedOrigins: []string{
			"*.brank.as",
		},
		AllowedHeaders: []string{
			"Content-Type",
			"X-Grpc-Web",
			"Authorization",
		},
		Debug: config.GetBool("cors.debug"),
	})

	logger.Fatal(http.Serve(
		tls.NewListener(l, tlsConfig),
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Ver", version)
			switch {
			// Handle cors
			case r.Method == "OPTIONS":
				c.HandlerFunc(w, r)
			// grpc
			case r.ProtoMajor == 2 && r.Header.Get("Content-Type") == "application/grpc":
				grpcServer.ServeHTTP(w, r)
			// grpc-web
			case grpcWebServer.IsGrpcWebRequest(r):
				grpcWebServer.ServeHTTP(w, r)
			// grpc-gateway
			default:
				grpcGW.ServeHTTP(w, r)
			}
		})))
}

func setupGRPCServer(creds credentials.TransportCredentials) (*grpc.Server, error) {
	// TODO: setup auth interceptor
	// ts, err := token.NewService(nil)
	// if err != nil {
	// 	return nil, err
	// }

	options := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			logging.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logger),
			// TODO: setup auth interceptor
			// token.UnaryAuthInterceptor(ts.ServerAuth(token.OrgSvc)),
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc.Creds(creds),
		grpc.StatsHandler(&ocgrpc.ServerHandler{IsPublicEndpoint: true}),
	}

	srv := grpc.NewServer(options...)

	reflection.Register(srv)

	// Create a new storage layer
	inmemStorage := inmemory.NewStorage()

	services.RegisterServices(
		srv,
		accounts.RegisterService(inmemStorage.AccountStore),
		transactions.RegisterService(inmemStorage.TransactionStore),
	)

	return srv, nil
}

func setupGRPCGateway() (*runtime.ServeMux, error) {
	// Currently gunk generated files have json name
	// based on go naming convention. The workaround
	// is to use a custom marshaller.
	// See issue https://github.com/gunk/gunk/issues/33
	m := &runtime.JSONPb{}
	mo := runtime.WithMarshalerOption("*", m)
	mux := runtime.NewServeMux(mo)

	runtime.HTTPError = apierror.CustomHTTPError(config.GetKey("docs.base"))

	options := []grpc.DialOption{grpc.WithTransportCredentials(
		credentials.NewTLS(&tls.Config{
			ServerName: config.Host(),
		}),
	)}
	address := fmt.Sprintf("localhost:%d", config.Port())

	if err := services.RegisterGateway(
		mux,
		address,
		options,
		accounts.RegisterGateway,
		transactions.RegisterGateway,
	); err != nil {
		return nil, err
	}
	return mux, nil
}
