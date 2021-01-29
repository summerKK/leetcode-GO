package main

import (
	"fmt"

	"github.com/summerKK/leetcode-Go/algs4/utils"
)

/**
1.3.20
编写一个方法 delete()，接受一个 int 参数 k，删除链表的第 k 个元素(如果它存在的话)
*/
func main() {
	linkList := utils.NewLinkList()
	linkList.Add(1)
	linkList.Add(2)
	linkList.Add(3)

	k := 3
	fmt.Printf("delete %d %v,len:%d\n", k, linkList.DeleteK(k), linkList.Size())
}
