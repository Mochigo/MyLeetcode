package main

func findCircleNum(isConnected [][]int) int {
	ans := 0
	n := len(isConnected)
	used := make([]bool, n)
	// 通过dfs找到所有与当前城市相连的城市
	var dfs func(int)
	dfs = func(c int) {
		if used[c] {
			return
		}
		used[c] = true
		for cn, v := range isConnected[c] {
			if v == 1 {
				dfs(cn)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !used[i] {
			dfs(i)
			ans++
		}
	}

	return ans
}
