package dataparse

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Int returns the underlying data as a int.
func (v Value) Int() (int, error) {
	if v.Data == nil {
		return 0, ErrValueIsNil
	}
	switch typed := v.Data.(type) {
	case int:
		return typed, nil
	case int8:
		return int(typed), nil
	case int16:
		return int(typed), nil
	case int32:
		return int(typed), nil
	case int64:
		return int(typed), nil
	case uint:
		return int(typed), nil
	case uint8:
		return int(typed), nil
	case uint16:
		return int(typed), nil
	case uint32:
		return int(typed), nil
	case uint64:
		return int(typed), nil
	case float32:
		return int(typed), nil
	case float64:
		return int(typed), nil
	case string:
		if strings.Contains(typed, ".") {
			parsed, err := strconv.ParseFloat(typed, 64)
			if err != nil {
				return 0, fmt.Errorf("dataparse: error parsing %q as Int: %w", typed, err)
			}
			return int(parsed), nil
		}
		parsed, err := strconv.ParseInt(typed, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("dataparse: error parsing %q as Int: %w", typed, err)
		}
		return int(parsed), nil
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case []byte:
		ret, numBytes := binary.Varint(typed)
		if numBytes <= 0 {
			return 0, fmt.Errorf("dataparse: error converting %v to Int: %d",
				typed, numBytes)
		}
		return int(ret), nil
		default:
		return 0, NewErrUnhandled(typed)
	}
}



// MustInt is the error-ignoring version of Int.
func (v Value) MustInt() int {
	if val, err := v.Int(); err == nil {
		return val
	}
	return 0
}



// Int8 returns the underlying data as a int8.
func (v Value) Int8() (int8, error) {
	if v.Data == nil {
		return 0, ErrValueIsNil
	}
	switch typed := v.Data.(type) {
	case int8:
		return typed, nil
	case int:
		return int8(typed), nil
	case int16:
		return int8(typed), nil
	case int32:
		return int8(typed), nil
	case int64:
		return int8(typed), nil
	case uint:
		return int8(typed), nil
	case uint8:
		return int8(typed), nil
	case uint16:
		return int8(typed), nil
	case uint32:
		return int8(typed), nil
	case uint64:
		return int8(typed), nil
	case float32:
		return int8(typed), nil
	case float64:
		return int8(typed), nil
	case string:
		if strings.Contains(typed, ".") {
			parsed, err := strconv.ParseFloat(typed, 64)
			if err != nil {
				return 0, fmt.Errorf("dataparse: error parsing %q as Int8: %w", typed, err)
			}
			return int8(parsed), nil
		}
		parsed, err := strconv.ParseInt(typed, 10, 8)
		if err != nil {
			return 0, fmt.Errorf("dataparse: error parsing %q as Int8: %w", typed, err)
		}
		return int8(parsed), nil
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case []byte:
		ret, numBytes := binary.Varint(typed)
		if numBytes <= 0 {
			return 0, fmt.Errorf("dataparse: error converting %v to Int8: %d",
				typed, numBytes)
		}
		return int8(ret), nil
		default:
		return 0, NewErrUnhandled(typed)
	}
}



// MustInt8 is the error-ignoring version of Int8.
func (v Value) MustInt8() int8 {
	if val, err := v.Int8(); err == nil {
		return val
	}
	return 0
}



