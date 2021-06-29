package problem14

import (
	"fmt"
	"testing"
)

func TestFindTargetSumWays1(t *testing.T) {
	data := []int{1,1,1,1,1}
	count := findTargetSumWays(data, 3)
	fmt.Printf("count: %d\n", count)
}

func TestFindTargetSumWays2(t *testing.T) {
	data := []int{1}
	count := findTargetSumWays(data, 1)
	fmt.Printf("count: %d\n", count)
}

func TestFindTargetSumWays3(t *testing.T) {
	data := []int{1,1,1,1,1}
	count := findTargetSumWays(data, 3)
	fmt.Printf("count: %d\n", count)
}


