package dataparse

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFrom_Json(t *testing.T) {
	maps, errs, err := From("./testdata/data.json")
	require.Nil(t, err)
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

func TestFromSingle_Json(t *testing.T) {
	m, err := FromSingle("./testdata/single.json")
	require.Nil(t, err)
	require.NotNil(t, m)

	assert.Equal(t, "Garrott", m.MustGet("first_name").String())
	assert.Equal(t, "Felgate", m.MustGet("last_name").String())
	assert.Equal(t, 1, m.MustGet("id").MustInt())

	assert.Equal(t,
		time.Date(2022, time.September, 28, 23, 9, 27, 0, time.UTC),
		m.MustGet("timestamp").MustTime(),
	)

	assert.Equal(t,
		net.ParseIP("77.111.249.225"),
		m.MustGet("ip_address").MustIP(),
	)
}

func TestFrom_Ndjson(t *testing.T) {
	maps, errs, err := From("./testdata/data.ndjson")
	require.Nil(t, err)
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

func TestFrom_Csv(t *testing.T) {
	maps, errs, err := From("./testdata/data.csv")
	require.Nil(t, err)
	for err := range errs {
		require.Nil(t, err)
	}

	m := <-maps
	require.NotNil(t, m)
	assert.Equal(t, "Robbie", m.MustGet("first_name").String())

	m = <-maps
	require.NotNil(t, m)
	assert.Equal(t, "Salvin", m.MustGet("last_name").String())

	m = <-maps
	require.NotNil(t, m)
	assert.Equal(t, 3, m.MustGet("id").MustInt())

	m = <-maps
	require.NotNil(t, m)
	assert.Equal(t,
		time.Date(2023, time.June, 26, 7, 22, 35, 0, time.UTC),
		m.MustGet("timestamp").MustTime(),
	)

	m = <-maps
	require.NotNil(t, m)
	assert.Equal(t,
		net.ParseIP("56.85.108.10"),
		m.MustGet("ip_address").MustIP(),
	)
}

func TestFromKVString(t *testing.T) {
	m, err := FromKVString("a=1,b=test,c,d=0x05")
	require.Nil(t, err)

	assert.Equal(t, 1, m.MustGet("a").MustInt())
	assert.Equal(t, "test", m.MustGet("b").String())
	assert.True(t, m.MustGet("c").IsNil())

	m, err = FromKVString("key1=value1|key2=value2", WithSeparator("|"))
	require.Nil(t, err)
	assert.Equal(t, "value1", m.MustGet("key1").String())
	assert.Equal(t, "value2", m.MustGet("key2").String())
}

func TestNewMap(t *testing.T) {
	// convert maps to Map
	m, err := NewMap(map[any]any{1: 1})
	require.Nil(t, err)
	assert.Equal(t, 1, m.MustGet(1).MustInt())

	m, err = NewMap(map[int]string{1: "lorem ipsum"})
	require.Nil(t, err)
	assert.Equal(t, "lorem ipsum", m.MustGet(1).String())

	_, err = NewMap(nil)
	require.NotNil(t, err)

	_, err = NewMap("test")
	require.NotNil(t, err)

	// handle pointers
	ptrMap := map[string]int{"dolor sit": 1}
	m, err = NewMap(&ptrMap)
	require.Nil(t, err)
	assert.Equal(t, 1, m.MustGet("dolor sit").MustInt())

	// handle structs
	type aStruct struct {
		A string
		B int
	}
	anInstance := aStruct{
		A: "amet consectetur",
		B: 5,
	}
	m, err = NewMap(anInstance)
	require.Nil(t, err)
	assert.Equal(t, "amet consectetur", m.MustGet("A").String())
	assert.Equal(t, 5, m.MustGet("B").MustInt())

	// handle pointer to struct
	m, err = NewMap(&anInstance)
	require.Nil(t, err)
	assert.Equal(t, "amet consectetur", m.MustGet("A").String())
	assert.Equal(t, 5, m.MustGet("B").MustInt())

}
