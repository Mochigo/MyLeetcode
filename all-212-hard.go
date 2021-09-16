package main

import "fmt"

// 补充一个新得到的感受：这里的word也好，isEnd也好，实际上都是为了确定是否是一个单词的末尾
// 当存在具有相同前缀，或者说其本身就是另一个单词的组成部分时，两者会在同一条路径上，所以使用word(isEnd)用来表示一个单词
// 所以word和isEnd不能作为实际上的结尾，只有node == nil时，才能判断退出
type Trie struct {
	children [26]*Trie
	word     string
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.word = word
}

func findWords(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// 这里因为题目只要求出现过就行，实际上用一个不重复的集合就可以了，这里使用了map来做 set
	appeared := map[string]bool{}

	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	ans := make([]string, 0, len(words))

	var dfs func(*Trie, int, int)
	dfs = func(t *Trie, x, y int) {
		ch := board[x][y]
		node := t.children[ch-'a']
		if node == nil {
			return
		}

		if node.word != "" {
			appeared[node.word] = true
		}

		vis[x][y] = true
		defer func() { vis[x][y] = false }()

		for _, dir := range dirs {
			if newX, newY := x-dir[0], y-dir[1]; newX >= 0 && newY >= 0 && newX < m && newY < n && !vis[newX][newY] {
				dfs(node, newX, newY)
			}
		}
	}

	for i, row := range board {
		for j := range row {
			dfs(t, i, j)
		}
	}

	for word := range appeared {
		ans = append(ans, word)
	}

	return ans
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}

	words := []string{"oath", "pea", "eat", "rain"}

	fmt.Println(findWords(board, words))
}
