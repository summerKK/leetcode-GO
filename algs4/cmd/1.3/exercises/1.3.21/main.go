package main

import (
	"fmt"

	"github.com/summerKK/leetcode-Go/algs4/utils"
)

/**
1.3.21
编写一个方法 find()，接受一条链表和一个字符串 key 作为参数。如果链表中的某个结点的
item 域的值为 key，则方法返回 true，否则返回 false。
*/
func main() {
	linkList := utils.NewLinkList()
	linkList.Add(1)
	linkList.Add(2)
	linkList.Add(3)

	item := 3
	result := linkList.Find(item)
	fmt.Printf("find %d,result:%v\n", item, result)
}
