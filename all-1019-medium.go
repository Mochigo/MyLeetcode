package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	stack := []int{}
	nums := []int{}
	for node := head; node != nil; node = node.Next {
		nums = append(nums, node.Val)
	}

	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			ans[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return ans
}
