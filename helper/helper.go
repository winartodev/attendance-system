package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type ResponseSuccess struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseFailed struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

// VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (ok bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	if err != nil {
		return false, ErrPasswordIncorrect
	}

	return true, err
}
