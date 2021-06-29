package problem10

import (
	"container/heap"
	"strings"
)

type Item struct {
	value string
	priority int
	index int
}
type PriorityQueue []*Item
func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].priority > pq[j].priority {
		return true
	}
	if pq[i].priority < pq[j].priority {
		return false
	}
	return strings.Compare(pq[i].value, pq[j].value) == -1
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i],pq[j] = pq[j],pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n - 1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n - 1]
	return item
}
func (pq *PriorityQueue) Update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func topKFrequent_pq(words []string, k int) []string {
	pq := make(PriorityQueue, 0)
	items := make(map[string]*Item, 0)
	for _, word := range words {
		item, ok := items[word]
		if !ok {
			item := &Item{
				value:    word,
				priority: 1,
			}
			heap.Push(&pq, item)
			items[word] = item
		} else {
			pq.Update(item, item.value, item.priority + 1)
		}
	}
	result := make([]string, 0)
	for i := 0;i < k;i ++ {
		item := heap.Pop(&pq)
		result = append(result, item.(*Item).value)
	}
	return result
}

