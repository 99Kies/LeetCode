package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func main() {
	//nums := []int{1,2,3,4,5}
	//fmt.Println()
	n := 0
	fmt.Scan(&n)
	head := &Node{
		Val:  1,
		Next: nil,
	}
	//phead := head
	for i := 2; i <= n; i++ {
		node := &Node{
			Val:  i,
			Next: nil,
		}
		head.Next = node
		head = head.Next
	}
	k := 0
	fmt.Scan(&k)
	res := make([][]int, k)
	L := 0
	R := 0
	for i := 0; i < k; i++ {
		fmt.Scan(&L)
		fmt.Scan(&R)
		res[i] = append(res[i])
	}

}

// 时间不足。。  使用链表进行操作。
