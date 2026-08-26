package response

type errorResponse struct {
	Code        string `json:"error,omitempty"`
	Description string `json:"error_description,omitempty"`
} //@name ErrorResponse

func newErrorResponse(code, description string) errorResponse {
	return errorResponse{
		Code:        code,
		Description: description,
	}
}
