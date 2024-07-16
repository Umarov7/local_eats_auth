package token

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

const (
	signingkey = "hello world"
)

func GenerateAccessToken(id, username, email string) (string, error) {
	token := *jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["username"] = username
	claims["email"] = email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()

	newToken, err := token.SignedString([]byte(signingkey))

	if err != nil {
		log.Println(err)
		return "", errors.Wrap(err, "failed to generate access token")
	}

	return newToken, nil
}

func ValidateAccessToken(tokenStr string) (bool, error) {
	_, err := ExtractAccessClaims(tokenStr)
	if err != nil {
		return false, errors.Wrap(err, "validation failure")
	}
	return true, nil
}

func ExtractAccessClaims(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingkey), nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse access token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid access token")
	}

	return &claims, nil
}
