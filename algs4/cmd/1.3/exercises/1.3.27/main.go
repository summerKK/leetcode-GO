package main

import (
	"fmt"

	"github.com/summerKK/leetcode-Go/algs4/utils"
)

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

	fmt.Printf("max:%v\n", Max(linkList0))
}

func Max(linkList *utils.LinkList) int {
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
