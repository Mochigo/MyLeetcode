package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InitBST(nums []int) *TreeNode {
	root := &TreeNode{}
	var createBST func(int, int, *TreeNode)
	createBST = func(left int, right int, root *TreeNode) {
		mid := (right-left)/2 + left
		root.Val = nums[mid]
		if left <= mid-1 {
			root.Left = &TreeNode{}
			createBST(left, mid-1, root.Left)
		}
		if right >= mid+1 {
			root.Right = &TreeNode{}
			createBST(mid+1, right, root.Right)
		}
	}

	createBST(0, len(nums)-1, root)
	return root
}

func printBST(root *TreeNode) []int {
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

	return inOrder
}

// func main() {
// 	nums := []int{1, 2, 3, 4, 5, 6}
// 	root := InitBST(nums)
// 	fmt.Println(printBST(root))
// }
