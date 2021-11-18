package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l, r := dfs(root.Left), dfs(root.Right)
		ans += abs(l - r)
		return l + r + root.Val
	}
	dfs(root)

	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
