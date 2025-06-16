package grpcclient

import ssopb "github.com/p1xray/pxr-sso-protos/gen/go/sso"

// GRPCClient provides gRPC clients.
type GRPCClient struct {
	Auth ssopb.SsoClient
}

// New creates new gRPC client instance.
func New(auth ssopb.SsoClient) *GRPCClient {
	return &GRPCClient{Auth: auth}
}
