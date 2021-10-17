package main

import "fmt"

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)

	return res
}

func main() {
	nums := []int{1, 2, 4, 8, 9, 10}
	root := InitBST(nums)
	fmt.Println(inorderTraversal(root))
}
