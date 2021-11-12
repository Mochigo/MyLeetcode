package main

import "math"

func getMoneyAmount(n int) int {
	// f[i][j] 表示从[i:j]区间内保证猜对的最小消费
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			minCost := math.MaxInt32
			// 这里是在列举每一个数是选中的数的可能性
			for x := i; x < j; x++ {
				cost := x + max(f[i][x-1], f[x+1][j])
				if cost < minCost {
					minCost = cost
				}
			}
			f[i][j] = minCost
		}
	}

	return f[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
