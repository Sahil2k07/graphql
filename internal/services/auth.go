package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Sahil2k07/graphql/internal/enums"
	errz "github.com/Sahil2k07/graphql/internal/errors"
	"github.com/Sahil2k07/graphql/internal/graphql/generated"
	"github.com/Sahil2k07/graphql/internal/interfaces"
	"github.com/Sahil2k07/graphql/internal/models"
	"github.com/Sahil2k07/graphql/internal/utils"
)

type authService struct {
	repo   interfaces.AuthRepository
	crypto interfaces.CryptoService
}

func NewAuthService(repo interfaces.AuthRepository) interfaces.AuthService {
	crypto := NewCryptoService()
	return &authService{
		repo:   repo,
		crypto: crypto,
	}
}

func (s *authService) Signup(ctx context.Context, input generated.SignupInput) (string, error) {
	exists, err := s.repo.CheckUserExist(input.Email)
	if err != nil {
		return "", err
	}
	if exists {
		return "", errz.NewAlreadyExists("user already exists")
	}

	hashed, err := s.crypto.HashPassword(input.Password)
	if err != nil {
		return "", err
	}

	user := models.User{
		Email:    input.Email,
		UserName: input.UserName,
		Role:     enums.GUEST,
		Profile: models.Profile{
			FirstName: input.FirstName,
			LastName:  input.LastName,
		},
		Password: hashed,
	}

	err = s.repo.AddUser(user)
	if err != nil {
		return "", err
	}

	return "Signup successfull", nil
}

func (s *authService) Signin(ctx context.Context, input generated.SigninInput) (*generated.AuthPayload, error) {
	user, err := s.repo.GetUser(input.Email)
	if err != nil {
		return nil, errz.NewValidation("invalid credentials")
	}

	if !s.crypto.VerifyPassword(user.Password, input.Password) {
		return nil, errz.NewValidation("invalid credentials")
	}

	token, err := s.crypto.GenerateJWT(ctx, &utils.UserClaims{
		ID:       user.ID,
		Email:    user.Email,
		Role:     user.Role,
		UserName: user.UserName,
	}, 48*time.Hour)
	if err != nil {
		return nil, err
	}

	return &generated.AuthPayload{
		Token: token,
		User: &generated.User{
			ID:        fmt.Sprintf("%d", user.ID),
			Email:     user.Email,
			UserName:  user.UserName,
			Role:      string(user.Role),
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}

func (s *authService) ForgotPassword(ctx context.Context, input generated.ForgotPasswordInput) (string, error) {
	exists, err := s.repo.CheckUserExist(input.Email)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errz.NewNotFound("user not found")
	}
	return "Password reset link sent (mock)", nil
}

func (s *authService) ResetPassword(ctx context.Context, input generated.ResetPasswordInput) (string, error) {
	claims, err := utils.GetUserClaims(ctx)
	if err != nil {
		return "", errz.NewUnauthorized("unauthorized")
	}

	hashed, err := s.crypto.HashPassword(input.NewPassword)
	if err != nil {
		return "", err
	}

	err = s.repo.UpdatePassword(claims.Email, hashed)
	if err != nil {
		return "", err
	}

	return "Password updated successfully", nil
}
