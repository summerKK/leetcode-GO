package leetcode

func intersect(nums1 []int, nums2 []int) []int {
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
