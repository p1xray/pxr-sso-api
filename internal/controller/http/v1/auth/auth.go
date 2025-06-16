package auth

import (
	"github.com/gin-gonic/gin"
	ssopb "github.com/p1xray/pxr-sso-protos/gen/go/sso"
	"pxr-sso-api/internal/server"
)

// Routes provides routes for authentication.
type Routes struct {
	grpcAuthClient ssopb.SsoClient
}

// InitRoutes initializes the routes for authentication.
func InitRoutes(api *gin.RouterGroup, grpcAuthClient ssopb.SsoClient) {
	ar := &Routes{grpcAuthClient: grpcAuthClient}

	auth := api.Group("/auth")
	{
		auth.POST("/login", ar.login)
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
//	@Param				input query LoginInput true "Input parameters for user login."
//	@Success			200	{object}	server.dataResponse[LoginOutput]
//	@Failure			500	{object}	server.dataResponse[LoginOutput]
//	@Router				/api/v1/auth/login [post]
func (a *Routes) login(c *gin.Context) {
	inp, err := server.GetInputFromQuery[LoginInput](c)
	if err != nil {
		server.ErrorResponse[LoginOutput](c, err.Error())
		return
	}

	grpcLoginRequest := &ssopb.LoginRequest{
		Username:    inp.Username,
		Password:    inp.Password,
		ClientCode:  inp.ClientCode,
		UserAgent:   server.GetUserAgent(c),
		Fingerprint: server.GetFingerprint(c),
		Issuer:      server.GetHost(c),
	}

	grpcLoginResponse, err := a.grpcAuthClient.Login(
		c.Request.Context(),
		grpcLoginRequest)
	if err != nil {
		// TODO: check error from gRPC server and return invalid credentials error

		server.ErrorResponse[LoginOutput](c, err.Error())
		return
	}

	response := &LoginOutput{
		AccessToken:  grpcLoginResponse.GetAccessToken(),
		RefreshToken: grpcLoginResponse.GetRefreshToken(),
	}

	server.SuccessResponse(c, response)
}
