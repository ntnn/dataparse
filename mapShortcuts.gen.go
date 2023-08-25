package dataparse

import (
	"time"
)

// Int is a shortcut to retrieve a value and call a function on
// the resulting Value.
//
// Calling this method is equivalent to:
// val, err := m.Get("a")
// if err != nil {
//	// error handling
// }
// parsed, err := val.Int()
// if err != nil {
//	// error handling
// }
// // use parsed
func (m Map) Int(keys ...any) (int, error) {
	v, err := m.Get(keys...)
	if err != nil {
		return 0, err
	}
	return v.Int()
}

// MustInt is the error-ignoring version of Int.
func (m Map) MustInt(keys ...any) int {
	v, _ := m.Int(keys...)
	return v
}

// Int64 is a shortcut to retrieve a value and call a function on
// the resulting Value.
//
// Calling this method is equivalent to:
// val, err := m.Get("a")
// if err != nil {
//	// error handling
// }
// parsed, err := val.Int64()
// if err != nil {
//	// error handling
// }
// // use parsed
func (m Map) Int64(keys ...any) (int64, error) {
	v, err := m.Get(keys...)
	if err != nil {
		return 0, err
	}
	return v.Int64()
}

// MustInt64 is the error-ignoring version of Int64.
func (m Map) MustInt64(keys ...any) int64 {
	v, _ := m.Int64(keys...)
	return v
}

// Uint is a shortcut to retrieve a value and call a function on
// the resulting Value.
//
// Calling this method is equivalent to:
// val, err := m.Get("a")
// if err != nil {
//	// error handling
// }
// parsed, err := val.Uint()
// if err != nil {
//	// error handling
// }
// // use parsed
func (m Map) Uint(keys ...any) (uint, error) {
	v, err := m.Get(keys...)
	if err != nil {
		return 0, err
	}
	return v.Uint()
}

// MustUint is the error-ignoring version of Uint.
func (m Map) MustUint(keys ...any) uint {
	v, _ := m.Uint(keys...)
	return v
}

// Uint64 is a shortcut to retrieve a value and call a function on
// the resulting Value.
//
// Calling this method is equivalent to:
// val, err := m.Get("a")
// if err != nil {
//	// error handling
// }
// parsed, err := val.Uint64()
// if err != nil {
//	// error handling
// }
// // use parsed
func (m Map) Uint64(keys ...any) (uint64, error) {
	v, err := m.Get(keys...)
	if err != nil {
		return 0, err
	}
	return v.Uint64()
}

// MustUint64 is the error-ignoring version of Uint64.
func (m Map) MustUint64(keys ...any) uint64 {
	v, _ := m.Uint64(keys...)
	return v
}

// String is a shortcut to retrieve a value and call a function on
// the resulting Value.
//
// Calling this method is equivalent to:
// val, err := m.Get("a")
// if err != nil {
//	// error handling
// }
// parsed, err := val.String()
// if err != nil {
//	// error handling
// }
// // use parsed
func (m Map) String(keys ...any) (string, error) {
	v, err := m.Get(keys...)
	if err != nil {
		return "", err
	}
	return v.String()
}

// MustString is the error-ignoring version of String.
func (m Map) MustString(keys ...any) string {
	v, _ := m.String(keys...)
	return v
}

// Time is a shortcut to retrieve a value and call a function on
// the resulting Value.
//
// Calling this method is equivalent to:
// val, err := m.Get("a")
// if err != nil {
//	// error handling
// }
// parsed, err := val.Time()
// if err != nil {
//	// error handling
// }
// // use parsed
func (m Map) Time(keys ...any) (time.Time, error) {
	v, err := m.Get(keys...)
	if err != nil {
		return time.Time{}, err
	}
	return v.Time()
}

// MustTime is the error-ignoring version of Time.
func (m Map) MustTime(keys ...any) time.Time {
	v, _ := m.Time(keys...)
	return v
}
