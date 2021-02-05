package main

import (
	"fmt"

	"github.com/summerKK/leetcode-Go/algs4/cmd/1.3"
)

/**
1.3.6
下面这段代码对队列 q 进行了什么操作?
       Stack<String> stack = new Stack<String>();
       while (!q.isEmpty())
          stack.push(q.dequeue());
       while (!stack.isEmpty())
          q.enqueue(stack.pop());
*/
func main() {
	S := &__3.MyStack{}
	S.Init(0)
	Q := &__3.MyQueue{}
	Q.Enqueue(1)
	Q.Enqueue(2)
	Q.Enqueue(3)
	Q.Enqueue(4)

	for !Q.IsEmpty() {
		S.Push(Q.Dequeue())
	}

	for !S.IsEmpty() {
		Q.Enqueue(S.Pop())
	}

	for !Q.IsEmpty() {
		fmt.Println(Q.Dequeue())
	}
}
