package dataparse

import (
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

//go:generate go run ./cmd/gen-map-shortcuts

// Map is one of the two central types in dataparse.
// It is used to store and retrieve data taken from various sources.
type Map map[any]any

// From returns maps parsed from a file.
//
// From utilizes other functions for various data types like JSON and
// CSV.
//
// From automatically unpacks the following archives based on their file
// extension:
//   - gzip: .gz
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
		ext = filepath.Ext(strings.TrimSuffix(path, filepath.Ext(path)))
	}

	var fn func(cfg *fromConfig) (chan Map, chan error)
	switch ext {
	case ".json", ".ndjson":
		fn = fromJson
	case ".csv":
		fn = fromCsv
	case ".tsv":
		// Default to tab as separator for .tsv
		c2 := newFromConfig(append([]FromOption{WithSeparator("\t")}, opts...)...)
		cfg.separator = c2.separator
		fn = fromCsv
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
	cfg.reader = reader
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

func fromJson(cfg *fromConfig) (chan Map, chan error) {
	mapCh, errCh := cfg.channels()

	decoder := json.NewDecoder(cfg.reader)

	go func() {
		defer cfg.Close()

		for decoder.More() {
			// decoder refuses to decode into Map or map[any]any
			var m any
			if err := decoder.Decode(&m); err != nil && !errors.Is(err, io.EOF) {
				errCh <- err
				return
			}
			val := reflect.ValueOf(m)
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

// FromCsv returns maps read from a CSV stream.
func FromCsv(reader io.Reader, opts ...FromOption) (chan Map, chan error) {
	cfg := newFromConfig(opts...)
	cfg.reader = reader
	return fromCsv(cfg)
}

func fromCsv(cfg *fromConfig) (chan Map, chan error) {
	mapCh, errCh := cfg.channels()

	if len(cfg.separator) != 1 {
		defer cfg.Close()
		errCh <- fmt.Errorf("dataparse: separator must be a string of length one for csv, got %q", cfg.separator)
		return mapCh, errCh
	}

	reader := csv.NewReader(cfg.reader)
	reader.Comma = rune(cfg.separator[0])
	reader.FieldsPerRecord = len(cfg.headers)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = cfg.trimSpace

	go func() {
		defer cfg.Close()

		if len(cfg.headers) == 0 {
			h, err := reader.Read()
			if err != nil {
				errCh <- err
				return
			}
			cfg.headers = h
		}

		for {
			elems, err := reader.Read()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				errCh <- err
				return
			}

			m := Map{}
			for i := range elems {
				m[cfg.headers[i]] = elems[i]
			}
			mapCh <- m
		}
	}()

	return mapCh, errCh
}

// FromKVString returns a map based on the passed string.
//
// Example:
//
//	input: a=1,b=test,c
//	output: {
//		a: 1,
//		b: "test",
//		c: nil,
//	}
func FromKVString(kv string, opts ...FromOption) (Map, error) {
	cfg := newFromConfig(opts...)

	m := Map{}
	for _, elem := range strings.Split(kv, cfg.separator) {
		split := strings.SplitN(elem, "=", 2)

		key := strings.TrimSpace(split[0])
		var value any
		if len(split) > 1 {
			if cfg.trimSpace {
				value = strings.TrimSpace(split[1])
			} else {
				value = split[1]
			}
		}

		m[key] = value
	}

	return m, nil
}

// NewMap creates a map from the passed value.
// Valid values are maps and structs.
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

// Has returns true if the map has an entry for any of the passed keys.
// The keys are checked in order.
func (m Map) Has(keys ...any) bool {
	return !m.MustGet(keys...).IsNil()
}

// Get checks for Value entries for each of the given keys in order and
// returns the first.
// If no Value is found a dataparse.Value `nil` and an error is
// returned.
//
// Nested value can be accessed by providing the keys separated with
// dots.
//
// Example:
//
//	m, err := NewMap(map[string]any{
//		"a": map[string]any{
//			"b": map[string]any{
//				"c": "lorem ipsum",
//			},
//		},
//	})
//	if err != nil {
//		return err
//	}
//	v, err := m.Get("a.b.c")
//	if err != nil {
//		return err
//	}
//	fmt.Printf(v.MustString())
//
// Will print "lorem ipsum".
//
// Note: Errors from attempting to convert Values to Maps are returned
// as stdlib multierrors and only when no match is found.
//
// Note: The entire key including dots is tested as well and the value
// returned if it exists.
// Example:
//
//	m, err := NewMap(map[string]any{
//		"a.b.c": "dolor sic amet",
//	})
//	if err != nil {
//		return err
//	}
//	m2, err := NewMap(map[string]any{
//		"a": map[string]any{
//			"b": map[string]any{
//				"c": "lorem ipsum",
//			},
//		},
//		"a.b.c": "dolor sic amet",
//	})
//	if err != nil {
//		return err
//	}
//	v, err := m.Get("a.b.c")
//	if err != nil {
//		return err
//	}
//	v2, err := m2.Get("a.b.c")
//	if err != nil {
//		return err
//	}
//	fmt.Printf(v.MustString())
//	fmt.Printf(v2.MustString())
//
// Will print "dolor sic amet" and "lorem ipsum".
func (m Map) Get(keys ...any) (Value, error) {
	var errs error
	for _, key := range keys {
		b, v, err := m.get(key)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		if b {
			return v, nil
		}

		if v, ok := m[key]; ok {
			return NewValue(v), nil
		}
	}

	return NewValue(nil), errors.Join(errs, NewErrNoValidKey(keys))
}

func (m Map) get(key any) (bool, Value, error) {
	switch typed := key.(type) {
	case string:
		if strings.Contains(typed, ".") {
			splitKey := strings.SplitN(typed, ".", 2)
			if len(splitKey) != 2 {
				return false, NewValue(nil), fmt.Errorf(
					"dataparse: splitn on %q with . did not produce exactly two values",
					typed,
				)
			}

			v, ok := m[splitKey[0]]
			if !ok {
				return false, NewValue(nil), nil
			}

			subM, err := NewValue(v).Map()
			if err != nil {
				return false, NewValue(nil), fmt.Errorf(
					"dataparse: key %q indicated nested maps but value at key %q cannot be converted to map: %#v",
					typed,
					splitKey[0],
					v,
				)
			}

			ret, err := subM.Get(splitKey[1])
			if err != nil {
				return false, NewValue(nil), nil
			}
			return true, ret, nil
		}
	}
	return false, NewValue(nil), nil
}

// MustGet is the error-ignoring version of Get.
func (m Map) MustGet(keys ...any) Value {
	v, _ := m.Get(keys...)
	return v
}

// Map works like Get but returns a Map.
func (m Map) Map(keys ...any) (Map, error) {
	for _, key := range keys {
		if v, ok := m[key]; ok {
			return NewMap(v)
		}
	}
	return Map{}, fmt.Errorf("dataparse: no valid keys: %v", keys)
}

// MustMap is the error-ignoring version of Map.
func (m Map) MustMap(keys ...any) Map {
	v, _ := m.Map(keys...)
	return v
}

type toConfig struct {
	lookupFieldName       bool
	skipFieldsWithoutTag  bool
	ignoreNoValidKeyError bool
	collectErrors         bool
}

func newToConfig() *toConfig {
	cfg := new(toConfig)
	cfg.lookupFieldName = true
	return cfg
}

type ToOption func(*toConfig)

func cfgFromOpts(opts ...ToOption) *toConfig {
	cfg := newToConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithLookupFieldName configures Map.To to try to lookup the field name
// in addition to any names given in the dataparse tag.
//
// If this option is set the field name will be looked up after any
// names in the dataparse tag.
//
// The default is true.
func WithLookupFieldName(lookupFieldName bool) ToOption {
	return func(cfg *toConfig) {
		cfg.lookupFieldName = lookupFieldName
	}
}

// WithSkipFieldsWithoutTag configures Map.To to skip fields without
// explicit tags.
//
// Note that this also skip fields without an explicit dataparse tag if
// WithLookupFieldName is set.
//
// The default is false.
func WithSkipFieldsWithoutTag() ToOption {
	return func(cfg *toConfig) {
		cfg.skipFieldsWithoutTag = true
	}
}

// WithIgnoreNoValidKeyError configures Map.To to ignore errors when no
// field could by found by the configured tags.
//
// This is primarily useful for inconsistent input or when using the
// same structure to parse data from different sources with different
// properties.
//
// The default is false.
func WithIgnoreNoValidKeyError() ToOption {
	return func(cfg *toConfig) {
		cfg.ignoreNoValidKeyError = true
	}
}

// WithCollectErrors configures Map.To to not return on the first
// encountered error when processing properties.
//
// Instead occurring errors are collected with errors.Join and returned
// after processing all fields.
//
// The default is false.
func WithCollectErrors() ToOption {
	return func(cfg *toConfig) {
		cfg.collectErrors = true
	}
}

// To reads the map into a struct similar to json.Unmarshal, utilizing Value.To.
// The passed variable must be a pointer to a struct.
//
// Multiple keys can be given, separated by a commata `,`:
//
//	type example struct {
//		Field string `dataparse:"field1,field2"`
//	}
//
// By default the field name is looked up if any fields in the dataparse
// tag are not found.
//
// By default it is an error if a struct field cannot be found in the
// Map.
// Fields without a dataparse tag can be skipped implicitly by passing
// the option WithSkipFieldsWithoutTag or explicitly by settings
// `dataparse:""`:
//
//	type example struct {
//		Field string `dataparse:""`
//	}
//
// Value.To uses the underlying field type to call the correct Value
// method to transform the source value into the targeted struct field
// type.
// E.g. if the field type is string and the map contains a number the
// field will contain a string with the number formatted in.
func (m Map) To(dest any, opts ...ToOption) error {
	cfg := cfgFromOpts(opts...)

	refV := reflect.ValueOf(dest)
	if refV.Kind() != reflect.Pointer {
		return ErrValueIsNotPointer
	}

	if refV.IsNil() {
		return ErrValueIsNil
	}

	for refV.Kind() == reflect.Pointer {
		refV = refV.Elem()
	}

	refT := refV.Type()
	var errs error

	for i := 0; i < refT.NumField(); i++ {
		fieldRefT := refT.Field(i)

		lookupKeys := []any{fieldRefT.Name}

		tags, ok := fieldRefT.Tag.Lookup("dataparse")
		if !ok && cfg.skipFieldsWithoutTag {
			continue
		}

		if ok {
			// skip the field on dataparse:""
			if len(tags) == 0 {
				continue
			}
			lookupKeys = ListToAny(strings.Split(tags, ","))
		}

		if cfg.lookupFieldName {
			lookupKeys = append(lookupKeys, fieldRefT.Name)
		}

		if len(lookupKeys) == 0 {
			err := fmt.Errorf("dataparse: no keys found to lookup for field %q, this error can be prevented by passing WithSkipFieldsWithoutTag",
				fieldRefT.Name)
			if !cfg.collectErrors {
				return err
			}
			errs = errors.Join(errs, err)
			continue
		}

		v, err := m.Get(lookupKeys...)
		if err != nil {
			if cfg.ignoreNoValidKeyError && errors.As(err, &ErrNoValidKey{}) {
				continue
			}
			err := fmt.Errorf("dataparse: error getting field %q from map: %w",
				fieldRefT.Name, err)
			if !cfg.collectErrors {
				return err
			}
			errs = errors.Join(errs, err)
			continue
		}

		fieldRefV := refV.Field(i)
		if !fieldRefV.CanAddr() {
			err := fmt.Errorf("dataparse: field %q is not addressable", fieldRefT.Name)
			if !cfg.collectErrors {
				return err
			}
			errs = errors.Join(errs, err)
			continue
		}

		if err := v.To(fieldRefV.Addr().Interface(), opts...); err != nil {
			err := fmt.Errorf("dataparse: error setting field %q from value %v: %w",
				fieldRefT.Name, v, err)
			if !cfg.collectErrors {
				return err
			}
			errs = errors.Join(errs, err)
		}
	}

	return errs
}
