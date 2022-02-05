package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func getMaximumGold1(grid [][]int) (ans int) {
	var dfs func(x, y, gold int)
	dfs = func(x, y, gold int) {
		gold += grid[x][y]
		if gold > ans {
			ans = gold
		}

		rec := grid[x][y]
		grid[x][y] = 0
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[nx]) && grid[nx][ny] > 0 {
				dfs(nx, ny, gold)
			}
		}
		grid[x][y] = rec
	}

	for i, row := range grid {
		for j, gold := range row {
			if gold > 0 {
				dfs(i, j, 0)
			}
		}
	}
	return
}

func MyGetMaximumGold(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	var dfs func(int, int, int)
	dfs = func(x, y, cur int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
			if cur > ans {
				ans = cur
			}
			return
		}
		before := grid[x][y]
		grid[x][y] = 0
		for _, d := range dirs {
			newX, newY, val := x+d.x, y+d.y, cur+before
			dfs(newX, newY, val)
		}
		grid[x][y] = before
	}

	for i, g := range grid {
		for j := range g {
			if grid[i][j] != 0 {
				dfs(i, j, 0)
			}
		}
	}

	return ans
}
