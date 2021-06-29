package problem12

func myAtoi(s string) int {
	len := len(s)
	flag := 0
	data := 0
	start := 0
	for i := 0;i < len; i ++ {
		item := s[i]
		if item >= '0' && item <= '9' {
			if flag == 0 {
				flag = 1
			}
			if start == 0 {
				start = 1
			}
			data = data * 10 + int(item - '0')
			continue
		}
		if item == '-' && flag == 0 {
			flag = -1
			start = 1
			continue
		}
		if item == '+' && flag == 0 {
			flag = 1
			start = 1
			continue
		}
		if item == ' ' && start == 0 {
			continue
		}
		break
	}
	if flag == -1 {
		data = 0 - data
	}
	return data
}
