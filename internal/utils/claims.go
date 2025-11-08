package utils

import (
	"context"
	"time"

	"github.com/Sahil2k07/graphql/internal/enums"
	errz "github.com/Sahil2k07/graphql/internal/errors"
)

type UserClaims struct {
	ID        uint
	Email     string
	Role      enums.Role
	UserName  string
	ExpiresAt *time.Time
}

type userCtxKey struct{}

var UserCtxKey = userCtxKey{}

func GetUserClaims(ctx context.Context) (*UserClaims, error) {
	raw := ctx.Value(UserCtxKey)
	if raw == nil {
		return nil, errz.NewUnauthorized("no user claims in context")
	}

	claims, ok := raw.(*UserClaims)
	if !ok {
		return nil, errz.NewUnauthorized("invalid claims type in context")
	}

	return claims, nil
}
