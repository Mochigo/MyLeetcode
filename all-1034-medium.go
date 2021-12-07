package main

// 找到联通分量，判断是否在矩阵边界上以及四周是否存在不属于该联通分量的色块
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func colorBorder1(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	borders := [][]int{}
	originColor := grid[row][col]
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(int, int)
	dfs = func(x, y int) {
		visited[x][y] = true
		isBorder := false
		// 这里是另一个条件， 如果是联通分量的一部分，且在位于矩形的边界上，那么它也应被上色
		if x == 0 || x == m-1 || y == 0 || y == n-1 {
			isBorder = true
		}
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			// 这里保证下标不越界
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				// 这里来确认作为联通分量的一部分，是否是边界
				if grid[nx][ny] != originColor {
					isBorder = true
				} else if !visited[nx][ny] {
					dfs(nx, ny)
				}
			}
		}

		if isBorder {
			borders = append(borders, []int{x, y})
		}
	}
	dfs(row, col)

	for _, p := range borders {
		grid[p[0]][p[1]] = color
	}

	return grid
}

type pair struct{ x, y int }

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	borders := []pair{}
	originColor := grid[row][col]
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	q := []pair{pair{row, col}}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		visited[p.x][p.y] = true
		isBorder := false

		if p.x == 0 || p.x == m-1 || p.y == 0 || p.y == n-1 {
			isBorder = true
		}

		for _, dir := range dirs {
			nx, ny := p.x+dir.x, p.y+dir.y
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				if grid[nx][ny] != originColor {
					isBorder = true
				} else if !visited[nx][ny] {
					q = append(q, pair{nx, ny})
				}
			}
		}

		if isBorder {
			borders = append(borders, pair{p.x, p.y})
		}
	}

	for _, p := range borders {
		grid[p.x][p.y] = color
	}

	return grid
}
