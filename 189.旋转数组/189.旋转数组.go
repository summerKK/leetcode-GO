package leetcode

// A [1,2,3,4,5,6,7], l=7 且 k=3 输出: [5,6,7,1,2,3,4]
// B [5,6,7] C [1,2,3,4]
// A 翻转后 [7,6,5,4,3,2,1] D [7,6,5] E [4,3,2,1]
// D、E 继续翻转就得到B
func rotate(nums []int, k int) {
	//  避免k小于nums
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(nums []int) {
	l := len(nums)
	for i := 0; i < l/2; i++ {
		// 避免数组越界 (l - i l = 7 i= 0)
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
}
