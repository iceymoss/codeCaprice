package binaryTree

func rightSideView(root *TreeNode) []int {
	//二叉树bfs从左到右的最后一个元素即可
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			top := queue[0]
			queue = queue[1:]
			if size == 1 {
				ans = append(ans, top.Val)
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			size--
		}
	}
	return ans
}
