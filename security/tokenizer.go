package security

import (
	"time"

	"github.com/aashisDevv/login-api/config"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	secret    string
	expiryMin int
}

func NewTokenService(authCfg *config.AuthConfig) *TokenService {
	return &TokenService{
		secret:    authCfg.Secret,
		expiryMin: authCfg.ExpiryMin,
	}
}

func (ts *TokenService) GenerateAccessToken(payload RequestClaims) (string, error) {
	now := time.Now()
	expiryTime := now.Add(time.Duration(ts.expiryMin) * time.Minute)

	payload.ExpiresAt = jwt.NewNumericDate(expiryTime)
	payload.IssuedAt = jwt.NewNumericDate(now)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString([]byte(ts.secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
