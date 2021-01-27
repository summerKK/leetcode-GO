package lib

import "fmt"

type Stack interface {
	Init(int)
	Push(item interface{})
	Pop() interface{}
	IsEmpty() bool
	Size() int
}

type FixedCapStack struct {
	a []interface{}
	N int
}

func (f *FixedCapStack) Init(n int) {
	f.a = make([]interface{}, n)
}

func (f *FixedCapStack) Push(item interface{}) {
	// 达到数组最大容量
	if f.N == len(f.a) {
		f.resize(len(f.a) * 2)
	}
	f.a[f.N] = item
	f.N++
}

func (f *FixedCapStack) Pop() interface{} {
	f.N--
	item := f.a[f.N]
	// 更好GC
	f.a[f.N] = nil
	// 数组容量过大,减小容量
	if f.N > 0 && len(f.a)/4 > f.N {
		// 容量减半
		f.resize(len(f.a) / 2)
	}

	return item
}

func (f *FixedCapStack) IsEmpty() bool {
	return f.N == 0
}

func (f *FixedCapStack) Size() int {
	return f.N
}

// 重置数组大小
func (f *FixedCapStack) resize(max int) {
	fmt.Printf("stack resize:%d\n", max)

	arr := make([]interface{}, max)
	for i := 0; i < f.Size(); i++ {
		arr[i] = f.a[i]
	}
	f.a = arr
}

func (f *FixedCapStack) Loop() <-chan interface{} {
	c := make(chan interface{})
	go func() {
		for !f.IsEmpty() {
			c <- f.next()
		}
		close(c)
	}()

	return c
}

func (f *FixedCapStack) next() interface{} {
	f.N--
	return f.a[f.N]
}