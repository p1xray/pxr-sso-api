package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	jwtmiddleware "github.com/p1xray/pxr-sso/pkg/jwt"
	"github.com/p1xray/pxr-sso/pkg/jwt/validator"
	"log"
	"net/http"
	"os"
)

var (
	keyFunc = func(ctx context.Context) ([]byte, error) {
		clientSecret := os.Getenv("CLIENT_SECRET")
		return []byte(clientSecret), nil
	}
)

func CheckJWT() gin.HandlerFunc {
	issuer := os.Getenv("ISSUER")
	audience := os.Getenv("AUDIENCE")

	jwtValidator, err := validator.New(
		keyFunc,
		issuer,
		[]string{audience},
	)
	if err != nil {
		log.Fatalf("error creating JWT validator: %v", err)
	}

	middleware := jwtmiddleware.New(jwtValidator.ValidateToken)

	return func(ctx *gin.Context) {
		encounteredError := true
		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			encounteredError = false
			ctx.Request = r
			ctx.Next()
		}

		middleware.ParseJWT(handler).ServeHTTP(ctx.Writer, ctx.Request)

		if encounteredError {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				map[string]string{"message": "JWT is invalid."},
			)
		}
	}
}
