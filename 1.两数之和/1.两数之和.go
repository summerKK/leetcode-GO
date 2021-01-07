package leetcode

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	var result []int
	m := make(map[int]int)
	for k, v := range nums {
		// 如果在hash表找到target-v即是结果.如果找不到就把当前元素加入到hash表
		if kk, exist := m[target-v]; exist {
			result = []int{k, kk}
			return result
		}
		m[v] = k
	}

	return result
}
