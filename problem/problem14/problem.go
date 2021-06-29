package problem14

func findTargetSumWays(nums []int, target int) int {
	results := findTargetSumWays_it(nums, 0)
	count := 0
	for _, result := range results {
		if result == target {
			count ++
		}
	}
	return count
}

func findTargetSumWays_it(nums []int, index int) []int {
	if index == len(nums) {
		return []int{0}
	}
	results := findTargetSumWays_it(nums, index + 1)
	newResults := make([]int, 0)
	for _, result := range results {
		newResults = append(newResults, nums[index] + result)
		newResults = append(newResults, 0 - nums[index] + result)
	}
	return newResults
}
