package main

func missingNumber(nums []int) int {
	sum := 0
	n := len(nums)
	for _, num := range nums {
		sum += num
	}

	return n*(n+1)/2 - sum
}
