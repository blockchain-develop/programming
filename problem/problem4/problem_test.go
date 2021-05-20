package problem4

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays1(t *testing.T) {
	nums1 := []int{1,3}
	nums2 := []int{2}
	{
		result := findMedianSortedArrays1(nums1, nums2)
		fmt.Printf("result: %f\n", result)
	}
}

func TestFindMedianSortedArrays2(t *testing.T) {
	nums1 := []int{1,2}
	nums2 := []int{3,4}
	{
		result := findMedianSortedArrays1(nums1, nums2)
		fmt.Printf("result: %f\n", result)
	}
}

func TestTwoSum3(t *testing.T) {
	nums1 := []int{0,0}
	nums2 := []int{0,0}
	{
		result := findMedianSortedArrays1(nums1, nums2)
		fmt.Printf("result: %f\n", result)
	}
}

func TestTwoSum4(t *testing.T) {
	nums1 := []int{}
	nums2 := []int{1}
	{
		result := findMedianSortedArrays1(nums1, nums2)
		fmt.Printf("result: %f\n", result)
	}
}
