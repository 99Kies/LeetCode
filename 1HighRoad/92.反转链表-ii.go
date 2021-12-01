package main

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head

	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	rightNode := prev

	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	leftNode := prev.Next
	curr := rightNode.Next

	// 切断中间链表的连接
	prev.Next = nil
	rightNode.Next = nil

	reverseList(leftNode)
	prev.Next = rightNode
	leftNode.Next = curr
	return dummy.Next

}

// @lc code=end
