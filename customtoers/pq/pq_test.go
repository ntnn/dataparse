package custompq

import (
	"testing"

	"github.com/lib/pq"
	"github.com/ntnn/dataparse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomTo_StringArray(t *testing.T) {
	v := dataparse.NewValue([]string{"a", "b", "c"})
	o := pq.StringArray{}

	require.Nil(t, v.To(&o))
	assert.Equal(t, "c", o[2])
}
