package main

import "strings"

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	cnt := 0
	flag := false

	var ch1, ch2 [26]int
	for i := range s {
		ch1[s[i]-'a']++
		ch2[goal[i]-'a']++
		if s[i] != goal[i] {
			cnt++
		} else if !flag && ch1[s[i]-'a'] > 1 {
			flag = true
		}
	}

	// 组成元素不同直接返回false
	if ch1 != ch2 {
		return false
	}

	return (strings.Compare(s, goal) == 0 && flag) || cnt == 2
}
