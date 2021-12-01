/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// if head == nil || head.Next == nil {
	// 	return head
	// }
	// nextHead := reverseList(head.Next)
	// head.Next.Next = head
	// head.Next = nil
	// return nextHead
	var prev *ListNode = nil
	var curr = head
	for curr != nil {
		var next = curr.Next
		curr.Next = prev // 倒转
		prev, curr = curr, next
	}
	return prev
}

// @lc code=end

