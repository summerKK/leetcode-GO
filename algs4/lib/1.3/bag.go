package lib

import "github.com/summerKK/leetcode-Go/algs4/utils"

type Bag interface {
	Init()
	Add(item interface{})
	IsEmpty() bool
	Size() int
}

type MyBag struct {
	first *utils.Node
	n     int
}

func (m *MyBag) Init() {
	m.first = &utils.Node{}
}

func (m *MyBag) Add(item interface{}) {
	node := utils.NewNode(item, nil)
	node.Next = m.first
	m.first = node
	m.n++
}

func (m *MyBag) IsEmpty() bool {
	return m.n <= 0
}

func (m *MyBag) Size() int {
	return m.n
}
