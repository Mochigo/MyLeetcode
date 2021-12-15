package main

// 这种一次遍历不成，通常需要多次遍历才能递推出完整的关系的情况，一般都要考虑拓扑排序，
// 拓扑排序要考虑入度和出度选一种，
func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	ans := make([]int, n)
	in := make([]int, n)
	g := make([][]int, n)

	for _, r := range richer {
		a, b := r[0], r[1]
		in[b]++
		g[a] = append(g[a], b)
	}

	// 有一种可能的结果，就是自己就是和自己一样或者更有钱的人里面最安静的
	for i := range ans {
		ans[i] = i
	}

	q := make([]int, 0, n)
	for i, deg := range in {
		if deg == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			//x比y都有钱，且比y安静，就更新y的值
			if quiet[ans[x]] < quiet[ans[y]] {
				ans[y] = ans[x]
			}
			in[y]--
			if in[y] == 0 {
				q = append(q, y)
			}
		}
	}

	return ans
}
