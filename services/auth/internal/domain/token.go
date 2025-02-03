package domain

import (
	"time"
)

type Token struct {
	UserID       uint64    `db:"user_id"`
	AccessToken  string    `db:"access_token"`
	RefreshToken string    `db:"refresh_token"`
	ExpiryAt     uint64    `db:"expiry_at"`
	CreatedAt    time.Time `db:"created_at"`
}
