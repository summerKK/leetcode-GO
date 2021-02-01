package main

import (
	"fmt"

	"github.com/summerKK/leetcode-Go/algs4/utils"
)

/**
1.3.27
编写一个方法 max()，接受一条链表的首结点作为参数，返回链表中键最大的节点的值。假设所
有键均为正整数，如果链表为空则返回 0。
1.3.28
用递归的方法解答上一道练习。
*/
func main() {
	linkList0 := utils.NewLinkList()
	// 4 3 2 1
	linkList0.Add(4)
	linkList0.Add(5)
	linkList0.Add(1)
	linkList0.Add(2)
	linkList0.Add(3)
	linkList0.Add(4)
	linkList0.Add(4)
	linkList0.Add(4)
	linkList0.Add(4)
	linkList0.Add(4)
	linkList0.Add(4)
	linkList0.Add(4)

	fmt.Printf("max:%v\n", Max0(linkList0))
	fmt.Printf("max:%v\n", Max1(linkList0.First, 0))
}

func Max0(linkList *utils.LinkList) int {
	if linkList.First == nil {
		return 0
	}

	pre := linkList.First
	max := pre.Item.(int)
	for node := linkList.First.Next; node != nil; node = node.Next {
		if node.Item.(int) > max {
			max = node.Item.(int)
		}
		pre = node
	}

	return max
}

func Max1(node *utils.Node, max int) int {
	if node == nil {
		return max
	}

	if node.Item.(int) > max {
		max = node.Item.(int)
	}

	return Max1(node.Next, max)
}
