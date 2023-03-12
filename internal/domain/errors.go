package domain

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrInernalServerError = errors.New("internal server error")
	ErrBadParamInput      = errors.New("given param is not valid")
)
