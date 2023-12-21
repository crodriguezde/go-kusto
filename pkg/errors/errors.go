package errors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidType = errors.New("invalid type")
)

func ErrWrapf(err error, format string, a ...any) error {
	return fmt.Errorf("%s: %w", fmt.Sprintf(format, a...), err)
}
