package main

func construct2DArray(original []int, m int, n int) [][]int {
	s := len(original)
	if s != m*n {
		return [][]int{}
	}

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i, val := range original {
		ans[i/n][i%n] = val

	}

	return ans
}
