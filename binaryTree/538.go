package binaryTree

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		//右中序遍历
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}
