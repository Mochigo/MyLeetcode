package main

func minimumDeletions(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	min, max := nums[0], nums[0]
	l1, l2 := 0, 0
	for i, n := range nums {
		if n < min {
			min = n
			l1 = i
		}
		if n > max {
			max = n
			l2 = i
		}
	}

	s1 := smaller(l1, n-1-l1) + smaller(l2, n-1-l2) + 2
	s2 := smaller(larger(l1, l2), larger(n-1-l1, n-1-l2)) + 1
	return smaller(s1, s2)
}

func smaller(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func larger(a, b int) int {
	if a < b {
		return b
	}
	return a
}
