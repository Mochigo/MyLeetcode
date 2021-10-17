package main

import "fmt"

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func kthSmallest(root *TreeNode, k int) int {
	inOrder := []int{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		inOrder = append(inOrder, root.Val)
		dfs(root.Right)
	}

	dfs(root)

	return inOrder[k-1]
}

func main() {
	nums := []int{1, 2, 4, 8, 9, 10}
	root := InitBST(nums)
	fmt.Println(kthSmallest(root, 6))
}
