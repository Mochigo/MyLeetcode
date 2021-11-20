package main

import "sort"

func findLHS1(nums []int) int {
	// 只关注数值，虽说按照子序列的定义，是不可以排序的，但考虑到题意，只需要返回长度，就先进行一次排序
	sort.Ints(nums)

	n := len(nums)
	max := 0
	for i, start := 0, 0; i < n; start = i {
		// 怎么保存中间的部分？它会两个为一组，两个为一组的跳跃
		for i < n && nums[i] == nums[start] {
			i++
		}

		j := i
		for j < n && nums[j] == nums[start]+1 {
			j++
		}

		if j == i {
			continue
		}

		if j-start > max {
			max = j - start
		}
	}
	// 这种合并导致了一种数字的跳过

	return max
}

func findLHS2(nums []int) int {
	cnt := map[int]int{}
	max := 0

	for _, num := range nums {
		cnt[num]++
	}

	for num, v1 := range cnt {
		if v2, ok := cnt[num+1]; ok && v1+v2 > max {
			max = v1 + v2
		}
	}

	return max
}
