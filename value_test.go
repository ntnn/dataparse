package dataparse

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Fuzz implements the interface exepcted by gofuzz to fuzz itself.
//
// Using testing-native fuzzing only allows base types to be fuzzed and
// can only be used by either writing or generating the relevant tests
// with explicit types.
//
// fuzz on the other hand doesn't handle any/interface types and panics
// when asked to fuzz them.
//
// By implementing the gofuzz self-fuzzing interface tests can be run
// multiple types.
func (v *Value) Fuzz(c fuzz.Continue) {
	switch c.Intn(18) {
	case 0:
		v.Data = fmt.Sprintf("%d", c.Int())
	case 1:
		v.Data = int8(c.Intn(math.MaxInt8))
	case 2:
		v.Data = int8(-c.Intn(math.MaxInt8))
	case 3:
		v.Data = int16(c.Intn(math.MaxInt16))
	case 4:
		v.Data = int16(-c.Intn(math.MaxInt16))
	case 5:
		v.Data = int32(c.Intn(math.MaxInt32))
	case 6:
		v.Data = int32(-c.Intn(math.MaxInt32))
	case 7:
		v.Data = int64(c.Intn(math.MaxInt64))
	case 8:
		v.Data = int64(-c.Intn(math.MaxInt64))
	case 9:
		v.Data = uint8(c.Intn(math.MaxUint8))
	case 10:
		v.Data = uint16(c.Intn(math.MaxUint16))
	case 11:
		v.Data = c.Uint32()
	case 12:
		v.Data = c.Uint64()
	case 13:
		v.Data = byte(c.Intn(math.MaxInt8))
	case 14:
		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, uint64(c.Int()))
		v.Data = bs
	case 15:
		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, c.Uint64())
		v.Data = bs
	case 16:
		v.Data = c.RandBool()
	case 17:
		v.Data = c.Float32()
	case 18:
		v.Data = c.Float64()
	}
}

type testValueToConst int

const (
	testValueToConst1 testValueToConst = iota
	testValueToConst2
	testValueToConst3
	testValueToConst4
	testValueToConst5
)

func (c *testValueToConst) From(v Value) error {
	i, err := v.Int()
	if err != nil {
		return err
	}

	*c = testValueToConst(i)
	return nil
}

func TestValue_To(t *testing.T) {
	// test passing pointer
	v := NewValue("test")
	s := ""
	require.Nil(t, v.To(&s))
	assert.Equal(t, "test", s)

	// test passing pointer of pointer
	var s2 *string
	require.Nil(t, v.To(&s2))
	assert.Equal(t, "test", *s2)

	// test some random types
	v = NewValue(5)
	var i int = 0
	require.Nil(t, v.To(&i))
	assert.Equal(t, 5, i)

	v = NewValue("5")
	i = 0
	require.Nil(t, v.To(&i))
	assert.Equal(t, 5, i)

	v = NewValue(uint(math.MaxInt64) + 5)
	var u uint = 0
	require.Nil(t, v.To(&u))
	assert.Equal(t, uint(math.MaxInt64)+5, u)

	v = NewValue("192.168.1.1")
	var ip net.IP
	require.Nil(t, v.To(&ip))
	assert.Equal(t, net.ParseIP("192.168.1.1"), ip)

	v = NewValue(2)
	var c testValueToConst
	require.Nil(t, v.To(&c))
	assert.Equal(t, testValueToConst3, c)
}

func TestValue_List(t *testing.T) {
	v, err := NewValue([]int{1, 2, 3}).List()
	assert.Nil(t, err)
	assert.Equal(t,
		[]Value{
			Value{Data: 1},
			Value{Data: 2},
			Value{Data: 3},
		},
		v,
	)

	v, err = NewValue([]any{1, "test", 3.56}).List()
	assert.Nil(t, err)
	assert.Equal(t,
		[]Value{
			Value{Data: 1},
			Value{Data: "test"},
			Value{Data: 3.56},
		},
		v,
	)
}
