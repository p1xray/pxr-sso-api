package profile

import (
	"github.com/gin-gonic/gin"
	ssoprofilepb "github.com/p1xray/pxr-sso-protos/gen/go/profile"
	"pxr-sso-api/internal/controller/http/middleware"
	"pxr-sso-api/internal/controller/http/v1/model"
	"pxr-sso-api/internal/server"
	"time"
)

// Routes provides routes for user profile.
type Routes struct {
	grpcProfileClient ssoprofilepb.SsoProfileClient
}

// InitRoutes initializes the routes for user profile.
func InitRoutes(api *gin.RouterGroup, grpcProfileClient ssoprofilepb.SsoProfileClient) {
	r := &Routes{grpcProfileClient: grpcProfileClient}

	profile := api.Group("/profile")
	profile.Use(middleware.CheckJWT())
	{
		profile.GET("", r.profile)
	}
}

// Current user profile.
//
//	@Summary		Current user profile
//	@Description	Current user profile
//	@Tags			Profile
//	@Id 			profile
//	@Produce		json
//	@Security 		ApiKeyAuth
//	@Success		200	{object}  server.dataResponse[ProfileOutput]
//	@Failure		500	{object}  server.dataResponse[ProfileOutput]
//	@Router			/api/v1/profile [get]
func (r *Routes) profile(c *gin.Context) {
	userID, err := server.GetUserID(c)
	if err != nil {
		server.ErrorResponse[ProfileOutput](c, err.Error())
		return
	}

	grpcProfileRequest := &ssoprofilepb.GetProfileRequest{UserId: userID}
	grpcProfileResponse, err := r.grpcProfileClient.GetProfile(c.Request.Context(), grpcProfileRequest)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[ProfileOutput](c, err.Error())
		return
	}

	var dateOfBirth *time.Time
	if grpcProfileResponse.GetDateOfBirth() != nil {
		dateOfBirthValue := grpcProfileResponse.GetDateOfBirth().AsTime()
		dateOfBirth = &dateOfBirthValue
	}

	var gender *model.GenderEnum
	if grpcProfileResponse.GetGender() != ssoprofilepb.Gender_GENDER_UNSPECIFIED {
		genderValue := model.GenderEnum(grpcProfileResponse.GetGender())
		gender = &genderValue
	}

	var avatarFileKey *string
	if grpcProfileResponse.GetAvatarFileKey() != nil {
		avatarFileKeyValue := grpcProfileResponse.GetAvatarFileKey().GetValue()
		avatarFileKey = &avatarFileKeyValue
	}

	output := &ProfileOutput{
		UserID:        grpcProfileResponse.GetUserId(),
		Username:      grpcProfileResponse.GetUsername(),
		Fio:           grpcProfileResponse.GetFio(),
		DateOfBirth:   dateOfBirth,
		Gender:        gender,
		AvatarFileKey: avatarFileKey,
	}
	server.SuccessResponse(c, output)
}
