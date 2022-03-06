package main

import "strings"

func cellsInRange(s string) []string {
	sigs := strings.Split(s, ":")
	col, row := sigs[0], sigs[1]
	a1, b1, a2, b2 := col[0], col[1], row[0], row[1]

	ans := []string{}
	for i := a1; i <= a2; i++ {
		for j := b1; j <= b2; j++ {
			ans = append(ans, string([]byte{i, j}))
		}
	}

	return ans
}
