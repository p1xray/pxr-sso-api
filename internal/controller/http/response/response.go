package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pxr-sso-api/internal/constants"
)

func InternalServerError(c *gin.Context, code, description, uri string) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		newErrorResponse(code, description, uri))
}

func BadRequest(c *gin.Context, code, description, uri string) {
	c.AbortWithStatusJSON(
		http.StatusBadRequest,
		newErrorResponse(code, description, uri))
}

func Unauthorized(c *gin.Context, code, description, uri string) {
	c.AbortWithStatusJSON(
		http.StatusUnauthorized,
		newErrorResponse(code, description, uri))
}

func Redirect(c *gin.Context, uri string) {
	c.Redirect(http.StatusFound, uri)
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func TryRedirect(c *gin.Context, res ServiceResponse[string]) {
	if res.Data() == "" {
		if res.ErrorCode() == constants.ErrorCodeInternalServerError {
			InternalServerError(c, res.ErrorCode(), res.ErrorDescription(), res.ErrorURI())
			return
		}

		BadRequest(c, res.ErrorCode(), res.ErrorDescription(), res.ErrorURI())
		return
	}

	Redirect(c, res.Data())
}

func TrySuccess[T any](c *gin.Context, res ServiceResponse[T]) {
	if !res.IsSuccess() {
		if res.ErrorCode() == constants.ErrorCodeInternalServerError {
			InternalServerError(c, res.ErrorCode(), res.ErrorDescription(), res.ErrorURI())
			return
		}

		BadRequest(c, res.ErrorCode(), res.ErrorDescription(), res.ErrorURI())
		return
	}

	Success(c, res.Data())
}
