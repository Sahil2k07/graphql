package configs

import "os"

type jwtConfig struct {
	CookieName    string `toml:"cookie_name"`
	SigningKey    string `toml:"signing_key"`
	EncryptionKey string `toml:"encryption_key"`
}

func loadProdJwtConfig() jwtConfig {
	return jwtConfig{
		CookieName:    os.Getenv("COOKIE_NAME"),
		SigningKey:    os.Getenv("JWT_SIGNING_KEY"),
		EncryptionKey: os.Getenv("JWT_ENCRYPTION_KEY"),
	}
}

func GetJWTConfig() jwtConfig {
	return jwtConfig{
		CookieName:    globalConfig.JWT.CookieName,
		SigningKey:    globalConfig.JWT.SigningKey,
		EncryptionKey: globalConfig.JWT.EncryptionKey,
	}
}
