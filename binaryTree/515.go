package binaryTree

import "math"

func largestValues(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		max := math.MinInt
		for size > 0 {
			top := queue[0]
			queue = queue[1:]
			if top.Val > max {
				max = top.Val
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			size--
		}
		ans = append(ans, max)
	}
	return ans
}
