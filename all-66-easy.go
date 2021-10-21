package main

import "fmt"

func plusOne(digits []int) []int {
	cur := 1
	index := len(digits) - 1
	for {
		digits[index] += cur
		cur = digits[index] / 10
		digits[index] %= 10
		index--
		if index < 0 || cur == 0 {
			break
		}
	}

	if cur == 1 {
		return append([]int{1}, digits...)
	}

	return digits
}

func main() {
	nums := []int{9, 9, 9}
	fmt.Println(plusOne(nums))
}
