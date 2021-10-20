package main

import "fmt"

func minMoves(nums []int) int {
	// 题解，n-1个元素加一，相当于一个元素减一
	ans := 0
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	for _, num := range nums {
		ans += num - min
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(minMoves(nums))
}
