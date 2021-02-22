package problem3

import "strings"

func LongestSubstring(s string) string {
	substring_so_far := ""
	substring_end_here := ""
	for i := 0;i < len(s);i ++ {
		item := s[i:i + 1]
		index := strings.Index(substring_end_here, item)
		if index != -1 {
			substring_end_here = substring_end_here[index + 1:len(substring_end_here)]
			substring_end_here += item
		} else {
			substring_end_here += item
		}
		if len(substring_so_far) < len(substring_end_here) {
			substring_so_far = substring_end_here
		}
	}
	return substring_so_far
}

