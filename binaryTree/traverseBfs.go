package binaryTree

func leveTraverse(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	//队列
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		Items := make([]int, 0)
		for size > 0 {
			top := queue[0]
			queue = queue[1:]
			Items = append(Items, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			size--
		}
		ans = append(ans, Items)
	}
	return ans
}

func levelOrder(root *TreeNode) [][]int {
	//将队列的元素，始终维护当前一层的数据
	var res [][]int
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		curLayerSize := len(queue)
		curLayerItems := make([]int, 0)
		for i := 0; i < curLayerSize; i++ {
			top := queue[0]
			queue = queue[1:]
			curLayerItems = append(curLayerItems, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		res = append(res, curLayerItems)
	}
	return res
}
