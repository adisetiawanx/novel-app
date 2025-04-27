package helper

import (
	"fmt"
	"github.com/adisetiawanx/novel-app/internal/app"
	"github.com/golang-jwt/jwt"
	"time"
)

type CustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func CreateAccessToken(userID string, userRole string) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 8).Unix()
	claims := &CustomClaims{
		UserID: userID,
		Role:   userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(app.Config.Token.AccessSecret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, exp, nil
}

func CreateRefreshToken(userID string, userRole string) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 720).Unix()
	claims := &CustomClaims{
		UserID: userID,
		Role:   userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(app.Config.Token.RefreshSecret))
	if err != nil {
		return "", 0, err
	}
	return tokenString, exp, nil
}

func VerifyRefreshToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.Config.Token.RefreshSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims, nil
}

func ParseAccessToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(app.Config.Token.AccessSecret), nil
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
		return []byte(app.Config.Token.RefreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
