package auth

import (
	"mime/multipart"
	"time"
)

// GenderEnum is type for gender enum.
type GenderEnum int16 // @name GenderEnum

// Gender enum.
const (
	MALE   GenderEnum = 1 // Male
	FEMALE GenderEnum = 2 // Female
)

// LoginInput is input model of user login request.
type LoginInput struct {
	Username   string `form:"username" binding:"required"`    // Username.
	Password   string `form:"password" binding:"required"`    // Password.
	ClientCode string `form:"client_code" binding:"required"` // Client code.
} // @name LoginInput

// LoginOutput is output model of user login request.
type LoginOutput struct {
	AccessToken  string `json:"access_token" binding:"required"`  // Access token.
	RefreshToken string `json:"refresh_token" binding:"required"` // Refresh token.
} // @name LoginOutput

// RegisterInput is input model of user register request.
type RegisterInput struct {
	Username    string                `form:"username" binding:"required"`            // Username.
	Password    string                `form:"password" binding:"required"`            // Password.
	ClientCode  string                `form:"client_code" binding:"required"`         // Client code.
	Fio         string                `form:"fio" binding:"required"`                 // Full name.
	DateOfBirth *time.Time            `form:"date_of_birth" time_format:"2006-01-02"` // Date of birth.
	Gender      *GenderEnum           `form:"gender"`                                 // Gender.
	AvatarFile  *multipart.FileHeader `form:"avatar_file" format:"file"`              // Avatar file.
} // @name RegisterInput

// RegisterOutput is output model of user register request.
type RegisterOutput struct {
	AccessToken  string `json:"access_token" binding:"required"`  // Access token.
	RefreshToken string `json:"refresh_token" binding:"required"` // Refresh token.
} // @name RegisterOutput
