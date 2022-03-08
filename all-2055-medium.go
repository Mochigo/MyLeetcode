package main

func platesBetweenCandles(s string, queries [][]int) []int {
	// 通过一次遍历，记录|的位置，并且记录当前|距离前一个|的距离
	n := len(s)
	data := []byte(s)
	cnt := 0
	preSum := make([]int, n)
	left := make([]int, n) // 记录每个 * 左侧最近的一个蜡烛位置
	l := -1
	for i, c := range data {
		if c == '*' {
			cnt++
		} else {
			l = i
		}

		left[i] = l
		preSum[i] = cnt
	}
	r := -1
	right := make([]int, n) // 记录每个 * 右侧最近的一个蜡烛位置
	for i := n - 1; i >= 0; i-- {
		if s[i] == '|' {
			r = i
		}
		right[i] = r
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		x, y := right[q[0]], left[q[1]]
		if x >= 0 && y >= 0 && x < y {
			ans[i] = preSum[y] - preSum[x]
		}
	}
	return ans
}
