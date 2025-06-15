package http

import (
	"github.com/gin-gonic/gin"
	v1 "pxr-sso-api/internal/controller/http/v1"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "pxr-sso-api/docs"
)

// Handler is handler for http server requests.
type Handler struct {
}

// New creates a new http server request handler.
func New() *Handler {
	return &Handler{}
}

// Init initializes the http server request handler.
func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	h.initAPI(router)
	initSwagger(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	v1Handler := v1.New()
	api := router.Group("/api")
	{
		v1Handler.Init(api)
	}
}

func initSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
