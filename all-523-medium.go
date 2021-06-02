package main

import (
	"fmt"
)

func checkSubarraySum(nums []int, k int) bool {
	m := len(nums)
	if m < 2 {
		return false
	}
	mp := map[int]int{0: -1} // 0:-1 能确保出现nums[0:1]被k整除的情况能够正确返回
	remainder := 0
	for i, num := range nums {
		remainder = (remainder + num) % k //这里是基于前缀和 模 k 和前缀和每个值模 k 再模 k 的值相同
		if prevIndex, has := mp[remainder]; has {
			if i-prevIndex >= 2 {
				return true
			}
		} else {
			mp[remainder] = i //记录该余数第一次出现的位置
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 3
	fmt.Println(checkSubarraySum(nums, k))
}
