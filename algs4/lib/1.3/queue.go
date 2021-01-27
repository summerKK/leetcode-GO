package lib

import "github.com/summerKK/leetcode-Go/algs4/utils"

type Queue interface {
	Init()
	Enqueue(item interface{})
	Dequeue() interface{}
	IsEmpty() bool
	Size() int
}

type MyQueue struct {
	// 指向最早添加的元素
	first *utils.Node
	// 指向最后添加的元素
	last *utils.Node
	n    int
}

func (m *MyQueue) Init() {
	m.first = &utils.Node{}
	m.last = m.first
}

func (m *MyQueue) Enqueue(item interface{}) {
	node := utils.NewNode(item, nil)
	oldLast := m.last
	m.last = node
	if m.IsEmpty() {
		m.first = m.last
	} else {
		oldLast.Next = node
	}
	m.n++
}

func (m *MyQueue) Dequeue() interface{} {
	if m.IsEmpty() {
		return nil
	}

	node := m.first
	m.first = m.first.Next
	m.n--

	if m.IsEmpty() {
		m.last = nil
	}

	return node.Item
}

func (m *MyQueue) IsEmpty() bool {
	return m.n <= 0
}

func (m *MyQueue) Size() int {
	return m.n
}
