package interfaces

import (
	"context"

	"github.com/Sahil2k07/graphql/internal/graphql/generated"
	"github.com/Sahil2k07/graphql/internal/models"
)

type AuthRepository interface {
	// Check if a user already exists by email
	CheckUserExist(email string) (bool, error)

	// Retrieve a user by email (with profile)
	GetUser(email string) (models.User, error)

	// Add a new user to the database
	AddUser(user models.User) error

	// Update the user’s password
	UpdatePassword(email, newPassword string) error
}

type AuthService interface {
	Signup(ctx context.Context, input generated.SignupInput) (string, error)
	Signin(ctx context.Context, input generated.SigninInput) (*generated.AuthPayload, error)
	ForgotPassword(ctx context.Context, input generated.ForgotPasswordInput) (string, error)
	ResetPassword(ctx context.Context, input generated.ResetPasswordInput) (string, error)
}
