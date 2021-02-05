package main

type UF struct {
	arr   []int
	count int
}

func NewUF(n int) *UF {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	return &UF{
		arr:   arr,
		count: n,
	}
}

func (u *UF) Find(p int) int {
	return u.arr[p]
}

func (u *UF) Union(p int, q int) {
	if u.Connected(p, q) {
		return
	}
	pV := u.Find(p)
	qV := u.Find(q)

	for i, v := range u.arr {
		if v == qV {
			u.arr[i] = pV
		}
	}

	u.count--
}

func (u *UF) Connected(p int, q int) bool {
	if u.Find(p) == u.Find(q) {
		return true
	}

	return false
}

func (u *UF) Count() int {
	return u.count
}
