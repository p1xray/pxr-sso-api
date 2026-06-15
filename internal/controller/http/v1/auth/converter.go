package auth

import (
	authpb "github.com/p1xray/pxr-sso-protos/gen/go/auth"
	authsessionpb "github.com/p1xray/pxr-sso-protos/gen/go/session"
	"pxr-sso-api/internal/controller/http/request"
)

func toLoginRequest(input LoginInput) *authpb.LoginRequest {
	loginRequest := &authpb.LoginRequest{
		RequestUri: input.RequestURI,
		Username:   input.Username,
		Password:   input.Password,
	}

	return loginRequest
}

func toLoginOutput(response *authpb.LoginResponse) LoginOutput {
	loginOutput := LoginOutput{
		RedirectURI: response.GetRedirectUri(),
	}

	return loginOutput
}

func toRegisterRequest(input RegisterInput) *authpb.RegisterRequest {
	registerRequest := &authpb.RegisterRequest{
		RequestUri: input.RequestURI,
		Username:   input.Username,
		Password:   input.Password,
		FullName:   input.FullName,
	}

	return registerRequest
}

func toRegisterOutput(response *authpb.RegisterResponse) RegisterOutput {
	registerOutput := RegisterOutput{
		RedirectURI: response.GetRedirectUri(),
	}

	return registerOutput
}

func toConsentRequest(input ConsentInput, sessionCookies []request.SessionCookie) *authpb.ConsentRequest {
	sessions := make([]*authsessionpb.Cookie, 0)
	for _, cookie := range sessionCookies {
		session := &authsessionpb.Cookie{Name: cookie.Name, Value: cookie.Value}
		sessions = append(sessions, session)
	}

	consentRequest := &authpb.ConsentRequest{
		RequestUri: input.RequestURI,
		Scopes:     input.Scopes,
		Sessions:   sessions,
	}

	return consentRequest
}

func toConsentOutput(response *authpb.ConsentResponse) ConsentOutput {
	consentOutput := ConsentOutput{
		RedirectURI: response.GetRedirectUri(),
	}

	return consentOutput
}
