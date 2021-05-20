package problem4

func findMedianSortedArrays1(nums1 []int, nums2 []int) float32 {
	nums := make([]int, 0, len(nums1) + len(nums2))
	i,j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			nums = append(nums, nums1[i])
			i ++
		} else {
			nums = append(nums, nums2[j])
			j ++
		}
	}
	for i < len(nums1) {
		nums = append(nums, nums1[i])
		i ++
	}
	for j < len(nums2) {
		nums = append(nums, nums2[j])
		j ++
	}
	total := len(nums1) + len(nums2)
	if total % 2 == 0 {
		return float32(nums[total / 2 - 1] + nums[total / 2]) / 2
	} else {
		return float32(nums[total / 2])
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float32 {
	return 0
}
