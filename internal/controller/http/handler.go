package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "pxr-sso-api/docs"
	grpcclient "pxr-sso-api/internal/client/grpc"
	v1 "pxr-sso-api/internal/controller/http/v1"
	"time"
)

// Handler is handler for http server requests.
type Handler struct {
	grpcClient *grpcclient.GRPCClient
}

// New creates a new http server request handler.
func New(grpcClient *grpcclient.GRPCClient) *Handler {
	return &Handler{grpcClient: grpcClient}
}

// Init initializes the http server request handler.
func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	router.Use(cors.New(config))

	h.initAPI(router)
	initSwagger(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	v1Handler := v1.New(h.grpcClient)
	api := router.Group("/api")
	{
		v1Handler.Init(api)
	}
}

func initSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
