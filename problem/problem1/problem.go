package problem1

func numWays(n int, start int, end int, relations [][]int, k int) int {
	if k == 0 {
		if start == end {
			return 1
		} else {
			return 0
		}
	}
	total := 0
	for _, relation := range relations {
		if relation[0] == start {
			total += numWays(n, relation[1], end, relations, k - 1)
		}
	}
	return total
}

func NumWays(n int, relations [][]int, k int) int {
	return numWays(n, 0, n - 1, relations, k)
}
