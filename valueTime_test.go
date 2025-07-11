package dataparse

import (
	"math"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestValue_Time(t *testing.T) {
	i8 := gofakeit.Int8()
	v, err := NewValue(i8).Time()
	require.Nil(t, err)
	require.Equal(t, i8, int8(v.Unix()))

	i32 := gofakeit.Int32()
	v, err = NewValue(i32).Time()
	require.Nil(t, err)
	require.Equal(t, i32, int32(v.Unix()))

	i64 := gofakeit.Int64()
	v, err = NewValue(i64).Time()
	require.Nil(t, err)
	require.Equal(t, i64, v.Unix())

	u8 := gofakeit.Uint8()
	v, err = NewValue(u8).Time()
	require.Nil(t, err)
	require.Equal(t, u8, uint8(v.Unix()))

	u64 := uint64(gofakeit.UintRange(0, math.MaxInt64))
	v, err = NewValue(u64).Time()
	require.Nil(t, err)
	require.Equal(t, u64, uint64(v.Unix()))

	u64 = uint64(gofakeit.UintRange(math.MaxInt64, math.MaxUint64))
	_, err = NewValue(u64).Time()
	require.NotNil(t, err)
}
