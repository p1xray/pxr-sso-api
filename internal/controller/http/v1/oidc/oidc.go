package oidc

import (
	"github.com/gin-gonic/gin"
	oidcpb "github.com/p1xray/pxr-sso-protos/gen/go/oidc"
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

// Authorize.
//
//	@Summary			Authorize
//	@Description		Authorize
//	@Tags				OIDC
//	@Id 				authorize
//	@Param				input query AuthorizeInput true "Input parameters for authorizer endpoint"
//	@Success			302
//	@Failure			400	{object}	response.errorResponse
//	@Failure			500	{object}	response.errorResponse
//	@Router				/api/v1/oidc/authorize [get]
func (r *Routes) authorize(c *gin.Context) {
	input, err := request.FromQuery[AuthorizeInput](c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	sessions := request.SessionFromCookie(c)
	authorizeRequest := toAuthorizeRequest(input, sessions)

	authorizeResponse, err := r.grpcClient.Authorize(c.Request.Context(), authorizeRequest)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Redirect(c, authorizeResponse.GetRedirectUri())
}

// Token.
//
//	@Summary			Token
//	@Description		Token
//	@Tags				OIDC
//	@Id 				token
//	@Accept				x-www-form-urlencoded
//	@Produce			json
//	@Param				input formData TokenInput true "Input parameters for token endpoint."
//	@Success			200	{object}	TokenOutput
//	@Failure			400	{object}	response.errorResponse
//	@Failure			500	{object}	response.errorResponse
//	@Router				/api/v1/oidc/token [post]
func (r *Routes) token(c *gin.Context) {
	input, err := request.FromForm[TokenInput](c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tokenRequest := toTokenRequest(input)

	tokenResponse, err := r.grpcClient.Token(c.Request.Context(), tokenRequest)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	output := toTokenOutput(tokenResponse)
	response.Success(c, output)
}
