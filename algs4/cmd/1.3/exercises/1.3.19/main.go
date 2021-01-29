package main

import (
	"fmt"

	"github.com/summerKK/leetcode-Go/algs4/utils"
)

/**
1.3.19
给出一段代码，删除链表的尾结点，其中链表的首结点为 first。
*/
func main() {
	linkList := utils.NewLinkList()
	linkList.Add(1)
	linkList.Add(2)
	linkList.Add(3)

	DeleteLast(linkList)

	fmt.Println(linkList)
}

func DeleteLast(linkList *utils.LinkList) {
	pre := linkList.First
	for node := linkList.First; node.Next != nil; node = node.Next {
		pre = node
	}

	pre.Next = nil
}
