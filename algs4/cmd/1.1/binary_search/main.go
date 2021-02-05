package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var arr []int
	rand.Seed(time.Now().UnixNano())
	key := 0
	for i := 0; i < 10; i++ {
		x := rand.Intn(100)
		if i == 5 {
			key = x
		}
		arr = append(arr, x)
	}
	sort.Ints(arr)

	search := BinarySearch(key, arr)
	fmt.Println(arr)
	fmt.Printf("key:%d index:%d\n", key, search)
}
