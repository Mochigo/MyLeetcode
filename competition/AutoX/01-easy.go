package main

func maxDistance(colors []int) int {
	max := 0
	n := len(colors)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if colors[i] != colors[j] && j-i > max {
				max = j - i
			}
		}
	}
	return max
}
