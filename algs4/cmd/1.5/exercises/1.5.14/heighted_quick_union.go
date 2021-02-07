package main

type WeightedQuickUnion struct {
	arr []int
	// 统计树的深度
	hz    []int
	count int
	// 树的最大深度
	maxHeight int
}

func NewWeightedQuickUnion(n int) *WeightedQuickUnion {
	arr := make([]int, n)
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		sz[i] = 1
	}
	return &WeightedQuickUnion{
		arr:   arr,
		hz:    sz,
		count: n,
	}
}

func (qu *WeightedQuickUnion) Union(p int, q int) {
	head0 := qu.Find(p)
	head1 := qu.Find(q)

	if head0 == head1 {
		return
	}

	if qu.hz[p] < qu.hz[q] {
		qu.arr[head0] = head1
	} else if qu.hz[p] > qu.hz[q] {
		qu.arr[head1] = head0
	} else {
		// 两个数高度相等的时候才增加
		qu.arr[head0] = head1
		qu.hz[head1]++

		if qu.hz[head1] > qu.maxHeight {
			qu.maxHeight = qu.hz[head1]
		}
	}

	qu.count--
}

func (qu *WeightedQuickUnion) Connected(p int, q int) bool {
	return qu.Find(p) == qu.Find(q)
}

func (qu *WeightedQuickUnion) Find(p int) int {
	for p != qu.arr[p] {
		p = qu.arr[p]
	}

	return p
}

func (qu *WeightedQuickUnion) Count() int {
	return qu.count
}

func (qu *WeightedQuickUnion) MaxHeight() int {
	return qu.maxHeight
}
