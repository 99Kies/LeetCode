package main

type Node struct {
	Next *Node
	Val  int
}

type Stack struct {
	top    *Node
	length int
}

func CreateStackNode(value int) *Node {
	return &Node{
		nil,
		value,
	}
}

func CreateStack() *Stack {
	return &Stack{
		&Node{
			nil,
			0,
		},
		0,
	}
}

func (stack *Stack) IsStackEmpty() bool {
	return stack.top.Next == nil
}

func (stack *Stack) Push(value int) {
	p := CreateStackNode(value)
	p.Next = stack.top.Next
	stack.top.Next = p
	stack.length++
}

func (stack *Stack) Pop() int {
	if stack.top.Next == nil {
		return -999
	}
	p := stack.top.Next
	value := p.Val
	stack.top.Next = p.Next
	p = nil
	stack.length--
	return value
}

func (stack *Stack) Top() int {
	if stack.top.Next == nil {
		return -999
	}
	return stack.top.Next.Val
}
