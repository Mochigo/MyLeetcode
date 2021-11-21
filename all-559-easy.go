/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

// dfs
func maxDepth1(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	for _, child := range root.Children {
		v := maxDepth(child)
		if v > max {
			max = v
		}
	}
	return max + 1
}

// bfs
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	trees := []*Node{root}
	for len(trees) > 0 {
		depth++
		tmp := []*Node{}
		for _, tree := range trees {
			for _, child := range tree.Children {
				tmp = append(tmp, child)
			}
		}

		trees = tmp
	}

	return depth
}

