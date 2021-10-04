package main

import (
	"fmt"
	"unicode"
)

func licenseKeyFormatting(s string, k int) string {
	ans := make([]byte, 0, len(s)) // 这里因为可以确定答案的长度不会超过s，所以可以直接指定容量，不需要后续重新分配内存，对时间和空间复杂度有很大的提升
	// 这里是逆序，所以最后有一个反转的操作
	for i, cnt := len(s)-1, 0; i >= 0; i-- {
		if s[i] != '-' {
			ans = append(ans, byte(unicode.ToUpper(rune(s[i]))))
			cnt++
			if cnt%k == 0 {
				ans = append(ans, '-')
			}
		}
	}
	// 这里是为了去掉，当正好分配完成的情况下多余的一个'-'
	if len(ans) > 0 && ans[len(ans)-1] == '-' {
		ans = ans[:len(ans)-1]
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
}
