package main

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	sum := 0
	for i, num := range nums {
		if num < 0 && k > 0 {
			nums[i] = -num
			k--
		}
		sum += nums[i]
	}
	sort.Ints(nums)
	if k > 0 {
		if k%2 == 1 {
			sum -= 2 * nums[0]
		}
	}
	return sum
}
