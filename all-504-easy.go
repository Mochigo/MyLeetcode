package main

import "math"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	ans := []byte{}
	left := int(math.Abs(float64(num)))
	for left != 0 {
		ans = append(ans, byte('0'+left%7))
		left /= 7
	}
	if num < 0 {
		ans = append(ans, '-')
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return string(ans)
}
