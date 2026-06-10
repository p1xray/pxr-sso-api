package auth

// LoginInput is input model of user login request.
type LoginInput struct {
	RequestURI string `json:"request_uri" binding:"required"`
	Username   string `form:"username" binding:"required"`
	Password   string `form:"password" binding:"required"`
}

// RegisterInput is input model of user register request.
type RegisterInput struct {
	RequestURI string `json:"request_uri" binding:"required"`
	Username   string `form:"username" binding:"required"`
	Password   string `form:"password" binding:"required"`
	FullName   string `form:"full_name" binding:"required"`
}

// ConsentInput is output model of confirm consent request.
type ConsentInput struct {
	RequestURI string   `json:"request_uri" binding:"required"`
	Scopes     []string `form:"scopes" binding:"required"`
}
