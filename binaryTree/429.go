package binaryTree

// Node 定义一个N叉数
type Node struct {
	Val      int
	Children []*Node
}

// N叉树的层序遍历
func levelOrders(root *Node) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		tmp := make([]int, 0)
		for size > 0 {
			top := queue[0]
			queue = queue[1:]
			tmp = append(tmp, top.Val)
			for _, v := range top.Children {
				queue = append(queue, v)
			}
			size--
		}
		ans = append(ans, tmp)
	}
	return ans
}
