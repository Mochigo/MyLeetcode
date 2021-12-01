package main

func maxPower(s string) int {
	max := 1
	n := len(s)
	for i, start := 0, 0; i < n; start = i {
		for i < n && s[i] == s[start] {
			i++
		}

		if i-start > max {
			max = i - start
		}
	}

	return max
}
