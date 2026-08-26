package grpcclient

import (
	"fmt"
	authpb "github.com/p1xray/pxr-sso-protos/gen/go/auth"
	oidcpb "github.com/p1xray/pxr-sso-protos/gen/go/oidc"
	ssoprofilepb "github.com/p1xray/pxr-sso-protos/gen/go/profile"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients interface {
	OIDC() oidcpb.OidcClient
	Auth() authpb.AuthClient
	Profile() ssoprofilepb.SsoProfileClient
}

// GRPCClient provides gRPC clients.
type clients struct {
	oidc    oidcpb.OidcClient
	auth    authpb.AuthClient
	profile ssoprofilepb.SsoProfileClient
}

// New creates new gRPC client instance.
func New(cfg Config) (Clients, error) {
	oidc, err := createOIDCClient(cfg.OIDC)
	if err != nil {
		return nil, err
	}

	auth, err := createAuthClient(cfg.Auth)
	if err != nil {
		return nil, err
	}

	profile, err := createProfileClient(cfg.Auth)
	if err != nil {
		return nil, err
	}

	return &clients{
		oidc:    oidc,
		auth:    auth,
		profile: profile,
	}, nil
}

func (c *clients) OIDC() oidcpb.OidcClient {
	return c.oidc
}

func (c *clients) Auth() authpb.AuthClient {
	return c.auth
}

func (c *clients) Profile() ssoprofilepb.SsoProfileClient {
	return c.profile
}

func createOIDCClient(addr string) (oidcpb.OidcClient, error) {
	con, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("create oidc grpc client: %w", err)
	}

	client := oidcpb.NewOidcClient(con)
	return client, nil
}

func createAuthClient(addr string) (authpb.AuthClient, error) {
	con, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("create auth grpc client: %w", err)
	}

	client := authpb.NewAuthClient(con)
	return client, nil
}

func createProfileClient(addr string) (ssoprofilepb.SsoProfileClient, error) {
	con, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("create profile grpc client: %w", err)
	}

	oauthClient := ssoprofilepb.NewSsoProfileClient(con)
	return oauthClient, nil
}
