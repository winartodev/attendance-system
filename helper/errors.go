package helper

import "errors"

var (
	ErrPasswordIncorrect = errors.New("login or passowrd is incorrect")
	ErrInvalidToken      = errors.New("the token is invalid")
	ErrTokenExpired      = errors.New("token is expired")
	ErrNoToken           = errors.New("user not loggedin")
)
