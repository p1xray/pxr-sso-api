package auth

// LoginOutput is output model of login request.
type LoginOutput struct {
	RedirectURI string `json:"redirect_uri"`
} //@name LoginOutput

// RegisterOutput is output model of register request.
type RegisterOutput struct {
	RedirectURI string `json:"redirect_uri"`
} //@name RegisterOutput

// ConsentCardOutput is output model of consent card request.
type ConsentCardOutput struct {
	Scopes []ScopeOutput `json:"scopes"`
} //@name ConsentCardOutput

// ScopeOutput is output model of scope in consent card.
type ScopeOutput struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsGranted   bool   `json:"is_granted"`
} //@name ScopeOutput

// ConsentOutput is output model of confirm consent request.
type ConsentOutput struct {
	RedirectURI string `json:"redirect_uri"`
} //@name ConsentOutput
