package main

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	sum := mean * (m + n)
	for _, roll := range rolls {
		sum -= roll
	}

	ans := make([]int, n)

	// 先取平均值
	avg := sum / n
	if avg >= 7 || avg <= 0 {
		return []int{}
	}
	// 先往n个里面，每个填 1， 总共能填 avg 次
	for i := 0; i < n; i++ {
		ans[i] = avg
	}

	sum -= avg * n
	if avg == 6 && sum > 0 {
		return []int{}
	}

	idx := 0
	for sum > 0 {
		sum--
		ans[idx]++
		idx++
	}
	return ans
}
