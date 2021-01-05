package leetcode

import "sort"

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

func intersect2(nums1 []int, nums2 []int) []int {
	var i, j, k int
	// 先对数组进行排序
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i < len(nums1) && j < len(nums2) {
		// 两个数组元素相同.放入到空白数组(这里使用nums1来存放,节省内存.没放一个元素k++)
		if nums1[i] == nums2[j] {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		} else if nums1[i] > nums2[j] {
			// 如果两个数组的元素不同,元素小的那个数组往前移动一位.元素大的那个数组保持不变
			j++
		} else {
			i++
		}
	}

	return nums1[0:k]
}
