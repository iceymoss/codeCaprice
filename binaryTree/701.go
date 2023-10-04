package binaryTree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	var insert func(node *TreeNode, val int) *TreeNode
	insert = func(node *TreeNode, val int) *TreeNode {
		if node == nil {
			return &TreeNode{
				Val: val,
			}
		}
		//如果当前节点值大于插入值
		if node.Val > val {
			node.Left = insert(node.Left, val)
		} else {
			node.Right = insert(node.Right, val)
		}
		return node
	}
	return insert(root, val)
}
