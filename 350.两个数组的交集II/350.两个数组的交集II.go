package leetcode

func intersect0(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, v := range nums1 {
		m[v] += 1
	}

	var inters []int
	for _, v := range nums2 {
		if m[v] > 0 {
			inters = append(inters, v)
			m[v] -= 1
		}
	}

	return inters
}

func intersect1(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, v := range nums1 {
		m[v] += 1
	}

	k := 0
	for _, v := range nums2 {
		if m[v] > 0 {
			m[v] -= 1
			// 把元素存在nums2中节省内存
			nums2[k] = v
			k++
		}
	}

	return nums2[0:k]
}
