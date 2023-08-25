package dataparse

import (
	"fmt"
	"reflect"
	"strings"
)

func (v Value) String() string {
	if v.Data == nil {
		return ""
	}

	switch typed := v.Data.(type) {
	case rune:
		return string(typed)
	default:
		return fmt.Sprintf("%v", v.Data)
	}
}

func (v Value) TrimString() string {
	return strings.TrimSpace(v.String())
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
		return strings.Split(v.String(), sep), nil
	}
}

func (v Value) MustListString(sep string) []string {
	val, _ := v.ListString(sep)
	return val
}
