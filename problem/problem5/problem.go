package problem5

func twoSum(nums []int, target int) (int, int) {
	numsLen := len(nums)
	for i := 0;i < numsLen;i ++ {
		for j := i + 1;j < numsLen;j ++ {
			if nums[i] + nums[j] == target {
				return i, j
			}
		}
	}
	return -1, -1
}

func twoSum1(nums []int, target int) (int, int) {
	data := make(map[int]int)
	for i, num := range nums {
		j, ok := data[target-num]
		if ok {
			return i, j
		} else {
			data[num] = i
		}
	}
	return -1, -1
}
