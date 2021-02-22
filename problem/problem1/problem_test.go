package problem1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProblem1(t *testing.T) {
	n := 5
	k := 3
	relations := make([][]int, 0)
	relations = append(relations, []int{0, 2})
	relations = append(relations, []int{2, 1})
	relations = append(relations, []int{3, 4})
	relations = append(relations, []int{2, 3})
	relations = append(relations, []int{1, 4})
	relations = append(relations, []int{2, 0})
	relations = append(relations, []int{0, 4})
	ways := NumWays(n, relations, k)
	assert.Equal(t, ways, 3)
}

func TestProblem2(t *testing.T) {
	n := 3
	k := 2
	relations := make([][]int, 0)
	relations = append(relations, []int{0, 2})
	relations = append(relations, []int{2, 1})
	ways := NumWays(n, relations, k)
	assert.Equal(t, ways, 0)
}
