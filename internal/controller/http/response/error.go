package response

type errorResponse struct {
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	ErrorURI         string `json:"error_uri,omitempty"`
}

func newErrorResponse(code, description, uri string) errorResponse {
	return errorResponse{
		Error:            code,
		ErrorDescription: description,
		ErrorURI:         uri,
	}
}
