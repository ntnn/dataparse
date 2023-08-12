package dataparse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
