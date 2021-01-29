package main

import (
	"fmt"

	"github.com/summerKK/leetcode-Go/algs4/utils"
)

/**
1.3.25
编写一个方法 insertAfter()，接受两个链表结点作为参数，将第二个结点插入链表并使之成
为第一个结点的后续结点(如果两个参数为空则什么也不做)。
*/
func main() {
	linkList0 := utils.NewLinkList()
	// 4 3 2 1
	linkList0.Add(1)
	linkList0.Add(2)
	linkList0.Add(3)
	linkList0.Add(4)

	linkList1 := utils.NewLinkList()
	// d c b a
	linkList1.Add("a")
	linkList1.Add("b")
	linkList1.Add("c")
	linkList1.Add("d")

	linkList0.InsertAfter(linkList1.First)

	for v := range linkList0.Loop() {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
