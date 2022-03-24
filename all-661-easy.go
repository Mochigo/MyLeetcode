package main

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	ans := make([][]int, m)
	for i, g := range img {
		ans[i] = make([]int, n)
		for j := range g {
			sum := 0
			cnt := 0
			for _, list := range img[max(0, i-1):min(i+2, m)] {
				for _, v := range list[max(0, j-1):min(j+2, n)] {
					sum += v
					cnt++
				}
			}
			ans[i][j] = sum / cnt
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
	if a < b {
		return a
	}
	return b
}
