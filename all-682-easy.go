package main

import "strconv"

func calPoints(ops []string) int {
	nums := make([]int, 0)
	for _, c := range ops {
		switch c {
		case "D":
			num := nums[len(nums)-1]
			nums = append(nums, num*2)
		case "C":
			nums = nums[:len(nums)-1]
		case "+":
			n1, n2 := nums[len(nums)-1], nums[len(nums)-2]
			nums = append(nums, n1+n2)
		default:
			num, _ := strconv.Atoi(c)
			nums = append(nums, num)
		}
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}
