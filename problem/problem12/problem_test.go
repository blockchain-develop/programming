package problem12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi1(t *testing.T) {
	data :="42"
	sub := myAtoi(data)
	assert.Equal(t, sub, 42)
}

func TestMyAtoi2(t *testing.T) {
	data := "   -42"
	sub := myAtoi(data)
	assert.Equal(t,sub, -42)
}

func TestMyAtoi3(t *testing.T) {
	data := "4193 with words"
	sub := myAtoi(data)
	assert.Equal(t,sub, 4193)
}

func TestMyAtoi4(t *testing.T) {
	data := "words and 987"
	sub := myAtoi(data)
	assert.Equal(t,sub, 0)
}

func TestMyAtoi5(t *testing.T) {
	data := "-91283472332"
	sub := myAtoi(data)
	assert.Equal(t,sub, -2147483648)
}

