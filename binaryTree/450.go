package binaryTree

func deleteNode(root *TreeNode, key int) *TreeNode {
	//递归终止条件
	if root == nil {
		return nil
	}

	if key < root.Val {
		//找左子树
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		//找右子树
		root.Right = deleteNode(root.Right, key)
	} else {
		//找到值
		//判断需要删除的节点是否存在左右子树

		//左子树为空
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		//找到右子树最小节点，或者找到左子树最大节点
		//找到右子树最小节点
		success := root.Right
		for success.Left != nil {
			success = success.Left
		}
		//同时需要移除用来替换root位置的节点
		success.Right = deleteNode(root.Right, success.Val)
		success.Left = root.Left
		return success
	}
	return root
}
