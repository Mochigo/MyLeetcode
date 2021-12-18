package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}

	lists := []*ListNode{}
	for node := dummyHead; node != nil; node = node.Next {
		lists = append(lists, node)
	}

	sz := len(lists)
	if n != 1 {
		lists[sz-n-1].Next = lists[sz-n+1]
	} else {
		lists[sz-n-1].Next = nil
	}
	return dummyHead.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{0, head}
	first, second := head, dummyHead
	for i := 0; i < n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummyHead.Next
}
