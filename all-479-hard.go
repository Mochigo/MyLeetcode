package main

import "math"

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	max := int(math.Pow10(n)) - 1 // 9876 6 7 8 9
	for cur := max; ; cur-- {
		ans := cur
		for x := cur; x > 0; x /= 10 {
			ans = ans*10 + x%10
		}

		for p := max; p*p >= ans; p-- {
			if ans%p == 0 {
				return ans % 1337
			}
		}
	}
}
