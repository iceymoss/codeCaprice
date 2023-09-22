package binaryTree

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var dfs func(*TreeNode, int) bool
	dfs = func(node *TreeNode, pathSum int) bool {
		//递归到底的条件: 找到叶子节点
		if node.Right == nil && node.Left == nil {
			fmt.Println(pathSum)
			return targetSum == pathSum+node.Val
		}
		//计算路径上的值
		pathSum += node.Val
		var lfind, rfind bool
		if node.Left != nil {
			lfind = dfs(node.Left, pathSum)
		}
		if node.Right != nil {
			rfind = dfs(node.Right, pathSum)
		}
		//是否找到
		if lfind || rfind {
			return true
		}
		return false
	}
	return dfs(root, 0)
}
