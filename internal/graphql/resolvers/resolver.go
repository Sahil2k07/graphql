package resolvers

import "github.com/Sahil2k07/graphql/internal/interfaces"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthService interfaces.AuthService
	TodoService interfaces.TodoService
}
