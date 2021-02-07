package main

import (
	"fmt"
)

/**
1.5.13
使用路径压缩的加权quick-union算法。修改加权quick-union算法(算法1.5)，实现如练习1.5.12 所述的路径压缩。给出一列输入，使该方法能够产生一棵高度为 4 的树。注意:该算法的所有 操作的均摊成本已知被限制在反 Ackermann 函数的范围之内，且对于实际应用中可能出现的所 有 N 值均小于 5。
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

	fmt.Println(weightedQuickUnion.Count())
}
