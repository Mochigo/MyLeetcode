package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	ans := "1"
	for i := 2; i <= n; i++ {
		tmp := ""
		for j, start := 0, 0; j < len(ans); start = j {
			for j < len(ans) && ans[start] == ans[j] {
				j++
			}
			tmp += strconv.Itoa(j-start) + string(ans[start])
		}
		ans = tmp
	}

	return ans
}

// 使用了strings.Builder{}的官方题解方法
func countAndSay2(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			cur.WriteString(strconv.Itoa(j - start))
			cur.WriteByte(prev[start])
		}
		prev = cur.String()
	}
	return prev
}

func main() {
	fmt.Println(countAndSay(15))
}
