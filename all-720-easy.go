package main

import "sort"

func longestWord(words []string) string {
	sort.Slice(words, func(a, b int) bool {
		w1, w2 := words[a], words[b]
		return len(w1) < len(w2) || len(w1) == len(w2) && w1 > w2
	})

	// 还有一个限制：必须从一个字符开始
	ans := ""
	set := map[string]struct{}{"": {}}
	for _, word := range words {
		if _, ok := set[word[:len(word)-1]]; ok {
			ans = word
			set[word] = struct{}{}
		}
	}

	return ans
}
