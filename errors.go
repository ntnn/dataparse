package dataparse

import (
	"errors"
	"fmt"
)

var (
	ErrValueIsNil        = errors.New("dataparse: value is nil")
	ErrValueIsNotPointer = errors.New("dataparse: value is not pointer")
)

// ErrUnhandled is returned as an error if the underlying type is not
// handled by dataparse.
type ErrUnhandled struct {
	Value any
}

// NewErrUnhandled returns an ErrUnhandled with the given value.
func NewErrUnhandled(value any) ErrUnhandled {
	return ErrUnhandled{Value: value}
}

func (e ErrUnhandled) Error() string {
	return fmt.Sprintf("dataparse: value type %T of value %q is not handled",
		e.Value, e.Value)
}

type ErrNoValidKey struct {
	keys []any
}

func NewErrNoValidKey(keys []any) ErrNoValidKey {
	return ErrNoValidKey{keys: keys}
}

func (e ErrNoValidKey) Error() string {
	return fmt.Sprintf("dataparse: no valid key: %v", e.keys)
}

func (e ErrNoValidKey) Keys() []any {
	return e.keys
}
