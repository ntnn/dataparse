package dataparse

import (
	"fmt"
	"net"
	"reflect"
	"time"
)

// Value is one of the two central types in dataparse.
// It is used to transform data between various representations.
type Value struct {
	Data any
}

// NewValue returns the passed data as a Value.
func NewValue(data any) Value {
	return Value{Data: data}
}

//go:generate go run ./cmd/gen-value-numbers

// IsNil returns true if the data Value stores is nil.
func (v Value) IsNil() bool {
	return v.Data == nil
}

// To transforms the stored data into the target type and returns any
// occurring errors.
//
// The passed value must be a pointer.
//
// To utilizes the various transformation methods and returns their
// errors.
func (v Value) To(other any) error {
	var err error
	switch typed := other.(type) {
	case *string:
		*typed, err = v.String()
	case **string:
		var p string
		p, err = v.String()
		*typed = &p
	case *[]string:
		*typed, err = v.ListString(",")
	case **[]string:
		var p []string
		p, err = v.ListString(",")
		*typed = &p
	case *int:
		*typed, err = v.Int()
	case **int:
		var p int
		p, err = v.Int()
		*typed = &p
	case *int8:
		*typed, err = v.Int8()
	case **int8:
		var p int8
		p, err = v.Int8()
		*typed = &p
	case *int16:
		*typed, err = v.Int16()
	case **int16:
		var p int16
		p, err = v.Int16()
		*typed = &p
	case *int32:
		*typed, err = v.Int32()
	case **int32:
		var p int32
		p, err = v.Int32()
		*typed = &p
	case *int64:
		*typed, err = v.Int64()
	case **int64:
		var p int64
		p, err = v.Int64()
		*typed = &p
	case *uint:
		*typed, err = v.Uint()
	case **uint:
		var p uint
		p, err = v.Uint()
		*typed = &p
	case *uint8:
		*typed, err = v.Uint8()
	case **uint8:
		var p uint8
		p, err = v.Uint8()
		*typed = &p
	case *uint16:
		*typed, err = v.Uint16()
	case **uint16:
		var p uint16
		p, err = v.Uint16()
		*typed = &p
	case *uint32:
		*typed, err = v.Uint32()
	case **uint32:
		var p uint32
		p, err = v.Uint32()
		*typed = &p
	case *uint64:
		*typed, err = v.Uint64()
	case **uint64:
		var p uint64
		p, err = v.Uint64()
		*typed = &p
	case *float32:
		*typed, err = v.Float32()
	case **float32:
		var p float32
		p, err = v.Float32()
		*typed = &p
	case *float64:
		*typed, err = v.Float64()
	case **float64:
		var p float64
		p, err = v.Float64()
		*typed = &p
	case *bool:
		*typed, err = v.Bool()
	case **bool:
		var p bool
		p, err = v.Bool()
		*typed = &p
	case *net.IP:
		*typed, err = v.IP()
	case **net.IP:
		var p net.IP
		p, err = v.IP()
		*typed = &p
	case *time.Time:
		*typed, err = v.Time()
	case **time.Time:
		var t time.Time
		t, err = v.Time()
		*typed = &t
	// case *byte:
	// 	*typed = v.MustByte()
	// case *[]byte:
	// 	*typed = v.MustBytes()
	default:
		// If the passed value is a pointer to a struct try
		// converting Value to map and call .To
		if reflect.ValueOf(other).Elem().Kind() == reflect.Struct {
			m, err := v.Map()
			if err != nil {
				return err
			}
			return m.To(other)
		}
		return fmt.Errorf("dataparse: unhandled type: %T", other)
	}
	return err
}

// List returns the underlying data as a slice of Values.
//
// Warning: This method is very simplistic and at the moment only
// returns a proper slice of values if the underlying data is a slice.
func (v Value) List() ([]Value, error) {
	if v.Data == nil {
		return []Value{}, ErrValueIsNil
	}

	switch reflect.TypeOf(v.Data).Kind() {
	case reflect.Slice:
		l := reflect.ValueOf(v.Data)
		ret := make([]Value, l.Len())
		for i := 0; i < l.Len(); i++ {
			ret[i] = NewValue(l.Index(i).Interface())
		}
		return ret, nil
	default:
		return []Value{v}, nil
	}
}

// MustList is the error-ignoring version of List.
func (v Value) MustList() []Value {
	l, _ := v.List()
	return l
}

// Map returns the underlying data as a Map.
func (v Value) Map() (Map, error) {
	return NewMap(v.Data)
}

// MustMap is the error-ignoring version of Map.
func (v Value) MustMap() Map {
	m, _ := v.Map()
	return m
}
