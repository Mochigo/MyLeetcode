package main

func longestSubsequence(arr []int, difference int) int {
	l := map[int]int{}
	max := 0
	for _, num := range arr {
		l[num] = l[num-difference] + 1
		if l[num] > max {
			max = l[num]
		}
	}

	return max
}
