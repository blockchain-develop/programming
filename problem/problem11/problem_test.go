package problem11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindromicSubstring_DP1(t *testing.T) {
	data :="hello"
	sub := longestPalindrome(data)
	assert.Equal(t, len(sub), 2)
}

func TestLongestPalindromicSubstring_DP2(t *testing.T) {
	data := "ABCDZJUDCBA"
	sub := longestPalindrome(data)
	assert.Equal(t,len(sub), 1)
}

func TestLongestPalindromicSubstring_DP3(t *testing.T) {
	data := "PATZJUJZTACCBCC"
	sub := longestPalindrome(data)
	assert.Equal(t,len(sub), 9)
}

func TestLongestPalindromicSubstring_DP4(t *testing.T) {
	data := "ABCDEEDCBA"
	sub := longestPalindrome(data)
	assert.Equal(t,len(sub), 10)
}

func TestLongestPalindromicSubstring_DP5(t *testing.T) {
	data := "ll"
	sub := longestPalindrome(data)
	assert.Equal(t,len(sub), 2)
}

func TestLongestPalindromicSubstring_DP6(t *testing.T) {
	data := "ab"
	sub := longestPalindrome(data)
	assert.Equal(t,len(sub), 1)
}

func TestLongestPalindromicSubstring_DP7(t *testing.T) {
	data := "a"
	sub := longestPalindrome(data)
	assert.Equal(t,len(sub), 1)
}

func TestLongestPalindromicSubstring_DP8(t *testing.T) {
	data := "baefeab"
	sub := longestPalindrome(data)
	assert.Equal(t,len(sub), 7)
}

func TestLongestPalindromicSubstring_DP9(t *testing.T) {
	data := "forgeeksskeegfor"
	sub := longestPalindrome(data)
	assert.Equal(t,len(sub), 10)
}

