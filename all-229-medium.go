package main

import "sort"

func majorityElement1(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	bound := n / 3
	ans := make([]int, 0, 2)
	for j, start := 0, 0; j < n; start = j {
		for j < n && nums[j] == nums[start] {
			j++
		}

		if j-start > bound {
			ans = append(ans, nums[start])
		}
	}

	return ans
}

func majorityElement2(nums []int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	ans := make([]int, 0, 2)

	bound := len(nums) / 3
	for num, val := range count {
		if val > bound {
			ans = append(ans, num)
		}
	}

	return ans
}

func majorityElement3(nums []int) []int {
	// 前提 求出现次数超过 n/k 次的元素，则最多有k-1个元素
	c1, c2 := 0, 0
	vote1, vote2 := 0, 0
	for _, num := range nums {
		if vote1 > 0 && num == c1 {
			vote1++
		} else if vote2 > 0 && num == c2 {
			vote2++
		} else if vote1 == 0 {
			c1 = num
			vote1 = 1
		} else if vote2 == 0 {
			c2 = num
			vote2 = 1
		} else {
			vote1--
			vote2--
		}
	}

	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		if num == c1 {
			cnt1++
		}
		if num == c2 {
			cnt2++
		}
	}

	ans := make([]int, 0, 2)
	if vote1 > 0 && cnt1*3 > len(nums) {
		ans = append(ans, c1)
	}
	if vote2 > 0 && cnt2*3 > len(nums) {
		ans = append(ans, c2)
	}

	return ans
}

func main() {

}
