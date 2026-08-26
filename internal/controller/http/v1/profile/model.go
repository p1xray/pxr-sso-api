package profile

import (
	"time"
)

// GenderEnum is type for gender enum.
type GenderEnum int16 // @name GenderEnum

// Gender enum.
const (
	MALE   GenderEnum = 1 // Male
	FEMALE GenderEnum = 2 // Female
)

// ProfileOutput is output model of user profile request.
type ProfileOutput struct {
	UserID        int64       `json:"user_id"`         // User ID.
	Username      string      `json:"username"`        // Username.
	Fio           string      `json:"fio"`             // User full name.
	DateOfBirth   *time.Time  `json:"date_of_birth"`   // User date of birth.
	Gender        *GenderEnum `json:"gender"`          // User gender.
	AvatarFileKey *string     `json:"avatar_file_key"` // User avatar file key.
} // @name ProfileOutput

// EditProfileInput is input model of editing user profile data request.
type EditProfileInput struct {
	FullName      string      `json:"full_name"`       // User full name.
	DateOfBirth   *time.Time  `json:"date_of_birth"`   // User date of birth.
	Gender        *GenderEnum `json:"gender"`          // User gender.
	AvatarFileKey *string     `json:"avatar_file_key"` // User avatar file key.
} // @name EditProfileInput
