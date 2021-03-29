package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			dfs(root.Right)
			if root.Val >= low && root.Val <= high {
				sum += root.Val
			}
		}
	}
	dfs(root)
	return sum
}

//func rangeSumBST(root *TreeNode, low int, high int) int {
//	// 给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。
//	if root == nil{
//		return 0
//	}
//	if root.Val < low {
//		return rangeSumBST(root.Right, low, high)
//	}
//
//	if root.Val > high {
//		return rangeSumBST(root.Left, low, high)
//	}
//
//	return root.Val + rangeSumBST(root.Right, low, high) + rangeSumBST(root.Left, low, high)
//}

func main() {

}
