package quantum

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadjustAndFairnessEquality(t *testing.T) {
	q := New()
	loads := []float64{1, 5, 9, 1}

	ok, _ := q.FairnessEquality(loads)
	assert.False(t, ok)

	out, err := q.Readjust(loads)
	require.NoError(t, err)
	assert.Equal(t, []float64{4, 4, 4, 4}, out)

	ok, dev := q.FairnessEquality(out)
	assert.True(t, ok)
	assert.Zero(t, dev)
}

func TestReadjustErrors(t *testing.T) {
	q := New()
	_, err := q.Readjust(nil)
	require.Error(t, err)
	_, err = q.Readjust([]float64{1, -1})
	require.Error(t, err)
}
