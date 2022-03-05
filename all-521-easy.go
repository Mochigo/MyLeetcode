package main

func findLUSlength(a string, b string) int {
	// 统计双方的最大不同值
	if a != b {
		return max(len(a), len(b))
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
