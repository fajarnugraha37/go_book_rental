package jwt

import (
	"backend/internal/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtService struct {
	cfg *config.Config
}

func NewJwtService(cfg *config.Config) *JwtService {
	return &JwtService{
		cfg: cfg,
	}
}

func (j *JwtService) Signed(input ClaimsInput) (string, error) {
	claims := NewClaims(input)
	claims.ExpiresAt = time.Now().Add(time.Duration(j.cfg.Jwt.ExpirationInSecond) * time.Second).Unix()
	claims.Issuer = j.cfg.Jwt.Issuer
	// Group     string
	// Audience  string
	// Subject   string

	token := jwt.NewWithClaims(j.signingMethod(), claims)
	signedToken, err := token.SignedString([]byte(j.cfg.Jwt.SignatureKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT token: %w", err)
	}

	return signedToken, nil
}

func (j *JwtService) Verify(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != j.signingMethod() {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte(j.cfg.Jwt.SignatureKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}
	claims := token.Claims.(Claims)
	if err = claims.Valid(); err != nil {
		return nil, fmt.Errorf("claims is invalid: %w", err)
	}

	return token, nil
}

func (j *JwtService) signingMethod() *jwt.SigningMethodHMAC {
	switch j.cfg.Jwt.SigningMethod {
	case "HS256":
		return jwt.SigningMethodHS256
	case "HS384":
		return jwt.SigningMethodHS384
	case "HS512":
		return jwt.SigningMethodHS512
	default:
		panic(fmt.Errorf("invalid signing method: %+v", j.cfg.Jwt.SigningMethod))
	}
}
