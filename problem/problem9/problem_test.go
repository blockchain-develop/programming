package problem9

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	result := lengthOfLongestSubstring(s)
	fmt.Printf("result: %s, len: %d\n", result, len(result))
}

func TestLengthOfLongestSubstring1(t *testing.T) {
	s := "bbbbb"
	result := lengthOfLongestSubstring(s)
	fmt.Printf("result: %s, len: %d\n", result, len(result))
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	s := "pwwkew"
	result := lengthOfLongestSubstring(s)
	fmt.Printf("result: %s, len: %d\n", result, len(result))
}
