package main

type pos struct{ x, y int }

var dirs = []pos{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	layer := 0
	poss := []pos{pos{0, 0}}
	for len(poss) > 0 {
		layer++
		tmp := []pos{}
		for _, p := range poss {
			x, y := p.x, p.y
			if x < 0 || x >= n || y < 0 || y >= n || grid[x][y] == 1 || visited[x][y] {
				continue
			}
			if x == n-1 && y == n-1 {
				return layer
			}
			visited[x][y] = true
			for _, dir := range dirs {
				newX, newY := x+dir.x, y+dir.y
				tmp = append(tmp, pos{newX, newY})
			}
		}
		poss = tmp
	}

	return -1
}
