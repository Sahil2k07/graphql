package utils

import (
	"context"
	"errors"

	errz "github.com/Sahil2k07/graphql/internal/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// HandleGraphQLError logs and converts Go errors to GraphQL errors.
func HandleGraphQLError(ctx context.Context, err error) *gqlerror.Error {
	if err == nil {
		return nil
	}

	LogGraphQLError(err)

	switch e := err.(type) {
	case *errz.NotFoundError:
		return &gqlerror.Error{
			Message: e.Msg,
			Extensions: map[string]any{
				"code": "NOT_FOUND",
			},
		}

	case *errz.ValidationError:
		return &gqlerror.Error{
			Message: e.Msg,
			Extensions: map[string]any{
				"code": "VALIDATION_FAILED",
			},
		}

	case *errz.UnauthorizedError:
		return &gqlerror.Error{
			Message: e.Msg,
			Extensions: map[string]any{
				"code": "UNAUTHORIZED",
			},
		}

	case *errz.ForbiddenError:
		return &gqlerror.Error{
			Message: e.Msg,
			Extensions: map[string]any{
				"code": "FORBIDDEN",
			},
		}

	case *errz.AlreadyExistsError:
		return &gqlerror.Error{
			Message: e.Msg,
			Extensions: map[string]any{
				"code": "ALREADY_EXISTS",
			},
		}

	case *errz.InternalError:
		return &gqlerror.Error{
			Message: "Internal server error",
			Extensions: map[string]any{
				"code": "INTERNAL_ERROR",
			},
		}
	}

	if errors.Is(err, context.Canceled) {
		return &gqlerror.Error{
			Message: "Request cancelled",
			Extensions: map[string]any{
				"code": "CANCELLED",
			},
		}
	}

	return &gqlerror.Error{
		Message: "Unexpected error occurred",
		Extensions: map[string]any{
			"code": "UNKNOWN",
		},
	}
}
