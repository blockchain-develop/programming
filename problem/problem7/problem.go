package problem7

import "container/list"

func addTwoNumbers(num1 *list.List, num2 *list.List) *list.List {
	item := 0
	result := list.New()
	element1 := num1.Front()
	element2 := num2.Front()
	for true {
		if element1 == nil && element2 == nil {
			if item > 0 {
				result.PushBack(item)
			}
			break
		}
		item1, item2 := 0, 0
		if element1 != nil {
			item1 = element1.Value.(int)
			element1 = element1.Next()
		}
		if element2 != nil {
			item2 = element2.Value.(int)
			element2 = element2.Next()
		}
		item = item1 + item2 + item
		item1 = item % 10
		item2 = item / 10
		item = item2
		result.PushBack(item1)
	}
	return result
}