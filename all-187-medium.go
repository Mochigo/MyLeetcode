package main

import "fmt"

// 利用了固定的滑动窗口，以及截取切片作为map的key的操作
func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	if n <= 10 {
		return nil
	}
	// 固定长度 -> 滑动窗口
	count := map[string]int{}
	ans := []string{}
	for l, r := 0, 10; r <= n; l, r = l+1, r+1 {
		tmp := s[l:r]
		count[tmp]++
		if count[tmp] == 2 {
			ans = append(ans, tmp)

		}
	}

	return ans
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequences(s))
}
