package main

import (
	"fmt"
)

type ObjectLeaf interface{}

type LeafNode struct {
	Val   ObjectLeaf
	left  *LeafNode
	right *LeafNode
}

func (this *LeafNode) Print() {
	fmt.Printf("%d ", this.Val)
}

func (this *LeafNode) SetValue(value int) {
	this.Val = value
}

func CreateNode(value int) *LeafNode {
	return &LeafNode{value, nil, nil}
}

func (this *LeafNode) FindNode(node *LeafNode, key int) *LeafNode {
	if node == nil {
		return nil
	} else if node.Val == key {
		return node
	} else {
		prev := this.FindNode(node.left, key)
		if prev != nil {
			return prev
		}
		return this.FindNode(node.right, key)
	}
}

func (this *LeafNode) GetTreeHigh(node *LeafNode) int {
	if node == nil {
		return 0
	} else {
		lhigh := this.GetTreeHigh(node.left)
		rhigh := this.GetTreeHigh(node.right)
		if lhigh > rhigh {
			return lhigh + 1
		} else {
			return rhigh + 1
		}
	}
}

func (this *LeafNode) PreOrder(node *LeafNode) {
	if node != nil {
		fmt.Printf("%d ", node.Val)
		this.PreOrder(node.left)
		this.PreOrder(node.right)
	}
}

func (this *LeafNode) InOrder(node *LeafNode) {
	if node != nil {
		this.InOrder(node.left)
		fmt.Printf("%d ", node.Val)
		this.InOrder(node.right)
	}
}

func (this *LeafNode) PostOrder(node *LeafNode) {
	if node != nil {
		this.PostOrder(node.left)
		this.PostOrder(node.right)
		fmt.Printf("%d ", node.Val)
	}
}

func postorderTraversal(root *LeafNode) []ObjectLeaf {
	result := make([]ObjectLeaf, 0)
	var postorder func(node *LeafNode)
	postorder = func(node *LeafNode) {
		if node != nil {
			postorder(node.left)
			postorder(node.right)
			result = append(result, node.Val)
		}
	}
	postorder(root)
	return result
}

func getrightNode(node *LeafNode) []ObjectLeaf {
	result := make([]ObjectLeaf, 0)
	result = append(result, node.Val)
	var noderight func(leafNode *LeafNode)

	noderight = func(leafNode *LeafNode) {
		if node.right != nil {
			noderight(node.left)
			noderight(node.right)
			result = append(result, node.right.Val)
		}
	}

	noderight(node)
	return result
}

func (this *LeafNode) GetLeafNode(node *LeafNode) {
	// 获取叶子节点
	if node != nil {
		if node.left == nil && node.right == nil {
			fmt.Printf("%d ", node.Val)
		}
		this.GetLeafNode(node.left)
		this.GetLeafNode(node.right)
	}
}

func main() {
	root := CreateNode(5)
	root.left = CreateNode(2)
	root.right = CreateNode(4)
	root.left.right = CreateNode(7)
	root.left.right.left = CreateNode(6)
	root.right.left = CreateNode(8)
	root.right.right = CreateNode(9)

	fmt.Printf("%d\n", root.FindNode(root, 4).Val)
	fmt.Printf("%d\n", root.GetTreeHigh(root))
	root.PreOrder(root)
	fmt.Println()
	root.InOrder(root)
	fmt.Println()
	root.PostOrder(root)
	fmt.Println()
	root.GetLeafNode(root)
	fmt.Println()

	fmt.Println(postorderTraversal(root))
	fmt.Println()
	root.PostOrder(root)
}
