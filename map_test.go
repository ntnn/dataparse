package dataparse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMap(t *testing.T) {
	m, err := NewMap(map[any]any{1: 1})
	require.Nil(t, err)
	assert.Equal(t, 1, m.MustGet(1).MustInt())

	m, err = NewMap(map[any]any{1: "test"})
	require.Nil(t, err)
	assert.Equal(t, "test", m.MustGet(1).MustString())

	_, err = NewMap(nil)
	require.NotNil(t, err)

	_, err = NewMap("test")
	require.NotNil(t, err)
}
