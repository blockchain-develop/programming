package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func MinHeapDump(mh *MinHeap) []int {
	len := mh.Len()
	data := make([]int, len)
	for len > 0 {
		data[len - 1] = mh.Remove()
		len = mh.Len()
	}
	return data
}

func TestMinHeap1(t *testing.T) {
	data := [...]int{1, 2, 3, 4, 5}
	mh := NewMinHeap()
	mh.Init(data[:])
	r := MinHeapDump(mh)
	for _, item := range r {
		fmt.Printf(" %d ", item)
	}
	fmt.Printf("\n")
	order := IsOrderly1(r[:])
	assert.Equal(t, order, byte(2))
}

func TestMinHeap2(t *testing.T) {
	data := [...]int{5, 4, 3, 2, 1}
	mh := NewMinHeap()
	mh.Init(data[:])
	r := MinHeapDump(mh)
	for _, item := range r {
		fmt.Printf(" %d ", item)
	}
	fmt.Printf("\n")
	order := IsOrderly1(r[:])
	assert.Equal(t, order, byte(2))
}

func TestMinHeap3(t *testing.T) {
	data := [...]int{5, 1, 3, 1, 2}
	mh := NewMinHeap()
	mh.Init(data[:])
	r := MinHeapDump(mh)
	for _, item := range r {
		fmt.Printf(" %d ", item)
	}
	fmt.Printf("\n")
	order := IsOrderly1(r[:])
	assert.Equal(t, order, byte(2))
}

func TestMinHeap4(t *testing.T) {
	size := 10240000
	data := make([]int, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int()
	}
	mh := NewMinHeap()
	mh.Init(data[:])
	r := MinHeapDump(mh)
	order := IsOrderly1(r[:])
	assert.Equal(t, order, byte(2))
}

func TestMinHeap5(t *testing.T) {
	mh := NewMinHeap()
	mh.Insert(0)
	mh.Insert(5)
	mh.Insert(1)
	mh.Insert(3)
	mh.Insert(4)
	mh.Insert(2)
	mh.Insert(9)
	r := MinHeapDump(mh)
	order := IsOrderly1(r[:])
	assert.Equal(t, order, byte(2))
}

func TestMinHeap6(t *testing.T) {
	size := 10240
	mh := NewMinHeap()
	for i := 0;i < size;i ++ {
		mh.Insert(rand.Int())
	}
	r := MinHeapDump(mh)
	order := IsOrderly1(r[:])
	assert.Equal(t, order, byte(2))
}
