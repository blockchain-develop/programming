package longestnorepeatstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestnorepeatstring_dp1(t *testing.T) {
	res := longestnorepeatstring_dp([]byte("abcabcbb"))
	assert.Equal(t, 3, len(res))
}

func TestLongestnorepeatstring_dp2(t *testing.T) {
	res := longestnorepeatstring_dp([]byte("bbbbb"))
	assert.Equal(t, 1, len(res))
}

func TestLongestnorepeatstring_dp3(t *testing.T) {
	res := longestnorepeatstring_dp([]byte("pwwkew"))
	assert.Equal(t, 3, len(res))
}

func TestLongestnorepeatstring_dp4(t *testing.T) {
	res := longestnorepeatstring_dp([]byte(""))
	assert.Equal(t, 0, len(res))
}

func TestLongestnorepeatstring_dp5(t *testing.T) {
	res := longestnorepeatstring_dp([]byte("abcdbefgh"))
	assert.Equal(t, 7, len(res))
}
