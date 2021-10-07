package main

import "fmt"

func countSegments(s string) int {
	ans := 0
	for i, c := range s {
		if c != ' ' && (i == 0 || s[i-1] == ' ') {
			ans++
		}
	}

	return ans
}

func main() {
	s := "Hello, my name is Jojo"
	fmt.Println(countSegments(s))
}
