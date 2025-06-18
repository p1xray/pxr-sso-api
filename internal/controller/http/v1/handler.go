package v1

import (
	"github.com/gin-gonic/gin"
	grpcclient "pxr-sso-api/internal/client/grpc"
	"pxr-sso-api/internal/controller/http/v1/auth"
	"pxr-sso-api/internal/controller/http/v1/ping"
	"pxr-sso-api/internal/controller/http/v1/profile"
)

// Handler is request handler for API v1.
type Handler struct {
	grpcClient *grpcclient.GRPCClient
}

// New creates new instance of the API v1 request handler.
func New(grpcClient *grpcclient.GRPCClient) *Handler {
	return &Handler{grpcClient: grpcClient}
}

// Init initializes the API v1 request handler.
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		ping.InitRoutes(v1)
		auth.InitRoutes(v1, h.grpcClient.Auth)
		profile.InitRoutes(v1, h.grpcClient.Profile)
	}
}
