package auth

import (
	"github.com/gin-gonic/gin"
	authpb "github.com/p1xray/pxr-sso-protos/gen/go/auth"
	"pxr-sso-api/internal/constants"
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
		auth.POST("/login", r.login)
		auth.POST("/register", r.register)
		auth.POST("/consent", r.consent)
	}
}

// Login.
//
//	@Summary			Login
//	@Description		Login
//	@Tags				Auth
//	@Id 				login
//	@Accept				mpfd
//	@Produce			json
//	@Param        		X-Fingerprint	  header    string    true   	"User browser fingerprint."
//	@Param				input formData LoginInput true "Input parameters for user login."
//	@Success			200	{object}	server.dataResponse[LoginOutput]
//	@Failure			500	{object}	server.dataResponse[LoginOutput]
//	@Router				/api/v1/auth/login [post]
func (r *Routes) login(c *gin.Context) {
	input, err := request.FromForm[LoginInput](c)
	if err != nil {
		response.BadRequest(c, constants.ErrorCodeInvalidRequest, err.Error(), "")
		return
	}

	loginRequest := toLoginRequest(input)

	loginResponse, err := r.grpcClient.Login(c.Request.Context(), loginRequest)
	if err != nil {
		response.InternalServerError(c, constants.ErrorCodeInvalidRequest, err.Error(), "")
		return
	}

	session := loginResponse.GetSession()
	response.SetSessionCookie(c, session.GetName(), session.GetValue())

	output := toLoginOutput(loginResponse)
	response.Success(c, output)
}

// Register.
//
//	@Summary			Register
//	@Description		Register
//	@Tags				Auth
//	@Id 				register
//	@Accept				mpfd
//	@Produce			json
//	@Param        		X-Fingerprint	  header    string    true   	"User browser fingerprint."
//	@Param				input formData RegisterInput true "Input parameters for user register."
//	@Param				avatar_file formData file false "Avatar file."
//	@Success			200	{object}	server.dataResponse[RegisterOutput]
//	@Failure			500	{object}	server.dataResponse[RegisterOutput]
//	@Router				/api/v1/auth/register [post]
func (r *Routes) register(c *gin.Context) {
	input, err := request.FromForm[RegisterInput](c)
	if err != nil {
		response.BadRequest(c, constants.ErrorCodeInvalidRequest, err.Error(), "")
		return
	}

	registerRequest := toRegisterRequest(input)

	registerResponse, err := r.grpcClient.Register(c.Request.Context(), registerRequest)
	if err != nil {
		response.InternalServerError(c, constants.ErrorCodeInvalidRequest, err.Error(), "")
		return
	}

	session := registerResponse.GetSession()
	response.SetSessionCookie(c, session.GetName(), session.GetValue())

	output := toRegisterOutput(registerResponse)
	response.Success(c, output)
}

func (r *Routes) consent(c *gin.Context) {
	input, err := request.FromForm[ConsentInput](c)
	if err != nil {
		response.BadRequest(c, constants.ErrorCodeInvalidRequest, err.Error(), "")
		return
	}

	consentRequest := toConsentRequest(input)

	consentResponse, err := r.grpcClient.Consent(c.Request.Context(), consentRequest)
	if err != nil {
		response.InternalServerError(c, constants.ErrorCodeInvalidRequest, err.Error(), "")
		return
	}

	output := toConsentOutput(consentResponse)
	response.Success(c, output)
}
