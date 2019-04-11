package utils

import "errors"

var (
	ErrNotFound             = errors.New("not found")
	ErrNotFoundOrganisation = errors.New("organisation not found")
	ErrNotValidCurrency     = errors.New("currency not valid")
	ErrInvalid              = errors.New("invalid request")
	ErrDuplicate            = errors.New("duplicate resource")
)
