package auth

// LoginOutput is output model of login request.
type LoginOutput struct {
	RedirectURI string `json:"redirect_uri"`
} //@name LoginOutput

// RegisterOutput is output model of register request.
type RegisterOutput struct {
	RedirectURI string `json:"redirect_uri"`
} //@name RegisterOutput

// ConsentOutput is output model of confirm consent request.
type ConsentOutput struct {
	RedirectURI string `json:"redirect_uri"`
} //@name ConsentOutput
