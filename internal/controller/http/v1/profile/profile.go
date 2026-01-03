package profile

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/wrappers"
	ssoprofilepb "github.com/p1xray/pxr-sso-protos/gen/go/profile"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		profile.GET("", middleware.HasScope("profile.read"), r.profile)
		profile.GET(":id", middleware.HasScope("profile.read"), r.profileByID)
		profile.POST("edit", middleware.HasScope("profile.edit"), r.editProfile)
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

	profile, err := r.profileFromGRPC(c.Request.Context(), userID)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[ProfileOutput](c, err.Error())
		return
	}

	server.SuccessResponse(c, &profile)
}

// User profile by ID.
//
//	@Summary		User profile by ID
//	@Description	User profile by ID
//	@Tags			Profile
//	@Id 			profileByID
//	@Produce		json
//	@Security 		ApiKeyAuth
//	@Param			id	path  int  true  "User ID"
//	@Success		200	{object}  server.dataResponse[ProfileOutput]
//	@Failure		500	{object}  server.dataResponse[ProfileOutput]
//	@Router			/api/v1/profile/{id} [get]
func (r *Routes) profileByID(c *gin.Context) {
	userID, err := server.GetIdFromRoute(c)
	if err != nil {
		server.ErrorResponse[ProfileOutput](c, err.Error())
		return
	}

	profile, err := r.profileFromGRPC(c.Request.Context(), userID)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[ProfileOutput](c, err.Error())
		return
	}

	server.SuccessResponse(c, &profile)
}

// Edit user profile data.
//
//	@Summary		Edit user profile data
//	@Description	Edit user profile data
//	@Tags			Profile
//	@Id 			editProfile
//	@Accept			json
//	@Produce		json
//	@Security 		ApiKeyAuth
//	@Param			input body EditProfileInput true "Input parameters for editing user profile data."
//	@Success		200	{object}  server.dataResponse[bool]
//	@Failure		500	{object}  server.dataResponse[bool]
//	@Router			/api/v1/profile/edit [post]
func (r *Routes) editProfile(c *gin.Context) {
	userID, err := server.GetUserID(c)
	if err != nil {
		server.ErrorResponse[bool](c, err.Error())
		return
	}

	inp, err := server.GetInputFromBody[EditProfileInput](c)
	if err != nil {
		server.ErrorResponse[bool](c, err.Error())
		return
	}

	var dateOfBirthPb *timestamppb.Timestamp
	if inp.DateOfBirth != nil {
		dateOfBirthPb = timestamppb.New(*inp.DateOfBirth)
	}

	var genderPb ssoprofilepb.Gender
	if inp.Gender != nil {
		genderPb = ssoprofilepb.Gender(*inp.Gender)
	}

	var avatarFileKeyPb *wrappers.StringValue
	if inp.AvatarFileKey != nil {
		avatarFileKeyPb = &wrappers.StringValue{Value: *inp.AvatarFileKey}
	}

	request := &ssoprofilepb.EditProfileRequest{
		UserId:        userID,
		FullName:      inp.FullName,
		DateOfBirth:   dateOfBirthPb,
		Gender:        genderPb,
		AvatarFileKey: avatarFileKeyPb,
	}

	response, err := r.grpcProfileClient.EditProfile(c.Request.Context(), request)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[ProfileOutput](c, err.Error())
		return
	}

	success := response.GetSuccess()
	server.SuccessResponse(c, &success)
}

func (r *Routes) profileFromGRPC(ctx context.Context, userID int64) (ProfileOutput, error) {
	grpcProfileRequest := &ssoprofilepb.GetProfileRequest{UserId: userID}
	grpcProfileResponse, err := r.grpcProfileClient.GetProfile(ctx, grpcProfileRequest)
	if err != nil {
		return ProfileOutput{}, err
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

	output := ProfileOutput{
		UserID:        grpcProfileResponse.GetUserId(),
		Username:      grpcProfileResponse.GetUsername(),
		Fio:           grpcProfileResponse.GetFio(),
		DateOfBirth:   dateOfBirth,
		Gender:        gender,
		AvatarFileKey: avatarFileKey,
	}

	return output, nil
}
