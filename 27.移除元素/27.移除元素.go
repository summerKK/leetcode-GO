package leetcode

// 3,2,2,3 val = 3
// 倒叙遍历 k初始化指向最后一个元素下标
// 第一次 nums[l] = 3  val 3  相等. 把 nums[l] 放在第k位置上面.此时应该把nums[k] 位置的元素放在第l位置上面
// 第二次 nums[l] = 2  val 3  不相等. 跳过
// 第三次 nums[l] = 2  val 3  不相等. 跳过
// 第四次 nums[l] = 3  val 3  相等. 把 nums[l](3) 放在第k(3-1)位置上面.此时应该把nums[k](2) 位置的元素放在第l(0)位置上面
func removeElement(nums []int, val int) int {
	l := len(nums) - 1
	k := l
	for l >= 0 {
		if nums[l] == val {
			// 交换k和l对应元素的顺序.这样相等的元素都排在最后面
			nums[k], nums[l] = val, nums[k]
			k--
		}
		l--
	}

	return k + 1
}
