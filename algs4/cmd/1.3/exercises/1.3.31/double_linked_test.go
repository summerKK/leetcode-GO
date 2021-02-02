package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDoubleLinked_InsertByHead(t *testing.T) {
	linked := NewDoubleLinked()
	linked.InsertByHead(1)
	linked.InsertByHead(2)
	linked.InsertByHead(3)
	linked.InsertByHead(4)
	linked.InsertByHead(5)
	linked.InsertByHead(6)

	fmt.Println(linked)
}

func TestDoubleLinked_InsertByTail(t *testing.T) {
	linked := NewDoubleLinked()
	linked.InsertByTail(1)
	linked.InsertByTail(2)
	linked.InsertByTail(3)
	linked.InsertByTail(4)
	linked.InsertByTail(5)
	linked.InsertByTail(6)

	fmt.Println(linked)
}

func TestDoubleLinked_Insert(t *testing.T) {
	// 6 5 4 1 2 3
	linked := NewDoubleLinked()
	linked.InsertByTail(1)
	linked.InsertByTail(2)
	linked.InsertByTail(3)
	linked.InsertByHead(4)
	linked.InsertByHead(5)
	linked.InsertByHead(6)

	except := []interface{}{
		6, 5, 4, 1, 2, 3,
	}
	var results []interface{}
	for v := range linked.Loop() {
		results = append(results, v)
	}

	assert.Equal(t, results, except)
}

func TestDoubleLinked_DeleteByHead(t *testing.T) {
	linked := NewDoubleLinked()
	linked.InsertByTail(1)
	linked.InsertByTail(2)
	linked.InsertByTail(3)
	linked.InsertByHead(4)
	linked.InsertByHead(5)
	linked.InsertByHead(6)

	linked.DeleteByHead()

	fmt.Println(linked)
}

func TestDoubleLinked_DeleteByTail(t *testing.T) {
	linked := NewDoubleLinked()
	linked.InsertByTail(1)
	linked.InsertByTail(2)
	linked.InsertByTail(3)
	linked.InsertByHead(4)
	linked.InsertByHead(5)
	linked.InsertByHead(6)

	linked.DeleteByTail()

	fmt.Println(linked)
}
