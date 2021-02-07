package main

import (
	"fmt"
	"math"
)

/**
1.5.14
根据高度加权的 quick-Union 算法。给出 UF 的一个实现，使用和加权 quick-Union 算法相同的策略， 但记录的是树的高度并总是将较矮的树连接到较高的树上。用算法证明 N 个触点的树的高度不 会超过其大小的对数级别。
*/
func main() {

	weightedQuickUnion := NewWeightedQuickUnion(19)

	weightedQuickUnion.Union(0, 1)
	weightedQuickUnion.Union(0, 2)
	weightedQuickUnion.Union(0, 3)
	weightedQuickUnion.Union(6, 7)
	weightedQuickUnion.Union(8, 9)
	weightedQuickUnion.Union(6, 8)
	weightedQuickUnion.Union(0, 6)
	weightedQuickUnion.Union(10, 11)
	weightedQuickUnion.Union(10, 12)
	weightedQuickUnion.Union(10, 13)
	weightedQuickUnion.Union(10, 14)
	weightedQuickUnion.Union(10, 15)
	weightedQuickUnion.Union(10, 16)
	weightedQuickUnion.Union(10, 17)
	weightedQuickUnion.Union(10, 18)
	weightedQuickUnion.Union(0, 10)

	fmt.Printf("一共存在%d个分量\n", weightedQuickUnion.Count())
	fmt.Printf("树的最大高度为:%d,%d个触点的对数为:%.2f\n", weightedQuickUnion.MaxHeight(),
		len(weightedQuickUnion.arr), math.Log2(float64(len(weightedQuickUnion.arr))))
}
