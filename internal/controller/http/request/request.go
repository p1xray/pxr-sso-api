package request

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	jwtmiddleware "github.com/p1xray/pxr-sso/pkg/jwt"
	jwtclaims "github.com/p1xray/pxr-sso/pkg/jwt/claims"
	"pxr-sso-api/internal/constants"
	"strconv"
	"strings"
)

var (
	ErrInvalidInputQuery     = errors.New("invalid input query")
	ErrInvalidInputForm      = errors.New("invalid input form")
	ErrGetTokenClaims        = errors.New("error getting token claims from request")
	ErrConvertStringToNumber = errors.New("error converting string value to number")
)

type SessionCookie struct {
	Name  string
	Value string
}

func FromQuery[T any](c *gin.Context) (T, error) {
	var inp T
	if err := c.BindQuery(&inp); err != nil {
		return inp, fmt.Errorf("%w: %w", ErrInvalidInputQuery, err)
	}

	return inp, nil
}

func FromForm[T any](c *gin.Context) (T, error) {
	var inp T
	if err := c.Bind(&inp); err != nil {
		return inp, fmt.Errorf("%w: %w", ErrInvalidInputForm, err)
	}

	return inp, nil
}

func SessionFromCookie(c *gin.Context) []SessionCookie {
	sessions := make([]SessionCookie, 0)
	for _, cookie := range c.Request.Cookies() {
		if strings.HasPrefix(cookie.Name, constants.SessionCookieNamePrefix) {
			session := SessionCookie{Name: cookie.Name, Value: cookie.Value}
			sessions = append(sessions, session)
		}
	}

	return sessions
}

func SubjectFromToken(c *gin.Context) (int64, error) {
	claims, err := getTokenClaims(c)
	if err != nil {
		return 0, err
	}

	userID, err := strconv.ParseInt(claims.RegisteredClaims.Subject, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrConvertStringToNumber, err)
	}

	return userID, nil
}

func UserHasScope(c *gin.Context, expectedScope string) (bool, error) {
	claims, err := getTokenClaims(c)
	if err != nil {
		return false, err
	}

	scopes := strings.Split(claims.RegisteredClaims.Scope, " ")
	for _, scope := range scopes {
		if scope == expectedScope {
			return true, nil
		}
	}

	return false, nil
}

func getTokenClaims(c *gin.Context) (jwtclaims.ValidatedClaims, error) {
	ctx := c.Request.Context()
	claims, ok := ctx.Value(jwtmiddleware.ContextKey{}).(jwtclaims.ValidatedClaims)
	if !ok {
		return jwtclaims.ValidatedClaims{}, ErrGetTokenClaims
	}

	return claims, nil
}
