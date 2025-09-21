package counter

import "errors"

var (
	ErrNotFound = errors.New("counter not found")
)

func NewNotFoundError() error {
	return ErrNotFound
}
