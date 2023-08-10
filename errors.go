package dataparse

import (
	"errors"
	"fmt"
)

var (
	ErrValueIsNil = errors.New("dataparse: value is nil")
)

type ErrUnhandled struct {
	Value any
}

func NewErrUnhandled(value any) ErrUnhandled {
	return ErrUnhandled{Value: value}
}

func (e ErrUnhandled) Error() string {
	return fmt.Sprintf("dataparse: value type %T of value %q is not handled",
		e.Value, e.Value)
}
