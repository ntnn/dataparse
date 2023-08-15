package dataparse

import (
	"fmt"
	"reflect"
)

type Map map[any]any

func NewMap(in any) (Map, error) {
	if in == nil {
		return Map{}, ErrValueIsNil
	}

	val := reflect.ValueOf(in)
	switch val.Kind() {
	case reflect.Map:
		m := Map{}
		iter := val.MapRange()
		for iter.Next() {
			m[iter.Key().Interface()] = iter.Value().Interface()
		}
		return m, nil
	default:
		return Map{}, fmt.Errorf("dataparse: cannot be transformed to map: %T", in)
	}
}

func (m Map) Get(keys ...any) (Value, error) {
	for _, key := range keys {
		if v, ok := m[key]; ok {
			return NewValue(v), nil
		}
	}
	return NewValue(nil), fmt.Errorf("dataparse: no valid key: %v", keys)
}

func (m Map) MustGet(keys ...any) Value {
	v, _ := m.Get(keys...)
	return v
}

func (m Map) Map(keys ...any) (Map, error) {
	for _, key := range keys {
		if v, ok := m[key]; ok {
			return NewMap(v)
		}
	}
	return Map{}, fmt.Errorf("dataparse: no valid keys: %v", keys)
}

func (m Map) MustMap(keys ...any) Map {
	v, _ := m.Map(keys...)
	return v
}
