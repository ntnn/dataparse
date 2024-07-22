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

type Fromer interface {
	From(Value) error
}

type CustomToFunc func(source Value, other any) (any, bool, error)

var customTo = []CustomToFunc{stdlibToer}

func AddCustomToFunc(fn CustomToFunc) {
	customTo = append(customTo, fn)
}

// To transforms the stored data into the target type and returns any
// occurring errors.
//
// The passed value must be a pointer.
//
// To utilizes the various transformation methods and returns their
// errors.
//
// If the parameter satisfies the Fromer interface it will be used to
// set the value.
func (v Value) To(other any, opts ...ToOption) error {
	if fromer, ok := other.(Fromer); ok {
		return fromer.From(v)
	}

	target := reflect.ValueOf(other)

	if target.Kind() != reflect.Pointer {
		return ErrValueIsNotPointer
	}

	// dereference until the target is a pointer but the value pointer
	// to is not
	// for target.Kind() == reflect.Pointer && target.Elem().Kind() == reflect.Pointer {
	for target.Kind() == reflect.Pointer {
		if target.IsNil() {
			// initialize pointer with a valid value
			target.Set(reflect.New(target.Type().Elem()))
		}
		// handle pointers to constants or structs that satisfy the
		// Fromer interface
		if fromer, ok := target.Interface().(Fromer); ok {
			return fromer.From(v)
		}
		target = target.Elem()
	}

	// handle slices but skip named types (like net.IP which is
	// a []byte)
	if target.Type().Name() == "" && target.Kind() == reflect.Slice || target.Kind() == reflect.Array {
		vs, err := v.List()
		if err != nil {
			return fmt.Errorf("dataparse: target is a slice, error converting %T to slice: %w",
				v.Data, err)
		}

		converts := reflect.MakeSlice(
			target.Type(),
			len(vs),
			len(vs),
		)

		for i, v := range vs {
			if err := v.To(converts.Index(i).Addr().Interface(), opts...); err != nil {
				return err
			}
		}

		target.Set(converts)
		return nil
	}

	for _, fn := range customTo {
		customVal, ok, err := fn(v, target.Interface())
		if err != nil {
			return fmt.Errorf("dataparse: error in custom toer: %w", err)
		}
		if !ok {
			continue
		}
		target.Set(reflect.ValueOf(customVal))
		return nil
	}

	// If the passed value is a pointer to a struct try
	// converting Value to map and call .To
	if target.Kind() == reflect.Struct {
		m, err := v.Map()
		if err != nil {
			return err
		}
		return m.To(other, opts...)
	}

	return fmt.Errorf("dataparse: unhandled type: %T", other)
}

func stdlibToer(v Value, other any) (any, bool, error) {
	switch other.(type) {
	case string:
		newValue, err := v.String()
		return newValue, true, err
	case int:
		newValue, err := v.Int()
		return newValue, true, err
	case int8:
		newValue, err := v.Int8()
		return newValue, true, err
	case int16:
		newValue, err := v.Int16()
		return newValue, true, err
	case int32:
		newValue, err := v.Int32()
		return newValue, true, err
	case int64:
		newValue, err := v.Int64()
		return newValue, true, err
	case uint:
		newValue, err := v.Uint()
		return newValue, true, err
	case uint8:
		newValue, err := v.Uint8()
		return newValue, true, err
	case uint16:
		newValue, err := v.Uint16()
		return newValue, true, err
	case uint32:
		newValue, err := v.Uint32()
		return newValue, true, err
	case uint64:
		newValue, err := v.Uint64()
		return newValue, true, err
	case float32:
		newValue, err := v.Float32()
		return newValue, true, err
	case float64:
		newValue, err := v.Float64()
		return newValue, true, err
	case bool:
		newValue, err := v.Bool()
		return newValue, true, err
	case net.IP:
		newValue, err := v.IP()
		return newValue, true, err
	case time.Time:
		newValue, err := v.Time()
		return newValue, true, err
	default:
		return nil, false, nil
	}
}

// List returns the underlying data as a slice of Values.
//
// The passed separators are passed to .ListString if the underlying
// value is a string.
//
// Warning: This method is very simplistic and at the moment only
// returns a proper slice of values if the underlying data is a slice.
func (v Value) List(seps ...string) ([]Value, error) {
	if v.Data == nil {
		return []Value{}, ErrValueIsNil
	}

	switch reflect.TypeOf(v.Data).Kind() {
	case reflect.String:
		s, err := v.ListString(seps...)
		if err != nil {
			return nil, err
		}
		vs := make([]Value, len(s))
		for i := range s {
			vs[i] = NewValue(s[i])
		}
		return vs, nil
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
func (v Value) Map() (*Map, error) {
	return NewMap(v.Data)
}

// MustMap is the error-ignoring version of Map.
func (v Value) MustMap() *Map {
	m, _ := v.Map()
	return m
}
