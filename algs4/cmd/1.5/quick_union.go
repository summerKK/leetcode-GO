package lib

type QuickUnion struct {
	arr   []int
	count int
}

func NewQuickUnion(n int) *QuickUnion {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return &QuickUnion{
		arr:   arr,
		count: n,
	}
}

func (qu *QuickUnion) Union(p int, q int) {
	head0 := qu.Find(p)
	head1 := qu.Find(q)

	if head0 == head1 {
		return
	}

	qu.arr[head0] = head1

	qu.count--
}

func (qu *QuickUnion) Connected(p int, q int) bool {
	return qu.Find(p) == qu.Find(q)
}

func (qu *QuickUnion) Find(p int) int {
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

func (qu *QuickUnion) Count() int {
	return qu.count
}
