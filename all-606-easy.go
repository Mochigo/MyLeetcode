package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
	if root == nil {
		return "()"
	}

	ret := &strings.Builder{}
	ret.WriteString(strconv.Itoa(root.Val))
	if root.Left != nil {
		ret.WriteRune('(')
		ret.WriteString(tree2str(root.Left))
		ret.WriteRune(')')
	} else if root.Right != nil {
		ret.WriteString("()")
	}

	if root.Right != nil {
		ret.WriteRune('(')
		ret.WriteString(tree2str(root.Right))
		ret.WriteRune(')')
	}

	return ret.String()
}
