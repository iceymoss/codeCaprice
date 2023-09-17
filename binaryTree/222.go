package binaryTree

func countNodes(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans++
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
