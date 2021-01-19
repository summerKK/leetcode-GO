package leetcode

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	// 当 n = 0 机器人只能往右走.只有一种结果
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	// 当 m = 0 机器人只能往下走.只有一种结果
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	//
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m-1][n-1]
}
