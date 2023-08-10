package dataparse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue_String(t *testing.T) {
	// string to string
	s, err := NewValue("test").String()
	assert.Nil(t, err)
	assert.Equal(t, "test", s)

	// integer to string
	s, err = NewValue(1).String()
	assert.Nil(t, err)
	assert.Equal(t, "1", s)

	// negative integer to string
	s, err = NewValue(-1).String()
	assert.Nil(t, err)
	assert.Equal(t, "-1", s)

	// int8 to string
	s, err = NewValue(int8(1)).String()
	assert.Nil(t, err)
	assert.Equal(t, "1", s)

	// unsigned integer to string
	s, err = NewValue(uint(1)).String()
	assert.Nil(t, err)
	assert.Equal(t, "1", s)

	// uint8 to string
	s, err = NewValue(uint8(1)).String()
	assert.Nil(t, err)
	assert.Equal(t, "1", s)

	// float to string
	s, err = NewValue(1.0).String()
	assert.Nil(t, err)
	assert.Equal(t, "1", s)

	s, err = NewValue(1.4).String()
	assert.Nil(t, err)
	assert.Equal(t, "1.4", s)

	// rune to string
	s, err = NewValue('c').String()
	assert.Nil(t, err)
	assert.Equal(t, "c", s)
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
}
