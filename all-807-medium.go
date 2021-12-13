package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	colMax := make([]int, n)
	rowMax := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			colMax[i] = max(colMax[i], grid[j][i])
			rowMax[i] = max(rowMax[i], grid[i][j])
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans += min(colMax[j], rowMax[i]) - grid[i][j]
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
