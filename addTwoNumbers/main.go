package main

import "fmt"

func addTwoNumbershuifeng(l1 *List, l2 *List) *List {
	result := &List{}
	headList := new(ListNode)
	head := headList
	result.headNode = head
	prev1 := l1.headNode
	prev2 := l2.headNode
	res := 0
	for prev1 != nil && prev2 != nil {
		headList.Next = new(ListNode)
		headList = headList.Next
		res = prev1.Val.(int) + prev2.Val.(int) + res
		if res >= 10 {
			headList.Val = res % 10
			res = res / 10
		} else {
			headList.Val = res % 10
			res = 0
		}
		//res = prev1.Val.(int) + prev2.Val.(int) + res
		prev2 = prev2.Next
		prev1 = prev1.Next
	}
	return result

}

func addTwoNumbers(l1 *List, l2 *List) *List {
	result := &List{}
	prev1 := l1.headNode
	prev2 := l2.headNode
	headList := new(ListNode)
	head := headList
	result.headNode = head
	res := 0
	for prev1 != nil || prev2 != nil || res > 0 {
		headList.Next = new(ListNode)
		headList = headList.Next
		if prev1 != nil {
			res = prev1.Val.(int) + res
			fmt.Print("prev1: ")
			fmt.Println(res % 10)
			prev1 = prev1.Next
		}

		if prev2 != nil {
			res = prev2.Val.(int) + res
			fmt.Print("prev2: ")
			fmt.Println(res % 10)
			prev2 = prev2.Next
		}

		//res = prev1.Val.(int) + prev2.Val.(int) + res
		headList.Val = res % 10
		res = res / 10
	}
	return result

}

func main() {
	list := List{}
	list.Append(2)
	list.Append(4)
	list.Append(3)
	list.ShowList()
	fmt.Println()
	list2 := List{}
	list2.Append(5)
	list2.Append(6)
	list2.Append(4)
	list2.Append(3)
	list2.ShowList()
	fmt.Println()
	ans := addTwoNumbers(&list, &list2)
	ans.ShowList()
}