// Int16 returns the underlying data as a int16.
func (v Value) Int16() (int16, error) {
	if v.Data == nil {
		return 0, ErrValueIsNil
	}
	switch typed := v.Data.(type) {
	case int16:
		return typed, nil
	case int:
		return int16(typed), nil
	case int8:
		return int16(typed), nil
	case int32:
		return int16(typed), nil
	case int64:
		return int16(typed), nil
	case uint:
		return int16(typed), nil
	case uint8:
		return int16(typed), nil
	case uint16:
		return int16(typed), nil
	case uint32:
		return int16(typed), nil
	case uint64:
		return int16(typed), nil
	case float32:
		return int16(typed), nil
	case float64:
		return int16(typed), nil
	case string:
		if strings.Contains(typed, ".") {
			parsed, err := strconv.ParseFloat(typed, 64)
			if err != nil {
				return 0, fmt.Errorf("dataparse: error parsing %q as Int16: %w", typed, err)
			}
			return int16(parsed), nil
		}
		parsed, err := strconv.ParseInt(typed, 10, 16)
		if err != nil {
			return 0, fmt.Errorf("dataparse: error parsing %q as Int16: %w", typed, err)
		}
		return int16(parsed), nil
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case []byte:
		ret, numBytes := binary.Varint(typed)
		if numBytes <= 0 {
			return 0, fmt.Errorf("dataparse: error converting %v to Int16: %d",
				typed, numBytes)
		}
		return int16(ret), nil
		default:
		return 0, NewErrUnhandled(typed)
	}
}



// MustInt16 is the error-ignoring version of Int16.
func (v Value) MustInt16() int16 {
	if val, err := v.Int16(); err == nil {
		return val
	}
	return 0
}



// Int32 returns the underlying data as a int32.
func (v Value) Int32() (int32, error) {
	if v.Data == nil {
		return 0, ErrValueIsNil
	}
	switch typed := v.Data.(type) {
	case int32:
		return typed, nil
	case int:
		return int32(typed), nil
	case int8:
		return int32(typed), nil
	case int16:
		return int32(typed), nil
	case int64:
		return int32(typed), nil
	case uint:
		return int32(typed), nil
	case uint8:
		return int32(typed), nil
	case uint16:
		return int32(typed), nil
	case uint32:
		return int32(typed), nil
	case uint64:
		return int32(typed), nil
	case float32:
		return int32(typed), nil
	case float64:
		return int32(typed), nil
	case string:
		if strings.Contains(typed, ".") {
			parsed, err := strconv.ParseFloat(typed, 64)
			if err != nil {
				return 0, fmt.Errorf("dataparse: error parsing %q as Int32: %w", typed, err)
			}
			return int32(parsed), nil
		}
		parsed, err := strconv.ParseInt(typed, 10, 32)
		if err != nil {
			return 0, fmt.Errorf("dataparse: error parsing %q as Int32: %w", typed, err)
		}
		return int32(parsed), nil
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case []byte:
		ret, numBytes := binary.Varint(typed)
		if numBytes <= 0 {
			return 0, fmt.Errorf("dataparse: error converting %v to Int32: %d",
				typed, numBytes)
		}
		return int32(ret), nil
		default:
		return 0, NewErrUnhandled(typed)
	}
}



// MustInt32 is the error-ignoring version of Int32.
func (v Value) MustInt32() int32 {
	if val, err := v.Int32(); err == nil {
		return val
	}
	return 0
}



// Int64 returns the underlying data as a int64.
func (v Value) Int64() (int64, error) {
	if v.Data == nil {
		return 0, ErrValueIsNil
	}
	switch typed := v.Data.(type) {
	case int64:
		return typed, nil
	case int:
		return int64(typed), nil
	case int8:
		return int64(typed), nil
	case int16:
		return int64(typed), nil
	case int32:
		return int64(typed), nil
	case uint:
		return int64(typed), nil
	case uint8:
		return int64(typed), nil
	case uint16:
		return int64(typed), nil
	case uint32:
		return int64(typed), nil
	case uint64:
		return int64(typed), nil
	case float32:
		return int64(typed), nil
	case float64:
		return int64(typed), nil
	case string:
		if strings.Contains(typed, ".") {
			parsed, err := strconv.ParseFloat(typed, 64)
			if err != nil {
				return 0, fmt.Errorf("dataparse: error parsing %q as Int64: %w", typed, err)
			}
			return int64(parsed), nil
		}
		parsed, err := strconv.ParseInt(typed, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("dataparse: error parsing %q as Int64: %w", typed, err)
		}
		return int64(parsed), nil
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case []byte:
		ret, numBytes := binary.Varint(typed)
		if numBytes <= 0 {
			return 0, fmt.Errorf("dataparse: error converting %v to Int64: %d",
				typed, numBytes)
		}
		return int64(ret), nil
		default:
		return 0, NewErrUnhandled(typed)
	}
}



