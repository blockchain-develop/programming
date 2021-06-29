package problem10

import (
	"sort"
	"strings"
)

type MyWord struct {
	Word string
	Frequent int
}

type MyWordSlice []*MyWord

func (this MyWordSlice) Len() int {
	return len(this)
}
func (this MyWordSlice) Swap(i, j int) {
	key := this[i]
	this[i] = this[j]
	this[j] = key
}
func (this MyWordSlice) Less(i, j int) bool {
	if this[i].Frequent > this[j].Frequent {
		return true
	}
	if this[i].Frequent < this[j].Frequent {
		return false
	}
	return strings.Compare(this[i].Word, this[j].Word) == -1
}

func topKFrequent(words []string, k int) []string {
	keep := make(map[string]int, 0)
	for _, word := range words {
		_, ok := keep[word]
		if ok {
			keep[word] = keep[word] + 1
		} else {
			keep[word] = 1
		}
	}
	result := make(MyWordSlice, 0)
	for k, v := range keep {
		result = append(result, &MyWord{
			Word:     k,
			Frequent: v,
		})
	}
	sort.Sort(result)
	count := 0
	kWords := make([]string, 0)
	for _, item := range result {
		//fmt.Printf("work: %s, frequent: %d\n", item.Word, item.Frequent)
		kWords = append(kWords, item.Word)
		count ++
		if count >= k {
			break
		}
	}
	return kWords
}
