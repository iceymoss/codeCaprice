package binaryTree

// 求二叉树的所有左节点值
func sumOfLeft(root *TreeNode) int {
	var ans int
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		sum := 0
		for size != 0 {
			top := queue[0]
			queue = queue[1:]
			if top.Left != nil {
				queue = append(queue, top.Left)
				sum += top.Left.Val
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			size--
		}
		ans += sum
	}
	return ans
}

// 求二叉树的所有左叶子节点值
func sumOfLeftLeaves(root *TreeNode) int {
	var ans int
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		sum := 0
		for size != 0 {
			top := queue[0]
			queue = queue[1:]
			if top.Left != nil {
				queue = append(queue, top.Left)
				if top.Left.Right == nil && top.Left.Left == nil {
					sum += top.Left.Val
				}
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			size--
		}
		ans += sum
	}
	return ans
}
