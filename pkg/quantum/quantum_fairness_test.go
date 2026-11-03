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

func TestUnfairnessAndQuality(t *testing.T) {
	q := New()
	assert.Zero(t, q.Unfairness([]float64{2, 2, 2}))
	assert.InDelta(t, 0.5, q.Unfairness([]float64{0, 4}), 1e-12)
	assert.Zero(t, q.Unfairness(nil))

	_, err := q.AdjustNodes(1, 1)
	require.NoError(t, err)
	state, err := q.EntangleNodes(nil)
	require.NoError(t, err)

	assert.InDelta(t, 1, q.Quality(state, []float64{3, 3}), 1e-12)
	assert.InDelta(t, 0.5, q.Quality(state, []float64{0, 4}), 1e-12)
	assert.Zero(t, q.Quality(nil, []float64{3, 3}))
}
