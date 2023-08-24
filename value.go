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
	switch typed := other.(type) {
	case *string:
		*typed = v.String()
	case *int:
		*typed = v.MustInt()
	case *int8:
		*typed = v.MustInt8()
	case *int16:
		*typed = v.MustInt16()
	case *int32:
		*typed = v.MustInt32()
	case *int64:
		*typed = v.MustInt64()
	case *uint:
		*typed = v.MustUint()
	case *uint8:
		*typed = v.MustUint8()
	case *uint16:
		*typed = v.MustUint16()
	case *uint32:
		*typed = v.MustUint32()
	case *uint64:
		*typed = v.MustUint64()
	case *float32:
		*typed = v.MustFloat32()
	case *float64:
		*typed = v.MustFloat64()
	case *bool:
		*typed = v.MustBool()
	// case *byte:
	// 	*typed = v.MustByte()
	// case *[]byte:
	// 	*typed = v.MustBytes()
	default:
		return fmt.Errorf("dataparse: unhandled type: %T", other)
	}
	return nil
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
