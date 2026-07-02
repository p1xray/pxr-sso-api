package auth

import (
	"github.com/gin-gonic/gin"
	authpb "github.com/p1xray/pxr-sso-protos/gen/go/auth"
	"pxr-sso-api/internal/controller/http/request"
	"pxr-sso-api/internal/controller/http/response"
)

// Routes provides routes for authentication.
type Routes struct {
	grpcClient authpb.AuthClient
}

// InitRoutes initializes the routes for authentication.
func InitRoutes(api *gin.RouterGroup, grpcClient authpb.AuthClient) {
	r := &Routes{grpcClient: grpcClient}

	auth := api.Group("/auth")
	{
		auth.POST("/signin", r.signin)
		auth.POST("/signup", r.signup)
		auth.POST("/signout", r.signout)
		auth.POST("/consent", r.consent)
	}
}

// Sign in.
//
//	@Summary			Sign in
//	@Description		Sign in
//	@Tags				Auth
//	@Id 				signin
//	@Accept				x-www-form-urlencoded
//	@Produce			json
//	@Param				input formData LoginInput true "Input parameters for sign in endpoint."
//	@Success			200	{object}	LoginOutput
//	@Failure			400	{object}	response.errorResponse
//	@Failure			500	{object}	response.errorResponse
//	@Router				/api/v1/auth/login [post]
func (r *Routes) signin(c *gin.Context) {
	input, err := request.FromForm[LoginInput](c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	loginRequest := toLoginRequest(input)

	loginResponse, err := r.grpcClient.Login(c.Request.Context(), loginRequest)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	session := loginResponse.GetSession()
	response.SetSessionCookie(c, session.GetName(), session.GetValue())

	output := toLoginOutput(loginResponse)
	response.Success(c, output)
}

// Sign up.
//
//	@Summary			Sign up
//	@Description		Sign up
//	@Tags				Auth
//	@Id 				signup
//	@Accept				x-www-form-urlencoded
//	@Produce			json
//	@Param				input formData RegisterInput true "Input parameters for sign up endpoint."
//	@Success			200	{object}	RegisterOutput
//	@Failure			400	{object}	response.errorResponse
//	@Failure			500	{object}	response.errorResponse
//	@Router				/api/v1/auth/register [post]
func (r *Routes) signup(c *gin.Context) {
	input, err := request.FromForm[RegisterInput](c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	registerRequest := toRegisterRequest(input)

	registerResponse, err := r.grpcClient.Register(c.Request.Context(), registerRequest)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	session := registerResponse.GetSession()
	response.SetSessionCookie(c, session.GetName(), session.GetValue())

	output := toRegisterOutput(registerResponse)
	response.Success(c, output)
}

// Sign out.
//
//	@Summary			Sign out
//	@Description		Sign out
//	@Tags				Auth
//	@Id 				signout
//	@Accept				x-www-form-urlencoded
//	@Produce			json
//	@Param				input formData SignoutInput true "Input parameters for sign out endpoint."
//	@Success			200	{object}	SignoutOutput
//	@Failure			400	{object}	response.errorResponse
//	@Failure			500	{object}	response.errorResponse
//	@Router				/api/v1/auth/signout [post]
func (r *Routes) signout(c *gin.Context) {
	response.Success(c, true)
}

// Consent.
//
//	@Summary			Consent
//	@Description		Consent
//	@Tags				Auth
//	@Id 				consent
//	@Accept				x-www-form-urlencoded
//	@Produce			json
//	@Param				input formData ConsentInput true "Input parameters for consent endpoint."
//	@Success			200	{object}	ConsentOutput
//	@Failure			400	{object}	response.errorResponse
//	@Failure			500	{object}	response.errorResponse
//	@Router				/api/v1/auth/consent [post]
func (r *Routes) consent(c *gin.Context) {
	input, err := request.FromForm[ConsentInput](c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	sessions := request.SessionFromCookie(c)
	consentRequest := toConsentRequest(input, sessions)

	consentResponse, err := r.grpcClient.Consent(c.Request.Context(), consentRequest)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	output := toConsentOutput(consentResponse)
	response.Success(c, output)
}
