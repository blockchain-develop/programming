package problem13

func letterCombinations(digits string) []string {
	maps := make(map[uint8][]uint8, 0)
	maps['2'] = []uint8{'a', 'b', 'c'}
	maps['3'] = []uint8{'d', 'e', 'f'}
	maps['4'] = []uint8{'g', 'h', 'i'}
	maps['5'] = []uint8{'j', 'k', 'l'}
	maps['6'] = []uint8{'m', 'n', 'o'}
	maps['7'] = []uint8{'p', 'q', 'r', 's'}
	maps['8'] = []uint8{'t', 'u', 'v'}
	maps['9'] = []uint8{'w', 'x', 'y', 'z'}
	return letterCombinations_it(digits, 0, maps)
}

func letterCombinations_it(digits string, index int, maps map[uint8][]uint8) []string {
	if index >= len(digits) {
		return []string{""}
	}
	digit := digits[index]
	results := make([]string, 0)
	for _, item := range maps[digit] {
		datas := letterCombinations_it(digits, index + 1, maps)
		for _, data := range datas {
			results = append(results, string(item) + data)
		}
	}
	return results
}
