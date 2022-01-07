package main

import "strings"

func simplifyPath(path string) string {
	s := strings.Split(path, "/")
	stack := make([]string, 0, len(s))
	for _, word := range s {
		if len(word) == 0 {
			continue
		}
		if word == "." {
			continue
		} else if word == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, word)
		}
	}

	ret := &strings.Builder{}
	for _, word := range stack {
		ret.WriteByte('/')
		ret.WriteString(word)
	}
	if len(stack) == 0 {
		ret.WriteByte('/')
	}
	return ret.String()
}
