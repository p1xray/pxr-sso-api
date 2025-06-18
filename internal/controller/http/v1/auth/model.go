package auth

import (
	"mime/multipart"
	"pxr-sso-api/internal/controller/http/v1/model"
	"time"
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
	Gender      *model.GenderEnum     `form:"gender"`                                 // Gender.
	AvatarFile  *multipart.FileHeader `form:"avatar_file" format:"file"`              // Avatar file.
} // @name RegisterInput

// RegisterOutput is output model of user register request.
type RegisterOutput struct {
	AccessToken  string `json:"access_token" binding:"required"`  // Access token.
	RefreshToken string `json:"refresh_token" binding:"required"` // Refresh token.
} // @name RegisterOutput

// RefreshTokensInput is input model of refresh tokens request.
type RefreshTokensInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"` // Refresh token.
	ClientCode   string `json:"client_code" binding:"required"`   // Client code.
} // @name RefreshTokensInput

// RefreshTokensOutput is output model of refresh tokens request.
type RefreshTokensOutput struct {
	AccessToken  string `json:"access_token" binding:"required"`  // Access token.
	RefreshToken string `json:"refresh_token" binding:"required"` // Refresh token.
} // @name RefreshTokensOutput

// LogoutInput is input model of user logout request.
type LogoutInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"` // Refresh token.
	ClientCode   string `json:"client_code" binding:"required"`   // Client code.
} // @name LogoutInput
