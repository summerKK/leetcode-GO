package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type RandomQueue struct {
	arr []interface{}
	n   int
}

func NewRandomQueue(cap int) *RandomQueue {
	rand.Seed(time.Now().UnixNano())
	return &RandomQueue{
		arr: make([]interface{}, cap),
	}
}

func (r *RandomQueue) IsEmpty() bool {
	return r.Size() == 0
}

func (r *RandomQueue) Size() int {
	return r.n
}

func (r *RandomQueue) Enqueue(item interface{}) {
	r.arr[r.n] = item
	r.n++
	if r.n >= len(r.arr)-1 {
		r.resize(len(r.arr) * 2)
	}
}

func (r *RandomQueue) resize(size int) {
	fmt.Printf("resize:%d to %d\n", len(r.arr), size)
	newArr := make([]interface{}, size)
	for i := 0; i < r.n; i++ {
		newArr[i] = r.arr[i]
	}
	r.arr = newArr
}

func (r *RandomQueue) Dequeue() interface{} {
	if r.IsEmpty() {
		return nil
	}

	r.shuffle()
	item := r.arr[r.n-1]
	r.arr[r.n-1] = nil
	r.n--

	if r.n < len(r.arr)/2 {
		r.resize(len(r.arr) / 2)
	}

	return item
}

func (r *RandomQueue) shuffle() {
	if r.n <= 1 {
		return
	}

	randIndex := rand.Intn(r.n - 1)
	r.arr[randIndex], r.arr[r.n-1] = r.arr[r.n-1], r.arr[randIndex]
}

func (r *RandomQueue) sample() interface{} {
	if r.IsEmpty() {
		return nil
	}
	r.shuffle()

	return r.arr[r.n]
}

func (r *RandomQueue) Loop() <-chan interface{} {
	c := make(chan interface{})
	go func() {
		for i := 0; i < r.n; i++ {
			c <- r.arr[i]
		}

		close(c)
	}()
	return c
}

func (r *RandomQueue) RandomLoop() <-chan interface{} {
	c := make(chan interface{})
	go func() {
		arr := r.arr
		l := r.n - 1
		randIndex := 0
		for i := 0; i < r.n; i++ {
			if l <= 0 {
				randIndex = 0
			} else {
				randIndex = rand.Intn(l)
			}
			arr[randIndex], arr[l] = arr[l], arr[randIndex]
			c <- arr[l]
			l--
		}
		close(c)
	}()
	return c
}

func (r *RandomQueue) String() string {
	var buf bytes.Buffer
	for v := range r.Loop() {
		buf.WriteString(fmt.Sprintf("%v", v))
		buf.WriteByte(' ')
	}

	return buf.String()
}
