package dataparse

import (
	"fmt"
	"reflect"
	"strings"
)

// String returns the underlying value as a string.
//
// Note that String never returns an error and is identical to
// MustString. String and MustString are only kept to follow the same
// conventions as all other transformation methods follow.
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

// MustString is the error-ignoring version of String.
func (v Value) MustString() string {
	s, _ := v.String()
	return s
}

// TrimString returns the result of String with spaces trimmed.
func (v Value) TrimString() string {
	return strings.TrimSpace(v.MustString())
}

// DefaultStringSeparators is used when no separators are passed,
var DefaultStringSeparators = []string{
	",",
	"\n",
}

// ListString returns the underlying data as a slice of strings.
//
// If the underlying data is a slice each member is transformed into
// a string using the Value.String method.
//
// If the underlying data is a string the string is split using the
// passed separator. If not separators are passed
// DefaultStringSeparators is used.
func (v Value) ListString(seps ...string) ([]string, error) {
	val := reflect.ValueOf(v.Data)
	switch val.Kind() {
	case reflect.Slice:
		ret := make([]string, val.Len())
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i)
			if elem.Kind() == reflect.Interface {
				elem = elem.Elem()
			}
			ret[i] = NewValue(elem.Interface()).MustString()
			if v.cfg.trimSpace {
				ret[i] = strings.TrimSpace(ret[i])
			}
		}
		return ret, nil
	default:
		s, err := v.String()
		if err != nil {
			return nil, fmt.Errorf("dataparse: error turning %q into string to split: %w", v.Data, err)
		}

		if len(seps) == 0 {
			seps = DefaultStringSeparators
		}

		for _, sep := range seps {
			split := strings.Split(s, sep)
			if v.cfg.trimSpace {
				for i := range split {
					split[i] = strings.TrimSpace(split[i])
				}
			}
			if len(split) == 1 {
				// walk through separators until more than one element
				// is present
				continue
			}
			return split, nil
		}

		// default to returning the singular element if no separator
		// split the element.
		return []string{s}, nil
	}
}

// MustListString is the error-ignoring version of ListString.
func (v Value) MustListString(sep string) []string {
	val, _ := v.ListString(sep)
	return val
}
