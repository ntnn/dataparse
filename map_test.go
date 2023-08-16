package dataparse

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromNDJSONFile(t *testing.T) {
	maps, errs := FromNDJSONFile("./testdata/data.ndjson")
	for err := range errs {
		require.Nil(t, err)
	}

	m := <-maps
	require.NotNil(t, m)
	assert.Equal(t, "Garrott", m.MustGet("first_name").String())

	m = <-maps
	require.NotNil(t, m)
	assert.Equal(t, "Vasovic", m.MustGet("last_name").String())

	m = <-maps
	require.NotNil(t, m)
	assert.Equal(t, 3, m.MustGet("id").MustInt())

	m = <-maps
	require.NotNil(t, m)
	assert.Equal(t,
		time.Date(2023, time.June, 20, 23, 34, 57, 0, time.UTC),
		m.MustGet("timestamp").MustTime(),
	)

	m = <-maps
	require.NotNil(t, m)
	assert.Equal(t,
		net.ParseIP("166.215.142.79"),
		m.MustGet("ip_address").MustIP(),
	)
}

func TestNewMap(t *testing.T) {
	// convert maps to Map
	m, err := NewMap(map[any]any{1: 1})
	require.Nil(t, err)
	assert.Equal(t, 1, m.MustGet(1).MustInt())

	m, err = NewMap(map[int]string{1: "test"})
	require.Nil(t, err)
	assert.Equal(t, "test", m.MustGet(1).String())

	_, err = NewMap(nil)
	require.NotNil(t, err)

	_, err = NewMap("test")
	require.NotNil(t, err)

	// handle pointers
	ptrMap := map[string]int{"test": 1}
	m, err = NewMap(&ptrMap)
	require.Nil(t, err)
	assert.Equal(t, 1, m.MustGet("test").MustInt())

	// handle structs
	type aStruct struct {
		A string
		B int
	}
	anInstance := aStruct{
		A: "test",
		B: 5,
	}
	m, err = NewMap(anInstance)
	require.Nil(t, err)
	assert.Equal(t, "test", m.MustGet("A").String())
	assert.Equal(t, 5, m.MustGet("B").MustInt())

	// handle pointer to struct
	m, err = NewMap(&anInstance)
	require.Nil(t, err)
	assert.Equal(t, "test", m.MustGet("A").String())
	assert.Equal(t, 5, m.MustGet("B").MustInt())

}
