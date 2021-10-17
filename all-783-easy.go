package main

import (
	"fmt"
	"math"
)

func minDiffInBST(root *TreeNode) int {
	ans, pre := math.MaxInt32, -1
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		if pre != -1 && root.Val-pre < ans {
			ans = root.Val - pre
		}
		pre = root.Val
		dfs(root.Right)
	}

	dfs(root)

	return ans
}

func main() {
	nums := []int{1, 4, 7, 10, 20, 30}
	root := InitBST(nums)
	fmt.Println(minDiffInBST(root))
}
