package main

import (
	"fmt"
	"testing"
)

func TestRandomQueue_Enqueue(t *testing.T) {
	queue := NewRandomQueue(1)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Enqueue(7)

	fmt.Println(queue)
}

func TestRandomQueue_Dequeue(t *testing.T) {
	queue := NewRandomQueue(1)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Enqueue(7)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}

func TestRandomQueue_RandomLoop(t *testing.T) {
	queue := NewRandomQueue(1)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Enqueue(7)

	for v := range queue.RandomLoop() {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
