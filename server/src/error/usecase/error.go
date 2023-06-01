package usecase

import "errors"

var (
	IdEmptyError    = errors.New("id empty")
	TitleEmptyError = errors.New("title empty")
)
