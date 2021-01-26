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
		fmt.Printf("stack resize:%d\n", len(f.a)*2)
	}
	f.a[f.N] = item
	f.N++
}

func (f *FixedCapStack) Pop() interface{} {
	f.N--
	item := f.a[f.N]

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
	arr := make([]interface{}, max)
	for i := 0; i < f.Size(); i++ {
		arr[i] = f.a[i]
	}
	f.a = arr
}
