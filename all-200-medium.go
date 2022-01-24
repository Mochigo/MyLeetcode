package main

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n {
			return
		}
		if grid[r][c] == '0' {
			return
		}
		grid[r][c] = '0'
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	ans := 0
	for i, g := range grid {
		for j := range g {
			if grid[i][j] == '1' {
				dfs(i, j)
				ans++
			}
		}
	}

	return ans
}
