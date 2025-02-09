package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Claims struct {
	jwt.StandardClaims
	Sub      string   `json:"sub"`
	Username string   `json:"username"`
	Group    string   `json:"group"`
	Roles    []string `json:"roles"`
}

type ClaimsInput struct {
	UserID    string
	Username  string
	Group     string
	Roles     []string
	Audience  string
	Issuer    string
	Subject   string
	ExpiresAt int
}

func NewClaims(input ClaimsInput) *Claims {
	return &Claims{
		Sub:      input.UserID,
		Username: input.Username,
		Group:    input.Group,
		Roles:    input.Roles,
		StandardClaims: jwt.StandardClaims{
			Audience: input.Audience,
			// The "exp" (expiration time) claim identifies the expiration time on
			// or after which the JWT MUST NOT be accepted for processing.  The
			// processing of the "exp" claim requires that the current date/time
			// MUST be before the expiration date/time listed in the "exp" claim.
			ExpiresAt: int64(input.ExpiresAt),
			Id:        uuid.NewString(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    input.Issuer,
			// The "nbf" (not before) claim identifies the time before which the JWT
			// MUST NOT be accepted for processing.  The processing of the "nbf"
			// claim requires that the current date/time MUST be after or equal to
			// the not-before date/time listed in the "nbf" claim.
			NotBefore: int64(input.ExpiresAt),
			Subject:   input.Subject,
		},
	}
}
