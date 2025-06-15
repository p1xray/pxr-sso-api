package grpcapp

import (
	"fmt"
	ssopb "github.com/p1xray/pxr-sso-protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	grpcclient "pxr-sso-api/internal/client/grpc"
	"pxr-sso-api/internal/config"
	"pxr-sso-api/internal/lib/logger/sl"
)

// App is gRPC client application.
type App struct {
	log    *slog.Logger
	config *config.Config
}

// New creates new instance of the gRPC client application.
func New(log *slog.Logger, cfg *config.Config) *App {
	return &App{
		log:    log,
		config: cfg,
	}
}

// CreateGRPCClient creates new gRPC clients.
func (a *App) CreateGRPCClient() *grpcclient.GRPCClient {
	auth, err := a.createAuthClient()
	if err != nil {
		a.log.Error("failed creating auth grpc client", sl.Err(err))
	}

	return grpcclient.New(auth)
}

func (a *App) createAuthClient() (ssopb.SsoClient, error) {
	const op = "grpcapp.createAuthClient"

	con, err := grpc.NewClient(
		a.config.GRPCClients.Auth.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	authClient := ssopb.NewSsoClient(con)
	return authClient, nil
}
