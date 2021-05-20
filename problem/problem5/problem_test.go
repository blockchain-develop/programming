package problem5

import (
	"fmt"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	nums := []int{2,7,11,15}
	target := 9
	{
		i, j := twoSum(nums, target)
		fmt.Printf("result: (%d, %d)\n", i, j)
	}
	{
		i, j := twoSum1(nums, target)
		fmt.Printf("result: (%d, %d)\n", i, j)
	}
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3,2,4}
	target := 6
	{
		i, j := twoSum(nums, target)
		fmt.Printf("result: (%d, %d)\n", i, j)
	}
	{
		i, j := twoSum1(nums, target)
		fmt.Printf("result: (%d, %d)\n", i, j)
	}
}

func TestTwoSum3(t *testing.T) {
	nums := []int{3,3}
	target := 6
	{
		i, j := twoSum(nums, target)
		fmt.Printf("result: (%d, %d)\n", i, j)
	}
	{
		i, j := twoSum1(nums, target)
		fmt.Printf("result: (%d, %d)\n", i, j)
	}
}
