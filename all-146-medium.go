/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
type LNode struct {
	Key  int
	Val  int
	Prev *LNode
	Next *LNode
}

type LRUCache struct {
	Cap        int
	Head, Tail *LNode
	Size       int
	Dict       map[int]*LNode
}

func initLNode(key, value int) *LNode {
	return &LNode{
		Key: key,
		Val: value,
	}
}

func Constructor(capacity int) LRUCache {
	ret := LRUCache{
		Cap:  capacity,
		Head: initLNode(0, 0),
		Tail: initLNode(0, 0),
		Size: 0,
		Dict: make(map[int]*LNode),
	}

	ret.Head.Next = ret.Tail
	ret.Tail.Prev = ret.Head

	return ret
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.Dict[key]
	if !ok {
		return -1
	}
	value := node.Val
	l.removeNode(node)
	l.addTailNode(key, value)
	return value
}

func (l *LRUCache) Put(key int, value int) {
	// 如果值且仍在链表中，我们就更新这个值，并且，调整节点到结尾。
	if node, ok := l.Dict[key]; ok {
		l.removeNode(node)
		l.addTailNode(key, value)
		return
	}

	// 如果满了，就把队头弹出
	if l.Size == l.Cap {
		l.removeHeadNode()
	}

	l.addTailNode(key, value)
}

func (l *LRUCache) removeHeadNode() {
	delete(l.Dict, l.Head.Next.Key)
	l.removeNode(l.Head.Next)
}

func (l *LRUCache) removeNode(node *LNode) {
	l.Size--
	prev, next := node.Prev, node.Next
	prev.Next, next.Prev = next, prev
}

func (l *LRUCache) addTailNode(key, value int) {
	l.Size++
	newNode := initLNode(key, value)
	prev := l.Tail.Prev
	prev.Next, newNode.Prev = newNode, prev
	newNode.Next, l.Tail.Prev = l.Tail, newNode
	l.Dict[key] = newNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

