package main

type Trie struct {
	children [26]*Trie
	val      int
}

func (t *Trie) Insert(key string, v int) {
	node := t
	for _, c := range key {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &Trie{}
		}
		node = node.children[c-'a']
	}
	node.val = v
}

func (t *Trie) getSubTree(prefix string) *Trie {
	node := t
	for _, c := range prefix {
		if node.children[c-'a'] == nil {
			return nil
		}
		node = node.children[c-'a']
	}

	return node
}

type MapSum struct {
	root *Trie
}

func Constructor() MapSum {
	return MapSum{&Trie{}}
}

func (this *MapSum) Insert(key string, val int) {
	this.root.Insert(key, val)
}

func (this *MapSum) Sum(prefix string) int {
	root := this.root.getSubTree(prefix)
	sum := 0
	var dfs func(*Trie)
	dfs = func(root *Trie) {
		if root == nil {
			return
		}
		if root.val != 0 {
			sum += root.val
		}

		for _, child := range root.children {
			dfs(child)
		}
	}

	dfs(root)

	return sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
