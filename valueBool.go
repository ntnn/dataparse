package dataparse

import (
	"fmt"
	"strconv"
)

func (v Value) Bool() (bool, error) {
	if v.Data == nil {
		return false, ErrValueIsNil
	}

	var s string
	switch typed := v.Data.(type) {
	case string:
		s = typed
	default:
		s = fmt.Sprintf("%v", v.Data)
	}

	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, fmt.Errorf("dataparse: error parsing bool from %q: %w", s, err)
	}
	return b, nil
}

func (v Value) MustBool() bool {
	b, _ := v.Bool()
	return b
}
