package binaryTree

import "math"

// bfs
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			top := queue[0]
			queue = queue[1:]
			if top.Left == nil && top.Right == nil {
				return res
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			size--
		}
		res++
	}
	return res
}

func minDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		//终止条件
		if node.Left == nil && node.Right == nil {
			return 1
		}

		//当前层处理的逻辑
		minD := math.MaxInt
		if node.Left != nil {
			minD = min(minD, dfs(node.Left))
		}
		if node.Right != nil {
			minD = min(minD, dfs(node.Right))
		}
		return minD + 1
	}
	return dfs(root)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
