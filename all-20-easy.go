package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	brackets := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	for _, ch := range []byte(s) {
		if len(stack) > 0 && stack[len(stack)-1] == brackets[ch] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}

func main() {
	s := "[]{}()"
	fmt.Println(isValid(s))
}
