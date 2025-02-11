package command

import (
	"context"
	"errors"
	"log/slog"

	"github.com/nidea1/go-gavel/services/auth/internal/domain"
	"github.com/nidea1/go-gavel/services/auth/internal/repository"
	"github.com/nidea1/go-gavel/services/auth/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUseCase struct {
	userRepo   repository.UserRepository
	JWTManager *jwt.JWTManager
}

func NewRegisterUseCase(userRepo repository.UserRepository, JWTManager *jwt.JWTManager) *RegisterUseCase {
	return &RegisterUseCase{
		userRepo,
		JWTManager,
	}
}

func (uc *RegisterUseCase) Register(ctx context.Context, firstName, lastName, email, password string) (string, string, error) {
	if len(password) < 8 {
		slog.Error("Password must be at least 8 characters long")
		return "", "", errors.New("Password must be at least 8 characters long")
	}

	if existingUser, _ := uc.userRepo.GetUserByEmail(ctx, email); existingUser != nil {
		slog.Error("Email is already in use")
		return "", "", errors.New("Email is already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error(err.Error())
		return "", "", err
	}

	user := domain.NewUser(firstName, lastName, email, string(hashedPassword), false)
	if err := user.Validate(); err != nil {
		slog.Error(err.Error())
		return "", "", err
	}

	if err := uc.userRepo.CreateUser(ctx, user); err != nil {
		slog.Error(err.Error())
		return "", "", err
	}

	// Will add kafka producer here to send register event

	accessToken, err := uc.JWTManager.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		slog.Error(err.Error())
		return "", "", err
	}

	refreshToken, err := uc.JWTManager.GenerateRefreshToken(user.ID, user.Email)
	if err != nil {
		slog.Error(err.Error())
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
