package main

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func islandPerimeter(grid [][]int) int {
	// 从岛屿走向非岛屿的部分，+1即可
	ans := 0
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 {
			ans++
			return
		}
		if grid[r][c] == 2 {
			return
		}
		grid[r][c] = 2
		for _, dir := range dirs {
			dfs(r+dir.x, c+dir.y)
		}
	}

	for i, g := range grid {
		for j := range g {
			if grid[i][j] == 1 {
				dfs(i, j)
				return ans
			}
		}
	}

	return ans
}
