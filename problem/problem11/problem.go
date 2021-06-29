package problem11

func longestPalindrome(s string) string {
	len := len(s)
	keeps := make([][]int, len)
	for i := 0;i < len;i ++ {
		keeps[i] = make([]int, len)
	}
	for i := 0;i < len;i ++ {
		keeps[i][i] = 1
	}
	for i := 0;i < len;i ++ {
		for j := i + 1;j < len && 2*i - j > 0; j ++ {
			if s[j] == s[2 * i - j] && keeps[i][j - 1] != 0 {
				keeps[i][j] = keeps[i][j - 1] + 2
			}
		}
	}
	max := 0
	index := 0
	for i := 0;i < len;i ++ {
		for j := 0;j < len;j ++ {
			if keeps[i][j] > max {
				max = keeps[i][j]
				index = j
			}
		}
	}
	return s[index - max + 1 : index]
}
