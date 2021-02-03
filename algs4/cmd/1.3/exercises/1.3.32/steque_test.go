package main

import (
	"fmt"
	"testing"
)

func TestSteque_Enqueue(t *testing.T) {
	steque := NewSteque()
	steque.Enqueue(1)
	steque.Enqueue(2)
	steque.Enqueue(3)
	steque.Enqueue(4)

	steque.Pop()

	fmt.Println(steque)
}

func TestSteque_Push(t *testing.T) {
	steque := NewSteque()
	steque.Push(1)
	steque.Push(2)
	steque.Push(3)
	steque.Push(4)

	fmt.Println(steque)
}

func TestStequePushEnqueue(t *testing.T) {
	steque := NewSteque()
	steque.Push(1)
	steque.Push(2)
	steque.Push(3)
	steque.Push(4)

	steque.Enqueue(5)
	steque.Enqueue(6)
	steque.Enqueue(7)

	steque.Pop()

	fmt.Println(steque)
	fmt.Println(steque.Size())
}
