package helper

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type SignedDetails struct {
	Username string
	Role     int64
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateToken(username string, role int64) (signedToken string, err error) {
	claims := &SignedDetails{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	signedToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return signedToken, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return claims, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return claims, ErrInvalidToken
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return claims, ErrTokenExpired
	}

	return claims, err
}
