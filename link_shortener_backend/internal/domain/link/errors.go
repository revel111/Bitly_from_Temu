package link

import "errors"

var (
	ErrServiceUnavailable = errors.New("service is unavailable")
	ErrInvalidURL         = errors.New("invalid URL format")
	ErrNotFound           = errors.New("link not found")
	ErrExists             = errors.New("link already exists")
)

func NewServiceUnavailableError() error {
	return ErrServiceUnavailable
}

func NewInvalidUrlError() error {
	return ErrInvalidURL
}

func NewNotFoundError() error {
	return ErrNotFound
}

func NewExistsError() error {
	return ErrExists
}
