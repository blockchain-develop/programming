package problem10

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	works := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	result := topKFrequent(works, k)
	fmt.Printf("%d words: ", k)
	for _, item := range result {
		fmt.Printf("%s, ", item)
	}
	fmt.Printf("\n")
}

func TestTopKFrequent1(t *testing.T) {
	works := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	k := 4
	result := topKFrequent(works, k)
	fmt.Printf("%d words: ", k)
	for _, item := range result {
		fmt.Printf("%s, ", item)
	}
	fmt.Printf("\n")
}

