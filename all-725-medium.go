package main

import (
	"fmt"
	"strconv"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	ans := make([]*ListNode, k) // k 用来表示分几段， 所以这里可以直接确定ans的容量或者说长度

	// 因为需要根据长度来划分长度相对均等的区间，所以这里通过依次遍历将 node 的信息存放在一个切片中
	list := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		list = append(list, node)
	}

	n, i := len(list), 0
	left, leftK := n, k //left 表示剩下的多少， leftK表示还剩下几个区间可以放置
	index := 0
	avg := n / k // 表示每一段应该有的长度，如果 left % leftK != 0，就比其他段多给一个

	for i < n && index < k {
		if left%leftK != 0 {
			ans[index] = list[i]
			i += avg + 1
			left -= avg + 1
		} else {
			ans[index] = list[i]
			i += avg
			left -= avg
		}
		index++
		leftK--
		if i < n {
			list[i-1].Next = nil
		} else {
			list[n-1].Next = nil
		}
	}

	return ans
}

func getList(head *ListNode) string {
	node := head.Next
	res := strconv.Itoa(head.Val)
	for node != nil {
		res += "->" + strconv.Itoa(node.Val)
		node = node.Next
	}

	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 3
	index := 0
	head := &ListNode{}
	for node := head; index < len(nums); node = node.Next {
		node.Val = nums[index]
		node.Next = &ListNode{}
		index++
		if index == len(nums) {
			node.Next = nil
			break
		}
	}

	for _, list := range splitListToParts(head, k) {
		fmt.Println(getList(list))
	}
}
