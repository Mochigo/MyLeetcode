package main

func deleteNode(node *ListNode) {
	// 从被删除的节点开始。。。让自己变成下一个节点
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
