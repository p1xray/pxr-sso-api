package converter

import (
	oauthpb "github.com/p1xray/pxr-sso-protos/gen/go/oauth"
	"pxr-sso-api/internal/constants"
	"pxr-sso-api/internal/controller/http/response"
	"pxr-sso-api/internal/controller/http/v1/oauth/model"
)

// TODO: wrap redirect URI in output model for each request
func ToAuthorizeServiceResponse(data *oauthpb.AuthorizeResponse, err error) response.ServiceResponse[string] {
	if data == nil && err == nil {
		return response.NewServiceResponse(
			"",
			constants.ErrorCodeInternalServerError,
			constants.ErrorDescriptionInternalServerError,
			"")
	}

	if data != nil {
		responseErr := data.GetError()
		if data.GetRedirectUri() == "" && responseErr == nil {
			return response.NewServiceResponse(
				"",
				constants.ErrorCodeInternalServerError,
				constants.ErrorDescriptionEmptyRedirectURI,
				"")
		}

		return response.NewServiceResponse(
			data.GetRedirectUri(),
			responseErr.GetCode(),
			responseErr.GetDescription(),
			responseErr.GetUri())
	}

	return response.NewServiceResponse("", constants.ErrorCodeInternalServerError, err.Error(), "")
}

func ToLoginServiceResponse(data *oauthpb.LoginResponse, err error) response.ServiceResponse[string] {
	if data == nil && err == nil {
		return response.NewServiceResponse(
			"",
			constants.ErrorCodeInternalServerError,
			constants.ErrorDescriptionInternalServerError,
			"")
	}

	if data != nil {
		responseErr := data.GetError()
		if data.GetRedirectUri() == "" && responseErr == nil {
			return response.NewServiceResponse(
				"",
				constants.ErrorCodeInternalServerError,
				constants.ErrorDescriptionEmptyRedirectURI,
				"")
		}

		return response.NewServiceResponse(
			data.GetRedirectUri(),
			responseErr.GetCode(),
			responseErr.GetDescription(),
			responseErr.GetUri())
	}

	return response.NewServiceResponse("", constants.ErrorCodeInternalServerError, err.Error(), "")
}

func ToRegisterServiceResponse(data *oauthpb.RegisterResponse, err error) response.ServiceResponse[string] {
	if data == nil && err == nil {
		return response.NewServiceResponse(
			"",
			constants.ErrorCodeInternalServerError,
			constants.ErrorDescriptionInternalServerError,
			"")
	}

	if data != nil {
		responseErr := data.GetError()
		if data.GetRedirectUri() == "" && responseErr == nil {
			return response.NewServiceResponse(
				"",
				constants.ErrorCodeInternalServerError,
				constants.ErrorDescriptionEmptyRedirectURI,
				"")
		}

		return response.NewServiceResponse(
			data.GetRedirectUri(),
			responseErr.GetCode(),
			responseErr.GetDescription(),
			responseErr.GetUri())
	}

	return response.NewServiceResponse("", constants.ErrorCodeInternalServerError, err.Error(), "")
}

func ToConsentServiceResponse(data *oauthpb.ConsentResponse, err error) response.ServiceResponse[string] {
	if data == nil && err == nil {
		return response.NewServiceResponse(
			"",
			constants.ErrorCodeInternalServerError,
			constants.ErrorDescriptionInternalServerError,
			"")
	}

	if data != nil {
		responseErr := data.GetError()
		if data.GetRedirectUri() == "" && responseErr == nil {
			return response.NewServiceResponse(
				"",
				constants.ErrorCodeInternalServerError,
				constants.ErrorDescriptionEmptyRedirectURI,
				"")
		}

		return response.NewServiceResponse(
			data.GetRedirectUri(),
			responseErr.GetCode(),
			responseErr.GetDescription(),
			responseErr.GetUri())
	}

	return response.NewServiceResponse("", constants.ErrorCodeInternalServerError, err.Error(), "")
}

func ToTokenServiceResponse(data *oauthpb.TokenResponse, err error) response.ServiceResponse[model.TokenOutput] {
	if (data == nil || data.GetError() == nil) && err == nil {
		return response.NewServiceResponse(
			model.TokenOutput{},
			constants.ErrorCodeInternalServerError,
			constants.ErrorDescriptionInternalServerError,
			"")
	}

	if data != nil {
		responseErr := data.GetError()
		tokens := model.TokenOutput{
			AccessToken:  data.GetAccessToken(),
			TokenType:    data.GetTokenType(),
			ExpiresIn:    data.GetExpiresIn(),
			RefreshToken: data.GetRefreshToken(),
			IDToken:      data.GetIdToken(),
		}

		return response.NewServiceResponse(
			tokens,
			responseErr.GetCode(),
			responseErr.GetDescription(),
			responseErr.GetUri())
	}

	return response.NewServiceResponse(model.TokenOutput{}, constants.ErrorCodeInternalServerError, err.Error(), "")
}
