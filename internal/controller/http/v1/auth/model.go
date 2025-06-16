package auth

// LoginInput is input model of user login request.
type LoginInput struct {
	Username   string `form:"username"`    // Username
	Password   string `form:"password"`    // Password
	ClientCode string `form:"client_code"` // Client code
} // @name LoginInput

// LoginOutput is output model of user login request.
type LoginOutput struct {
	AccessToken  string `json:"access_token" binding:"required"`  // Access token
	RefreshToken string `json:"refresh_token" binding:"required"` // Refresh token
} // @name LoginOutput
