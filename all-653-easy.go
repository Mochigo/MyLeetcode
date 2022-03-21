package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	set := map[int]struct{}{}
	var dfs func(*TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if _, ok := set[k-root.Val]; ok {
			return true
		}
		set[root.Val] = struct{}{}
		return dfs(root.Left) || dfs(root.Right)

	}

	return dfs(root)
}
