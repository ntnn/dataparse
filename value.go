package dataparse

import "reflect"

type Value struct {
	Data any
}

func NewValue(data any) Value {
	return Value{Data: data}
}

//go:generate go run ./value_numbers.go

func (v Value) IsNil() bool {
	return v.Data == nil
}

func (v Value) List() ([]Value, error) {
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
	if l, err := v.List(); err != nil {
		return l
	}
	return []Value{v}
}
