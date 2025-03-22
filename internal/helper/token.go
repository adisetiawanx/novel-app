package helper

import (
	"fmt"
	"github.com/adisetiawanx/novel-app/internal/app"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateAccessToken(userID string, userRole string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    userRole,
		"exp":     time.Now().Add(time.Hour * 8).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(app.App.Token.AccessSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateRefreshToken(userID string, userRole string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    userRole,
		"exp":     time.Now().Add(time.Hour * 720).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(app.App.Token.RefreshSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseAccessToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(app.App.Token.AccessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ParseRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(app.App.Token.RefreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
