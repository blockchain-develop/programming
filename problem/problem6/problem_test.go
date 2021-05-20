package problem6

import (
	"fmt"
	"testing"
)

func TestFind132pattern1(t *testing.T) {
	nums := []int{1,2,3,4}
	{
		result := find132pattern(nums)
		fmt.Printf("result: %v\n", result)
	}
}

func TestFind132pattern2(t *testing.T) {
	nums := []int{3,1,4,2}
	{
		result := find132pattern(nums)
		fmt.Printf("result: %v\n", result)
	}
}

func TestFind132pattern3(t *testing.T) {
	nums := []int{-1,3,2,0}
	{
		result := find132pattern(nums)
		fmt.Printf("result: %v\n", result)
	}
}

