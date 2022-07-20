package pkg

import "errors"

const (
	errEmptyPassword = "1000-External User I/O Error"
	errWrongPassword = "1001-External User I/O Error"
)

const (
	errInvalidCost = "900-Invalid Cost"
)

var (
	ErrInternalEmptyPassword = errors.New(errEmptyPassword)
	ErrInternalWrongPassword = errors.New(errWrongPassword)
)

var (
	ErrInternalInvalidCost = errors.New(errInvalidCost)
)
