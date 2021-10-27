package main

import "fmt"

func removeInvalidParentheses(s string) []string {
	ans := []string{}
	lremove, rremove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lremove++
		} else if ch == ')' {
			if lremove > 0 {
				lremove--
			} else {
				rremove++
			}
		}
	}

	var dfs func(string, int, int, int)
	dfs = func(str string, start, lremove, rremove int) {
		if lremove == 0 && rremove == 0 {
			if isValid(str) {
				ans = append(ans, str)
				return
			}
		}

		for i := start; i < len(str); i++ {
			// 用于去重，如果当前字符和前一个字符相同，就跳过，避免出现同样的情况，达到事实上的去重
			if i != start && str[i] == str[i-1] {
				continue
			}
			// 如果剩余需要删除的左右括号数量多余字符串剩余长度，直接退出
			if lremove+rremove > len(str)-i {
				return
			}
			// 尝试去除左右括号
			if lremove > 0 && str[i] == '(' {
				dfs(str[:i]+str[i+1:], i, lremove-1, rremove)
			}
			if rremove > 0 && str[i] == ')' {
				dfs(str[:i]+str[i+1:], i, lremove, rremove-1)
			}
		}
	}

	dfs(s, 0, lremove, rremove)
	return ans
}

func isValid(s string) bool {
	cnt := 0
	for _, ch := range s {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			if cnt == 0 {
				return false
			}

			cnt--
		}
	}
	return cnt == 0
}

func main() {
	s := "(a)())()"

	fmt.Println(removeInvalidParentheses(s))
}
