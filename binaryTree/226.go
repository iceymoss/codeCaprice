package binaryTree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var invert func(*TreeNode)
	invert = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Right, node.Left = node.Left, node.Right
		invert(node.Left)
		invert(node.Right)
	}
	invert(root)
	return root
}
