package dataparse

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

func (v Value) Time() (time.Time, error) {
	switch typed := v.Data.(type) {
	// assume epoch for int/uint
	case int, int8, int16, int32, int64:
		return time.Unix(reflect.ValueOf(v.Data).Int(), 0), nil
	case uint, uint8, uint16, uint32, uint64:
		u := reflect.ValueOf(v.Data).Uint()
		if u > math.MaxInt64 {
			return time.Time{}, fmt.Errorf("dataparse: value is too big to parse as time: %q", u)
		}
		return time.Unix(int64(u), 0), nil
	case float32, float64:
		// TODO There's probably more useful ways to interpret a float
		// as time than simply converting to int.
		return time.Unix(int64(reflect.ValueOf(v.Data).Float()), 0), nil
	case string:
		return ParseTime(typed)
	default:
		return time.Time{}, nil
	}
}

func (v Value) MustTime() time.Time {
	val, _ := v.Time()
	return val
}
