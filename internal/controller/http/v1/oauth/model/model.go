package model

// AuthorizeInput is input model of oauth authorize request.
type AuthorizeInput struct {
	ResponseType        []string `form:"response_type"`
	ClientID            []string `form:"client_id"`
	RedirectURI         []string `form:"redirect_uri"`
	CodeChallenge       []string `form:"code_challenge"`
	CodeChallengeMethod []string `form:"code_challenge_method"`
	State               []string `form:"state"`
	Audience            []string `form:"audience"`
	Scope               []string `form:"scope"`
}

// AuthorizeOutput is output model of oauth authorize request.
type AuthorizeOutput struct {
	RedirectURI string `json:"redirect_uri"`
}

// LoginInput is input model of user login request.
type LoginInput struct {
	FlowID       string `form:"flow_id" binding:"required"`
	ResponseType string `form:"response_type" binding:"required"`
	ClientID     string `form:"client_id" binding:"required"`
	RedirectURI  string `form:"redirect_uri" binding:"required"`
	State        string `form:"state" binding:"required"`
	Scope        string `form:"scope"`
	Username     string `form:"username" binding:"required"`
	Password     string `form:"password" binding:"required"`
}

// LoginOutput is output model of login request.
type LoginOutput struct {
	RedirectURI string `json:"redirect_uri"`
}

// RegisterInput is input model of user register request.
type RegisterInput struct {
	FlowID       string `form:"flow_id" binding:"required"`
	ResponseType string `form:"response_type" binding:"required"`
	ClientID     string `form:"client_id" binding:"required"`
	RedirectURI  string `form:"redirect_uri" binding:"required"`
	State        string `form:"state" binding:"required"`
	Scope        string `form:"scope"`
	Username     string `form:"username" binding:"required"`
	Password     string `form:"password" binding:"required"`
	FullName     string `form:"full_name" binding:"required"`
}

// RegisterOutput is output model of register request.
type RegisterOutput struct {
	RedirectURI string `json:"redirect_uri"`
}

// ConsentInput is output model of confirm consent request.
type ConsentInput struct {
	FlowID       string `form:"flow_id" binding:"required"`
	ResponseType string `form:"response_type" binding:"required"`
	ClientID     string `form:"client_id" binding:"required"`
	RedirectURI  string `form:"redirect_uri" binding:"required"`
	State        string `form:"state" binding:"required"`
	Scope        string `form:"scope"`
}

// ConsentOutput is output model of confirm consent request.
type ConsentOutput struct {
	RedirectURI string `json:"redirect_uri"`
}

// TokenInput is output model of oauth exchange token request.
type TokenInput struct {
	GrantType    string `form:"grant_type" binding:"required"`
	ClientID     string `form:"client_id" binding:"required"`
	Code         string `form:"code" binding:"required"`
	RedirectURI  string `form:"redirect_uri" binding:"required"`
	CodeVerifier string `form:"code_verifier" binding:"required"`
}

// TokenOutput is output model of oauth exchange token request.
type TokenOutput struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
}
