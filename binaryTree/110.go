package binaryTree

// 使用递归+回溯：将每一个节点都看作一个root, 然后判断其左右子树高度差值是否大于1
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node.Left == nil && node.Right == nil {
			return 1
		}
		l, r := 0, 0
		if node.Left != nil {
			l = dfs(node.Left)
		}
		if node.Right != nil {
			r = dfs(node.Right)
		}
		if l == -1 || r == -1 {
			return -1
		}
		if abs(l, r) > 1 {
			return -1
		}
		return max(l, r) + 1
	}
	if dfs(root) == -1 {
		return false
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
