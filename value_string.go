package dataparse

import (
	"fmt"
	"strings"
)

func (v Value) String() (string, error) {
	if v.Data == nil {
		return "", ErrValueIsNil
	}
	switch typed := v.Data.(type) {
	case rune:
		return string(typed), nil
	default:
		return fmt.Sprintf("%v", v.Data), nil
	}
}

func (v Value) ListString(sep string) ([]string, error) {
	val, err := v.String()
	if err != nil {
		return nil, err
	}

	return strings.Split(val, sep), nil
}

func (v Value) MustString() string {
	return fmt.Sprintf("%v", v.Data)
}

func (v Value) MustListString(sep string) []string {
	if val, err := v.ListString(sep); err == nil {
		return val
	}
	return []string{}
}
