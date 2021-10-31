package main

import "unicode"

func findWords(words []string) (ans []string) {
	const rowIdx = "12210111011122000010020202"
	for _, word := range words {
		idx := rowIdx[unicode.ToLower(rune(word[0]))-'a']
		flag := false
		for _, ch := range word[1:] {
			if rowIdx[unicode.ToLower(ch)-'a'] != idx {
				flag = true
				break
			}
		}
		if !flag {
			ans = append(ans, word)
		}
	}
	return
}
