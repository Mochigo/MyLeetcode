package main

import "fmt"

type Trie struct {
	ch    [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, c := range []byte(word) {
		if node.ch[c-'a'] == nil {
			node.ch[c-'a'] = &Trie{}
		}
		node = node.ch[c-'a']
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range []byte(word) {
		if node.ch[c-'a'] == nil {
			return false
		}
		node = node.ch[c-'a']
	}

	return node.isEnd == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range []byte(prefix) {
		if node.ch[c-'a'] == nil {
			return false
		}
		node = node.ch[c-'a']
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	obj := Constructor()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.StartsWith("app"))

	fmt.Println(obj.Search("apply"))
	fmt.Println(obj.StartsWith("ab"))
}
