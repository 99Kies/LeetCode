package main

import (
	"fmt"
	//"DataStruct/Node"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	sum := 0
	for head != nil {
		sum = sum*2 + head.Val.(int)
		head = head.Next
	}
	return sum
}

func main() {

	list := List{}
	list.Append(1)
	list.Append(0)
	list.Append(0)
	list.ShowList()
	fmt.Println()
	result := getDecimalValue(list.headNode)
	fmt.Println(result)
}
