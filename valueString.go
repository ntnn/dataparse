package dataparse

import (
	"fmt"
	"reflect"
	"strings"
)

// String returns the underlying value as a string.
func (v Value) String() (string, error) {
	if v.Data == nil {
		return "", nil
	}

	switch typed := v.Data.(type) {
	case rune:
		return string(typed), nil
	default:
		return fmt.Sprintf("%v", v.Data), nil
	}
}

func (v Value) MustString() string {
	s, _ := v.String()
	return s
}

func (v Value) TrimString() string {
	return strings.TrimSpace(v.MustString())
}

func (v Value) ListString(sep string) ([]string, error) {
	val := reflect.ValueOf(v.Data)
	switch val.Kind() {
	case reflect.Slice:
		ret := make([]string, val.Len())
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i)
			if elem.Kind() == reflect.Interface {
				elem = elem.Elem()
			}
			ret[i] = elem.String()
		}
		return ret, nil
	default:
		s, err := v.String()
		if err != nil {
			return nil, fmt.Errorf("dataparse: error turning %q into string to split: %w", v.Data, err)
		}
		return strings.Split(s, sep), nil
	}
}

func (v Value) MustListString(sep string) []string {
	val, _ := v.ListString(sep)
	return val
}
