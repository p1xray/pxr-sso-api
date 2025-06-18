package profile

import (
	"pxr-sso-api/internal/controller/http/v1/model"
	"time"
)

// ProfileOutput is output model of user profile request.
type ProfileOutput struct {
	UserID        int64             `json:"user_id"`         // User ID.
	Username      string            `json:"username"`        // Username.
	Fio           string            `json:"fio"`             // User full name.
	DateOfBirth   *time.Time        `json:"date_of_birth"`   // User date of birth.
	Gender        *model.GenderEnum `json:"gender"`          // User gender.
	AvatarFileKey *string           `json:"avatar_file_key"` // User avatar file key.
} // @name ProfileOutput
