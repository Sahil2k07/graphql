package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/Sahil2k07/graphql/internal/services"
	"github.com/Sahil2k07/graphql/internal/utils"
	"github.com/labstack/echo/v4"
)

func JWTContext() echo.MiddlewareFunc {
	crypto := services.NewCryptoService()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return next(c) // no token → unauthenticated but allowed
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid authorization header format"})
			}

			tokenStr := parts[1]

			claims, err := crypto.DecryptAndVerifyJWT(c.Request().Context(), tokenStr)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
			}

			ctx := context.WithValue(c.Request().Context(), utils.UserCtxKey, claims)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
