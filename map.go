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

	"github.com/k0kubun/pp/v3"
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
	case ".json", ".ndjson":
		fn = fromJson
	default:
		return nil, nil, fmt.Errorf("dataparse: unhandled file extension: %q", ext)
	}

	chMap, chErr := fn(cfg)
	return chMap, chErr, nil
}

// FromSingle is a wrapper around From and returns the first map and
// error in the result set.
// It is only intended for instances where it is already known that the
// input can only contain a single document.
func FromSingle(path string, opts ...FromOption) (Map, error) {
	chMap, chErr, err := From(path, append(opts, WithChannelSize(1))...)
	if err != nil {
		return nil, err
	}
	return <-chMap, <-chErr
}

// FromJson returns maps parsed from a stream which may consist of:
// 1. A single JSON document
// 2. A stream of JSON documents
// 3. An array of JSON documents
func FromJson(reader io.Reader, opts ...FromOption) (chan Map, chan error) {
	cfg := newFromConfig(opts...)
	return fromJson(cfg)
}

// FromJsonSingle is a wrapper around FromJson and returns the first map
// and error in the result set.
// It is only intended for instances where it is already known that the
// input can only contain a single document.
func FromJsonSingle(reader io.Reader, opts ...FromOption) (Map, error) {
	mapCh, mapErr := FromJson(reader, append(opts, WithChannelSize(1))...)
	return <-mapCh, <-mapErr
}

func fromJson(cfg *FromConfig) (chan Map, chan error) {
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
			val := reflect.ValueOf(m)
			pp.Println("kind", val.Kind())
			switch val.Kind() {
			case reflect.Slice:
				for i := 0; i < val.Len(); i++ {
					elem, err := NewMap(val.Index(i).Interface())
					if err != nil {
						errCh <- fmt.Errorf("dataparse: error parsing element %d: %w", i, err)
						return
					}
					mapCh <- elem
				}
			case reflect.Struct, reflect.Map:
				mMap, err := NewMap(m)
				if err != nil {
					errCh <- err
					return
				}
				mapCh <- mMap
			default:
				errCh <- fmt.Errorf("dataparse: unhandled type %q in file", val.Kind())
				return
			}
		}
	}()

	return mapCh, errCh
}

func NewMap(in any) (Map, error) {
	if in == nil {
		return Map{}, ErrValueIsNil
	}

	val := reflect.ValueOf(in)
	for val.Kind() == reflect.Pointer {
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
			field := val.Field(i)
			if field.CanInterface() {
				m[val.Type().Field(i).Name] = field.Interface()
			}
		}
		return m, nil
	default:
		return Map{}, fmt.Errorf("dataparse: cannot be transformed to map: %T", in)
	}
}

func (m Map) Has(keys ...any) bool {
	return !m.MustGet(keys...).IsNil()
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
