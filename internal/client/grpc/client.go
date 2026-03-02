package grpcclient

import (
	oauthpb "github.com/p1xray/pxr-sso-protos/gen/go/oauth"
	ssoprofilepb "github.com/p1xray/pxr-sso-protos/gen/go/profile"
	ssopb "github.com/p1xray/pxr-sso-protos/gen/go/sso"
)

// GRPCClient provides gRPC clients.
type GRPCClient struct {
	Auth    ssopb.SsoClient
	Profile ssoprofilepb.SsoProfileClient
	OAuth   oauthpb.OauthClient
}

// New creates new gRPC client instance.
func New(auth ssopb.SsoClient, profile ssoprofilepb.SsoProfileClient, oauth oauthpb.OauthClient) *GRPCClient {
	return &GRPCClient{
		Auth:    auth,
		Profile: profile,
		OAuth:   oauth,
	}
}
