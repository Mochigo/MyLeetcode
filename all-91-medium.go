package main

import "fmt"

func numDecodings(s string) int {
	n := len(s)
	// f[i] = cur, f[i-1] = p1, f[i-2] = p2
	cur, p1, p2 := 0, 1, 0
	for i := 1; i <= n; i++ {
		// 这里是因为原来用 f[i] 的时候初始的默认值为 0
		cur = 0
		if s[i-1] != '0' {
			cur += p1
		}

		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+s[i-1]-'0') <= 26 {
			cur += p2
		}

		p1, p2 = cur, p1
	}

	return cur
}

func main() {
	fmt.Println(numDecodings("11106"))
}
