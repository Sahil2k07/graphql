package web

import (
	"github.com/Sahil2k07/graphql/internal/graphql/resolvers"
	"github.com/Sahil2k07/graphql/internal/repositories"
	"github.com/Sahil2k07/graphql/internal/services"
)

// Resolvers wires all dependencies for gqlgen resolvers.
func Resolvers() *resolvers.Resolver {
	authRepo := repositories.NewAuthRepository()
	authService := services.NewAuthService(authRepo)

	return &resolvers.Resolver{
		AuthService: authService,
	}
}
