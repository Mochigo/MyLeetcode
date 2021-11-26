/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST1(root *TreeNode, val int) *TreeNode {
	node := root
	for node != nil {
		if node.Val == val {
			return node
		} else if node.Val > val {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return nil
}

func searchBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	}

	return searchBST(root.Left, val)
}

func searchBST3(root *TreeNode, val int) *TreeNode {
	var dfs func(*TreeNode, int) bool
	var ans *TreeNode
	dfs = func(root *TreeNode, val int) bool {
		if root == nil {
			return false
		}
		if root.Val == val {
			ans = root
			return true
		}
		if dfs(root.Left, val) || dfs(root.Right, val) {
			return true
		}
		return false
	}

	dfs(root, val)
	return ans
}