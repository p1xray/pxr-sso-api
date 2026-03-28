package oauth

import (
	"github.com/gin-gonic/gin"
	oauthpb "github.com/p1xray/pxr-sso-protos/gen/go/oauth"
	"pxr-sso-api/internal/constants"
	"pxr-sso-api/internal/controller/http/request"
	"pxr-sso-api/internal/controller/http/response"
	"pxr-sso-api/internal/controller/http/v1/oauth/converter"
	"pxr-sso-api/internal/controller/http/v1/oauth/model"
	"strings"
)

// Routes provides routes for OAuth2.1.
type Routes struct {
	grpcOAuthClient oauthpb.OauthClient
}

// InitRoutes initializes the routes for OAuth2.1.
func InitRoutes(api *gin.RouterGroup, grpcOAuthClient oauthpb.OauthClient) {
	r := &Routes{grpcOAuthClient: grpcOAuthClient}

	oauth := api.Group("/oauth")
	{
		oauth.GET("/authorize", r.authorize)
		oauth.POST("/login", r.login)
		oauth.POST("/register", r.register)
		oauth.POST("/consent", r.consent)
		oauth.POST("/token", r.token)
	}
}

func (r *Routes) authorize(c *gin.Context) {
	input, err := request.FromQuery[model.AuthorizeInput](c)
	if err != nil {
		response.BadRequest(c, constants.ErrorCodeInvalidRequest, err.Error(), "")
		return
	}

	authorizeRequest := &oauthpb.AuthorizeRequest{
		ResponseType:        input.ResponseType,
		ClientId:            input.ClientID,
		RedirectUri:         input.RedirectURI,
		CodeChallenge:       input.CodeChallenge,
		CodeChallengeMethod: input.CodeChallengeMethod,
		State:               input.State,
		Audience:            input.Audience,
		Scope:               input.Scope,
	}

	authorizeResponse, err := r.grpcOAuthClient.Authorize(c.Request.Context(), authorizeRequest)
	serviceResponse := converter.ToAuthorizeServiceResponse(authorizeResponse, err)

	response.TryRedirect(c, serviceResponse)
}

func (r *Routes) login(c *gin.Context) {
	input, err := request.FromForm[model.LoginInput](c)
	if err != nil {
		response.BadRequest(c, constants.ErrorCodeInvalidRequest, err.Error(), "")
		return
	}

	loginRequest := &oauthpb.LoginRequest{
		FlowId:       input.FlowID,
		ResponseType: input.ResponseType,
		ClientId:     input.ClientID,
		RedirectUri:  input.RedirectURI,
		State:        input.State,
		Scope:        strings.Split(input.Scope, " "),
		Username:     input.Username,
		Password:     input.Password,
	}

	loginResponse, err := r.grpcOAuthClient.Login(c.Request.Context(), loginRequest)
	serviceResponse := converter.ToLoginServiceResponse(loginResponse, err)

	response.TrySuccess(c, serviceResponse)
}

func (r *Routes) register(c *gin.Context) {
	input, err := request.FromForm[model.RegisterInput](c)
	if err != nil {
		response.BadRequest(c, constants.ErrorCodeInvalidRequest, err.Error(), "")
		return
	}

	registerRequest := &oauthpb.RegisterRequest{
		FlowId:       input.FlowID,
		ResponseType: input.ResponseType,
		ClientId:     input.ClientID,
		RedirectUri:  input.RedirectURI,
		State:        input.State,
		Scope:        strings.Split(input.Scope, " "),
		Username:     input.Username,
		Password:     input.Password,
		FullName:     input.FullName,
	}

	registerResponse, err := r.grpcOAuthClient.Register(c.Request.Context(), registerRequest)
	serviceResponse := converter.ToRegisterServiceResponse(registerResponse, err)

	response.TrySuccess(c, serviceResponse)
}

func (r *Routes) consent(c *gin.Context) {
	input, err := request.FromForm[model.ConsentInput](c)
	if err != nil {
		response.BadRequest(c, constants.ErrorCodeInvalidRequest, err.Error(), "")
		return
	}

	consentRequest := &oauthpb.ConsentRequest{
		FlowId:       input.FlowID,
		ResponseType: input.ResponseType,
		ClientId:     input.ClientID,
		RedirectUri:  input.RedirectURI,
		State:        input.State,
		Scope:        strings.Split(input.Scope, " "),
	}

	consentResponse, err := r.grpcOAuthClient.Consent(c.Request.Context(), consentRequest)
	serviceResponse := converter.ToConsentServiceResponse(consentResponse, err)

	response.TrySuccess(c, serviceResponse)
}

func (r *Routes) token(c *gin.Context) {
	input, err := request.FromForm[model.TokenInput](c)
	if err != nil {
		response.BadRequest(c, constants.ErrorCodeInvalidRequest, err.Error(), "")
		return
	}

	tokenRequest := &oauthpb.TokenRequest{
		GrantType:    input.GrantType,
		ClientId:     input.ClientID,
		Code:         input.Code,
		RedirectUri:  input.RedirectURI,
		CodeVerifier: input.CodeVerifier,
	}

	tokenResponse, err := r.grpcOAuthClient.Token(c.Request.Context(), tokenRequest)
	serviceResponse := converter.ToTokenServiceResponse(tokenResponse, err)

	response.TrySuccess(c, serviceResponse)
}
