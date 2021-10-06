package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	max, cur := nums[n-1], nums[n-1]
	k := 1
	for i := n - 2; k < 3 && i >= 0; i-- {
		if nums[i] != cur {
			k++
			cur = nums[i]
		}
	}

	if k != 3 {
		return max
	}
	return cur
}

func main() {
	nums := []int{1, 2, 2, 3, 4, 4, 4, 5, 6}
	fmt.Println(thirdMax(nums))
}
