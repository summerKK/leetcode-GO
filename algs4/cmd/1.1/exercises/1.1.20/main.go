package main

import "fmt"

// 编写一个递归的静态方法计算 ln(n!) 的值。
func main() {
	r := ln(5)
	fmt.Printf("ln result:%d", r)
}

func ln(N int) int {
	if N <= 1 {
		return 1
	}

	return N * ln(N-1)
}
