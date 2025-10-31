package utils

import (
	"context"
	"time"

	"github.com/Sahil2k07/graphql/internal/enums"
	"github.com/Sahil2k07/graphql/internal/errors"
)

type UserClaims struct {
	ID        uint
	Email     string
	Role      enums.Role
	UserName  string
	ExpiresAt *time.Time
}

func GetUserClaims(ctx context.Context) (*UserClaims, error) {
	raw := ctx.Value("user")
	if raw == nil {
		return nil, errors.NewUnauthorized("no user claims in context")
	}

	claims, ok := raw.(*UserClaims)
	if !ok {
		return nil, errors.NewUnauthorized("invalid claims type in context")
	}

	return claims, nil
}
