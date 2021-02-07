package main

/**
quick-union 实现路径压缩算法.树的总体高度不会超过3
*/
type QuickUnionPathCompression struct {
	arr   []int
	count int
}

func NewQuickUnionPathCompression(n int) *QuickUnionPathCompression {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return &QuickUnionPathCompression{
		arr:   arr,
		count: n,
	}
}

func (qu *QuickUnionPathCompression) Union(p int, q int) {
	head0 := qu.Find(p)
	head1 := qu.Find(q)

	if head0 == head1 {
		return
	}

	qu.arr[head0] = head1

	qu.count--
}

func (qu *QuickUnionPathCompression) Connected(p int, q int) bool {
	return qu.Find(p) == qu.Find(q)
}

// 不关心节点的关系.每个节点都指向父节点
func (qu *QuickUnionPathCompression) Find(p int) int {
	if p == qu.arr[p] {
		return p
	}

	qu.arr[p] = qu.Find(qu.arr[p])

	return qu.arr[p]
}

func (qu *QuickUnionPathCompression) Count() int {
	return qu.count
}
