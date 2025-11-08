package middlewares

import (
	"context"
	"net/http"

	"github.com/Sahil2k07/graphql/internal/configs"
	"github.com/Sahil2k07/graphql/internal/services"
	"github.com/Sahil2k07/graphql/internal/utils"
	"github.com/labstack/echo/v4"
)

func JWTContext() echo.MiddlewareFunc {
	jwtConfig := configs.GetJWTConfig()
	crypto := services.NewCryptoService()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenCookie, err := c.Cookie(jwtConfig.CookieName)
			if err != nil {
				return next(c) // no cookie → unauthenticated but allowed
			}

			tokenStr := tokenCookie.Value

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
