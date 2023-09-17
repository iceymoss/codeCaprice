package binaryTree

func maxDepth(root *TreeNode) int {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		//终止条件
		if node == nil {
			return 0
		}
		//当前层需要处理的逻辑
		l := dfs(node.Left)
		r := dfs(node.Right)
		if l > r {
			return l + 1
		}
		return r + 1
	}
	return dfs(root)
}

func maxDepthV2(root *TreeNode) int {
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
