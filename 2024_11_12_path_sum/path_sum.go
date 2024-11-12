package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	var cond bool
	if root.Left != nil {
		cond = HasPathSum(root.Left, targetSum-root.Val)
	}
	if !cond && root.Right != nil {
		cond = HasPathSum(root.Right, targetSum-root.Val)
	}
	return cond
}
