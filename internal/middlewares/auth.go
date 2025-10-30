package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/Sahil2k07/graphql/internal/configs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type contextKey string

const UserKey contextKey = "user"

func JWTContext() echo.MiddlewareFunc {
	jwtConfig := configs.GetJWTConfig()

	secret := []byte(jwtConfig.Secret)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenCookie, err := c.Cookie(jwtConfig.CookieName)
			if err != nil {
				return next(c)
			}
			tokenString := tokenCookie.Value

			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				return secret, nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
			}

			ctx := context.WithValue(c.Request().Context(), UserKey, token.Claims)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
