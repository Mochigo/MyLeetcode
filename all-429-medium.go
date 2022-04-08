package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	trees := []*Node{root}
	var ans [][]int
	for len(trees) != 0 {
		tmp := []*Node{}
		vals := make([]int, 0, len(trees))
		for _, n := range trees {
			vals = append(vals, n.Val)
			for _, child := range n.Children {
				tmp = append(tmp, child)
			}
		}

		ans = append(ans, vals)
		trees = tmp
	}

	return ans
}
