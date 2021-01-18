package leetcode

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	n0 := 1
	n1 := 2
	result := 2
	// 上 n 阶台阶，到达第n阶的方法总数就是到第 (n-1) 阶和第 (n-2) 阶的方法数之和。
	for i := 3; i <= n; i++ {
		result = n1 + n0
		n0, n1 = n1, result
	}

	return result
}
