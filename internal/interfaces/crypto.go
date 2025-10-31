package interfaces

import (
	"context"
	"time"

	"github.com/Sahil2k07/graphql/internal/utils"
)

type CryptoService interface {
	GenerateJWT(ctx context.Context, claims *utils.UserClaims, ttl time.Duration) (string, error)
	DecryptAndVerifyJWT(ctx context.Context, tokenStr string) (*utils.UserClaims, error)
	HashPassword(plain string) (string, error)
	VerifyPassword(hashed, plain string) bool
}
