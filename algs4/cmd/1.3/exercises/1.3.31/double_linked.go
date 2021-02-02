package main

import (
	"bytes"
	"fmt"
)

type DoubleNode struct {
	pre  *DoubleNode
	next *DoubleNode
	item interface{}
}

func NewDoubleNode(item interface{}) *DoubleNode {
	return &DoubleNode{
		item: item,
	}
}

type DoubleLinked struct {
	head    *DoubleNode
	current *DoubleNode
	n       int
}

func NewDoubleLinked() *DoubleLinked {
	linked := &DoubleLinked{
		head: NewDoubleNode(nil),
	}

	return linked
}

func (d *DoubleLinked) InsertByHead(item interface{}) {
	node := NewDoubleNode(item)

	if d.IsEmpty() {
		d.current = node
		d.head.next = node
	} else {
		node.next = d.head.next
		d.head.next.pre = node
		d.head.next = node
	}

	node.pre = d.head
	d.n++
}

func (d *DoubleLinked) InsertByTail(item interface{}) {
	node := NewDoubleNode(item)

	if d.IsEmpty() {
		d.current = node
		d.head.next = node
		node.pre = d.head
	} else {
		node.pre = d.current
		d.current.next = node
		d.current = node
	}

	d.n++
}

func (d *DoubleLinked) DeleteByHead() interface{} {
	if d.IsEmpty() {
		return nil
	}

	item := d.head.next
	d.head.next = item.next

	return item
}

func (d *DoubleLinked) DeleteByTail() interface{} {
	if d.IsEmpty() {
		return nil
	}

	item := d.current
	d.current.pre.next = nil
	d.current = item.pre

	return item
}

func (d *DoubleLinked) Size() int {
	return d.n
}

func (d *DoubleLinked) IsEmpty() bool {
	return d.Size() == 0
}

func (d *DoubleLinked) Loop() chan interface{} {
	c := make(chan interface{})
	go func() {
		for node := d.head.next; node != nil; node = node.next {
			c <- node.item
		}
		close(c)
	}()

	return c
}

func (d *DoubleLinked) String() string {
	var buf bytes.Buffer
	for v := range d.Loop() {
		buf.WriteString(fmt.Sprintf("%v", v))
		buf.WriteByte(' ')
	}

	return buf.String()
}
