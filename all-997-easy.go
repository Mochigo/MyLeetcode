package main

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	beTrust := map[int]int{}

	for _, t := range trust {
		b := t[1]
		beTrust[b]++
		if beTrust[b] == n-1 {
			for _, t1 := range trust {
				if t1[0] == b {
					return -1
				}
			}
			return b
		}
	}

	return -1
}
