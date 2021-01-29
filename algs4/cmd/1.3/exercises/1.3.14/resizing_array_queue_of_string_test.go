package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

var Q *ResizingArrayQueueOfString

func TestMain(m *testing.M) {
	Q = &ResizingArrayQueueOfString{}
	Q.Init(1)
	m.Run()
}

func TestResizingArrayQueueOfString_Enqueue(t *testing.T) {
	input := "to be or not to - be - - that - - - is"
	strs := strings.Split(input, " ")
	for _, str := range strs {
		if str == "-" {
			fmt.Println(Q.Dequeue())
		} else {
			Q.Enqueue(str)
		}
	}
}

func TestResizingArrayQueueOfString_Dequeue(t *testing.T) {
	strs := []string{"to", "be", "or", "not", "to", "be", "that", "is"}
	for _, str := range strs {
		Q.Enqueue(str)
	}

	for v := Q.Dequeue(); v != ""; v = Q.Dequeue() {
		fmt.Println(v)
	}

	assert.Equal(t, 0, Q.Size())
}
