package oidc

import (
	oidcpb "github.com/p1xray/pxr-sso-protos/gen/go/oidc"
	authsessionpb "github.com/p1xray/pxr-sso-protos/gen/go/session"
	"pxr-sso-api/internal/controller/http/request"
)

func toAuthorizeRequest(input AuthorizeInput, audience string, sessionCookies []request.SessionCookie) *oidcpb.AuthorizeRequest {
	sessions := make([]*authsessionpb.Cookie, 0)
	for _, cookie := range sessionCookies {
		session := &authsessionpb.Cookie{Name: cookie.Name, Value: cookie.Value}
		sessions = append(sessions, session)
	}

	authorizeRequest := &oidcpb.AuthorizeRequest{
		ResponseType:        input.ResponseType,
		Prompt:              input.Prompt,
		ClientId:            input.ClientID,
		RedirectUri:         input.RedirectURI,
		CodeChallenge:       input.CodeChallenge,
		CodeChallengeMethod: input.CodeChallengeMethod,
		State:               input.State,
		Audience:            []string{audience},
		Scope:               input.Scope,
		Sessions:            sessions,
	}

	return authorizeRequest
}

func toTokenRequest(input TokenInput) *oidcpb.TokenRequest {
	tokenRequest := &oidcpb.TokenRequest{
		GrantType:    input.GrantType,
		ClientId:     input.ClientID,
		Code:         input.Code,
		RedirectUri:  input.RedirectURI,
		CodeVerifier: input.CodeVerifier,
	}

	return tokenRequest
}

func toTokenOutput(response *oidcpb.TokenResponse) TokenOutput {
	tokenOutput := TokenOutput{
		AccessToken:  response.GetAccessToken(),
		TokenType:    response.GetTokenType(),
		ExpiresIn:    response.GetExpiresIn(),
		RefreshToken: response.GetRefreshToken(),
		IDToken:      response.GetIdToken(),
	}

	return tokenOutput
}
