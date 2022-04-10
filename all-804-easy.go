package main

import "strings"

func uniqueMorseRepresentations(words []string) int {
	code := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	cnt := map[string]struct{}{}
	var res strings.Builder
	for _, word := range words {
		for _, c := range word {
			res.WriteString(code[c-'a'])
		}
		cnt[res.String()] = struct{}{}
		res.Reset()
	}

	return len(cnt)
}
