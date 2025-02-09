package repository

import (
	"context"
	"time"

	"github.com/nidea1/go-gavel/services/auth/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type TokenRepository interface {
	AddToBlacklist(ctx context.Context, token string, expiresAt time.Time) error
	IsBlacklisted(ctx context.Context, token string) (bool, error)
	SaveRefreshToken(ctx context.Context, userID uint64, refreshToken string, expiresAt time.Time) error
	GetRefreshToken(ctx context.Context, userID uint64) (string, error)
	DeleteRefreshToken(ctx context.Context, userID uint64) error
}
