package jwt

import (
	"log/slog"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey       string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewJWTManager(secretKey string, accessTokenTTL, refreshTokenTTL time.Duration) *JWTManager {
	return &JWTManager{secretKey, accessTokenTTL, refreshTokenTTL}
}

func (m *JWTManager) Generate(userID uint64, email string) (accessTokenStr string, refreshTokenStr string, err error) {
	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(m.accessTokenTTL).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenStr, err = accessToken.SignedString([]byte(m.secretKey))
	if err != nil {
		slog.Error(err.Error())
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(m.refreshTokenTTL).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenStr, err = refreshToken.SignedString([]byte(m.secretKey))
	if err != nil {
		slog.Error(err.Error())
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}

func (m *JWTManager) Verify(accessTokenStr string) (*jwt.Token, error) {
	return jwt.Parse(accessTokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.secretKey), nil
	})
}
