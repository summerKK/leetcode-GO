package main

import (
	"fmt"

	"github.com/summerKK/leetcode-Go/algs4/utils"
)

/**
1.3.26
编写一个方法 remove()，接受一条链表和一个字符串 key 作为参数，删除链表中所有 item 域
为 key 的结点
*/
func main() {
	linkList0 := utils.NewLinkList()
	// 4 3 2 1
	linkList0.Add(4)
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

	linkList0.Remove(4)

	fmt.Println(linkList0)
}
