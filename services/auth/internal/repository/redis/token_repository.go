package redis

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/nidea1/go-gavel/services/auth/internal/repository"
	"github.com/redis/go-redis/v9"
)

type RedisTokenRepository struct {
	client *redis.Client
}

func NewRedisTokenRepository(addr, password string, db uint8) repository.TokenRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return &RedisTokenRepository{client}
}

func (r *RedisTokenRepository) AddToBlacklist(ctx context.Context, token string, expiresAt time.Time) error {
	ttl := time.Until(expiresAt)
	if ttl < 0 {
		ttl = 0
	}
	key := fmt.Sprintf("blacklist:%s", token)
	return r.client.Set(ctx, key, 1, ttl).Err()
}

func (r *RedisTokenRepository) IsBlacklisted(ctx context.Context, token string) (bool, error) {
	key := fmt.Sprintf("blacklist:%s", token)
	return r.client.Exists(ctx, key).Val() == 1, nil
}

func (r *RedisTokenRepository) SaveRefreshToken(ctx context.Context, userID uint64, refreshToken string, expiresAt time.Time) error {
	key := fmt.Sprintf("refresh:%d", userID)
	ttl := time.Until(expiresAt)
	if ttl < 0 {
		ttl = 0
	}
	return r.client.Set(ctx, key, refreshToken, ttl).Err()
}

func (r *RedisTokenRepository) GetRefreshToken(ctx context.Context, userID uint64) (string, error) {
	key := fmt.Sprintf("refresh:%d", userID)
	return r.client.Get(ctx, key).Result()
}

func (r *RedisTokenRepository) DeleteRefreshToken(ctx context.Context, userID uint64) error {
	key := fmt.Sprintf("refresh:%d", userID)
	return r.client.Del(ctx, key).Err()
}
