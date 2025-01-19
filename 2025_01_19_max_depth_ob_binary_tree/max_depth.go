package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := 1+maxDepth(root.Left), 1+maxDepth(root.Right)
	if left > right {
		return left
	}
	return right
}
