package main

import (
	"bytes"
	"fmt"
)

type node struct {
	item interface{}
	next *node
}

type CircleQueue struct {
	head    *node
	current *node
	n       int
}

func NewQueue() *CircleQueue {
	q := &CircleQueue{}
	q.head = &node{}
	q.current = q.head

	return q
}

func (q *CircleQueue) Enqueue(item interface{}) {
	n := &node{
		item: item,
	}
	// 尾部指向第一个节点(circle)
	n.next = q.head.next

	q.current.next = n
	q.current = n
	q.n++
}

func (q *CircleQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	item := q.head.next.item
	q.head.next = q.head.next.next
	// current.是tail节点.tail节点指向第一个节点(circle)
	q.current.next = q.head.next

	q.n--

	return item
}

func (q *CircleQueue) IsEmpty() bool {
	return q.n == 0
}

func (q *CircleQueue) Size() int {
	return q.n
}

func (q *CircleQueue) String() string {
	var buf bytes.Buffer
	for node := q.head.next; node != nil; node = node.next {
		_, _ = buf.WriteString(fmt.Sprintf("%v", node.item))
		_ = buf.WriteByte(' ')
		// 尾部
		if node.next == q.head.next {
			break
		}
	}

	return buf.String()

}