// MustInt64 is the error-ignoring version of Int64.
func (v Value) MustInt64() int64 {
	if val, err := v.Int64(); err == nil {
		return val
	}
	return 0
}



// Uint returns the underlying data as a uint.
func (v Value) Uint() (uint, error) {
	if v.Data == nil {
		return 0, ErrValueIsNil
	}
	switch typed := v.Data.(type) {
	case uint:
		return typed, nil
	case int:
		return uint(typed), nil
	case int8:
		return uint(typed), nil
	case int16:
		return uint(typed), nil
	case int32:
		return uint(typed), nil
	case int64:
		return uint(typed), nil
	case uint8:
		return uint(typed), nil
	case uint16:
		return uint(typed), nil
	case uint32:
		return uint(typed), nil
	case uint64:
		return uint(typed), nil
	case float32:
		return uint(typed), nil
	case float64:
		return uint(typed), nil
	case string:
		if strings.Contains(typed, ".") {
			parsed, err := strconv.ParseFloat(typed, 64)
			if err != nil {
				return 0, fmt.Errorf("dataparse: error parsing %q as Uint: %w", typed, err)
			}
			return uint(parsed), nil
		}
		parsed, err := strconv.ParseUint(typed, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("dataparse: error parsing %q as Uint: %w", typed, err)
		}
		return uint(parsed), nil
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case []byte:
		ret, numBytes := binary.Uvarint(typed)
		if numBytes <= 0 {
			return 0, fmt.Errorf("dataparse: error converting %v to Uint: %d",
				typed, numBytes)
		}
		return uint(ret), nil
		default:
		return 0, NewErrUnhandled(typed)
	}
}



// MustUint is the error-ignoring version of Uint.
func (v Value) MustUint() uint {
	if val, err := v.Uint(); err == nil {
		return val
	}
	return 0
}



// Uint8 returns the underlying data as a uint8.
func (v Value) Uint8() (uint8, error) {
	if v.Data == nil {
		return 0, ErrValueIsNil
	}
	switch typed := v.Data.(type) {
	case uint8:
		return typed, nil
	case int:
		return uint8(typed), nil
	case int8:
		return uint8(typed), nil
	case int16:
		return uint8(typed), nil
	case int32:
		return uint8(typed), nil
	case int64:
		return uint8(typed), nil
	case uint:
		return uint8(typed), nil
	case uint16:
		return uint8(typed), nil
	case uint32:
		return uint8(typed), nil
	case uint64:
		return uint8(typed), nil
	case float32:
		return uint8(typed), nil
	case float64:
		return uint8(typed), nil
	case string:
		if strings.Contains(typed, ".") {
			parsed, err := strconv.ParseFloat(typed, 64)
			if err != nil {
				return 0, fmt.Errorf("dataparse: error parsing %q as Uint8: %w", typed, err)
			}
			return uint8(parsed), nil
		}
		parsed, err := strconv.ParseUint(typed, 10, 8)
		if err != nil {
			return 0, fmt.Errorf("dataparse: error parsing %q as Uint8: %w", typed, err)
		}
		return uint8(parsed), nil
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case []byte:
		ret, numBytes := binary.Uvarint(typed)
		if numBytes <= 0 {
			return 0, fmt.Errorf("dataparse: error converting %v to Uint8: %d",
				typed, numBytes)
		}
		return uint8(ret), nil
		default:
		return 0, NewErrUnhandled(typed)
	}
}



// MustUint8 is the error-ignoring version of Uint8.
func (v Value) MustUint8() uint8 {
	if val, err := v.Uint8(); err == nil {
		return val
	}
	return 0
}



