package problem2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProblem1(t *testing.T) {
	x := 5
	nums := []int{1, 1, 4, 2, 3}
	ops := MinOperations(nums, x)
	assert.Equal(t, 2, ops)
}

func TestProblem2(t *testing.T) {
	x := 4
	nums := []int{5, 6, 7, 8, 9}
	ops := MinOperations(nums, x)
	assert.Equal(t, -1, ops)
}

func TestProblem3(t *testing.T) {
	x := 10
	nums := []int{3,2,20,1,1,3}
	ops := MinOperations(nums, x)
	assert.Equal(t, 5, ops)
}
