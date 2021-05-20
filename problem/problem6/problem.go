package problem6

func find132pattern(nums []int) bool {
	numsLen := len(nums)
	if numsLen < 3 {
		return false
	}
	leftMin := nums[0]
	for i := 1;i < numsLen;i ++ {
		if nums[i] <= leftMin {
			leftMin = nums[i]
			continue
		}
		for j := i + 1;j < numsLen;j ++ {
			if nums[j] < nums[i] && nums[j] > leftMin {
				return true
			}
		}
	}
	return false
}