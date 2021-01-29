package main

import (
	"fmt"

	"github.com/summerKK/leetcode-Go/algs4/utils"
)

/**
1.3.24
编写一个方法 removeAfter()，接受一个链表结点作为参数并删除该结点的后续结点(如果参
数结点或参数结点的后续结点为空则什么也不做)。
*/
func main() {
	linkList := utils.NewLinkList()
	// 4 3 2 1
	linkList.Add(1)
	linkList.Add(2)
	linkList.Add(3)
	linkList.Add(4)

	// 3
	linkList.RemoveAfter(linkList.First.Next)

	for v := range linkList.Loop() {
		fmt.Println(v)
	}
}
