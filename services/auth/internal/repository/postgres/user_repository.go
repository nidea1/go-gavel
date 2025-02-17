package postgres

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nidea1/go-gavel/services/auth/internal/domain"
	"github.com/nidea1/go-gavel/services/auth/internal/repository"
)

type PostgresUserRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresUserRepository(pool *pgxpool.Pool) repository.UserRepository {
	return &PostgresUserRepository{pool: pool}
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (first_name, last_name, password, email, is_admin, is_active, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.pool.Exec(ctx, query, user.FirstName, user.LastName, user.Password, user.Email, user.IsAdmin, user.IsActive, time.Now())
	if err != nil {
		slog.Error("failed to create user", "error", err)
		return err
	}

	return nil
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT *
		FROM users
		WHERE email = $1
	`

	user := domain.User{}
	err := r.pool.QueryRow(ctx, query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Email, &user.IsAdmin, &user.IsActive, &user.LastLogin, &user.CreatedAt)
	if err != nil {
		slog.Error("failed to get user by email", "error", err)
		return nil, err
	}

	return &user, nil
}
