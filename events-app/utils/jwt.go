package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "SECRET"

type UserClaims struct {
	UserId int64  `bindings:"required" json:"user_id"`
	Email  string `bindings:"required" json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		userId,
		email,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "NikunjGoServer",
			Subject:   "User",
			Audience: jwt.ClaimStrings{
				"someone",
			},
			ID: "1",
		},
	})

	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyUserToken(token string) (*UserClaims, error) {

	jwt_token, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, errors.New("could not parse token")
	}

	if !jwt_token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := jwt_token.Claims.(*UserClaims)

	if !ok {
		return nil, errors.New("invalid claims type, cannot proceed")
	}

	return claims, nil
}
