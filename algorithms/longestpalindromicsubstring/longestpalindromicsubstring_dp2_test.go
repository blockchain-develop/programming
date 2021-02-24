package longestpalindromicsubstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindromicSubstring_DP21(t *testing.T) {
	data := []byte("hello")
	sub := LongestPalindromicSubstring_DP2(data[:])
	assert.Equal(t, sub, 2)
}

func TestLongestPalindromicSubstring_DP22(t *testing.T) {
	data := []byte("ABCDZJUDCBA")
	sub := LongestPalindromicSubstring_DP2(data[:])
	assert.Equal(t, sub, 1)
}

func TestLongestPalindromicSubstring_DP23(t *testing.T) {
	data := []byte("PATZJUJZTACCBCC")
	sub := LongestPalindromicSubstring_DP2(data[:])
	assert.Equal(t, sub, 9)
}

func TestLongestPalindromicSubstring_DP24(t *testing.T) {
	data := []byte("ABCDEEDCBA")
	sub := LongestPalindromicSubstring_DP2(data[:])
	assert.Equal(t, sub, 10)
}

func TestLongestPalindromicSubstring_DP25(t *testing.T) {
	data := []byte("ll")
	sub := LongestPalindromicSubstring_DP2(data[:])
	assert.Equal(t, sub, 2)
}

func TestLongestPalindromicSubstring_DP26(t *testing.T) {
	data := []byte("ab")
	sub := LongestPalindromicSubstring_DP2(data[:])
	assert.Equal(t, sub, 1)
}

func TestLongestPalindromicSubstring_DP27(t *testing.T) {
	data := []byte("a")
	sub := LongestPalindromicSubstring_DP2(data[:])
	assert.Equal(t, sub, 1)
}

func TestLongestPalindromicSubstring_DP28(t *testing.T) {
	data := []byte("baefeab")
	sub := LongestPalindromicSubstring_DP2(data[:])
	assert.Equal(t, sub, 7)
}

func TestLongestPalindromicSubstring_DP29(t *testing.T) {
	data := []byte("forgeeksskeegfor")
	sub := LongestPalindromicSubstring_DP2(data[:])
	assert.Equal(t, sub, 10)
}