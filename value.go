package dataparse

import "reflect"

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
