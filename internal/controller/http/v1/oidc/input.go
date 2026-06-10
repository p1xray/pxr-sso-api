package oidc

// AuthorizeInput is input model of oauth authorize request.
type AuthorizeInput struct {
	ResponseType        []string `form:"response_type"`
	Prompt              []string `form:"prompt"`
	ClientID            []string `form:"client_id"`
	RedirectURI         []string `form:"redirect_uri"`
	CodeChallenge       []string `form:"code_challenge"`
	CodeChallengeMethod []string `form:"code_challenge_method"`
	State               []string `form:"state"`
	Audience            []string `form:"audience"`
	Scope               []string `form:"scope"`
} //@name AuthorizeInput

// TokenInput is output model of oauth exchange token request.
type TokenInput struct {
	GrantType    string `form:"grant_type" binding:"required"`
	ClientID     string `form:"client_id" binding:"required"`
	Code         string `form:"code" binding:"required"`
	RedirectURI  string `form:"redirect_uri" binding:"required"`
	CodeVerifier string `form:"code_verifier" binding:"required"`
} //@name TokenInput
