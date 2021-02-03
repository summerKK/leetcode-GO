package main

import (
	"bytes"
	"fmt"
)

type Deque struct {
	arr   []interface{}
	n     int
	first int
}

func NewDeque(cap int) *Deque {
	if cap < 4 {
		cap = 4
	}

	q := &Deque{
		arr: make([]interface{}, cap),
		// 从数组中间插入值.因为支持左右两边插入
		first: cap / 2,
	}

	return q
}

func (d *Deque) isEmpty() bool {
	return d.Size() == 0
}

func (d *Deque) Size() int {
	return d.n
}

func (d *Deque) PushLeft(item interface{}) {
	d.arr[d.first] = item
	d.first--
	d.n++

	if d.first < 0 {
		if d.Size() > len(d.arr)/2 {
			d.resize(len(d.arr) * 2)
		} else {
			// 在原数组居中
			d.resize(len(d.arr))
		}
	}
}

func (d *Deque) PushRight(item interface{}) {
	lastIndex := d.first + 1 + d.n
	d.arr[lastIndex] = item
	d.n++
	if lastIndex == len(d.arr)-1 {
		if d.Size() > len(d.arr)/2 {
			d.resize(len(d.arr) * 2)
		} else {
			d.resize(len(d.arr))
		}
	}
}

func (d *Deque) resize(size int) {
	fmt.Printf("resize:%d to %d\n", len(d.arr), size)
	newArr := make([]interface{}, size)
	// 新的数组长度 - 已经存在的元素
	startPoint := (size - d.n) / 2
	startPoint1 := startPoint
	for i := d.first + 1; i < d.first+1+d.n; i++ {
		newArr[startPoint] = d.arr[i]
		startPoint++
	}

	d.arr = newArr
	d.first = startPoint1 - 1
}

func (d *Deque) PopLeft() interface{} {
	if d.isEmpty() {
		return nil
	}

	d.first++
	item := d.arr[d.first]
	d.arr[d.first] = nil
	d.n--

	if d.Size() <= len(d.arr)/4 {
		d.resize(len(d.arr) / 2)
	}

	return item
}

func (d *Deque) PopRight() interface{} {
	if d.isEmpty() {
		return nil
	}
	lastIndex := d.first + 1 + d.n
	item := d.arr[lastIndex]
	d.arr[lastIndex] = nil
	d.n--

	if d.Size() <= len(d.arr)/4 {
		d.resize(len(d.arr) / 2)
	}

	return item
}

func (d *Deque) Loop() chan interface{} {
	c := make(chan interface{})
	go func() {
		for i := d.first + 1; i < d.first+1+d.n; i++ {
			c <- d.arr[i]
		}
		close(c)
	}()

	return c
}

func (d *Deque) String() string {
	var buf bytes.Buffer
	for v := range d.Loop() {
		buf.WriteString(fmt.Sprintf("%v", v))
		buf.WriteByte(' ')
	}

	return buf.String()
}
