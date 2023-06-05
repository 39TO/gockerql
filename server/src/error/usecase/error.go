package usecase

import "errors"

var (
	// todo
	IdEmptyError    = errors.New("id empty")
	TitleEmptyError = errors.New("title empty")

	//user
	NameEmptyError     = errors.New("name empty")
	EmailEmptyError    = errors.New("email empty")
	PasswordEmptyError = errors.New("password empty")
)
