package main

func maxValue1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += max(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}

func maxValue2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// dp表示达到该点能够取得的最大价值的礼物
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func maxValue3(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dpNext := make([]int, n)
	start := make([]int, m)

	start[0] = grid[0][0]
	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		start[i] = start[i-1] + grid[i][0]
		dpNext[0] = start[i]
		for j := 1; j < n; j++ {
			dpNext[j] = max(dpNext[j-1], dp[j]) + grid[i][j]
		}

		dp = dpNext
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
