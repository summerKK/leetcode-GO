package leetcode

import "sort"

// 数组是已经排过序的,拿数组i和i-1作对比.如果重复就放在末尾.
func removeDuplicates0(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	l = l - 1
	k := l
	c := 0
	j := l + 1
	for l > 0 {
		if nums[l] == nums[l-1] {
			nums[k], nums[l] = nums[l], nums[k]
			k--
			c++
		}
		l--
	}

	newLen := j - c
	sort.Ints(nums[:newLen])

	return newLen
}

// 数组是已经排过序的,拿数组i和i-1作对比.如果重复就放在前面.
func removeDuplicates1(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	k := 0
	for i := 1; i < l; i++ {
		if nums[i] == nums[i-1] {
			// 这里要注意.把k和i-1的元素做变换.如果拿k和i的元素作替换并且i和i+1也相等的时候就会出问题
			nums[k], nums[i-1] = nums[i-1], nums[k]
			k++
		}
	}

	// 把数组反转,重复数据放在末尾
	for i := 0; i < k; i++ {
		nums[i], nums[l-1-i] = nums[l-1-i], nums[i]
	}

	// 对数组进行排序
	sort.Ints(nums[:l-k])

	return l - k
}

// 0,0,1,1,1,2,2,3,3,4
// 第一次 nums[k] = 0 nums[i] = 0 不满足条件
// 第二次 nums[k] = 0 nums[i] = 1 满足条件 使用 nums[k+1] 保存不重复元素
// 第三次 nums[k] = 1 nums[i] = 1 不满足条件
// 第四次 nums[k] = 1 nums[i] = 1 不满足条件
// 第五次 nums[k] = 1 nums[i] = 2 满足条件 使用 nums[k+1] 保存不重复元素
// 第六次 nums[k] = 2 nums[i] = 2 不满足条件
// .....
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[k] != nums[i] {
			k++
			nums[k], nums[i] = nums[i], nums[k]
		}
	}

	return k + 1
}
