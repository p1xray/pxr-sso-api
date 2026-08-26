package auth

// LoginInput is input model of user login request.
type LoginInput struct {
	RequestURI string `form:"request_uri" binding:"required"`
	Username   string `form:"username" binding:"required"`
	Password   string `form:"password" binding:"required"`
} //@name LoginInput

// RegisterInput is input model of user register request.
type RegisterInput struct {
	RequestURI string `form:"request_uri" binding:"required"`
	Username   string `form:"username" binding:"required"`
	Password   string `form:"password" binding:"required"`
	FullName   string `form:"full_name" binding:"required"`
} //@name RegisterInput

// ConsentCardInput is input model of consent card request.
type ConsentCardInput struct {
	RequestURI string `form:"request_uri" binding:"required"`
} //@name ConsentCardInput

// ConsentInput is input model of confirm consent request.
type ConsentInput struct {
	RequestURI string `form:"request_uri" binding:"required"`
	Scopes     string `form:"scopes" binding:"required"`
} //@name ConsentInput
