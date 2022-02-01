package main

import "unicode"

func longestNiceSubstring(s string) string {
	n := len(s)
	var ans string
	for i := 0; i < n; i++ {
		lower, upper := 0, 0
		for j := i; j < n; j++ {
			if unicode.IsLower(rune(s[j])) {
				lower |= 1 << (s[j] - 'a')
			} else {
				upper |= 1 << (s[j] - 'A')
			}

			if lower == upper && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}

	return ans
}
