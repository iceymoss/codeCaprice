package binaryTree

func averageOfLevels(root *TreeNode) []float64 {
	var ans []float64
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		num := size
		var sum float64
		for size > 0 {
			top := queue[0]
			queue = queue[1:]
			sum += float64(top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			size--
		}
		ans = append(ans, sum/float64(num))
	}
	return ans
}
