package dataparse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue_String(t *testing.T) {
	// string to string
	assert.Equal(t, "test", NewValue("test").MustString())

	// integer to string
	assert.Equal(t, "1", NewValue(1).MustString())

	// negative integer to string
	assert.Equal(t, "-1", NewValue(-1).MustString())

	// int8 to string
	assert.Equal(t, "1", NewValue(int8(1)).MustString())

	// unsigned integer to string
	assert.Equal(t, "1", NewValue(uint(1)).MustString())

	// uint8 to string
	assert.Equal(t, "1", NewValue(uint8(1)).MustString())

	// float to string
	assert.Equal(t, "1", NewValue(1.0).MustString())

	assert.Equal(t, "1.4", NewValue(1.4).MustString())

	// rune to string
	assert.Equal(t, "c", NewValue('c').MustString())

	assert.Equal(t,
		"test string",
		NewValue(any("test string")).MustString(),
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
	assert.Equal(t, []string{"test1", "test2"}, s)

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
