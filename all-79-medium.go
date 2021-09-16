package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	dirs := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	vis := make([][]bool, m)

	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var dfs func(int, int, int) bool
	dfs = func(x, y, k int) bool {
		if board[x][y] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		vis[x][y] = true
		defer func() { vis[x][y] = false }() // 特别的是 defer只能跟着函数调用 记得再最后面加一个()表示启用
		for _, dir := range dirs {
			// 这里保证了新的点位，无重复使用，且下标符合规范
			if newX, newY := x-dir[0], y-dir[1]; newX >= 0 && newX < m && newY >= 0 && newY < n && !vis[newX][newY] {
				if dfs(newX, newY, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i, row := range board {
		for j := range row {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"

	fmt.Println(exist(board, word))
}
