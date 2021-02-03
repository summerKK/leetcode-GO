package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDeque_PushLeft(t *testing.T) {
	deque := NewDeque(1)
	deque.PushLeft(1)
	deque.PushLeft(2)
	deque.PushLeft(3)
	deque.PushLeft(4)
	deque.PushLeft(5)

	fmt.Println(deque)
}

func TestDeque_PushRight(t *testing.T) {
	deque := NewDeque(1)
	deque.PushRight(1)
	deque.PushRight(2)
	deque.PushRight(3)
	deque.PushRight(4)
	deque.PushRight(5)

	fmt.Println(deque)
}

func TestDeque_Push(t *testing.T) {
	deque := NewDeque(1)

	deque.PushRight(1)
	deque.PushLeft(2)
	deque.PushRight(3)
	deque.PushLeft(4)
	deque.PushLeft(5)
	deque.PushLeft(6)
	deque.PushRight(7)

	assert.Equal(t, deque.String(), "6 5 4 2 1 3 7 ")
}

func TestDeque_PopLeft(t *testing.T) {
	deque := NewDeque(1)
	deque.PushRight(1)
	deque.PushRight(2)
	deque.PushRight(3)

	deque.PopLeft()

	assert.Equal(t, deque.String(), "2 3 ")
}

func TestDeque_PopRight(t *testing.T) {
	deque := NewDeque(1)

	deque.PushLeft(1)
	deque.PushLeft(2)
	deque.PushLeft(3)
	deque.PushLeft(4)
	deque.PushLeft(5)

	deque.PopRight()
	deque.PopRight()
	deque.PopRight()
	deque.PopRight()

	assert.Equal(t, deque.String(), "5 ")
}
