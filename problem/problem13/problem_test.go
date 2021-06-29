package problem13

import (
	"fmt"
	"testing"
)

func TestLetterCombinations1(t *testing.T) {
	data := "23"
	subs := letterCombinations(data)
	for _, sub := range subs {
		fmt.Printf("%s ", sub)
	}
	fmt.Printf("\n")
}

func TestLetterCombinations2(t *testing.T) {
	data := ""
	subs := letterCombinations(data)
	for _, sub := range subs {
		fmt.Printf("%s ", sub)
	}
	fmt.Printf("\n")
}

func TestLetterCombinations3(t *testing.T) {
	data := "2"
	subs := letterCombinations(data)
	for _, sub := range subs {
		fmt.Printf("%s ", sub)
	}
	fmt.Printf("\n")
}


