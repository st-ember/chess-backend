package tools

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Error("No .env file found")
	}
}

// generate access and refresh tokens with user claims
func GenerateTokens(username string) (accessToken, refreshToken string, refreshExpiresAt *jwt.NumericDate, err error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	if len(key) == 0 {
		err = errors.New("jwt secret not set")
		return "", "", refreshExpiresAt, err
	}

	accessClaims := UserAuthClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			Issuer:    os.Getenv("JWT_ISSUER"),
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(key)

	if err != nil {
		return "", "", refreshExpiresAt, err
	}

	refreshExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7))

	refreshClaims := jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: refreshExpiresAt,
		Issuer:    os.Getenv("JWT_ISSUER"),
		ID:        fmt.Sprintf("%d", time.Now().Unix()),
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(key)

	if err != nil {
		return "", "", refreshExpiresAt, err
	}

	return accessToken, refreshToken, refreshExpiresAt, nil
}

type UserAuthClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
