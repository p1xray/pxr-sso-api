package app

import (
	"log/slog"
	grpcclient "pxr-sso-api/internal/client/grpc"
	"pxr-sso-api/internal/controller/http"
	"pxr-sso-api/pkg/logger"
)

// diContainer is a lazy-initialized dependency container.
type diContainer struct {
	// Infrastructure:
	// 	config, logger, metrics, etc
	cfg *Config
	log *slog.Logger

	// API:
	// 	gRPC
	grpcClients grpcclient.Clients

	// 	HTTP
	handler http.Handler
}

// newDIContainer creates a new empty DI container.
// All fields are nil, dependencies will be created lazily on the first access.
func newDIContainer() *diContainer {
	return &diContainer{}
}

// Config returns the application configuration.
func (d *diContainer) Config() *Config {
	if d.cfg == nil {
		cfgLoader := newConfigLoader()
		cfg := cfgLoader.MustLoad()

		d.cfg = cfg
	}

	return d.cfg
}

// Logger returns the logger configured depending on the environment specified in the config.
func (d *diContainer) Logger() *slog.Logger {
	if d.log == nil {
		d.log = logger.SetupLogger(d.Config().Env)
	}

	return d.log
}

// GRPCClients returns the gRPC clients.
func (d *diContainer) GRPCClients() grpcclient.Clients {
	if d.grpcClients == nil {
		clients, err := grpcclient.New(d.Config().GRPCClients)
		if err != nil {
			panic(err)
		}

		d.grpcClients = clients
	}

	return d.grpcClients
}

// Handler returns HTTP-handler.
func (d *diContainer) Handler() http.Handler {
	if d.handler == nil {
		d.handler = http.NewHandler(d.GRPCClients())
	}

	return d.handler
}
