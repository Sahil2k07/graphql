package middlewares

import (
	"context"
	"net/http"

	"github.com/Sahil2k07/graphql/internal/configs"
	"github.com/Sahil2k07/graphql/internal/services"
	"github.com/labstack/echo/v4"
)

type contextKey string

const UserKey contextKey = "user"

func JWTContext() echo.MiddlewareFunc {
	jwtConfig := configs.GetJWTConfig()
	crypto := services.NewCryptoService()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenCookie, err := c.Cookie(jwtConfig.CookieName)
			if err != nil {
				// No token — continue as anonymous
				return next(c)
			}

			tokenStr := tokenCookie.Value

			claims, err := crypto.DecryptAndVerifyJWT(c.Request().Context(), tokenStr)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
			}

			// Store claims in context for later use
			ctx := context.WithValue(c.Request().Context(), UserKey, claims)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
