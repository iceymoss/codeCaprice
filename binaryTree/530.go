package binaryTree

import (
	"math"
)

// 二分搜索树中序遍历-有序的
func getMinimumDifference(root *TreeNode) int {
	minInt := math.MaxInt
	var per *TreeNode
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if per != nil && node.Val-per.Val < minInt {
			minInt = node.Val - per.Val
		}
		per = node
		dfs(node.Right)
	}
	dfs(root)
	return minInt
}
