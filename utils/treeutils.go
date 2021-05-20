package utils

import (
	"math"
)

type BSTNode struct {
	Left *BSTNode
	Right *BSTNode
	Data int
}

func NewBSTNode(data int) *BSTNode {
	node := &BSTNode{
		Left:  nil,
		Right: nil,
		Data:  data,
	}
	return node
}

func Slice2BST(data []int) *BSTNode {
	dataLen := len(data)
	xxx := math.Log2(float64(dataLen + 1))
	xxx = math.Ceil(xxx)
	height := int(xxx)
	fullNodeNumber := int(math.Exp2(float64(height))) - 1
	if fullNodeNumber != dataLen {
		return nil
	}
	q := NewQueue()
	root := NewBSTNode(data[0])
	node := root
	left := true
	for i := 1;i < len(data);i ++ {
		if left == true {
			var leftChild *BSTNode
			if data[i] == int(math.MinInt32) {
				leftChild = nil
			} else {
				leftChild = NewBSTNode(data[i])
			}
			q.Push(leftChild)
			if node != nil {
				node.Left = leftChild
			}
			left = false
		} else {
			var rightChild *BSTNode
			if data[i] == int(math.MinInt32) {
				rightChild = nil
			} else {
				rightChild = NewBSTNode(data[i])
			}
			q.Push(rightChild)
			if node != nil {
				node.Right = rightChild
			}
			left = true
			node = q.Pop()
		}
	}
	return root
}
