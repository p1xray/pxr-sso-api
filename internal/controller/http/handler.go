package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "pxr-sso-api/docs"
	grpcclient "pxr-sso-api/internal/client/grpc"
	v1 "pxr-sso-api/internal/controller/http/v1"
	"time"
)

// Handler is HTTP-handler.
type Handler interface {
	Routes() http.Handler
}

type handler struct {
	grpcClients grpcclient.Clients
}

// NewHandler creates a new http server request handler.
func NewHandler(grpcClients grpcclient.Clients) Handler {
	return &handler{grpcClients: grpcClients}
}

// Routes returns a router with all registered handlers.
func (h *handler) Routes() http.Handler {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = false
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	router.Use(cors.New(config))

	h.initAPI(router)
	initSwagger(router)

	return router
}

func (h *handler) initAPI(router *gin.Engine) {
	v1Handler := v1.New(h.grpcClients)
	api := router.Group("/api")
	{
		v1Handler.Init(api)
	}
}

func initSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
