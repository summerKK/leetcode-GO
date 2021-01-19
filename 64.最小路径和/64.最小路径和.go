package leetcode

func minPathSum(grid [][]int) int {
	if len(grid) <= 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	// 对边界进行初始化(m)
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 对边界进行初始化(n)
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			//   2(3) 3(6)
			//   5(8) x
			// ----------
			// 括号里面是 (min(dp[i-1][j], dp[i][j-1]) + grid[i][j])
			// 取i、j最小的路径
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func min(m, n int) int {
	if m > n {
		return n
	}

	return m
}
