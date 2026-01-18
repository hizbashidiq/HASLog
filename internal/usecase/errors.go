package usecase

import(
	"errors"
)

var(
	// registration
	ErrEmailAlreadyExists = errors.New("email already exist")
	ErrUsernameAlreadyExists = errors.New("username already exist")

	//login
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken = errors.New("invalid token")

	//Other
	ErrUnknown = errors.New("unknown error")
)