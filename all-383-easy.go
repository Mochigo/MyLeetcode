package main

func canConstruct(ransomNote string, magazine string) bool {
	var c [26]int

	for _, ch := range []byte(magazine) {
		c[ch-'a']++
	}

	for _, ch := range []byte(ransomNote) {
		c[ch-'a']--
		if c[ch-'a'] < 0 {
			return false
		}
	}

	return true
}
