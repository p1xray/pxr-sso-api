package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jwtmiddleware "github.com/p1xray/pxr-sso/pkg/jwt"
	jwtclaims "github.com/p1xray/pxr-sso/pkg/jwt/claims"
	"strconv"
)

func GetUserID(c *gin.Context) (int64, error) {
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

func getTokenClaims(c *gin.Context) (jwtclaims.ValidatedClaims, error) {
	ctx := c.Request.Context()
	claims, ok := ctx.Value(jwtmiddleware.ContextKey{}).(jwtclaims.ValidatedClaims)
	if !ok {
		return jwtclaims.ValidatedClaims{}, ErrGetTokenClaims
	}

	return claims, nil
}
