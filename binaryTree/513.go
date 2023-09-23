package binaryTree

func findBottomLeftValue(root *TreeNode) int {
	//层序遍历
	if root == nil {
		return 0
	}
	ans := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		ans = queue[0].Val
		for size != 0 {
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
	}
	return ans
}
