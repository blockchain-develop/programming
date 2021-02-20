package maximumsubarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumSubarray_DP1(t *testing.T) {
	data := [...]int{1, 2, 3, 4, 5}
	max := MaximumSubarray_DP(data[:])
	assert.Equal(t, max, 15)
}

func TestMaximumSubarray_DP2(t *testing.T) {
	data := [...]int{-1, -2, 3, -4, -5}
	max := MaximumSubarray_DP(data[:])
	assert.Equal(t, max, 3)
}

func TestMaximumSubarray_DP3(t *testing.T) {
	data := [...]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	max:= MaximumSubarray_DP(data[:])
	assert.Equal(t, max, 6)
}

func TestMaximumSubarray_DP4(t *testing.T) {
	data := [...]int{-1, -2, -3, -4, -5}
	max := MaximumSubarray_DP(data[:])
	assert.Equal(t, -1, max)
}
