package command

import (
	"context"
	"log/slog"

	"github.com/nidea1/go-gavel/pkg/common"
	"github.com/nidea1/go-gavel/services/auth/internal/domain"
	"github.com/nidea1/go-gavel/services/auth/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUseCase struct {
	userRepo repository.UserRepository
}

func NewRegisterUseCase(userRepo repository.UserRepository) *RegisterUseCase {
	return &RegisterUseCase{userRepo: userRepo}
}

func (uc *RegisterUseCase) Register(ctx context.Context, firstName, lastName, email, password string) (bool, string, *common.ValidationErrors, error) {
	if len(password) < 8 {
		slog.Error("Password must be at least 8 characters long")
		valErrs := &common.ValidationErrors{
			Errors: []common.SingleValidationError{
				{Field: "password", Code: "min_length", Message: "Password must be at least 8 characters long"},
			},
		}
		return false, "Validation Error", valErrs, nil
	}

	if existingUser, _ := uc.userRepo.GetUserByEmail(ctx, email); existingUser != nil {
		slog.Error("Email is already in use")
		valErrs := &common.ValidationErrors{
			Errors: []common.SingleValidationError{
				{Field: "email", Code: "unique", Message: "Email is already in use"},
			},
		}
		return false, "Validation Error", valErrs, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error(err.Error())
		return false, "Internal Server Error", nil, err
	}

	user := domain.NewUser(firstName, lastName, email, string(hashedPassword), false)
	if err := user.Validate(); err != nil {
		slog.Error(err.Error())
		return false, "Validation Error", nil, err
	}

	if err := uc.userRepo.CreateUser(ctx, user); err != nil {
		slog.Error(err.Error())
		return false, "Internal Server Error", nil, err
	}

	// Will add kafka producer here to send register event

	return true, "Success", nil, nil
}
