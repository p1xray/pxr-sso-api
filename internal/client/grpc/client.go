package grpcclient

import (
	authpb "github.com/p1xray/pxr-sso-protos/gen/go/auth"
	oidcpb "github.com/p1xray/pxr-sso-protos/gen/go/oidc"
	ssoprofilepb "github.com/p1xray/pxr-sso-protos/gen/go/profile"
)

// GRPCClient provides gRPC clients.
type GRPCClient struct {
	OIDC    oidcpb.OidcClient
	Auth    authpb.AuthClient
	Profile ssoprofilepb.SsoProfileClient
}

// New creates new gRPC client instance.
func New(oidc oidcpb.OidcClient, auth authpb.AuthClient, profile ssoprofilepb.SsoProfileClient) *GRPCClient {
	return &GRPCClient{
		OIDC:    oidc,
		Auth:    auth,
		Profile: profile,
	}
}
