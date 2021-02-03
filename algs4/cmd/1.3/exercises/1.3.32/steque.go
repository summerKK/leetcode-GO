package main

import (
	"bytes"
	"fmt"
)

type Node struct {
	next *Node
	prev *Node
	item interface{}
}

type Steque struct {
	head    *Node
	current *Node
	n       int
}

func NewSteque() *Steque {
	return &Steque{
		head: &Node{},
	}
}

func (s *Steque) Push(item interface{}) {
	node := &Node{item: item}
	if s.IsEmpty() {
		s.current = node
		s.head.next = node
	} else {
		node.next = s.head.next
		node.prev = s.head
	}
	s.head.next = node

	s.n++
}

func (s *Steque) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	item := s.head.next
	s.head.next = item.next
	s.n--

	return item.item
}

func (s *Steque) Enqueue(item interface{}) {
	node := &Node{item: item}
	if s.IsEmpty() {
		node.prev = s.head
		s.head.next = node
		s.current = node
	} else {
		s.current.next = node
		s.current = node
	}

	s.n++
}

func (s *Steque) Size() int {
	return s.n
}

func (s *Steque) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Steque) Loop() chan interface{} {
	c := make(chan interface{})
	go func() {
		for node := s.head.next; node != nil; node = node.next {
			c <- node.item
		}
		close(c)
	}()

	return c
}

func (s *Steque) String() string {
	var buf bytes.Buffer
	for v := range s.Loop() {
		buf.WriteString(fmt.Sprintf("%v", v))
		buf.WriteByte(' ')
	}

	return buf.String()
}
