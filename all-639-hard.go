package main

import "fmt"

func get1Digit(c1 byte) int {
	if c1 != '0' {
		if c1 == '*' {
			return 9
		} else {
			return 1
		}
	}
	return 0
}

func get2Digit(c1, c2 byte) int {
	// 两位
	if c2 == '*' {
		if c1 == '*' {
			return 15 // 11~19 21~26
		} else {
			if c1-'0' <= 6 {
				return 2
			} else {
				return 1
			}
		}
	} else if c2 == '1' {
		if c1 == '*' {
			return 9
		} else {
			return 1
		}
	} else if c2 == '2' {
		if c1 == '*' {
			return 6
		} else if c1-'0' <= 6 {
			return 1
		}
	}
	return 0
}

func numDecodings(s string) int {
	const mod = 1e9 + 7
	n := len(s)

	cur, p1, p2 := 0, 1, 0
	for i := 1; i <= n; i++ {
		cur = get1Digit(s[i-1]) * p1 % mod

		if i > 1 {
			cur = (cur + p2*get2Digit(s[i-1], s[i-2])) % mod
		}

		p1, p2 = cur, p1
	}

	return cur
}

func main() {
	fmt.Println(numDecodings("1*"))
}
