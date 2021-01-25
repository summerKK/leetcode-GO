package main

import (
	"fmt"
	"sort"
)

// 1.1.29 等值键。为 BinarySearch 类添加一个静态方法 rank()，它接受一个键和一个整型有序数组(可 能存在重复键)作为参数并返回数组中小于该键的元素数量，
// 以及一个类似的方法 count() 来 返回数组中等于该键的元素的数量。注意:如果 i 和 j 分别是 rank(key,a) 和 count(key,a) 的返回值，
// 那么 a[i..i+j-1] 就是数组中所有和 key 相等的元素。
func main() {
	arrs := []int{0, 0, 1, 1, 1, 2, 3, 4, 4, 6, 6, 7, 8, 7, 5}
	key := 7
	sort.Ints(arrs)
	r0 := rank(key, arrs)
	fmt.Printf("rank result:%d\n", r0)
	r1 := count(key, arrs)
	fmt.Printf("count result:%d\n", r1)

	fmt.Printf("重复元素%d %v\n", key, arrs[r0:r1+r0])
}

func rank(key int, arrs []int) int {
	c := 0
	for i := 0; i < len(arrs); i++ {
		if arrs[i] >= key {
			return c
		}
		c++
	}

	return c
}

func count(key int, arrs []int) int {
	c := 0
	for i := 0; i < len(arrs); i++ {
		if arrs[i] == key {
			c++
		}
		if arrs[i] > key {
			break
		}
	}

	return c
}
