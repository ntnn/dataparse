package dataparse

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
)

type Map map[any]any

// FromNDJSON returns maps parsed from a stream of newline delimited
// JSON.
func FromNDJSON(reader io.Reader, opts ...ReadOption) (chan Map, chan error) {
	cfg := newReadConfig(opts...)
	mapCh, errCh := cfg.channels()

	decoder := json.NewDecoder(reader)

	go func() {
		defer close(mapCh)
		defer close(errCh)
		defer cfg.closeFinishChannel()

		for decoder.More() {
			// decoder refuses to decode into Map or map[any]any
			var m any
			if err := decoder.Decode(&m); err != nil && !errors.Is(err, io.EOF) {
				errCh <- err
				return
			}
			mMap, err := NewMap(m)
			if err != nil {
				errCh <- err
				return
			}
			mapCh <- mMap
		}
	}()

	return mapCh, errCh
}

// FromNDJSONFile returns maps parsed from a file containing newline
// delimited JSON objects.
func FromNDJSONFile(path string, opts ...ReadOption) (chan Map, chan error) {
	f, err := os.Open(path)
	if err != nil {
		errCh := make(chan error, 1)
		errCh <- err
		close(errCh)
		return nil, errCh
	}

	finishCh := make(chan struct{}, 1)
	opts = append(opts, withFinishChannel(finishCh))
	mapCh, errCh := FromNDJSON(f, opts...)

	go func() {
		<-finishCh
		f.Close()
	}()

	return mapCh, errCh
}

func NewMap(in any) (Map, error) {
	if in == nil {
		return Map{}, ErrValueIsNil
	}

	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Map:
		m := Map{}
		iter := val.MapRange()
		for iter.Next() {
			m[iter.Key().Interface()] = iter.Value().Interface()
		}
		return m, nil
	case reflect.Struct:
		m := Map{}
		for i := 0; i < val.NumField(); i++ {
			m[val.Type().Field(i).Name] = val.Field(i).Interface()
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
