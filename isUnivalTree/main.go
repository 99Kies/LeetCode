// 二叉树
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (this *TreeNode) Print() {
	fmt.Printf("%d ", this.Val)
}

func (this *TreeNode) SetValue(value int) {
	this.Val = value
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

func (this *TreeNode) FindNode(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return nil
	} else if node.Val == key {
		return node
	} else {
		prev := this.FindNode(node.Left, key)
		if prev != nil {
			return prev
		}
		return this.FindNode(node.Right, key)
	}
}

func (this *TreeNode) GetTreeHigh(node *TreeNode) int {
	if node == nil {
		return 0
	} else {
		lhigh := this.GetTreeHigh(node.Left)
		rhigh := this.GetTreeHigh(node.Right)
		if lhigh > rhigh {
			return lhigh + 1
		} else {
			return rhigh + 1
		}
	}
}

func (this *TreeNode) PreOrder(node *TreeNode) {
	if node != nil {
		fmt.Printf("%d ", node.Val)
		this.PreOrder(node.Left)
		this.PreOrder(node.Right)
	}
}

func (this *TreeNode) InOrder(node *TreeNode) {
	if node != nil {
		this.InOrder(node.Left)
		fmt.Printf("%d ", node.Val)
		this.InOrder(node.Right)
	}
}

func (this *TreeNode) PostOrder(node *TreeNode) {
	if node != nil {
		this.PostOrder(node.Left)
		this.PostOrder(node.Right)
		fmt.Printf("%d ", node.Val)
	}
}

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node != nil {
			postorder(node.Left)
			postorder(node.Right)
			result = append(result, node.Val)
		}
	}
	postorder(root)
	return result
}

func getRightNode(node *TreeNode) []int {
	result := make([]int, 0)
	result = append(result, node.Val)
	var nodeRight func(TreeNode *TreeNode)

	nodeRight = func(TreeNode *TreeNode) {
		if node.Right != nil {
			nodeRight(node.Left)
			nodeRight(node.Right)
			result = append(result, node.Right.Val)
		}
	}

	nodeRight(node)
	return result
}

func (this *TreeNode) GetTreeNode(node *TreeNode) {
	// 获取叶子节点
	if node != nil {
		if node.Left == nil && node.Right == nil {
			fmt.Printf("%d ", node.Val)
		}
		this.GetTreeNode(node.Left)
		this.GetTreeNode(node.Right)
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	cur := root.Val
	return checkNumsisSame(root, cur)
}

func checkNumsisSame(root *TreeNode, key int) bool {
	if root == nil {
		return true
	}
	if root.Val != key {
		return false
	}
	return checkNumsisSame(root.Left, key) && checkNumsisSame(root.Right, key)
}

func main() {
	//root1 := CreateNode(1)
	//root1.Left = CreateNode(2)
	//root1.Right = CreateNode(2)
	//root1.Left.Left = CreateNode(3)
	//root1.Left.Right = CreateNode(4)
	//root1.Right.Left = CreateNode(4)
	//root1.Right.Right = CreateNode(3)

	root1 := CreateNode(1)
	root1.Left = CreateNode(1)
	root1.Right = CreateNode(1)
	root1.Left.Left = CreateNode(1)
	root1.Left.Right = CreateNode(1)
	root1.Right.Left = CreateNode(1)
	root1.Right.Right = CreateNode(1)

	res := isUnivalTree(root1)
	fmt.Println(res)
}
