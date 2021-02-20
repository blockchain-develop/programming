package utils

import (
	"math"
)

type Queue struct {
	items []*BSTNode
}

func NewQueue() *Queue {
	queue := &Queue{
		items : make([]*BSTNode, 0),
	}
	return queue
}

func (q *Queue) Push(node *BSTNode) {
	q.items = append(q.items, node)
}

func (q *Queue) Pop() *BSTNode {
	if len(q.items) == 0 {
		return nil
	}
	node := q.items[0]
	q.items = q.items[1:]
	return node
}

func (q *Queue) Size() int {
	return len(q.items)
}

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
