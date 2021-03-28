// 栈
package main

import "fmt"

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
	// 下标0 是空Node，所以Top的时候是 Stack.top.Next.Val
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

func main() {
	aStack := CreateStack()
	fmt.Println("top ", aStack.Top())
	fmt.Println(aStack.top.Val)

	aNode := CreateStackNode(1)
	aStack.top.Next = aNode

	fmt.Println(aStack.Top())
	fmt.Println(aStack.top.Val)
	//aStack.Push(1)
	//fmt.Println(aStack.Top())
	//aStack.Push(2)
	//aStack.Push(3)
	//aStack.Push(4)
	//aStack.Push(5)
	//fmt.Println(aStack.Top())
	//aStack.Pop()
	//fmt.Println(aStack.Top())

}
