package dataparse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue_String(t *testing.T) {
	// string to string
	assert.Equal(t, "test", NewValue("test").String())

	// integer to string
	assert.Equal(t, "1", NewValue(1).String())

	// negative integer to string
	assert.Equal(t, "-1", NewValue(-1).String())

	// int8 to string
	assert.Equal(t, "1", NewValue(int8(1)).String())

	// unsigned integer to string
	assert.Equal(t, "1", NewValue(uint(1)).String())

	// uint8 to string
	assert.Equal(t, "1", NewValue(uint8(1)).String())

	// float to string
	assert.Equal(t, "1", NewValue(1.0).String())

	assert.Equal(t, "1.4", NewValue(1.4).String())

	// rune to string
	assert.Equal(t, "c", NewValue('c').String())

	assert.Equal(t,
		"test string",
		NewValue(any("test string")).String(),
	)
}

func TestValue_ListString(t *testing.T) {
	// string to string
	s, err := NewValue("test").ListString(",")
	assert.Nil(t, err)
	assert.Equal(t, []string{"test"}, s)

	// strings to string
	s, err = NewValue("test1,test2").ListString(",")
	assert.Nil(t, err)
	assert.Equal(t, []string{"test1", "test2"}, s)

	s, err = NewValue("test1 | test2").ListString(" | ")
	assert.Nil(t, err)
	assert.Equal(t, []string{"test1", "test2"}, s)

	s, err = NewValue("test1 |test2").ListString("|")
	assert.Nil(t, err)
	assert.Equal(t, []string{"test1 ", "test2"}, s)

	// integer to string
	s, err = NewValue(1).ListString(",")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1"}, s)

	// negative integer to string
	s, err = NewValue(-1).ListString("|")
	assert.Nil(t, err)
	assert.Equal(t, []string{"-1"}, s)

	// unsigned integer to string
	s, err = NewValue(uint(1)).ListString(",")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1"}, s)

	// float to string
	s, err = NewValue(1.0).ListString(",")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1"}, s)

	s, err = NewValue(1.0).ListString(".")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1"}, s)

	s, err = NewValue(1.4).ListString(",")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.4"}, s)

	s, err = NewValue(1.4).ListString(".")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "4"}, s)

	s, err = NewValue(1.6).ListString(",")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.6"}, s)

	s, err = NewValue(1.6).ListString(".")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "6"}, s)

	// rune to string
	s, err = NewValue('c').ListString(",")
	assert.Nil(t, err)
	assert.Equal(t, []string{"c"}, s)

	// string slice to string slice
	s, err = NewValue([]string{"test1", "test2", "test3"}).ListString(",")
	assert.Nil(t, err)
	assert.Equal(t, []string{"test1", "test2", "test3"}, s)

	// any slice to string slice
	s, err = NewValue([]any{"test1", "test2", "test3"}).ListString(",")
	assert.Nil(t, err)
	assert.Equal(t, []string{"test1", "test2", "test3"}, s)
}
