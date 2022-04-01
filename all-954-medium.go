package main

import "sort"

func canReorderDoubled(arr []int) bool {
	// 统计能否找到 len(arr)个2倍的数对
	cnt := map[int]int{}
	vals := make([]int, 0, len(arr))

	for _, n := range arr {
		cnt[n]++
	}

	// 当 0 为奇数个，必定找不到
	if cnt[0]%2 == 1 {
		return false
	}

	delete(cnt, 0)
	for num := range cnt {
		vals = append(vals, num)
	}
	sort.Slice(vals, func(a, b int) bool { return abs(vals[a]) < abs(vals[b]) })

	for _, num := range vals {
		if cnt[num] > cnt[2*num] {
			return false
		}
		cnt[2*num] -= cnt[num]
	}

	return true
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
