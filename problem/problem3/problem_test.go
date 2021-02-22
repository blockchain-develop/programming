package problem3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProblem1(t *testing.T) {
	res := LongestSubstring("abcabcbb")
	assert.Equal(t, 3, len(res))
}

func TestProblem2(t *testing.T) {
	res := LongestSubstring("bbbbb")
	assert.Equal(t, 1, len(res))
}

func TestProblem3(t *testing.T) {
	res := LongestSubstring("pwwkew")
	assert.Equal(t, 3, len(res))
}

func TestProblem4(t *testing.T) {
	res := LongestSubstring("")
	assert.Equal(t, 0, len(res))
}

func TestProblem5(t *testing.T) {
	res := LongestSubstring("abcdbefgh")
	assert.Equal(t, 7, len(res))
}
