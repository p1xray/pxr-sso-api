package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	errorCodeInternalServerError = "server_error"
	errorCodeInvalidRequest      = "invalid_request"
	errorCodeUnauthorized        = "unauthorized"
)

func InternalServerError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		newErrorResponse(errorCodeInternalServerError, message))
}

func BadRequest(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusBadRequest,
		newErrorResponse(errorCodeInvalidRequest, message))
}

func Unauthorized(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusUnauthorized,
		newErrorResponse(errorCodeUnauthorized, message))
}

func Redirect(c *gin.Context, uri string) {
	c.Redirect(http.StatusFound, uri)
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func SetSessionCookie(c *gin.Context, name string, value string) {
	maxAge := 30 * time.Minute
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(name, value, int(time.Now().Add(maxAge).Unix()), "/", "", true, true)
}
