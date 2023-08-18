package dataparse

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
)

type Map map[any]any

// From returns maps parsed from a file.
func From(path string, opts ...FromOption) (chan Map, chan error, error) {
	cfg := newFromConfig(opts...)
	defer cfg.Close()

	reader, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("dataparse: error opening file: %w", err)
	}
	cfg.reader = reader
	cfg.closers = append(cfg.closers, reader.Close)

	ext := filepath.Ext(path)

	switch ext {
	case ".gz":
		gzReader, err := gzip.NewReader(reader)
		if err != nil {
			return nil, nil, fmt.Errorf("dataparse: error creating gzip reader: %w", err)
		}
		cfg.reader = gzReader
		cfg.closers = append(cfg.closers, gzReader.Close)
	}

	if reader != cfg.reader {
		ext = filepath.Ext(filepath.Ext(path))
	}

	var fn func(cfg *FromConfig) (chan Map, chan error)
	switch ext {
	case ".ndjson":
		fn = fromNdjson
	default:
		return nil, nil, fmt.Errorf("dataparse: unhandled file extension: %q", ext)
	}

	chMap, chErr := fn(cfg)
	return chMap, chErr, nil
}

// FromNDJSON returns maps parsed from a stream of newline delimited
// JSON.
func FromNDJSON(reader io.Reader, opts ...FromOption) (chan Map, chan error) {
	cfg := newFromConfig(opts...)
	return fromNdjson(cfg)
}

func fromNdjson(cfg *FromConfig) (chan Map, chan error) {
	mapCh, errCh := cfg.channels()

	decoder := json.NewDecoder(cfg.reader)

	go func() {
		defer cfg.Close()
		defer close(mapCh)
		defer close(errCh)

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
