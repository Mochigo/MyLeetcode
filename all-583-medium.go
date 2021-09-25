package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	// 求最长公共子序列 -- 1143就是一个基础解法题，考察最长公共子序列的写法
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return m + n - 2*dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	word1 := "sea"
	word2 := "eat"
	fmt.Println(minDistance(word1, word2))
}
