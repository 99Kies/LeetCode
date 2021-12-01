/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 迭代
	// dummy := &ListNode{}
	// prev := dummy

	// for l1 != nil && l2 != nil {
	// 	if l1.Val <= l2.Val {
	// 		prev.Next = l1
	// 		l1 = l1.Next
	// 	} else {
	// 		prev.Next = l2
	// 		l2 = l2.Next
	// 	}
	// 	prev = prev.Next
	// }
	// if l1 == nil {
	// 	prev.Next = l2
	// } else {
	// 	prev.Next = l1
	// }
	// return dummy.Next

	// 递归
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

// @lc code=end

