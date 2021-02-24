package longestnorepeatstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubstring_Native1(t *testing.T) {
	res := LongestSubstring_Native("abcabcbb")
	assert.Equal(t, 3, len(res))
}

func TestLongestSubstring_Native2(t *testing.T) {
	res := LongestSubstring_Native("bbbbb")
	assert.Equal(t, 1, len(res))
}

func TestLongestSubstring_Native3(t *testing.T) {
	res := LongestSubstring_Native("pwwkew")
	assert.Equal(t, 3, len(res))
}

func TestLongestSubstring_Native4(t *testing.T) {
	res := LongestSubstring_Native("")
	assert.Equal(t, 0, len(res))
}

func TestLongestSubstring_Native5(t *testing.T) {
	res := LongestSubstring_Native("abcdbefgh")
	assert.Equal(t, 7, len(res))
}
