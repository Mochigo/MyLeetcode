package main

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	l := sort.SearchInts(nums, target)
	ans := []int{}
	if l == len(nums) || nums[l] != target {
		return ans
	}
	ans = append(ans, l)
	for i := l + 1; i < len(nums); i++ {
		if nums[i] == target {
			ans = append(ans, i)
		}
	}
	return ans
}
