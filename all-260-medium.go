package main

import "sort"

func singleNumber1(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	num1, num2 := 0, 0
	i := 0
	for i < n {
		if i+1 < n && nums[i] == nums[i+1] {
			i += 2
		} else {
			num1, num2 = nums[i], num1
			i++
		}
	}

	return []int{num1, num2}
}

func singleNumber2(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	lsb := xorSum & -xorSum
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	return []int{type1, type2}
}
