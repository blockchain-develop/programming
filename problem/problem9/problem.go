package problem9

func lengthOfLongestSubstring(s string) string {
	if len(s) <= 1 {
		return s
	}
	left := 0
	right := 0
	len := len(s)
	records := make(map[uint8]int, 0)
	longestSubstring := 0
	longestSubstringLen := 0
	for right < len {
		item := s[right]
		record, ok := records[item]
		if !ok || record == -1 {
			records[item] = right
			if right - left + 1 > longestSubstringLen {
				longestSubstringLen = right - left + 1
				longestSubstring = left
			}
			right ++
			continue
		}
		{
			for left < record {
				index := s[left]
				records[index] = -1
				left ++
			}
			left ++
			records[item] = right
			right ++

		}
	}
	return s[longestSubstring: longestSubstring + longestSubstringLen]
}
