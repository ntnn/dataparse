package dataparse

import (
	"fmt"
	"reflect"
)

type Value struct {
	Data any
}

func NewValue(data any) Value {
	return Value{Data: data}
}

//go:generate go run ./cmd/gen-value-numbers

func (v Value) IsNil() bool {
	return v.Data == nil
}

func (v Value) To(other any) error {
	var err error
	switch typed := other.(type) {
	case *string:
		*typed, err = v.String()
	case *int:
		*typed, err = v.Int()
	case *int8:
		*typed, err = v.Int8()
	case *int16:
		*typed, err = v.Int16()
	case *int32:
		*typed, err = v.Int32()
	case *int64:
		*typed, err = v.Int64()
	case *uint:
		*typed, err = v.Uint()
	case *uint8:
		*typed, err = v.Uint8()
	case *uint16:
		*typed, err = v.Uint16()
	case *uint32:
		*typed, err = v.Uint32()
	case *uint64:
		*typed, err = v.Uint64()
	case *float32:
		*typed, err = v.Float32()
	case *float64:
		*typed, err = v.Float64()
	case *bool:
		*typed, err = v.Bool()
	// case *byte:
	// 	*typed = v.MustByte()
	// case *[]byte:
	// 	*typed = v.MustBytes()
	default:
		return fmt.Errorf("dataparse: unhandled type: %T", other)
	}
	return err
}

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

func (v Value) MustList() []Value {
	l, _ := v.List()
	return l
}

func (v Value) Map() (Map, error) {
	return NewMap(v.Data)
}

func (v Value) MustMap() Map {
	m, _ := v.Map()
	return m
}
