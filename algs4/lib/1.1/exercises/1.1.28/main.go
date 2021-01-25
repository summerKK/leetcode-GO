package main

import (
	"fmt"
	"sort"
)

func main() {
	whitelist := []int{0, 0, 1, 1, 1, 2, 3, 4, 4, 6, 6, 7, 8, 7, 5}
	sort.Ints(whitelist)

	repeat := 0
	// 查看有多少重复数据.数组已经排序(GO其实不需要这一步.直接用slice)
	for i := 1; i < len(whitelist); i++ {
		if whitelist[i] == whitelist[i-1] {
			repeat++
		}
	}

	unique := make([]int, 0, len(whitelist)-repeat)
	// unique是之前一个i-1 和当前i比较,所以i=0的元素不会加入到unique(unique = append(unique, whitelist[i])).这里手动添加上
	unique = append(unique, whitelist[0])
	for i := 1; i < len(whitelist); i++ {
		if whitelist[i] != whitelist[i-1] {
			unique = append(unique, whitelist[i])
		}
	}

	fmt.Printf("unique arr:%v\n", unique)
}
