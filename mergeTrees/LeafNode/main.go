// 二叉树
package main

import (
	"fmt"
)

type ObjectLeaf interface{}

type TreeNode struct {
	Val   ObjectLeaf
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

func postorderTraversal(root *TreeNode) []ObjectLeaf {
	result := make([]ObjectLeaf, 0)
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

func getRightNode(node *TreeNode) []ObjectLeaf {
	result := make([]ObjectLeaf, 0)
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

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val = root1.Val.(int) + root2.Val.(int)
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func (this *TreeNode) rangeSumBST(root *TreeNode, low int, high int) int {
	// 给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。
	if root == nil {
		return 0
	}
	if root.Val.(int) < low {
		return this.rangeSumBST(root.Right, low, high)
	}

	if root.Val.(int) > high {
		return this.rangeSumBST(root.Left, low, high)
	}

	return root.Val.(int) + this.rangeSumBST(root.Right, low, high) + this.rangeSumBST(root.Left, low, high)
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			dfs(root.Right)
			if root.Val.(int) >= low && root.Val.(int) <= high {
				sum += root.Val.(int)
			}
		}
	}

	dfs(root)
	return sum

}

func main() {
	root1 := CreateNode(1)
	root1.Left = CreateNode(3)
	root1.Right = CreateNode(2)
	root1.Left.Left = CreateNode(5)

	root2 := CreateNode(10)
	root2.Left = CreateNode(5)
	root2.Right = CreateNode(15)
	root2.Left.Left = CreateNode(3)
	root2.Left.Right = CreateNode(7)
	root2.Right.Right = CreateNode(18)

	a := root2.rangeSumBST(root2, 7, 15)
	fmt.Println()
	fmt.Println(a)
	b := rangeSumBST(root2, 7, 15)
	fmt.Println(b)

}
