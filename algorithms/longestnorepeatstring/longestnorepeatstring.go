package longestnorepeatstring

import "strings"

func longestnorepeatstring_dp(data []byte) []byte {
	dataLen := len(data)
	keep := make([][]int, dataLen)
	for i, _ := range keep {
		keep[i] = make([]int, dataLen)
	}
	for i := 0;i < dataLen;i ++ {
		for j := i + 1;j < dataLen;j ++ {
			if keep[i][j - 1] == 0 && !strings.Contains(string(data[i:j]), string(data[j:j + 1])) {
				keep[i][j] = 0
			} else {
				keep[i][j] = 1
			}
		}
	}
	for m := dataLen - 1; m >= 0;m -- {
		for i := 0;i < dataLen && i + m < dataLen;i ++ {
			if keep[i][i + m] == 0 {
				return data[i:i + m + 1]
			}
		}
	}
	return []byte{}
}

func LongestSubstring_Native(s string) string {
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
