package main

import "fmt"

type Node struct {
	Next *Node
	Val  int
	min  int
}

type MinStack struct {
	top    *Node
	length int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		&Node{
			nil,
			0,
			0,
		},
		0,
	}
}

func (this *MinStack) Push(x int) {
	if this.top.Next == nil {
		node := &Node{
			Val:  x,
			min:  x,
			Next: nil,
		}
		this.top.Next = node
	} else {
		node := &Node{
			Val:  x,
			min:  this.top.Next.min,
			Next: this.top.Next,
		}
		if node.min > x {
			node.min = x
		}
		this.top.Next = node
	}
}

func (this *MinStack) Pop() {
	if this.top.Next == nil {
		return
	}
	p := this.top.Next
	this.top.Next = p.Next
	this.length--
}

func (this *MinStack) Top() int {
	return this.top.Next.Val

}

func (this *MinStack) GetMin() int {
	return this.top.Next.min
}

func (this *MinStack) ShowStack() {
	p := this.top.Next
	for p != nil {
		fmt.Println(p.Val, p.min)
		p = p.Next
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	a := Constructor()
	a.Push(1)
	a.Push(2)
	a.Push(3)
	a.Push(4)
	a.Push(5)
	a.ShowStack()
	fmt.Println()
	a.Pop()
	a.Pop()
	a.Pop()
	a.ShowStack()
	//fmt.Println(a.Top())
	//a.Pop()
	//fmt.Println(a.Top())
}
