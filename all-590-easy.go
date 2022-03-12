package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	ans := []int{}
	var dfs func(*Node)
	dfs = func(root *Node) {
		if root == nil {
			return
		}
		for _, ch := range root.Children {
			dfs(ch)
		}
		ans = append(ans, root.Val)
	}

	dfs(root)
	return ans
}
