package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/wrappers"
	ssopb "github.com/p1xray/pxr-sso-protos/gen/go/sso"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		auth.POST("/register", ar.register)
		auth.POST("/refresh-tokens", ar.refreshTokens)
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
func (a *Routes) login(c *gin.Context) {
	inp, err := server.GetInputFromForm[LoginInput](c)
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
func (a *Routes) register(c *gin.Context) {
	inp, err := server.GetInputFromForm[RegisterInput](c)
	if err != nil {
		server.ErrorResponse[RegisterOutput](c, err.Error())
		return
	}

	var dateOfBirthPb *timestamppb.Timestamp
	if inp.DateOfBirth != nil {
		dateOfBirthPb = timestamppb.New(*inp.DateOfBirth)
	}

	var genderPb ssopb.Gender
	if inp.Gender != nil {
		genderPb = ssopb.Gender(*inp.Gender)
	}

	var avatarFileKeyPb *wrappers.StringValue
	if inp.AvatarFile != nil {
		// TODO: save file to files storage.
		avatarFileKeyPb = &wrappers.StringValue{Value: inp.AvatarFile.Filename}
	}

	grpcRegisterRequest := &ssopb.RegisterRequest{
		Username:      inp.Username,
		Password:      inp.Password,
		ClientCode:    inp.ClientCode,
		Fio:           inp.Fio,
		DateOfBirth:   dateOfBirthPb,
		Gender:        genderPb,
		AvatarFileKey: avatarFileKeyPb,
		UserAgent:     server.GetUserAgent(c),
		Fingerprint:   server.GetFingerprint(c),
		Issuer:        server.GetHost(c),
	}

	grpcRegisterResponse, err := a.grpcAuthClient.Register(
		c.Request.Context(),
		grpcRegisterRequest)
	if err != nil {
		// TODO: check error from gRPC server and return invalid credentials error

		server.ErrorResponse[RegisterOutput](c, err.Error())
		return
	}

	response := &RegisterOutput{
		AccessToken:  grpcRegisterResponse.GetAccessToken(),
		RefreshToken: grpcRegisterResponse.GetRefreshToken(),
	}

	server.SuccessResponse(c, response)
}

// Refresh tokens.
//
//	@Summary			Refresh tokens
//	@Description		Refresh tokens
//	@Tags				Auth
//	@Id 				refreshTokens
//	@Accept				json
//	@Produce			json
//	@Param        		X-Fingerprint	  header    string    true   	"User browser fingerprint."
//	@Param				input body RefreshTokensInput true "Input parameters for refresh tokens."
//	@Success			200	{object}	server.dataResponse[RefreshTokensOutput]
//	@Failure			500	{object}	server.dataResponse[RefreshTokensOutput]
//	@Router				/api/v1/auth/refresh-tokens [post]
func (a *Routes) refreshTokens(c *gin.Context) {
	inp, err := server.GetInputFromBody[RefreshTokensInput](c)
	if err != nil {
		server.ErrorResponse[RefreshTokensOutput](c, err.Error())
		return
	}

	grpcRefreshTokensRequest := &ssopb.RefreshTokensRequest{
		RefreshToken: inp.RefreshToken,
		ClientCode:   inp.ClientCode,
		UserAgent:    server.GetUserAgent(c),
		Fingerprint:  server.GetFingerprint(c),
		Issuer:       server.GetHost(c),
	}

	grpcRefreshTokensResponse, err := a.grpcAuthClient.RefreshTokens(
		c.Request.Context(),
		grpcRefreshTokensRequest)
	if err != nil {
		// TODO: check error from gRPC server and return invalid credentials error

		server.ErrorResponse[RefreshTokensOutput](c, err.Error())
		return
	}

	response := &RefreshTokensOutput{
		AccessToken:  grpcRefreshTokensResponse.GetAccessToken(),
		RefreshToken: grpcRefreshTokensResponse.GetRefreshToken(),
	}

	server.SuccessResponse(c, response)
}
