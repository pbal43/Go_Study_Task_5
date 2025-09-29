package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"

	"net/http"
	"time"
	authErrors "toDoList/internal/server/auth/auth_errors"
	auth "toDoList/internal/server/auth/user_auth"
)

func AuthMiddleware(signer auth.HS256Signer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := ctx.Cookie("access_token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": authErrors.ErrorMissingAccessToken})
			ctx.Abort() // чтобы дальше на след функцию хендлера не ушло
			return
		}

		claims, err := signer.ParseAccessToken(accessToken, auth.ParseOptions{
			ExpectedIssuer:   signer.Issuer,
			ExpectedAudience: signer.Audience,
			AllowMethods:     []string{"HS256"},
			Leeway:           60 * time.Second,
		})

		if err != nil {
			if errors.Is(err, authErrors.ErrorInvalidAccessToken) {
				refreshToken, err := ctx.Cookie("refresh_token")
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, gin.H{"error": authErrors.ErrorMissingRefreshToken})
					ctx.Abort()
					return
				}

				refreshClaims, err := signer.ParseRefreshToken(refreshToken, auth.ParseOptions{
					ExpectedIssuer:   signer.Issuer,
					ExpectedAudience: signer.Audience,
					AllowMethods:     []string{"HS256"},
					Leeway:           5 * time.Minute,
				})
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
					ctx.Abort()
					return
				}

				newAccessToken, err := signer.NewAccessToken(refreshClaims.Subject)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					ctx.Abort()
					return
				}

				ctx.SetCookie("access_token", newAccessToken, 3600*24, "/", "127.0.0.1:8080", false, true)

				claims, err = signer.ParseAccessToken(newAccessToken, auth.ParseOptions{
					ExpectedIssuer:   signer.Issuer,
					ExpectedAudience: signer.Audience,
					AllowMethods:     []string{"HS256"},
					Leeway:           60 * time.Second,
				})
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": authErrors.ErrorFailToParseNewAccessToken})
					ctx.Abort()
					return
				}

			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				ctx.Abort()
				return
			}
		}
		
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
