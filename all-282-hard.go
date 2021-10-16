package main

import (
	"fmt"
	"strconv"
)

func addOperators(num string, target int) []string {
	ans := []string{}
	n := len(num)
	var dfs func(string, int, int, int)
	// expr 用来记录当前的表达式
	dfs = func(expr string, index, cur, mul int) {
		if index == n {
			if cur == target {
				ans = append(ans, expr)
			}
			return
		}
		for i, start := index, index; i < n; i++ {
			// 排除了枚举从当前位所有长度大于等于二的数字 且 第一个数字为0的情况
			if num[start] == '0' && i > start {
				break
			}
			val, _ := strconv.Atoi(string(num[start : i+1]))

			// 如果是 第一位 就把数字直接作为表达式的开始
			if index == 0 {
				dfs(expr+num[start:i+1], i+1, val, val)
			} else {
				dfs(expr+"+"+num[start:i+1], i+1, cur+val, val)
				dfs(expr+"-"+num[start:i+1], i+1, cur-val, -val)
				// 特别地, 因为乘法运算是高优先级的运算，所以用mul来记录最后的连乘结果，对于+和-运算，mul = val 或者 -val
				dfs(expr+"*"+num[start:i+1], i+1, cur-mul+mul*val, val*mul)
			}
		}
	}

	dfs("", 0, 0, 0)

	return ans
}

// 题解版，极大的性能提升！
func addOperators2(num string, target int) []string {
	ans := []string{}
	n := len(num)

	var dfs func([]byte, int, int, int) // 用[]byte代替string，对于expr的修改更为简单方便
	dfs = func(expr []byte, index, cur, mul int) {
		if index == n {
			if cur == target {
				ans = append(ans, string(expr)) //利用了[]byte切片可以方便的转化成string数组
			}
			return
		}
		opPos := len(expr)
		if index > 0 {
			expr = append(expr, '0')
		}
		for i, start, val := index, index, 0; i < n; i++ {
			// 排除了枚举从当前位所有可能长度数字第一个为0的情况
			if num[start] == '0' && i > start {
				break
			}
			val = val*10 + int(num[i]-'0')
			expr = append(expr, num[i]) // 这里不需要把这个弹出，是因为每一次都是一个一个增加，且枚举的长度就是从小到大
			if index == 0 {
				dfs(expr, i+1, val, val)
			} else {
				expr[opPos] = '+'
				dfs(expr, i+1, cur+val, val)

				expr[opPos] = '-'
				dfs(expr, i+1, cur-val, -val)

				expr[opPos] = '*'
				dfs(expr, i+1, cur-mul+mul*val, val*mul)
			}
		}
	}

	dfs(make([]byte, 0, n*2-1), 0, 0, 0) // 一个式子可能的最大容量是2*n-1，避免了append操作销毁，转移，重分配，赋值所造成的性能损失

	return ans
}

func main() {
	num := "1234"
	target := 10
	fmt.Println(addOperators(num, target))
}
