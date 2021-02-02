package main

import "fmt"

/**
1.3.31
]实现一个嵌套类 DoubleNode 用来构造双向链表，其中每个结点都含有一个指向前驱元素的引用
和一项指向后续元素的引用(如果不存在则为 null)。为以下任务实现若干静态方法:在表头 插入结点、在表尾插入结点、从表头删除结点、从表尾删除结点、在指定结点之前插入新结点、 在指定结点之后插入新结点、删除指定结点。
*/
func main() {
	linked := NewDoubleLinked()
	// 2 1 3 4
	linked.InsertByHead(1)
	linked.InsertByHead(2)
	linked.InsertByTail(3)
	linked.InsertByTail(4)

	fmt.Println(linked)

	// 1
	node := linked.head.next.next
	fmt.Printf("node(%v) pre:%v\n", node.item, node.pre.item)
	fmt.Printf("node(%v) next:%v\n", node.item, node.next.item)
}
