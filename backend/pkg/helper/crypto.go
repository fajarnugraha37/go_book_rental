package helper

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHashed(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func PasswordVerify(hashed, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainPassword))
}

func GenerateOTP(issuer, username, secret string) (*otp.Key, string, time.Time, error) {
	var (
		otpExpiry   time.Time
		otpCode     string
		otpUrl      = fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s&algorithm=sha256&digits=8", issuer, username, secret, issuer)
		otpKey, err = otp.NewKeyFromURL(otpUrl)
	)
	if err != nil {
		return nil, otpCode, otpExpiry, fmt.Errorf("failed to generate OTP key: %w", err)
	}

	otpCode, err = totp.GenerateCode(secret, time.Now())
	if err != nil {
		return nil, otpCode, otpExpiry, fmt.Errorf("failed to generate OTP code: %w", err)
	}
	otpExpiry = time.Now().Add(time.Minute * 5)

	return otpKey, otpCode, otpExpiry, nil
}

func VerifyOTP(username, otpCode, otpSecret string) error {
	valid := totp.Validate(otpCode, otpSecret) // Validate OTP code
	if !valid {
		return fmt.Errorf("invalid OTP code")
	}

	return nil
}

func GenerateMagicLinkToken() (string, time.Time) {
	var (
		magicLinkToken  = uuid.NewString()
		magicLinkExpiry = time.Now().Add(time.Minute * 15)
	)

	return magicLinkToken, magicLinkExpiry
}
