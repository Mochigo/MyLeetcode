package main

import (
	"math/bits"
	"sort"
)

func countExcellentPairs1(nums []int, k int) int64 {
	// a | b + a　& b = a + b
	// 可以转换成 求 a的1个数和b的1个数之和
	v := -1
	sort.Ints(nums)
	cnts := make(map[int]int) // 统计 1个数相同的数量
	for _, n := range nums {  //O(N)
		if v != n {
			v = n // 排序+记录值变化=去重，去重是因为题干把相同的数字看成一种处理
			cnts[bits.OnesCount(uint(v))]++
		}
	}
	// 10 ^ 9 > 1 << 29
	// 10 ^ 9 < 1 << 30
	// 这里因为一个数其一的个数是有范围的，例如int32，我们可以把1个个数范围看成[1, 31]，map不会很大，这里可以把复杂度看成O(30^2)
	var ans int64
	for x, cx := range cnts {
		for y, cy := range cnts {
			if x+y >= k {
				ans += int64(cx * cy)
			}
		}
	}

	return ans
}

// 前缀和优化
func countExcellentPairs2(nums []int, k int) int64 {
	// a | b + a　& b = a + b
	// 可以转换成 求 a的1个数和b的1个数之和
	v := -1
	sort.Ints(nums)
	// 10 ^ 9 > 1 << 29
	// 10 ^ 9 < 1 << 30
	var cnts [30]int         // 上等式可知，1个数不会超过30
	for _, n := range nums { //O(N)
		if v != n {
			v = n // 排序+记录值变化=去重，去重是因为题干把相同的数字看成一种处理
			cnts[bits.OnesCount(uint(v))]++
		}
	}
	var s int
	if k <= 30 {
		s = sum(cnts[k:]) // 求1个数就能满足题意得所有个数, 排除k>30导致panic的情况
	}

	var ans int64
	for x, cx := range cnts {
		ans += int64(cx * s)
		if k-x-1 >= 0 && k-x-1 < 30 {
			s += cnts[k-x-1]
		}
	}

	return ans
}

func sum(nums []int) int {
	ret := 0
	for _, v := range nums {
		ret += v
	}

	return ret
}
