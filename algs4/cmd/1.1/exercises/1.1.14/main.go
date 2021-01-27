package main

import "fmt"

// 编写一个静态方法 lg()，接受一个整型参数 n，返回不大于 log2N 的最大整数。不要使用 Math 库。
func main() {
	r := lg(16)
	fmt.Printf("lg result:%d\n", r)
}

func lg(N int) int {
	c := 0
	exp := 1
	for exp <= N {
		exp *= 2
		c++
	}
	return c - 1
}
