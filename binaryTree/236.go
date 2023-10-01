package binaryTree

// 使用回溯法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(node, p, q *TreeNode) *TreeNode
	dfs = func(node, p, q *TreeNode) *TreeNode {
		//递归终止条件
		if node == nil || node == q || node == p {
			return node
		}

		//向层找，找到回溯
		l, r := dfs(node.Left, p, q), dfs(node.Right, p, q)

		if l == nil {
			return r
		}
		if r == nil {
			return l
		}
		return node
	}
	return dfs(root, p, q)
}
