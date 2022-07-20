package pkg

import "errors"

const (
	errInvalidInput = "invalid input"
)

var (
	ErrInvalidInput = errors.New(errInvalidInput)
)
