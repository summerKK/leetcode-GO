package main

import (
	"fmt"
	"testing"
)

func TestCircleQueue_Enqueue(t *testing.T) {
	circleQueue := NewQueue()
	circleQueue.Enqueue(1)
	circleQueue.Enqueue(2)
	circleQueue.Enqueue(3)
	circleQueue.Enqueue(4)

	circleQueue.Dequeue()

	fmt.Println(circleQueue)
}
