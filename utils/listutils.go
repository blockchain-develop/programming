package utils

import (
	"container/list"
	"fmt"
	"strings"
)

func Slice2List(nums []int) *list.List {
	l := list.New()
	for _, num := range nums {
		l.PushBack(num)
	}
	return l
}

func PrintList(l *list.List) string {
	data := make([]string, 0, l.Len())
	for e := l.Front();e != nil;e = e.Next() {
		data = append(data, fmt.Sprintf("%d", e.Value.(int)))
	}
	result := strings.Join(data, " ")
	return result
}
