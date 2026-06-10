package auth

// LoginOutput is output model of login request.
type LoginOutput struct {
	RedirectURI string `json:"redirect_uri"`
}

// RegisterOutput is output model of register request.
type RegisterOutput struct {
	RedirectURI string `json:"redirect_uri"`
}

// ConsentOutput is output model of confirm consent request.
type ConsentOutput struct {
	RedirectURI string `json:"redirect_uri"`
}
