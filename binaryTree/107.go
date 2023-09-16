package binaryTree

type stack struct {
	arr []int
}

func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	arr := make([]stack, 0)
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
		s := stack{arr: Items}
		arr = append(arr, s)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		ans = append(ans, arr[i].arr)
	}
	return ans
}
