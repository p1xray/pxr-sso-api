package grpcclient

import (
	ssoprofilepb "github.com/p1xray/pxr-sso-protos/gen/go/profile"
	ssopb "github.com/p1xray/pxr-sso-protos/gen/go/sso"
)

// GRPCClient provides gRPC clients.
type GRPCClient struct {
	Auth    ssopb.SsoClient
	Profile ssoprofilepb.SsoProfileClient
}

// New creates new gRPC client instance.
func New(auth ssopb.SsoClient, profile ssoprofilepb.SsoProfileClient) *GRPCClient {
	return &GRPCClient{
		Auth:    auth,
		Profile: profile,
	}
}
