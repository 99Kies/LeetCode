package Node

import "fmt"

type Object interface{}

type ListNode struct {
	Val  Object
	Next *ListNode
}

type List struct {
	headNode *ListNode
}

func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

func (this *List) Length() int {
	cur := this.headNode
	count := 0

	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

func (this *List) Add(val Object) *ListNode {
	node := &ListNode{Val: val}
	node.Next = this.headNode
	this.headNode = node
	return node
}

func (this *List) Append(val Object) {
	node := &ListNode{Val: val}
	if this.IsEmpty() {
		this.headNode = node
	} else {
		cur := this.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

func (this *List) Insert(index int, val Object) {
	// index为下标 下标从零开始
	if index < 0 {
		this.Add(val)
	} else if index > this.Length() {
		this.Append(val)
	} else {
		count := 0
		prev := this.headNode
		for count < (index - 1) {
			count++
			prev = prev.Next
		}
		node := &ListNode{Val: val}
		node.Next = prev.Next
		prev.Next = node
	}
}

func (this *List) Remove(val Object) {
	prev := this.headNode
	if prev.Val == val {
		this.headNode = prev.Next
	} else {
		for prev.Next != nil {
			if prev.Next.Val == val {
				prev.Next = prev.Next.Next
			} else {
				prev = prev.Next
			}
		}
	}
}

func (this *List) RemoveAtIndex(index int) {
	prev := this.headNode
	if index <= 0 {
		this.headNode = prev.Next
	} else if index > this.Length() {
		fmt.Println("超出链表长度")
		return
	} else {
		count := 0
		for count < (index-1) && prev.Next != nil {
			count++
			prev = prev.Next
		}
		prev.Next = prev.Next.Next
	}
}

func (this *List) Contain(val Object) bool {
	cur := this.headNode
	for cur != nil {
		if cur.Val == val {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (this *List) ShowList() {
	if !this.IsEmpty() {
		cur := this.headNode
		for cur != nil {
			fmt.Println(cur.Val)
			cur = cur.Next
		}
	}
}

func main() {
	list := List{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Add("a")
	list.ShowList()
}
