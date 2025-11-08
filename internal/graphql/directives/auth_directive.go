package directives

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Sahil2k07/graphql/internal/utils"
)

func AuthDirective() func(ctx context.Context, obj any, next graphql.Resolver) (any, error) {
	return func(ctx context.Context, obj any, next graphql.Resolver) (any, error) {
		fieldCtx := graphql.GetFieldContext(ctx)
		for _, d := range fieldCtx.Field.Definition.Directives {
			if d.Name == "public" {
				return next(ctx)
			}
		}

		user := ctx.Value(utils.UserCtxKey)
		if user == nil {
			return nil, errors.New("unauthorized: missing or invalid token")
		}

		return next(ctx)
	}
}
