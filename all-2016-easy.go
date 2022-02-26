package main

import "math"

func maximumDifference(nums []int) int {
	min := math.MaxInt32
	ans := -1
	for _, num := range nums {
		if num <= min {
			min = num
			continue
		}

		if num-min > ans {
			ans = num - min
		}
	}

	return ans
}
