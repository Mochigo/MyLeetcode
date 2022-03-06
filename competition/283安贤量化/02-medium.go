package main

import "sort"

func minimalKSum(nums []int, k int) int64 {
	a := nums
	a = append(a, 0, 2e9)

	var ans int64
	sort.Ints(a)
	for i := 1; i < len(a); i++ {
		fill := a[i] - a[i-1] - 1
		if fill <= 0 {
			continue
		}

		if fill >= k {
			ans += int64((a[i-1]*2 + k + 1) * k / 2)
			break
		}

		ans += int64((a[i-1] + a[i]) * fill / 2)
		k -= fill
	}

	return ans
}
