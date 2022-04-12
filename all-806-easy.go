package main

func numberOfLines(widths []int, s string) []int {
	cur := 0
	line := 1
	for _, c := range s {
		val := widths[c-'a']
		if cur+val > 100 {
			cur = val
			line++
			continue
		}
		cur += val
	}
	return []int{line, cur}
}
