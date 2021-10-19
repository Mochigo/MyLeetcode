package main

import "fmt"

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) Insert(word string) {
	for _, c := range word {
		if t.children[c-'a'] == nil {
			t.children[c-'a'] = &Trie{}
		}
		t = t.children[c-'a']
	}
	t.isEnd = true
}

func (t *Trie) Search(word string) bool {
	for _, c := range word {
		if t.children[c-'a'] == nil {
			return false
		}
		t = t.children[c-'a']
	}
	return t.isEnd == true
}

type WordDictionary struct {
	trie *Trie
}

func Constructor() WordDictionary {
	return WordDictionary{trie: &Trie{}}
}

func (this *WordDictionary) AddWord(word string) {
	this.trie.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	found := false
	n := len(word)
	var dfs func(*Trie, int)
	dfs = func(node *Trie, index int) {
		if found {
			return
		}
		if index == n {
			if node.isEnd {
				found = true
			}
			return
		}
		ch := byte(word[index])
		if ch != '.' && node.children[ch-'a'] != nil {
			dfs(node.children[ch-'a'], index+1)
		}
		if ch == '.' {
			for i := 0; i <= 25; i++ {
				if node.children[i] != nil {
					dfs(node.children[i], index+1)
				}
			}
		}
	}

	dfs(this.trie, 0)

	return found
}

func (this *WordDictionary) Search2(word string) bool {
	n := len(word)
	var dfs func(*Trie, int) bool
	dfs = func(node *Trie, index int) bool {
		if index == n {
			return node.isEnd == true
		}
		ch := word[index]
		if ch != '.' {
			child := node.children[ch-'a']
			if child != nil && dfs(node.children[ch-'a'], index+1) {
				return true
			}
		} else {
			for i := range node.children {
				child := node.children[i]
				if child != nil && dfs(child, index+1) {
					return true
				}
			}
		}
		return false
	}

	return dfs(this.trie, 0)
}

func main() {
	obj := Constructor()
	obj.AddWord("bad")
	fmt.Println(obj.Search(".ad"))
	fmt.Println(obj.Search("..d"))
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
