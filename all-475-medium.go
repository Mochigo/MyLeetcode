package main

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	ans := 0

	// 这里循环是为了找到每个房子与最近加热器的距离，然后在这些距离中的最大值就一定能让所有房子被加热器笼罩
	for _, h := range houses {
		j := sort.SearchInts(heaters, h+1) // 加一是为了确保，j如果在房子后面，一定是房子在所有加热器的后面，不然房子后面一定有一个加热器
		minDis := math.MaxInt32

		// 如果j < len(headers)说明，存在一个加热器在当前房子的后面
		if j < len(heaters) {
			minDis = heaters[j] - h
		}
		i := j - 1 // 前一个 加热器的位子
		// 如果存在前一个加热器的位子，就比较前一个加热器和房子的距离
		if i >= 0 && h-heaters[i] < minDis {
			minDis = h - heaters[i]
		}
		if minDis > ans {
			ans = minDis
		}
	}

	return ans
}
