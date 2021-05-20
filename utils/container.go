package utils

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
