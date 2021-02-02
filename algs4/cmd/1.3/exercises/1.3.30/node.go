package main

type Node struct {
	next *Node
	item interface{}
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) Add(item interface{}) *Node {
	node := &Node{
		item: item,
	}
	n.next = node

	return node
}
