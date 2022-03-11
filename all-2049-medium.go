package main

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parents[i]
		children[p] = append(children[p], i)
	}

	cnt := 0
	max := 0

	// 非常巧妙的dfs
	// 这里的 dfs 用于统计子树节点个数，
	// 由于 总结点数一定，抛开自己，也就是剩下 n-1个节点，通过 n-1 - 所有子树相连的节点数之和可以得到与当前节点相连的父节点相连通的节点数
	var dfs func(int) int
	dfs = func(node int) int {
		score, size := 1, n-1
		for _, c := range children[node] {
			sz := dfs(c)
			score *= sz
			size -= sz
		}

		if size > 0 {
			score *= size
		}

		if score > max {
			max = score
			cnt = 1
		} else if score == max {
			cnt++
		}

		return n - size
	}

	dfs(0)
	return cnt
}
