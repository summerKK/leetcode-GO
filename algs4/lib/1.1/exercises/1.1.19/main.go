package main

import "fmt"

// 1.1.19
func main() {
	deep0 := 0
	r0 := fibonacci0(20, &deep0)
	fmt.Printf("fibonacci0 result:%d,deep:%d\n", r0, deep0)

	deep1 := 0
	r1 := fibonacci1(20, &deep1)
	fmt.Printf("fibonacci1 result:%d,deep:%d\n", r1, deep1)
}

var cache map[int]int = make(map[int]int, 1024)

func fibonacci0(N int, deep *int) int {
	if N == 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	if _, ok := cache[N-1]; !ok {
		*deep++
		cache[N-1] = fibonacci0(N-1, deep)
	}

	if _, ok := cache[N-2]; !ok {
		*deep++
		cache[N-2] = fibonacci0(N-2, deep)
	}

	return cache[N-1] + cache[N-2]
}

func fibonacci1(N int, deep *int) int {
	*deep++
	if N == 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	return fibonacci1(N-1, deep) + fibonacci1(N-2, deep)
}
