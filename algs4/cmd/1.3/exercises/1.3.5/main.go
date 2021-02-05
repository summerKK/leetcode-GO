package main

import (
	"fmt"

	"github.com/summerKK/leetcode-Go/algs4/cmd/1.3"
)

/**
1.3.5
当 N 为 50 时下面这段代码会打印什么?从较高的抽象层次描述给定正整数 N 时这段代码的行为。
	Stack<Integer> stack = new Stack<Integer>();
	while (N > 0)
	{
	   stack.push(N % 2);
	N = N / 2; }
	for (int d : stack) StdOut.print(d);
	StdOut.println();
*/
func main() {
	N := 50
	S := &__3.MyStack{}
	S.Init(0)
	for N > 0 {
		S.Push(N % 2)
		N = N / 2
	}

	for s := range S.Loop() {
		fmt.Print(s)
	}
}
