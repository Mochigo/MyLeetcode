package main

func isEvenOddTree(root *TreeNode) bool {
	trees := []*TreeNode{root}
	isEven := true
	for len(trees) > 0 {
		tmp := []*TreeNode{}
		pre := trees[len(trees)-1].Val + 1
		if isEven {
			for i := len(trees) - 1; i >= 0; i-- {
				t := trees[i]
				if t.Val%2 == 1 && pre > t.Val {
					if t.Right != nil {
						tmp = append(tmp, t.Right)
					}
					if t.Left != nil {
						tmp = append(tmp, t.Left)
					}
					pre = t.Val
				} else {
					return false
				}
			}
		} else {
			for i := len(trees) - 1; i >= 0; i-- {
				t := trees[i]
				if t.Val%2 == 0 && pre > t.Val {
					if t.Left != nil {
						tmp = append(tmp, t.Left)
					}
					if t.Right != nil {
						tmp = append(tmp, t.Right)
					}
					pre = t.Val
				} else {
					return false
				}
			}
		}
		isEven = !isEven
		trees = tmp
	}

	return true
}
