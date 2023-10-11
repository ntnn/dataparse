package dataparse

import (
	"compress/gzip"
	"io"
	"net"
	"os"
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
	assert.Equal(t, "Garrott", m.MustGet("first_name").MustString())

	m = <-maps
	require.NotNil(t, m)
	assert.Equal(t, "Vasovic", m.MustGet("last_name").MustString())

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

func gzipFile(t *testing.T, target string) {
	// Skip zipping a file that was already zipped
	if _, err := os.Stat(target + ".gz"); err == nil {
		return
	}

	out, err := os.Create(target + ".gz")
	require.Nil(t, err)
	defer func() {
		require.Nil(t, out.Close())
	}()

	gzipOut := gzip.NewWriter(out)
	defer func() {
		assert.Nil(t, gzipOut.Flush())
		require.Nil(t, gzipOut.Close())
	}()

	reader, err := os.Open(target)
	require.Nil(t, err)

	_, err = io.Copy(gzipOut, reader)
	require.Nil(t, err)
}

func TestFrom_JsonGz(t *testing.T) {
	gzipFile(t, "./testdata/data.json")

	maps, errs, err := From("./testdata/data.json.gz")
	require.Nil(t, err)
	for err := range errs {
		require.Nil(t, err)
	}

	m := <-maps
	require.NotNil(t, m)
	assert.Equal(t, "Garrott", m.MustGet("first_name").MustString())

	m = <-maps
	require.NotNil(t, m)
	assert.Equal(t, "Vasovic", m.MustGet("last_name").MustString())

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

	assert.Equal(t, "Garrott", m.MustGet("first_name").MustString())
	assert.Equal(t, "Felgate", m.MustGet("last_name").MustString())
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
	assert.Equal(t, "Garrott", m.MustGet("first_name").MustString())

	m = <-maps
	require.NotNil(t, m)
	assert.Equal(t, "Vasovic", m.MustGet("last_name").MustString())

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
	assert.Equal(t, "Robbie", m.MustGet("first_name").MustString())

	m = <-maps
	require.NotNil(t, m)
	assert.Equal(t, "Salvin", m.MustGet("last_name").MustString())

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
	assert.Equal(t, "test", m.MustGet("b").MustString())
	assert.True(t, m.MustGet("c").IsNil())

	m, err = FromKVString("key1=value1|key2=value2", WithSeparator("|"))
	require.Nil(t, err)
	assert.Equal(t, "value1", m.MustGet("key1").MustString())
	assert.Equal(t, "value2", m.MustGet("key2").MustString())
}

func TestNewMap(t *testing.T) {
	// convert maps to Map
	m, err := NewMap(map[any]any{1: 1})
	require.Nil(t, err)
	assert.Equal(t, 1, m.MustGet(1).MustInt())

	m, err = NewMap(map[int]string{1: "lorem ipsum"})
	require.Nil(t, err)
	assert.Equal(t, "lorem ipsum", m.MustGet(1).MustString())

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
	assert.Equal(t, "amet consectetur", m.MustGet("A").MustString())
	assert.Equal(t, 5, m.MustGet("B").MustInt())

	// handle pointer to struct
	m, err = NewMap(&anInstance)
	require.Nil(t, err)
	assert.Equal(t, "amet consectetur", m.MustGet("A").MustString())
	assert.Equal(t, 5, m.MustGet("B").MustInt())

}

func TestMap_To(t *testing.T) {
	type testStruct struct {
		A int
		B string
		C string
		D uint32
	}

	m, err := NewMap(map[string]any{
		"A": 5,
		"B": "lorem ipsum",
		"C": 15,
		"D": "6622",
	})
	require.Nil(t, err)

	var ts testStruct
	require.Nil(t, m.To(&ts))
	assert.Equal(t, 5, ts.A)
	assert.Equal(t, "lorem ipsum", ts.B)
	assert.Equal(t, "15", ts.C)
	assert.Equal(t, uint32(6622), ts.D)

	var ts2 testStruct
	require.NotNil(t, m.To(ts2))
}

func TestMap_To_Tag(t *testing.T) {
	type testStruct struct {
		Number  int    `dataparse:"a"`
		Message string `dataparse:"b"`
		Varying string `dataparse:"varying,second_varying"`
	}

	m, err := NewMap(map[string]any{
		"a":       5,
		"b":       "lorem ipsum",
		"varying": "sic dolor amet",
	})
	require.Nil(t, err)

	var ts testStruct
	require.Nil(t, m.To(&ts))
	assert.Equal(t, 5, ts.Number)
	assert.Equal(t, "lorem ipsum", ts.Message)
	assert.Equal(t, "sic dolor amet", ts.Varying)

	m2, err := NewMap(map[string]any{
		"a":              5,
		"b":              "lorem ipsum",
		"second_varying": "sic dolor amet",
	})
	require.Nil(t, err)

	var ts2 testStruct
	require.Nil(t, m2.To(&ts2))
	assert.Equal(t, 5, ts2.Number)
	assert.Equal(t, "lorem ipsum", ts2.Message)
	assert.Equal(t, "sic dolor amet", ts2.Varying)
}

func TestMap_To_Embedded(t *testing.T) {
	type testStruct struct {
		A int
		B string `dataparse:"msg"`
		C struct {
			CA int
			CB string `dataparse:"submsg"`
		} `dataparse:"sub"`
	}

	m, err := NewMap(map[string]any{
		"A":   3,
		"msg": "outer",
		"sub": map[string]any{
			"CA":     15,
			"submsg": "inner",
		},
	})
	require.Nil(t, err)

	var ts testStruct
	require.Nil(t, m.To(&ts))
	assert.Equal(t, 3, ts.A)
	assert.Equal(t, "outer", ts.B)
	assert.Equal(t, 15, ts.C.CA)
	assert.Equal(t, "inner", ts.C.CB)
}

func TestMap_To_DotNotation(t *testing.T) {
	type testStruct struct {
		A int    `dataparse:"a.b"`
		B string `dataparse:"msg.short"`
	}

	m, err := NewMap(map[string]any{
		"a": map[string]any{
			"b": 5,
		},
		"msg": map[string]any{
			"short": "lorem ipsum",
		},
	})
	require.Nil(t, err)

	var ts testStruct
	require.Nil(t, m.To(&ts))
	assert.Equal(t, 5, ts.A)
	assert.Equal(t, "lorem ipsum", ts.B)
}
