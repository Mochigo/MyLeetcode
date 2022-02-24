package main

func findBall(grid [][]int) []int {
	// 如何能滚到下一层，判断同一层的朝向，当当前为 1 时，其右侧必须也为1，否则小球无法下落，当小球
	// 这个问题，也是一个多源dfs
	n := len(grid[0])
	m := len(grid)
	ans := make([]int, n)
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x == m-1 {
			if y > 0 && grid[x][y] == -1 && grid[x][y-1] == -1 {
				return y - 1
			} else if y < n-1 && grid[x][y] == 1 && grid[x][y+1] == 1 {
				return y + 1
			} else {
				return -1
			}
		}

		if y == 0 && grid[x][y] == -1 || y == n-1 && grid[x][y] == 1 {
			return -1
		}

		if grid[x][y] == 1 && grid[x][y+1] == 1 {
			return dfs(x+1, y+1)
		} else if grid[x][y] == -1 && grid[x][y-1] == -1 {
			return dfs(x+1, y-1)
		}

		return -1
	}

	for i := 0; i < n; i++ {
		ans[i] = dfs(0, i)
	}

	return ans
}
