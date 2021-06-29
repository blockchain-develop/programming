package problem10

import (
	"fmt"
	"testing"
)

func TestTopKFrequent_pq(t *testing.T) {
	works := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	result := topKFrequent_pq(works, k)
	fmt.Printf("%d words: ", k)
	for _, item := range result {
		fmt.Printf("%s, ", item)
	}
	fmt.Printf("\n")
}

func TestTopKFrequent_pq1(t *testing.T) {
	works := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	k := 4
	result := topKFrequent_pq(works, k)
	fmt.Printf("%d words: ", k)
	for _, item := range result {
		fmt.Printf("%s, ", item)
	}
	fmt.Printf("\n")
}

