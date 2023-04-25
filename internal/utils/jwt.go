package utils

import (
	"errors"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

// @TODO : make function create jwt token and validate

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

var jwtSecretKey = ""

// SetJWTSecretKey sets the jwt secret key
func SetJWTSecretKey(key string) {
	jwtSecretKey = key
}

// GenerateNewJWT generates a JWT token with the given claims
func GenerateNewJWT(claims *Claims) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	signedToken, err = token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// GetJWTUserId returns the userid contained in the JWT token
func GetJWTUserId(tokenString string) (res uint, err error) {
	claims, err := GetJWTClaims(tokenString)
	if err != nil {
		return res, err
	}

	userId, err := strconv.Atoi(claims.UserId)
	if err != nil {
		return res, err
	}
	return uint(userId), nil
}

// GetJWTUserIdString returns the userid contained in the JWT token in the string format
func GetJWTUserIdString(tokenString string) (res string, err error) {
	claims, err := GetJWTClaims(tokenString)
	if err != nil {
		return res, err
	}

	return claims.UserId, nil
}

// GetJWTClaims returns the claims contained in the jwt token
func GetJWTClaims(tokenString string) (claims *Claims, err error) {
	claims = &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("you're Unauthorized")
	}
	return claims, nil
}
