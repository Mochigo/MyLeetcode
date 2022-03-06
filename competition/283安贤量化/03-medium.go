package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	// 在父节点出现，没有在子节点出现就是 根节点
	// 希望能在一次遍历中找到，并完成初始化

	childs := make(map[int]bool)
	nodes := make(map[int]*TreeNode)
	for _, des := range descriptions {
		childs[des[1]] = true
		nodes[des[0]] = &TreeNode{Val: des[0]}
		if _, ok := nodes[des[1]]; !ok {
			nodes[des[1]] = &TreeNode{Val: des[1]}
		}
	}

	rootVal := 0
	for _, des := range descriptions {
		if !childs[des[0]] {
			rootVal = des[0]
		}
		if des[2] == 1 {
			nodes[des[0]].Left = nodes[des[1]]
		} else {
			nodes[des[0]].Right = nodes[des[1]]
		}
	}

	return nodes[rootVal]
}
