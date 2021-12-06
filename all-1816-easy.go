package main

import "strings"

func truncateSentence(s string, k int) string {
	words := strings.Split(s, " ")
	ans := &strings.Builder{}
	for i := 0; i < k; i++ {
		ans.WriteString(words[i])
		if i < k-1 {
			ans.WriteByte(' ')
		}
	}

	return ans.String()
}

func truncateSentence1(s string, k int) string {
	n, end, count := len(s), 0, 0
	for i := 1; i <= n; i++ {
		if i == n || s[i] == ' ' {
			count++
			if count == k {
				end = i
				break
			}
		}
	}
	return s[:end]
}
