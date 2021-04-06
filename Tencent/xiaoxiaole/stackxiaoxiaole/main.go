// 栈
package main

import (
	"fmt"
)

type Node struct {
	Next *Node
	Val  int
}

type Stack struct {
	top    *Node
	length int
}

func Length(stack *Stack) int {
	return stack.length
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

func InsideUp(stack *Node) {
	if stack != nil {
		if stack.Next != nil {
			InsideUp(stack.Next)
		}
		fmt.Printf("%d ", stack.Val)
	}
}

//func main() {
//	aStack := CreateStack()
//	//fmt.Println(aStack.Top())
//	aStack.Push(1)
//	aStack.Push(2)
//	aStack.Push(3)
//	aStack.Push(4)
//	aStack.Push(5)
//	//aStack.UpsideDown()
//	fmt.Println()
//	//InsideUp(aStack.top)
//	fmt.Println(aStack.Top())
//	//aStack.Pop()
//	//fmt.Println(aStack.Top())
//
//}

func main() {
	input := ""
	fmt.Scanln(&input)
	aStack := CreateStack()
	for _, i := range []byte(input) {
		num := i - '0'
		if aStack.IsStackEmpty() {
			aStack.Push(int(num))
		} else {
			if aStack.Top()+int(num) == 10 {
				aStack.Pop()
			} else {
				aStack.Push(int(num))
			}
		}
	}

	fmt.Println()
	cnt := aStack.length
	fmt.Println(cnt)

}
