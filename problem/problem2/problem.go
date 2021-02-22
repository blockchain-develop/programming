package problem2

func minOperations(nums []int, start int, end int, x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	a := minOperations(nums, start + 1, end, x - nums[start])
	b := minOperations(nums, start, end - 1, x - nums[end])
	r := 0
	if a == -1 && b == -1 {
		r = -1
	} else if a != -1 && b == -1 {
		r = a + 1
	} else if a == -1 && b != -1 {
		r = b + 1
	} else if a < b {
		r = a + 1
	} else {
		r = b + 1
	}
	return r
}

func MinOperations(nums []int, x int) int {
	return minOperations(nums, 0, len(nums) - 1, x)
}
