package main

import "fmt"

type Node struct {
	Next *Node
	Val  string
}

type Stack struct {
	top    *Node
	length int
}

func Length(stack *Stack) int {
	return stack.length
}

func CreateStackNode(value string) *Node {
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
			"0",
		},
		0,
	}
}

func (stack *Stack) IsStackEmpty() bool {
	return stack.top.Next == nil
}

func (stack *Stack) Push(value string) {
	p := CreateStackNode(value)
	p.Next = stack.top.Next
	stack.top.Next = p
	stack.length++
}

func (stack *Stack) Pop() string {
	if stack.top.Next == nil {
		return "-999"
	}
	p := stack.top.Next
	value := p.Val
	stack.top.Next = p.Next
	p = nil
	stack.length--
	return value
}

func (stack *Stack) Top() string {
	if stack.top.Next == nil {
		return "-999"
	}
	return stack.top.Next.Val
}

func (stack *Stack) UpsideDown() {
	// 倒着输出
	if stack.top.Next == nil {
		fmt.Println("This is Empty Stack")
		return
	}
	p := stack.top.Next
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
}

func isValid(s string) bool {
	aStack := CreateStack()
	for _, ch := range s {
		if aStack.IsStackEmpty() {
			aStack.Push(string(ch))
		} else {
			if aStack.Top() == "(" && string(ch) == ")" {
				aStack.Pop()
			} else if aStack.Top() == "{" && string(ch) == "}" {
				aStack.Pop()
			} else if aStack.Top() == "[" && string(ch) == "]" {
				aStack.Pop()
			} else {
				aStack.Push(string(ch))
			}
		}
	}
	return aStack.length == 0

}

func main() {
	a := "[]{()"
	res := isValid(a)
	fmt.Println(res)
}
