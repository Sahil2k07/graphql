package services

import (
	"context"
	"time"

	"github.com/Sahil2k07/graphql/internal/configs"
	"github.com/Sahil2k07/graphql/internal/enums"
	"github.com/Sahil2k07/graphql/internal/interfaces"
	"github.com/Sahil2k07/graphql/internal/utils"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"golang.org/x/crypto/bcrypt"
)

type cryptoService struct {
	signingKey    []byte
	encryptionKey []byte
}

func NewCryptoService() interfaces.CryptoService {
	cfg := configs.GetJWTConfig()
	signingKey := []byte(cfg.SigningKey)
	encryptionKey := []byte(cfg.EncryptionKey)

	if len(encryptionKey) != 32 {
		panic("encryptionKey must be exactly 32 bytes for A256GCM")
	}

	return &cryptoService{
		signingKey:    signingKey,
		encryptionKey: encryptionKey,
	}
}

func (s *cryptoService) GenerateJWT(ctx context.Context, claims *utils.UserClaims, ttl time.Duration) (string, error) {
	tok := jwt.New()

	_ = tok.Set("id", claims.ID)
	_ = tok.Set("email", claims.Email)
	_ = tok.Set("role", claims.Role)
	_ = tok.Set("username", claims.UserName)

	now := time.Now().UTC()
	_ = tok.Set(jwt.IssuedAtKey, now)
	_ = tok.Set(jwt.ExpirationKey, now.Add(ttl))

	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.HS256, s.signingKey))
	if err != nil {
		return "", err
	}

	encrypted, err := jwe.Encrypt(
		signed,
		jwe.WithKey(jwa.DIRECT, s.encryptionKey),
		jwe.WithContentEncryption(jwa.A256GCM),
	)
	if err != nil {
		return "", err
	}

	return string(encrypted), nil
}

func (s *cryptoService) DecryptAndVerifyJWT(ctx context.Context, tokenStr string) (*utils.UserClaims, error) {
	decrypted, err := jwe.Decrypt([]byte(tokenStr), jwe.WithKey(jwa.DIRECT, s.encryptionKey))
	if err != nil {
		return nil, err
	}

	tok, err := jwt.Parse(decrypted, jwt.WithKey(jwa.HS256, s.signingKey))
	if err != nil {
		return nil, err
	}

	if err := jwt.Validate(tok); err != nil {
		return nil, err
	}

	claims := &utils.UserClaims{}

	if id, ok := tok.Get("id"); ok {
		switch v := id.(type) {
		case float64:
			claims.ID = uint(v)
		case int:
			claims.ID = uint(v)
		case uint:
			claims.ID = v
		}
	}

	if email, ok := tok.Get("email"); ok {
		if str, ok := email.(string); ok {
			claims.Email = str
		}
	}

	if role, ok := tok.Get("role"); ok {
		if str, ok := role.(string); ok {
			claims.Role = enums.Role(str)
		}
	}

	if username, ok := tok.Get("username"); ok {
		if str, ok := username.(string); ok {
			claims.UserName = str
		}
	}

	if exp, ok := tok.Get(jwt.ExpirationKey); ok {
		if t, ok := exp.(time.Time); ok {
			claims.ExpiresAt = &t
		}
	}

	return claims, nil
}

func (s *cryptoService) HashPassword(plain string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func (s *cryptoService) VerifyPassword(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
