package helpers

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/procode2/structio/models"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// Extend this to also generate refresh tokens
// https://dev.to/mdfaizan7/creating-jwt-s-and-signup-route-part-2-3-of-go-authentication-series-5339
type JWTClaims struct {
	Id string
	jwt.StandardClaims
}

func GetJWTKey(user *models.User) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	jwtClaim := &JWTClaims{
		Id: fmt.Sprint(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)

	fmt.Printf("JWTKeY %s", jwtKey)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValiadateTokenString(tokenString string) (claim *JWTClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(*JWTClaims)

	if !ok {
		return nil, errors.New("Could not parse claims")
	}

	if claim.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("The token Expired")
	}

	return claim, nil
}
