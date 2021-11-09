package main

import (
	"fmt"
	"strings"
)

func findMinStep(board string, hand string) int {
	queue := make([][]string, 0)
	queue = append(queue, []string{board, hand})

	set := make(map[string]bool)
	steps := 0

	for len(queue) > 0 {
		// 弹出第一个元素
		tmp := [][]string{}
		steps++
		for _, item := range queue {
			b, h := item[0], item[1]
			m, n := len(b), len(h)
			// 遍历手中的每个球
			for i := 0; i < n; i++ {
				// 遍历board中的每个位置
				for j := 0; j < m; j++ {
					// 剪枝：手中的球与上一个球颜色相同，
					if i >= 1 && h[i] == h[i-1] {
						continue
					}
					// 剪枝：仅当有可能优化的时候决定插入新球
					// 情况一：当前球颜色和即将插入的位置颜色相同
					// 情况二：当前球 和 两侧的球颜色不同，但是两侧球颜色相同
					isPlace := false
					if h[i] == b[j] {
						isPlace = true
					}
					if j >= 1 && b[j] == b[j-1] && h[i] != b[j] {
						isPlace = true
					}

					if isPlace {
						newB := getNewBoard(b[:j] + string(h[i]) + b[j:])

						// 如果已经是结果就返回
						if newB == "" {
							return steps
						}

						newH := &strings.Builder{}
						newH.WriteString(h[:i])
						newH.WriteString(h[i+1:])

						situation := fmt.Sprintf("%v_%v", newB, newH.String())
						if set[situation] {
							continue
						}
						tmp = append(tmp, []string{newB, newH.String()})
					}
				}
			}
		}

		queue = tmp
	}

	return -1
}

func getNewBoard(s string) string {
	stack := []byte{}
	l := []int{}
	n := len(s)

	for i, start := 0, 0; i < n; i++ {
		for i+1 < n && s[i+1] == s[start] {
			i++
		}
		count := i - start + 1
		if count < 3 {
			// 如果栈中最后的颜色和当前颜色，相同，且相加数量 > 2
			if len(stack) > 0 && stack[len(stack)-1] == s[start] {
				if l[len(l)-1]+count >= 3 {
					// 值够，出栈
					stack = stack[:len(stack)-1]
					l = l[:len(l)-1]
				} else {
					// 值不够 则更新
					l[len(l)-1] += count
				}
			} else {
				l = append(l, count)
				stack = append(stack, s[start])
			}
		}
		start = i + 1
	}

	res := &strings.Builder{}
	for i, c := range stack {
		for j := 0; j < l[i]; j++ {
			res.WriteByte(c)
		}
	}

	return res.String()
}

func main() {
	board := "RRWWRRBBRR"
	hand := "WB"
	fmt.Println(findMinStep(board, hand))
}
