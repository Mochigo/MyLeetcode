package main

import "fmt"

func reverseOnlyLetters(s string) string {
	chs := []byte(s)
	// letters := "abcdefghijklmnopqrstuvwxyz"
	// letters = letters + strings.ToUpper(letters)

	// 判断字母更适合另外写一个函数使用范围来比较，例如这里的isLetter
	// 因为string底层是一个结构体，其中len越长，会使变量占用更多的内存空间
	l, r := 0, len(chs)-1
	for l < r {
		for l < r && !isLetter(chs[l]) {
			l++
		}

		for l < r && !isLetter(chs[r]) {
			r--
		}

		if l < r {
			chs[l], chs[r] = chs[r], chs[l]
			l++
			r--
		}
	}

	return string(chs)
}

func isLetter(s byte) bool {
	return (s >= 'a' && s <= 'z') || (s <= 'Z' && s >= 'A')
}

func main() {
	word := "da-ab"
	fmt.Println(reverseOnlyLetters(word))
}
