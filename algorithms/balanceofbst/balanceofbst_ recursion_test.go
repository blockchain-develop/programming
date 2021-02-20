package balanceofbst

import (
	"github.com/programming/utils"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestIsBalance_Recursion1(t *testing.T) {
	data := []int{3, 9, 20, int(math.MinInt32), int(math.MinInt32), 15, 7}
	bst := utils.Slice2BST(data)
	result := IsBalance_Down2Up(bst)
	assert.Equal(t, result, true)
}

func TestIsBalance_Recursion2(t *testing.T) {
	data := []int{1, 2, 2, 3, 3, int(math.MinInt32), int(math.MinInt32), 4, 4, int(math.MinInt32), int(math.MinInt32), int(math.MinInt32), int(math.MinInt32), int(math.MinInt32), int(math.MinInt32)}
	bst := utils.Slice2BST(data)
	result := IsBalance_Down2Up(bst)
	assert.Equal(t, result, false)
}

func TestIsBalance_Recursion3(t *testing.T) {
	data := []int{3, 9, 20, int(math.MinInt32), int(math.MinInt32), 15, 7}
	bst := utils.Slice2BST(data)
	result := IsBalance_Up2Down(bst)
	assert.Equal(t, result, true)
}

func TestIsBalance_Recursion4(t *testing.T) {
	data := []int{1, 2, 2, 3, 3, int(math.MinInt32), int(math.MinInt32), 4, 4, int(math.MinInt32), int(math.MinInt32), int(math.MinInt32), int(math.MinInt32), int(math.MinInt32), int(math.MinInt32)}
	bst := utils.Slice2BST(data)
	result := IsBalance_Up2Down(bst)
	assert.Equal(t, result, false)
}

