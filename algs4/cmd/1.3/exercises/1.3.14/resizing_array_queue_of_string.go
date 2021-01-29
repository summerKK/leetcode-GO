package main

import "fmt"

type ResizingArrayQueueOfString struct {
	a    []string
	N    int
	head int
	tail int
}

func (s *ResizingArrayQueueOfString) Init(n int) {
	s.a = make([]string, n)
}

func (s *ResizingArrayQueueOfString) Enqueue(item string) {
	if s.a == nil {
		return
	}

	// 达到最大容量,扩容
	if s.N == len(s.a) {
		s.resize(s.N * 2)
	}

	s.a[s.tail] = item
	s.N++
	s.tail++
}

func (s *ResizingArrayQueueOfString) Dequeue() string {
	if s.IsEmpty() {
		return ""
	}

	item := s.a[s.head]
	s.a[s.head] = ""
	s.head++
	s.N--

	// 减小数组长度,避免内存浪费
	if s.N < len(s.a)/2 {
		s.resize(len(s.a) / 2)
	}

	return item
}

func (s *ResizingArrayQueueOfString) resize(n int) {
	fmt.Printf("queue resize: %d -> %d\n", s.N, n)
	arr := make([]string, n)
	c := 0
	// head记录的删除元素后数组下一个元素的下标
	// tail记录增加元素后数组下一个元素的下标
	for i := s.head; i < s.tail; i++ {
		arr[c] = s.a[i]
		c++
	}
	// tail - head 是增加元素和删除元素中间的差值.如果head为0.tail = 0 + (tail - head)
	s.tail -= s.head
	s.head = 0
	s.a = arr
}

func (s *ResizingArrayQueueOfString) IsEmpty() bool {
	return s.N == 0
}

func (s *ResizingArrayQueueOfString) Size() int {
	return s.N
}

func (s *ResizingArrayQueueOfString) Loop() <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < s.N; i++ {
			c <- s.a[i]
		}
		close(c)
	}()

	return c
}
