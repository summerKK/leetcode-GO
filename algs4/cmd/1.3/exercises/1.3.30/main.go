package main

import "fmt"

/**
1.3.30
编写一个函数，接受一条链表的首结点作为参数，(破坏性地)将链表反转并返回结果链表的
首结点
*/
func main() {
	node := NewNode()
	node.Add(1).Add(2).Add(3).Add(4)

	newNode := reverse(node)

	for node := newNode; node.item != nil; node = node.next {
		fmt.Println(node.item)
	}
}

func reverse(node *Node) *Node {
	newNode := NewNode()
	for node := node.next; node != nil; node = node.next {
		node0 := &Node{
			item: node.item,
			next: newNode,
		}
		newNode = node0
	}

	return newNode
}
