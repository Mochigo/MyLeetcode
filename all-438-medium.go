package main

func findAnagrams(s string, p string) []int {
	// 滑动窗口的大小由 p 决定
	if len(p) > len(s) {
		return nil
	}
	var c1, c2 [26]int
	tmp := []byte(s)
	for _, ch := range []byte(p) {
		c1[ch-'a']++
	}
	for _, ch := range tmp[:len(p)] {
		c2[ch-'a']++
	}

	ans := []int{}
	if c1 == c2 {
		ans = append(ans, 0)
	}

	for l, r := 0, len(p); r < len(s); l, r = l+1, r+1 {
		c2[tmp[l]-'a']--
		c2[tmp[r]-'a']++
		if c1 == c2 {
			ans = append(ans, l+1)
		}
	}

	return ans
}