// Uint16 returns the underlying data as a uint16.
func (v Value) Uint16() (uint16, error) {
	if v.Data == nil {
		return 0, ErrValueIsNil
	}
	switch typed := v.Data.(type) {
	case uint16:
		return typed, nil
	case int:
		return uint16(typed), nil
	case int8:
		return uint16(typed), nil
	case int16:
		return uint16(typed), nil
	case int32:
		return uint16(typed), nil
	case int64:
		return uint16(typed), nil
	case uint:
		return uint16(typed), nil
	case uint8:
		return uint16(typed), nil
	case uint32:
		return uint16(typed), nil
	case uint64:
		return uint16(typed), nil
	case float32:
		return uint16(typed), nil
	case float64:
		return uint16(typed), nil
	case string:
		if strings.Contains(typed, ".") {
			parsed, err := strconv.ParseFloat(typed, 64)
			if err != nil {
				return 0, fmt.Errorf("dataparse: error parsing %q as Uint16: %w", typed, err)
			}
			return uint16(parsed), nil
		}
		parsed, err := strconv.ParseUint(typed, 10, 16)
		if err != nil {
			return 0, fmt.Errorf("dataparse: error parsing %q as Uint16: %w", typed, err)
		}
		return uint16(parsed), nil
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case []byte:
		ret, numBytes := binary.Uvarint(typed)
		if numBytes <= 0 {
			return 0, fmt.Errorf("dataparse: error converting %v to Uint16: %d",
				typed, numBytes)
		}
		return uint16(ret), nil
		default:
		return 0, NewErrUnhandled(typed)
	}
}



// MustUint16 is the error-ignoring version of Uint16.
func (v Value) MustUint16() uint16 {
	if val, err := v.Uint16(); err == nil {
		return val
	}
	return 0
}



// Uint32 returns the underlying data as a uint32.
func (v Value) Uint32() (uint32, error) {
	if v.Data == nil {
		return 0, ErrValueIsNil
	}
	switch typed := v.Data.(type) {
	case uint32:
		return typed, nil
	case int:
		return uint32(typed), nil
	case int8:
		return uint32(typed), nil
	case int16:
		return uint32(typed), nil
	case int32:
		return uint32(typed), nil
	case int64:
		return uint32(typed), nil
	case uint:
		return uint32(typed), nil
	case uint8:
		return uint32(typed), nil
	case uint16:
		return uint32(typed), nil
	case uint64:
		return uint32(typed), nil
	case float32:
		return uint32(typed), nil
	case float64:
		return uint32(typed), nil
	case string:
		if strings.Contains(typed, ".") {
			parsed, err := strconv.ParseFloat(typed, 64)
			if err != nil {
				return 0, fmt.Errorf("dataparse: error parsing %q as Uint32: %w", typed, err)
			}
			return uint32(parsed), nil
		}
		parsed, err := strconv.ParseUint(typed, 10, 32)
		if err != nil {
			return 0, fmt.Errorf("dataparse: error parsing %q as Uint32: %w", typed, err)
		}
		return uint32(parsed), nil
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case []byte:
		ret, numBytes := binary.Uvarint(typed)
		if numBytes <= 0 {
			return 0, fmt.Errorf("dataparse: error converting %v to Uint32: %d",
				typed, numBytes)
		}
		return uint32(ret), nil
		default:
		return 0, NewErrUnhandled(typed)
	}
}



// MustUint32 is the error-ignoring version of Uint32.
func (v Value) MustUint32() uint32 {
	if val, err := v.Uint32(); err == nil {
		return val
	}
	return 0
}



