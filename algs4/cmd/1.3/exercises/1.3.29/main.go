package main

import "fmt"

/**
1.3.29
用环形链表实现 CircleQueue。环形链表也是一条链表，只是没有任何结点的链接为空，且只要链表非
空则 last.next 的值为 first。只能使用一个 node 类型的实例变量(last)。
*/
func main() {
	circleQueue := NewQueue()
	circleQueue.Enqueue(1)
	circleQueue.Enqueue(2)
	circleQueue.Enqueue(3)
	circleQueue.Enqueue(4)

	circleQueue.Dequeue()

	fmt.Println(circleQueue)
}
