package v1

import (
	"github.com/gin-gonic/gin"
	grpcclient "pxr-sso-api/internal/client/grpc"
	"pxr-sso-api/internal/controller/http/v1/auth"
	"pxr-sso-api/internal/controller/http/v1/oidc"
	"pxr-sso-api/internal/controller/http/v1/ping"
)

// Handler is request handler for API v1.
type Handler struct {
	grpcClients grpcclient.Clients
}

// New creates new instance of the API v1 request handler.
func New(grpcClients grpcclient.Clients) *Handler {
	return &Handler{grpcClients: grpcClients}
}

// Init initializes the API v1 request handler.
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		ping.InitRoutes(v1)
		oidc.InitRoutes(v1, h.grpcClients.OIDC())
		auth.InitRoutes(v1, h.grpcClients.Auth())
		// profile.InitRoutes(v1, h.grpcClient.Profile())
	}
}
