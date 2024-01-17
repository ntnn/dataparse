package dataparse

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValue_Int(t *testing.T) {
	parsed, err := NewValue("123").Int()
	require.Nil(t, err)
	assert.Equal(t, 123, parsed)

	parsed, err = NewValue("123.6").Int()
	require.Nil(t, err)
	assert.Equal(t, 123, parsed)

	parsed, err = NewValue("123.0").Int()
	require.Nil(t, err)
	assert.Equal(t, 123, parsed)

	parsed, err = NewValue(" 123.0").Int()
	require.Nil(t, err)
	assert.Equal(t, 123, parsed)

	parsed, err = NewValue("123.0 ").Int()
	require.Nil(t, err)
	assert.Equal(t, 123, parsed)

	parsed, err = NewValue(" 123.0 ").Int()
	require.Nil(t, err)
	assert.Equal(t, 123, parsed)

	parsed, err = NewValue(0).Int()
	require.Nil(t, err)
	assert.Equal(t, 0, parsed)

	parsed, err = NewValue(math.MaxInt64).Int()
	require.Nil(t, err)
	assert.Equal(t, math.MaxInt64, parsed)

	parsed, err = NewValue(0x00).Int()
	require.Nil(t, err)
	assert.Equal(t, 0, parsed)

	parsed, err = NewValue(0x08).Int()
	require.Nil(t, err)
	assert.Equal(t, 8, parsed)

	parsed, err = NewValue([]byte{0x9c, 0x85, 0xe3, 0xb, 0x0, 0x0, 0x0, 0x0}).Int()
	require.Nil(t, err)
	assert.Equal(t, 12345678, parsed)

	parsed, err = NewValue(true).Int()
	require.Nil(t, err)
	assert.Equal(t, 1, parsed)

	parsed, err = NewValue(false).Int()
	require.Nil(t, err)
	assert.Equal(t, 0, parsed)
}
