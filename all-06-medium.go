package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	grid := make([][]byte, numRows)
	idx := 0
	flag := 1
	for _, c := range []byte(s) {
		grid[idx] = append(grid[idx], c)
		if idx == 0 {
			flag = 1
		}
		if idx == numRows-1 {
			flag = -1
		}
		idx += flag
	}

	var ans string
	for _, g := range grid {
		ans += string(g)
	}

	return ans
}