// Uint64 returns the underlying data as a uint64.
func (v Value) Uint64() (uint64, error) {
	if v.Data == nil {
		return 0, ErrValueIsNil
	}
	switch typed := v.Data.(type) {
	case uint64:
		return typed, nil
	case int:
		return uint64(typed), nil
	case int8:
		return uint64(typed), nil
	case int16:
		return uint64(typed), nil
	case int32:
		return uint64(typed), nil
	case int64:
		return uint64(typed), nil
	case uint:
		return uint64(typed), nil
	case uint8:
		return uint64(typed), nil
	case uint16:
		return uint64(typed), nil
	case uint32:
		return uint64(typed), nil
	case float32:
		return uint64(typed), nil
	case float64:
		return uint64(typed), nil
	case string:
		if strings.Contains(typed, ".") {
			parsed, err := strconv.ParseFloat(typed, 64)
			if err != nil {
				return 0, fmt.Errorf("dataparse: error parsing %q as Uint64: %w", typed, err)
			}
			return uint64(parsed), nil
		}
		parsed, err := strconv.ParseUint(typed, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("dataparse: error parsing %q as Uint64: %w", typed, err)
		}
		return uint64(parsed), nil
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case []byte:
		ret, numBytes := binary.Uvarint(typed)
		if numBytes <= 0 {
			return 0, fmt.Errorf("dataparse: error converting %v to Uint64: %d",
				typed, numBytes)
		}
		return uint64(ret), nil
		default:
		return 0, NewErrUnhandled(typed)
	}
}



// MustUint64 is the error-ignoring version of Uint64.
func (v Value) MustUint64() uint64 {
	if val, err := v.Uint64(); err == nil {
		return val
	}
	return 0
}



// Float32 returns the underlying data as a float32.
func (v Value) Float32() (float32, error) {
	if v.Data == nil {
		return 0, ErrValueIsNil
	}
	switch typed := v.Data.(type) {
	case float32:
		return typed, nil
	case int:
		return float32(typed), nil
	case int8:
		return float32(typed), nil
	case int16:
		return float32(typed), nil
	case int32:
		return float32(typed), nil
	case int64:
		return float32(typed), nil
	case uint:
		return float32(typed), nil
	case uint8:
		return float32(typed), nil
	case uint16:
		return float32(typed), nil
	case uint32:
		return float32(typed), nil
	case uint64:
		return float32(typed), nil
	case float64:
		return float32(typed), nil
	case string:
		parsed, err := strconv.ParseFloat(typed, 32)
		if err != nil {
			return 0, fmt.Errorf("dataparse: error parsing %q as Float32: %w", typed, err)
		}
		return float32(parsed), nil
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case []byte:
		ret, numBytes := binary.Uvarint(typed)
		if numBytes <= 0 {
			return 0, fmt.Errorf("dataparse: error converting %v to Float32: %d",
				typed, numBytes)
		}
		return math.Float32frombits(uint32(ret)), nil
		default:
		return 0, NewErrUnhandled(typed)
	}
}



// MustFloat32 is the error-ignoring version of Float32.
func (v Value) MustFloat32() float32 {
	if val, err := v.Float32(); err == nil {
		return val
	}
	return 0
}



// Float64 returns the underlying data as a float64.
func (v Value) Float64() (float64, error) {
	if v.Data == nil {
		return 0, ErrValueIsNil
	}
	switch typed := v.Data.(type) {
	case float64:
		return typed, nil
	case int:
		return float64(typed), nil
	case int8:
		return float64(typed), nil
	case int16:
		return float64(typed), nil
	case int32:
		return float64(typed), nil
	case int64:
		return float64(typed), nil
	case uint:
		return float64(typed), nil
	case uint8:
		return float64(typed), nil
	case uint16:
		return float64(typed), nil
	case uint32:
		return float64(typed), nil
	case uint64:
		return float64(typed), nil
	case float32:
		return float64(typed), nil
	case string:
		parsed, err := strconv.ParseFloat(typed, 64)
		if err != nil {
			return 0, fmt.Errorf("dataparse: error parsing %q as Float64: %w", typed, err)
		}
		return float64(parsed), nil
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case []byte:
		ret, numBytes := binary.Uvarint(typed)
		if numBytes <= 0 {
			return 0, fmt.Errorf("dataparse: error converting %v to Float64: %d",
				typed, numBytes)
		}
		return math.Float64frombits(uint64(ret)), nil
		default:
		return 0, NewErrUnhandled(typed)
	}
}



// MustFloat64 is the error-ignoring version of Float64.
func (v Value) MustFloat64() float64 {
	if val, err := v.Float64(); err == nil {
		return val
	}
	return 0
}



