package utils

import "errors"

var (
	ErrInvalidSigningKey    = errors.New("invalid signing method")
	ErrInvalidTypeOfClaims  = errors.New("token claims are not type of `*tokenClaims`")
	ErrUserIdNotFound       = errors.New("user id not found")
	ErrTooMuchBalance       = errors.New("too much balance, your limit is under 5000")
	ErrInvalidAccountName   = errors.New("invalid account name")
	ErrUnsupportedMediaType = errors.New("unsupported media type, type must be 'jpg' or 'jpeg'")
	ErrInvalidName          = errors.New("invalid name")
	ErrInvalidType          = errors.New("invalid type")
)
