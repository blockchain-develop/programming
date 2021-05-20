package utils

func NewMinHeap() *MinHeap {
	mh := &MinHeap{data:make([]int, 0)}
	return mh
}

type MinHeap struct {
	data []int
}

func (mh *MinHeap) fix(node int) {
	leftChild, rightChild := node * 2, node * 2 + 1
	child := leftChild
	if leftChild > len(mh.data) {
		return
	}
	if rightChild <= len(mh.data) && mh.data[rightChild - 1] < mh.data[leftChild - 1] {
		child = rightChild
	}
	if mh.data[node - 1] < mh.data[child - 1] {
		return
	}
	mh.data[child - 1], mh.data[node - 1] = mh.data[node - 1], mh.data[child - 1]
	mh.fix(child)
}

func (mh *MinHeap) Init(data []int) {
	mh.data = append(mh.data, data...)
	total := len(mh.data)
	node := total / 2
	for i := node;i > 0;i -- {
		mh.fix(i)
	}
}

func (mh *MinHeap) Insert(item int) {
	mh.data = append(mh.data, item)
	total := len(mh.data)
	node := total / 2
	for node != 0 {
		mh.fix(node)
		node = node / 2
	}
}

func (mh *MinHeap) Remove() int {
	data := mh.data[0]
	mh.data[0] = mh.data[len(mh.data) - 1]
	mh.data = mh.data[0:len(mh.data)-1]
	mh.fix(1)
	return data
}

func (mh *MinHeap) Len() int {
	return len(mh.data)
}

