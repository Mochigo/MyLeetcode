package main

import (
	"strings"
	"unicode"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)
	var ch [26]int
	for _, c := range licensePlate {
		if unicode.IsLetter(c) {
			ch[c-'a']++
		}
	}

	ans := ""
	for _, word := range words {
		inValid := false
		var cnt [26]int
		for _, c := range word {
			cnt[c-'a']++
		}

		for i := 0; i < 26; i++ {
			if cnt[i] < ch[i] {
				inValid = true
				break
			}
		}

		if !inValid && (ans == "" || len(ans) > len(word)) {
			ans = word
		}
	}

	return ans
}
