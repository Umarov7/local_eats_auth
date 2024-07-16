package token

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

const (
	refreshsigningkey = "Here I am!"
)

func GenerateRefreshToken(userID string) (string, error) {
	token := *jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	newToken, err := token.SignedString([]byte(refreshsigningkey))
	if err != nil {
		log.Println(err)
		return "", errors.Wrap(err, "failed to generate refresh token")
	}

	return newToken, nil
}

func ValidateRefreshToken(tokenStr string) (bool, error) {
	claims, err := ExtractRefreshClaims(tokenStr)
	if err != nil {
		return false, errors.Wrap(err, "validation failure")
	}

	mp := *claims
	exp, ok := mp["exp"].(float64)
	if !ok {
		return false, errors.New("expiration not found")
	}

	if float64(time.Now().Unix()) > exp {
		return false, errors.New("token is expired")
	}

	return true, nil
}

func ExtractRefreshClaims(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(refreshsigningkey), nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid refresh token")
	}

	return &claims, nil
}

func GetUserIdFromRefreshToken(refreshTokenString string) (string, error) {
	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshsigningkey), nil
	})
	if err != nil || !refreshToken.Valid {
		return "", errors.Wrap(err, "invalid token")
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.Wrap(err, "invalid payload")
	}
	userID := claims["user_id"].(string)

	return userID, nil
}

func InvalidateRefreshToken(refreshTokenString string) (string, error) {
	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshsigningkey), nil
	})

	if err != nil {
		return "", errors.Wrap(err, "failed to parse token")
	}

	if !refreshToken.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.Wrap(err, "invalid payload")
	}

	claims["exp"] = time.Now().Unix()

	newToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(refreshsigningkey))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate new token")
	}

	return newToken, nil
}
