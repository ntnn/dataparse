package dataparse

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseTime(t *testing.T) {
	date := gofakeit.Date()
	s := date.Format(time.RFC3339Nano)
	result, err := ParseTime(s)
	require.Nil(t, err)
	assert.Equal(t, date, result)
}
