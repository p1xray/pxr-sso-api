package oidc

import (
	"github.com/gin-gonic/gin"
	oidcpb "github.com/p1xray/pxr-sso-protos/gen/go/oidc"
	"pxr-sso-api/internal/constants"
	"pxr-sso-api/internal/controller/http/request"
	"pxr-sso-api/internal/controller/http/response"
)

// Routes provides routes for OIDC.
type Routes struct {
	grpcClient oidcpb.OidcClient
}

// InitRoutes initializes the routes for OIDC.
func InitRoutes(api *gin.RouterGroup, grpcClient oidcpb.OidcClient) {
	r := &Routes{grpcClient: grpcClient}

	oidc := api.Group("/oidc")
	{
		oidc.GET("/authorize", r.authorize)
		oidc.POST("/token", r.token)
	}
}

func (r *Routes) authorize(c *gin.Context) {
	input, err := request.FromQuery[AuthorizeInput](c)
	if err != nil {
		response.BadRequest(c, constants.ErrorCodeInvalidRequest, err.Error(), "")
		return
	}

	sessions := request.SessionFromCookie(c)
	authorizeRequest := toAuthorizeRequest(input, sessions)

	authorizeResponse, err := r.grpcClient.Authorize(c.Request.Context(), authorizeRequest)
	if err != nil {
		response.InternalServerError(c, constants.ErrorCodeInternalServerError, err.Error(), "")
		return
	}

	response.Redirect(c, authorizeResponse.GetRedirectUri())
}

func (r *Routes) token(c *gin.Context) {
	input, err := request.FromForm[TokenInput](c)
	if err != nil {
		response.BadRequest(c, constants.ErrorCodeInvalidRequest, err.Error(), "")
		return
	}

	tokenRequest := toTokenRequest(input)

	tokenResponse, err := r.grpcClient.Token(c.Request.Context(), tokenRequest)
	if err != nil {
		response.InternalServerError(c, constants.ErrorCodeInternalServerError, err.Error(), "")
		return
	}

	output := toTokenOutput(tokenResponse)
	response.Success(c, output)
}
