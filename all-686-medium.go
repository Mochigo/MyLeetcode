package main

import "strings"

func repeatedStringMatch(a string, b string) int {
	s := &strings.Builder{}
	times := 0
	n := len(b)/len(a) + 1 // n是保证至少复制n次，可以使len(a) >= len(b)
	for times <= n+2 {
		s.WriteString(a)
		times++
		if strings.Contains(s.String(), b) {
			return times
		}
	}

	return -1
}
