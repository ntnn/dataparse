package dataparse

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// BoolStringsFalse are strings that will be interpreted as false by
// Value.Bool.
var BoolStringsFalse = []string{
	"",

	"0",
	"no",
	"n",
	"false",

	"na",
	"n/a",
}

// BoolStringsTrue are strings taht will be interpreted as true by
// Value.Bool.
var BoolStringsTrue = []string{
	"1",
	"yes",
	"y",
	"true",
}

// Bool returns a boolean for the underlying lower cased value.
//
// Strings in BoolStringsFalse and BoolStringsTrue are considered to be
// false and true respectively.
//
// If neither applies strconv.ParseBool is utilized.
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

	s = strings.ToLower(s)

	if slices.Contains(BoolStringsFalse, s) {
		return false, nil
	}

	if slices.Contains(BoolStringsTrue, s) {
		return true, nil
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
