package configs

import "os"

type jwtConfig struct {
	SigningKey    string `toml:"signing_key"`
	EncryptionKey string `toml:"encryption_key"`
}

func loadProdJwtConfig() jwtConfig {
	return jwtConfig{
		SigningKey:    os.Getenv("JWT_SIGNING_KEY"),
		EncryptionKey: os.Getenv("JWT_ENCRYPTION_KEY"),
	}
}

func GetJWTConfig() jwtConfig {
	return jwtConfig{
		SigningKey:    globalConfig.JWT.SigningKey,
		EncryptionKey: globalConfig.JWT.EncryptionKey,
	}
}
