package grpcclient

// Config is the configuration of gRPC clients.
type Config struct {
	OIDC    string `yaml:"oidc" env:"PXR_SSO_API_GRPC_CLIENT_OIDC" env-required:"true"`
	Auth    string `yaml:"auth" env:"PXR_SSO_API_GRPC_CLIENT_AUTH" env-required:"true"`
	Profile string `yaml:"profile" env:"PXR_SSO_API_GRPC_CLIENT_PROFILE" env-required:"true"`
}
