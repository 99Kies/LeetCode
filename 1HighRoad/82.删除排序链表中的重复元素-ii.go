package main

import "fmt"

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	temp := &ListNode{0, head}
	pre1 := temp
	pre2 := temp.Next

	for pre1.Next != nil && pre2.Next != nil {
		if pre1.Next.Val == pre2.Next.Val {
			val := pre1.Next.Val
			for pre1.Next.Val == val && pre1.Next != nil && pre2.Next != nil {
				if pre2 == nil {
					break
				}
				pre1.Next = pre2.Next
				pre2 = pre2.Next
			}
		} else {
			pre1 = pre1.Next
			pre2 = pre2.Next
		}
	}
	return temp.Next
}

// @lc code=end

func showListNode(node *ListNode) {
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func main() {
	// head7 := ListNode{
	// 	Val:  5,
	// 	Next: nil,
	// }

	// head6 := ListNode{
	// 	Val:  4,
	// 	Next: &head7,
	// }

	// head5 := ListNode{
	// 	Val:  3,
	// 	Next: nil,
	// }

	// head4 := ListNode{
	// 	Val:  2,
	// 	Next: nil,
	// }

	// head3 := ListNode{
	// 	Val:  1,
	// 	Next: &head4,
	// }

	head2 := ListNode{
		Val:  1,
		Next: nil,
	}
	head1 := &ListNode{
		Val:  1,
		Next: &head2,
	}
	// showListNode(head1)
	res := deleteDuplicates(head1)
	fmt.Println("============================")
	showListNode(res)
}
