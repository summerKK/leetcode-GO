package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	l := len(nums) - 1
	var result [][]int
	for i := 0; i < l; i++ {
		// 先找一个固定值.顺序在md文件中
		target := 0 - nums[i]
		right := l
		left := i + 1
		// 当固定值大于0的时候left和right值一定大于0(数组是排序过的).已经没有符合条件的值了
		if nums[i] > 0 {
			break
		}
		// 如果 nums[i] == nums[i-1]说明target值也是一样的.计算的left和right的值也是一样的.所以需要去重
		// i == 0 代表第一次.无需判断
		if i == 0 || nums[i] != nums[i-1] {
			// 循环中 left++ ,right--
			for left < right {
				// 目标值
				if nums[left]+nums[right] == target {
					result = append(result, []int{nums[i], nums[left], nums[right]})
					// 如果nums[left] == nums[left+1].下次循环已经没有意义.直接跳过
					// 注意left是往右.所以要判断left和left+1
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					// 如果nums[right] == nums[right-1].下次循环已经没有意义.直接跳过
					// 注意right是往左.所以要判断right和right-1
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if nums[right]+nums[left] < target {
					// nums[right]+nums[left] < target 代表值偏小.left++.数组是排过序的.左边的永远小于右边的值
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}
