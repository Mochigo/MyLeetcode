package main

func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	left := make([]int, n)
	right := make([]int, n)
	for i := range security {
		if i > 0 && security[i] <= security[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 0
		}

		if n-i-1 < n-1 && security[n-i-1] <= security[n-i] {
			right[n-i-1] = right[n-i] + 1
		} else {
			right[n-i-1] = 0
		}
	}

	ans := []int{}
	for i := time; i < n-time; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}

	return ans
}
