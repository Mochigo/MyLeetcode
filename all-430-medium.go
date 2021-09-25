package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	node := root
	for node != nil {
		// 这里是遍历到当前节点有子节点的情况。
		if node.Child != nil {
			node = dfs(node)
		} else {
			node = node.Next
		}
	}

	return root
}

// 除了有子节点的，别的自然迭代就可以了
func dfs(root *Node) *Node {
	// 记录下一个节点和子节点
	next := root.Next
	cur := root.Child

	// 把子节点和前一个节点相连接
	root.Next = cur
	root.Child = nil
	cur.Prev = root

	// 如果当前节点有下一个节点
	for cur.Next != nil {
		if cur.Child != nil {
			cur = dfs(cur)
		} else {
			cur = cur.Next
		}
	}

	if next != nil {
		cur.Next = next
		next.Prev = cur
	}

	return cur
}

func main() {

}
