package lib

type WeightedQuickUnion struct {
	arr []int
	// 记录每个数的大小
	sz    []int
	count int
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
		sz:    sz,
		count: n,
	}
}

func (qu *WeightedQuickUnion) Union(p int, q int) {
	head0 := qu.Find(p)
	head1 := qu.Find(q)

	if head0 == head1 {
		return
	}

	//  把小树连接到大树下面
	if qu.sz[p] < qu.sz[q] {
		qu.arr[head0] = head1
		qu.sz[head1] += qu.sz[head0]
	} else {
		qu.arr[head1] = head0
		qu.sz[head0] += qu.sz[head1]
	}

	qu.count--
}

func (qu *WeightedQuickUnion) Connected(p int, q int) bool {
	return qu.Find(p) == qu.Find(q)
}

func (qu *WeightedQuickUnion) Find(p int) int {
	for {
		index := qu.arr[p]
		if index == p {
			p = index
			break
		}
		p = index
	}

	return p
}

func (qu *WeightedQuickUnion) Count() int {
	return qu.count
}
